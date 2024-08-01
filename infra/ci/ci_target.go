package ci

import (
	"fmt"
	"os"
	"time"

	infutils "github.com/OKESTRO-AIDevOps/nkia/infra/utils"
	"gopkg.in/yaml.v3"
)

func LoadTargetsFromFile() (CITargets, error) {

	ci_targets := make(CITargets)

	var tmp_yaml map[string][]interface{}

	if _, err := os.Stat(".npia.infra/ci.yaml"); err != nil {

		return ci_targets, fmt.Errorf("npia ci.yaml not found")

	}

	file_b, err := os.ReadFile(".npia.infra/ci.yaml")

	if err != nil {

		return ci_targets, fmt.Errorf("failed to read nkia ci.yaml")

	}

	err = yaml.Unmarshal(file_b, &tmp_yaml)

	if err != nil {

		return ci_targets, fmt.Errorf("failed to unmarshal ci.yaml")
	}

	for k, v := range tmp_yaml {

		ci_targets[k] = v

	}

	return ci_targets, nil
}

func VerifyTargetProperty(ci_targets CITargets) error {

	var unique_names []string

	for k, v := range ci_targets {

		if k == "target.v1" {

			for i := 0; i < len(v); i++ {

				yaml_b, err := yaml.Marshal(v[i])

				if err != nil {

					return fmt.Errorf("failed to marshal at: version 1, target %d", i)

				}

				var target_ci TargetV1

				err = yaml.Unmarshal(yaml_b, &target_ci)

				if err != nil {

					return fmt.Errorf("failed to unmarshal at: version 1, target %d", i)

				}

				t_name := target_ci.GitPackage.Name

				flag := infutils.FindFromSlice[string](unique_names, t_name)

				if flag >= 0 {

					return fmt.Errorf("failed to verify: duplicate name: %s", t_name)

				}

				unique_names = append(unique_names, t_name)

			}
		}

	}

	return nil
}

func StartTargetsFromCIFile(ci_targets_ctl *CITargetsCtl, ci_cred *CICredStore) error {

	ci_targets, err := LoadTargetsFromFile()

	if err != nil {
		return fmt.Errorf("failed to start from ci file: %s", err.Error())
	}

	err = VerifyTargetProperty(ci_targets)

	if err != nil {
		return fmt.Errorf("failed to verify: %s", err.Error())
	}

	target_idx := 0

	for k, v := range ci_targets {

		if k == "target.v1" {

			for i := 0; i < len(v); i++ {

				yaml_b, err := yaml.Marshal(v[i])

				if err != nil {
					fmt.Println(err.Error())
					return fmt.Errorf("failed to marshal: %s: %s", k, err.Error())
				}

				cmd_to := make(chan int)

				stat_from := make(chan int)

				cit := CITargetCtl{
					CI_VERSION_ID: 1,
					CI_TARGET_ID:  i,
					CI_CMD_TO:     cmd_to,
					CI_STAT_FROM:  stat_from,
					CI_DONE:       0,
					CI_LOG_PTR:    0,
					CI_LOG:        []string{},
					CI_ERRLOG_PTR: 0,
					CI_ERRLOG:     []string{},
					CI_PLAYBOOK:   yaml_b,
				}

				if err := SetCredForCITarget(&cit, ci_cred); err != nil {
					return fmt.Errorf("failed to set cred: %s", err.Error())
				}

				ci_targets_ctl.CI_TARGETS = append(ci_targets_ctl.CI_TARGETS, cit)

				ci_targets_ctl.CI_ONGOING = append(ci_targets_ctl.CI_ONGOING, CITargetID{
					CI_VERSION_ID: 1,
					CI_TARGET_ID:  i,
				})

				time.Sleep(time.Millisecond * 100)

				go ActionV1(ci_targets_ctl, target_idx)

				target_idx += 1

			}

		} else {

			return fmt.Errorf("action not implemented for: %s", k)

		}

	}

	return nil
}

func SetCredForCITarget(cit *CITargetCtl, ci_cred *CICredStore) error {

	cred_store_len := len(ci_cred.CI_CRED)

	hit := 0

	for i := 0; i < cred_store_len; i++ {

		if cit.CI_VERSION_ID == ci_cred.CI_CRED[i].CI_VERSION_ID && cit.CI_TARGET_ID == ci_cred.CI_CRED[i].CI_TARGET_ID {

			cit.CI_USER_ID = ci_cred.CI_CRED[i].USER_ID

			cit.CI_USER_PW = ci_cred.CI_CRED[i].USER_PW

			cit.CI_USER_EMAIL = ci_cred.CI_CRED[i].USER_EMAIL

			hit = 1

			break
		}

	}

	if hit == 0 {

		return fmt.Errorf("cred not found for: version: %d, target: %d", cit.CI_VERSION_ID, cit.CI_TARGET_ID)
	}

	return nil
}

func TargetsResolver(target_ctl *CITargetsCtl) {

	ticker := time.NewTicker(100 * time.Millisecond)

	finish := 0

	step := -1

	tidx := 0

	for {

		finish, tidx = GetFinishSignalOrTargetIdx(target_ctl, &step)

		if finish == 1 {
			break
		}

		counter := 0

		for counter < 10 {

			select {

			case tc := <-ticker.C:

				_ = tc

				counter += 1

			case stat := <-target_ctl.CI_TARGETS[tidx].CI_STAT_FROM:

				v := target_ctl.CI_TARGETS[tidx].CI_VERSION_ID

				t := target_ctl.CI_TARGETS[tidx].CI_TARGET_ID

				if stat == SIG_FAILED {

					ResolveTarget(&target_ctl.CI_FAILED, v, t, &target_ctl.CI_ONGOING)

				}

				if stat == SIG_SUCCESS {

					ResolveTarget(&target_ctl.CI_SUCCESS, v, t, &target_ctl.CI_ONGOING)

				}

				target_ctl.CI_TARGETS[tidx].CI_DONE = 1

				counter = 10

			}

		}

	}

}

func GetFinishSignalOrTargetIdx(target_ctl *CITargetsCtl, last_step *int) (int, int) {

	CILOCK.Lock()

	target_len := len(target_ctl.CI_TARGETS)

	finish := 0

	done_count := 0

	tidx := 0

	for i := 0; i < target_len; i++ {

		if target_ctl.CI_TARGETS[i].CI_DONE != 1 && i <= *last_step {

			continue

		} else if target_ctl.CI_TARGETS[i].CI_DONE != 1 && i > *last_step {

			*last_step = i

			tidx = i

			break

		} else if target_ctl.CI_TARGETS[i].CI_DONE == 1 {

			done_count += 1

		}

	}

	if *last_step == target_len-1 {

		*last_step = -1
	}

	if done_count == target_len {

		finish = 1
	}

	CILOCK.Unlock()

	return finish, tidx

}

func ResolveTarget(result *[]CITargetID, vid int, tid int, on_going *[]CITargetID) {

	CILOCK.Lock()

	target_len := len(*on_going)

	idx := -1

	for i := 0; i < target_len; i++ {

		if (*on_going)[i].CI_VERSION_ID == vid && (*on_going)[i].CI_TARGET_ID == tid {

			idx = i

			break

		}

	}

	if idx == -1 {
		CILOCK.Unlock()
		return
	}

	(*result) = infutils.PushBackSlice[CITargetID](*result, (*on_going)[idx])

	(*on_going) = infutils.DeleteFromSlice[CITargetID](*on_going, idx)

	CILOCK.Unlock()
}

func TargetsControllerStdin(target_ctl *CITargetsCtl) error {

	go TargetsResolver(target_ctl)

	for {

		action := ""

		fmt.Printf("Which action do you want to perform?: ")

		fmt.Scanln(&action)

		if action == "read-all" {

			ReadAll(target_ctl)

		} else if action == "status" {

			Status(target_ctl)

		} else if action == "exit" {

			fmt.Println("EXIT.")

			break

		} else if action == "help" {

			helpComment := "" +

				"read-all: read all status including verbose log\n" +
				"status : read consise status of each target\n" +
				"exit: finish or abort"

			fmt.Println(helpComment)

		} else {

			fmt.Printf("invalid action: %s\n", action)

		}

	}

	return nil
}

func ReadAll(target_ctl *CITargetsCtl) {

	CILOCK.Lock()

	target_len := len(target_ctl.CI_TARGETS)

	for i := 0; i < target_len; i++ {

		normal_log := target_ctl.CI_TARGETS[i].CI_LOG

		err_log := target_ctl.CI_TARGETS[i].CI_ERRLOG

		normal_log_len := len(normal_log)

		err_log_len := len(err_log)

		v := target_ctl.CI_TARGETS[i].CI_VERSION_ID

		t := target_ctl.CI_TARGETS[i].CI_TARGET_ID

		fmt.Printf("LOG FOR: version: %d, target: %d \n", v, t)

		for j := 0; j < normal_log_len; j++ {

			fmt.Println(normal_log[j])

		}

		fmt.Printf("LOG END: version: %d, target: %d \n", v, t)

		fmt.Printf("ERR FOR: version: %d, target: %d \n", v, t)

		for j := 0; j < err_log_len; j++ {

			fmt.Println(err_log[j])

		}

		fmt.Printf("ERR END: version: %d, target: %d \n", v, t)

	}

	CILOCK.Unlock()
}

func Status(target_ctl *CITargetsCtl) {

	CILOCK.Lock()

	ongoing_len := len(target_ctl.CI_ONGOING)

	success_len := len(target_ctl.CI_SUCCESS)

	failed_len := len(target_ctl.CI_FAILED)

	for i := 0; i < ongoing_len; i++ {

		v := target_ctl.CI_ONGOING[i].CI_VERSION_ID

		t := target_ctl.CI_ONGOING[i].CI_TARGET_ID

		fmt.Printf("ONGOING : version: %d, target: %d \n", v, t)

	}

	for i := 0; i < success_len; i++ {

		v := target_ctl.CI_SUCCESS[i].CI_VERSION_ID

		t := target_ctl.CI_SUCCESS[i].CI_TARGET_ID

		fmt.Printf("SUCCESS : version: %d, target: %d \n", v, t)

	}

	for i := 0; i < failed_len; i++ {

		v := target_ctl.CI_FAILED[i].CI_VERSION_ID

		t := target_ctl.CI_FAILED[i].CI_TARGET_ID

		fmt.Printf("FAILED : version: %d, target: %d \n", v, t)

	}

	CILOCK.Unlock()
}
