package kaleidofs

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/OKESTRO-AIDevOps/nkia/pkg/runtimefs"
	goya "github.com/goccy/go-yaml"
)

func InitKaleidoRoot() error {

	var kube_config_path string

	var kube_config map[interface{}]interface{}

	contexts := make([]string, 0)

	multi_app_origin := make(MultiAppOrigin)

	var blank_app_origin runtimefs.AppOrigin

	var blank_record runtimefs.RecordInfo

	var blank_repo runtimefs.RepoInfo

	var blank_reg runtimefs.RegInfo

	blank_app_origin.RECORDS = append(blank_app_origin.RECORDS, blank_record)

	blank_app_origin.REPOS = append(blank_app_origin.REPOS, blank_repo)

	blank_app_origin.REGS = append(blank_app_origin.REGS, blank_reg)

	if !CheckNpiaRoot() {

		return fmt.Errorf("init multimod failed: %s", "npia has not been initiated")

	}

	if CheckKaleidoRoot() {

		return fmt.Errorf("init multimod failed: %s", ".kaleido already exists")
	}

	cmd := exec.Command("mkdir", "-p", ".kaleido")

	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("init multimod failed: %s", err.Error())
	}

	mode := CheckMode()

	if mode == ".npia" {

		kube_config_path, err = GetKubeConfigPath()

		if err != nil {
			return fmt.Errorf("init multimod failed: %s", err.Error())
		}

	} else {
		return fmt.Errorf("init multimod failed: %s", "not yet implemented for other mods")
	}

	kube_config_file_byte, err := os.ReadFile(kube_config_path)

	if err != nil {
		return fmt.Errorf("init multimod failed: %s", err.Error())
	}

	err = goya.Unmarshal(kube_config_file_byte, &kube_config)

	if err != nil {
		return fmt.Errorf("init multimod failed: %s", err.Error())
	}

	current_context := kube_config["current-context"].(string)

	contexts_len := len(kube_config["contexts"].([]interface{}))

	for i := 0; i < contexts_len; i++ {

		context_nm := kube_config["contexts"].([]interface{})[i].(map[string]interface{})["name"].(string)

		if current_context == context_nm {
			continue
		}

		contexts = append(contexts, context_nm)

	}

	current_app_origin, err := runtimefs.LoadAdmOrigin()

	if err != nil {
		return fmt.Errorf("init multimod failed: %s", err.Error())
	}

	multi_app_origin[current_context] = current_app_origin

	for _, c := range contexts {

		multi_app_origin[c] = blank_app_origin

	}

	file_b, err := json.Marshal(multi_app_origin)

	if err != nil {
		return fmt.Errorf("init multimod failed: %s", err.Error())
	}

	err = os.WriteFile(".kaleido/kube.json", file_b, 0644)

	if err != nil {
		return fmt.Errorf("init multimod failed: %s", err.Error())
	}

	return nil
}

func SaveAndSwitch(switch_to string) error {

	var kube_config_path string

	var kube_config map[interface{}]interface{}

	var err error

	context_hit := 0

	kaleido_kube := make(MultiAppOrigin)

	kaleido_kube, err = LoadMultiOrigin()

	var blank_app_origin runtimefs.AppOrigin

	var blank_record runtimefs.RecordInfo

	var blank_repo runtimefs.RepoInfo

	var blank_reg runtimefs.RegInfo

	blank_app_origin.RECORDS = append(blank_app_origin.RECORDS, blank_record)

	blank_app_origin.REPOS = append(blank_app_origin.REPOS, blank_repo)

	blank_app_origin.REGS = append(blank_app_origin.REGS, blank_reg)

	if err != nil {
		return fmt.Errorf("failed to s&s: %s", err.Error())
	}

	if !CheckAvailable() {
		return fmt.Errorf("failed to s&s: %s", "multimod requirements not satisfied")
	}

	mode := CheckMode()

	if mode == ".npia" {

		kube_config_path, err = GetKubeConfigPath()

		if err != nil {
			return fmt.Errorf("failed to s&s: %s", err.Error())
		}

	} else {
		return fmt.Errorf("failed to s&s: %s", "not yet implemented for other mods")
	}

	kube_config_file_byte, err := os.ReadFile(kube_config_path)

	if err != nil {
		return fmt.Errorf("failed to s&s: %s", err.Error())
	}

	err = goya.Unmarshal(kube_config_file_byte, &kube_config)

	if err != nil {
		return fmt.Errorf("failed to s&s: %s", err.Error())
	}

	current_context := kube_config["current-context"].(string)

	contexts_len := len(kube_config["contexts"].([]interface{}))

	for i := 0; i < contexts_len; i++ {

		context_nm := kube_config["contexts"].([]interface{})[i].(map[string]interface{})["name"].(string)

		if switch_to == context_nm {
			context_hit = 1
		}

	}

	if context_hit == 0 {
		return fmt.Errorf("failed to s&s: %s", "context not found")
	}

	current_app_origin, err := runtimefs.LoadAdmOrigin()

	if err != nil {
		return fmt.Errorf("failed to s&s: %s", err.Error())
	}

	kaleido_kube[current_context] = current_app_origin

	origin_switch_to, okay := kaleido_kube[switch_to]

	if !okay {

		kaleido_kube[switch_to] = blank_app_origin

		origin_switch_to = blank_app_origin
	}

	cmd := exec.Command("kubectl", "config", "use-context", switch_to)

	err = cmd.Run()

	if err != nil {
		return fmt.Errorf("failed to s&s: %s", err.Error())
	}

	err = UnloadMultiOrigin(kaleido_kube)

	if err != nil {

		cmd := exec.Command("kubectl", "config", "use-context", current_context)

		_ = cmd.Run()

		return fmt.Errorf("failed to s&s: %s", err.Error())

	}

	err = runtimefs.UnloadAdmOrigin(origin_switch_to)

	if err != nil {

		cmd := exec.Command("kubectl", "config", "use-context", current_context)

		_ = cmd.Run()

		return fmt.Errorf("failed to s&s: %s", err.Error())

	}

	return nil
}
