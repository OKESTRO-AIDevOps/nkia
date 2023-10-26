package runtimefs

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func LoadAdmOrigin() (AppOrigin, error) {

	var ao AppOrigin

	file_byte, err := os.ReadFile(".etc/ADM_origin.json")

	if err != nil {

		return ao, fmt.Errorf("failed to load admin origin: %s", err.Error())

	}

	err = json.Unmarshal(file_byte, &ao)

	if err != nil {
		return ao, fmt.Errorf("failed to load admin origin: %s", err.Error())
	}

	return ao, nil

}

func UnloadAdmOrigin(ao AppOrigin) error {

	new_ao, err := json.Marshal(ao)

	if err != nil {

		return fmt.Errorf("failed to unload admin origin: %s", err.Error())

	}

	err = os.WriteFile(".etc/ADM_origin.json", new_ao, 0644)

	if err != nil {

		return fmt.Errorf("failed to unload admin origin: %s", err.Error())

	}

	return nil

}

func CreateAdmOrigin() error {

	var ao AppOrigin

	var ri RecordInfo

	var rep RepoInfo

	var reg RegInfo

	if _, err := os.Stat(".etc/ADM_origin.json"); os.IsExist(err) {

		return fmt.Errorf("failed to create admin origin: %s", ".etc/ADM_origin.json already exists")

	}

	ao.RECORDS = append(ao.RECORDS, ri)

	ao.REPOS = append(ao.REPOS, rep)

	ao.REGS = append(ao.REGS, reg)

	new_ao, err := json.Marshal(ao)

	if err != nil {

		return fmt.Errorf("failed to create admin origin: %s", err.Error())
	}

	err = os.WriteFile(".etc/ADM_origin.json", new_ao, 0644)

	if err != nil {

		return fmt.Errorf("failed to create admin origin: %s", err.Error())
	}

	return nil

}

func SetAdminOriginNewNS(ns string, repo_url_in string, reg_url_in string) error {

	app_origin, err := LoadAdmOrigin()

	if err != nil {

		return fmt.Errorf("failed to set new record: %s", err.Error())

	}

	//repo_url, reg_url := GetRecordInfo(app_origin.RECORDS, ns)

	cmd := exec.Command("kubectl", "create", "namespace", ns)

	_, err = cmd.Output()

	if err != nil {

		return fmt.Errorf("failed to set new record: %s", err.Error())
	}

	app_origin.RECORDS = SetRecordInfo(app_origin.RECORDS, ns, repo_url_in, reg_url_in)

	err = UnloadAdmOrigin(app_origin)

	if err != nil {

		return fmt.Errorf("failed to set new record: %s", err.Error())
	}

	return nil

}

func GetRecordInfo(records []RecordInfo, ns string) (bool, string, string) {

	arr_leng := len(records)

	var repo_addr string = "N"

	var reg_addr string = "N"

	ns_found := false

	for i := 0; i < arr_leng; i++ {

		if records[i].NS == ns {

			ns_found = true

			repo_addr = records[i].REPO_ADDR

			reg_addr = records[i].REG_ADDR

			break

		}

	}

	return ns_found, repo_addr, reg_addr

}

func SetRecordInfo(records []RecordInfo, ns string, repo_addr string, reg_addr string) []RecordInfo {

	exists := 0

	arr_leng := len(records)

	var new_record_info RecordInfo

	for i := 0; i < arr_leng; i++ {

		if records[i].NS == ns {

			exists = 1

			records[i].REPO_ADDR = repo_addr

			records[i].REG_ADDR = reg_addr

			break

		}
	}

	if exists != 1 {

		new_record_info.NS = ns

		new_record_info.REPO_ADDR = repo_addr

		new_record_info.REG_ADDR = reg_addr

		records = append(records, new_record_info)

	}

	return records

}

func GetRepoInfo(repos []RepoInfo, addr string) (bool, string, string) {

	arr_leng := len(repos)

	var repo_id string = "N"

	var repo_pw string = "N"

	addr_found := false

	for i := 0; i < arr_leng; i++ {

		if repos[i].REPO_ADDR == addr {

			addr_found = true

			repo_id = repos[i].REPO_ID

			repo_pw = repos[i].REPO_PW

			break

		}
	}

	return addr_found, repo_id, repo_pw

}

func SetRepoInfo(repos []RepoInfo, addr string, id string, pw string) []RepoInfo {

	exists := 0

	arr_leng := len(repos)

	repo_id := id

	repo_pw := pw

	var new_repo_info RepoInfo

	for i := 0; i < arr_leng; i++ {

		if repos[i].REPO_ADDR == addr {

			exists = 1

			repos[i].REPO_ID = repo_id

			repos[i].REPO_PW = repo_pw

			break

		}
	}

	if exists != 1 {

		new_repo_info.REPO_ADDR = addr

		new_repo_info.REPO_ID = repo_id

		new_repo_info.REPO_PW = repo_pw

		repos = append(repos, new_repo_info)

	}

	return repos

}

func GetRegInfo(regs []RegInfo, addr string) (bool, string, string) {

	arr_leng := len(regs)

	var reg_id string = "N"

	var reg_pw string = "N"

	addr_found := false

	for i := 0; i < arr_leng; i++ {

		if regs[i].REG_ADDR == addr {

			addr_found = true

			reg_id = regs[i].REG_ID

			reg_pw = regs[i].REG_PW

			break

		}
	}

	return addr_found, reg_id, reg_pw

}

func SetRegInfo(regs []RegInfo, addr string, id string, pw string) []RegInfo {

	exists := 0

	arr_leng := len(regs)

	reg_id := id

	reg_pw := pw

	var new_reg_info RegInfo

	for i := 0; i < arr_leng; i++ {

		if regs[i].REG_ADDR == addr {

			exists = 1

			regs[i].REG_ID = reg_id

			regs[i].REG_PW = reg_pw

			break

		}
	}

	if exists != 1 {

		new_reg_info.REG_ADDR = addr

		new_reg_info.REG_ID = reg_id

		new_reg_info.REG_PW = reg_pw

		regs = append(regs, new_reg_info)

	}

	return regs
}

func CheckAppOrigin() (string, error) {

	app_origin, err := LoadAdmOrigin()

	if err != nil {

		return "ERRLOAD", fmt.Errorf("check failed: %s", err.Error())

	}

	cmd := exec.Command("kubectl", "get", "nodes")

	_, err = cmd.Output()

	if err != nil {

		return "ERRKUBE", fmt.Errorf("check failed: %s", err.Error())

	}

	if len(app_origin.RECORDS) == 0 {

		return "WARNRC", nil

	}

	if len(app_origin.REGS) == 0 || len(app_origin.REPOS) == 0 {

		return "WARNRE", nil

	}

	return "OKAY", nil
}

func CheckKubeNS(ns string) error {

	cmd := exec.Command("kubectl", "get", "namespace", ns)

	_, err := cmd.Output()

	if err != nil {

		return fmt.Errorf("check failed: %s", err.Error())

	}

	return nil

}
