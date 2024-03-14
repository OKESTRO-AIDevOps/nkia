package ci

import (
	"fmt"
	"syscall"

	"golang.org/x/term"
	"gopkg.in/yaml.v3"
)

func CITargetsFactory() *CITargetsCtl {

	return &CITargetsCtl{}

}

func CICredFactory() *CICredStore {

	return &CICredStore{}

}

func StoreCICredFromCIFile(ci_cred *CICredStore, mode string) error {

	if mode == "stdin" {

		ci_targets, err := LoadTargetsFromFile()

		if err != nil {
			return fmt.Errorf("failed to store ci cred: %s", err.Error())
		}

		err = StoreCICredStdin(ci_cred, ci_targets)

		if err != nil {
			return fmt.Errorf("failed to stroe ci cred: %s", err.Error())
		}

	} else {

		return fmt.Errorf("store cred mode not availabled for: %s", mode)

	}

	return nil
}

func StoreCICredStdin(ci_cred *CICredStore, ci_targets CITargets) error {

	all_the_way := 0

	q_in := ""

	fmt.Printf("\ndo all repos share the same credentials?: [y/n] ")

	fmt.Scanln(&q_in)

	if q_in == "y" || q_in == "Y" {

		all_the_way = 1

	}

	user_id := ""
	user_pw := ""

	index := 0

	for k, v := range ci_targets {

		if k == "target.v1" {

			for i := 0; i < len(v); i++ {

				ci_target_cred := CITargetCred{
					CI_VERSION_ID: 1,
					CI_TARGET_ID:  i,
				}

				var target_ci TargetV1

				yaml_b, err := yaml.Marshal(v[i])

				if err != nil {
					return fmt.Errorf("failed to marshal: %s", err.Error())
				}

				err = yaml.Unmarshal(yaml_b, &target_ci)

				if err != nil {
					index += 1
					return fmt.Errorf("failed to marshal: %s", err.Error())
				}

				if index > 0 && all_the_way == 1 {

					ci_target_cred.USER_ID = user_id

					ci_target_cred.USER_PW = user_pw

					ci_cred.CI_CRED = append(ci_cred.CI_CRED, ci_target_cred)

					index += 1

					fmt.Printf("\nid:pw automatically set for: %s\n", target_ci.GitPackage.Address)

					continue
				}

				fmt.Printf("\nid for %s: ", target_ci.GitPackage.Address)
				fmt.Scanln(&user_id)

				fmt.Printf("\npw for %s: ", target_ci.GitPackage.Address)

				pw_b, err := term.ReadPassword(int(syscall.Stdin))
				if err != nil {
					return fmt.Errorf("store cred error: %s", err.Error())
				}

				user_pw = string(pw_b)

				ci_target_cred.USER_ID = user_id

				ci_target_cred.USER_PW = user_pw

				ci_cred.CI_CRED = append(ci_cred.CI_CRED, ci_target_cred)

				index += 1

			}

		} else {

			return fmt.Errorf("store not implemented for: %s", k)

		}

	}

	return nil

}
