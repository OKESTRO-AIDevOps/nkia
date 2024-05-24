package runtimefs

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	goya "github.com/goccy/go-yaml"
)

func InitUsrTarget(repoaddr string) error {

	var app_origin AppOrigin

	if _, err := os.Stat(".usr/target"); err == nil {

		cmd := exec.Command("rm", "-r", ".usr/target")

		cmd.Stdout = os.Stdout

		cmd.Stderr = os.Stderr

		cmd.Run()

		cmd = exec.Command("mkdir", ".usr/target")

		cmd.Stdout = os.Stdout

		cmd.Stderr = os.Stderr

		cmd.Run()
	} else {
		cmd := exec.Command("mkdir", ".usr/target")

		cmd.Stdout = os.Stdout

		cmd.Stderr = os.Stderr

		cmd.Run()

	}

	cmd := exec.Command("git", "-C", ".usr/target", "init")

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	app_origin, err := LoadAdmOrigin()

	if err != nil {
		return fmt.Errorf("failed to init target: %s", err.Error())
	}

	addr_found, rec_repoid, rec_repopw := GetRepoInfo(app_origin.REPOS, repoaddr)

	if !addr_found {

		return fmt.Errorf("failed to init target: %s", "repo info not found")

	}

	insert := "%s:%s@"

	prt_idx := strings.Index(repoaddr, "://")

	prt_idx += 3

	repo_url := repoaddr[:prt_idx] + insert + repoaddr[prt_idx:]

	repo_url = fmt.Sprintf(repo_url, rec_repoid, rec_repopw)

	cmd = exec.Command("git", "-C", ".usr/target", "pull", repo_url)

	_, err = cmd.Output()

	if err != nil {
		return fmt.Errorf("failed to init target: %s", err.Error())
	}

	if _, err := os.Stat(".usr/target/docker-compose.yml"); err == nil {

		cmd = exec.Command("mv", ".usr/target/docker-compose.yml", ".usr/target/docker-compose.yaml")

		cmd.Run()

	}

	if _, err := os.Stat(".usr/target/docker-compose.yaml"); err != nil {

		cmd = exec.Command("rm", "-r", ".usr/target")

		cmd.Run()

		return fmt.Errorf("failed to init target: %s", err.Error())

	}

	return nil

}

func InitUsrTargetForPipeline(repoaddr string) error {

	var app_origin AppOrigin

	if _, err := os.Stat(".usr/target"); err == nil {

		cmd := exec.Command("rm", "-r", ".usr/target")

		cmd.Stdout = os.Stdout

		cmd.Stderr = os.Stderr

		cmd.Run()

		cmd = exec.Command("mkdir", ".usr/target")

		cmd.Stdout = os.Stdout

		cmd.Stderr = os.Stderr

		cmd.Run()
	} else {
		cmd := exec.Command("mkdir", ".usr/target")

		cmd.Stdout = os.Stdout

		cmd.Stderr = os.Stderr

		cmd.Run()

	}

	cmd := exec.Command("git", "-C", ".usr/target", "init")

	cmd.Stdout = os.Stdout

	cmd.Stderr = os.Stderr

	cmd.Run()

	app_origin, err := LoadAdmOrigin()

	if err != nil {
		return fmt.Errorf("failed to init target: %s", err.Error())
	}

	addr_found, rec_repoid, rec_repopw := GetRepoInfo(app_origin.REPOS, repoaddr)

	if !addr_found {

		return fmt.Errorf("failed to init target: %s", "repo info not found")

	}

	insert := "%s:%s@"

	prt_idx := strings.Index(repoaddr, "://")

	prt_idx += 3

	repo_url := repoaddr[:prt_idx] + insert + repoaddr[prt_idx:]

	repo_url = fmt.Sprintf(repo_url, rec_repoid, rec_repopw)

	cmd = exec.Command("git", "-C", ".usr/target", "pull", repo_url)

	_, err = cmd.Output()

	if err != nil {
		return fmt.Errorf("failed to init target: %s", err.Error())
	}

	return nil

}

func ClearUsrTarget() error {

	cmd := exec.Command("rm", "-r", ".usr/target")

	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("failed to clear target: %s", err.Error())
	}

	return nil
}

func CreateUsrTargetOperationSource(LIBIF_BIN_KOMPOSE string, regaddr string) (string, error) {

	var ops_src_list [][]byte
	var ops_src_file []byte

	regaddr_effective := strings.Split(regaddr, "://")[1]

	if _, err := os.Stat(".usr/target"); err != nil {

		return "", fmt.Errorf("failed to create ops src: %s", err.Error())

	}

	if _, err := os.Stat(".usr/target/docker-compose.yaml"); err != nil {

		return "", fmt.Errorf("failed to create ops src: %s", err.Error())
	}

	cmd := exec.Command(LIBIF_BIN_KOMPOSE, "convert", "-f", ".usr/target/docker-compose.yaml", "--stdout")

	out, err := cmd.Output()

	if err != nil {
		return "", fmt.Errorf("failed to create ops src: %s", err.Error())
	}

	var yaml_items []interface{}

	yaml_str := string(out)

	yaml_path_items := "$.items"

	ypath, err := goya.PathString(yaml_path_items)

	if err != nil {
		return "", fmt.Errorf("failed to create ops src: %s", err.Error())
	}

	err = ypath.Read(strings.NewReader(yaml_str), &yaml_items)

	if err != nil {
		return "", fmt.Errorf("failed to create ops src: %s", err.Error())
	}

	for _, val := range yaml_items {

		yaml_if := make(map[interface{}]interface{})

		resource_b, err := goya.Marshal(val)

		err = goya.Unmarshal(resource_b, &yaml_if)

		if err != nil {
			return "", fmt.Errorf("failed to create ops src: %s", err.Error())
		}

		if yaml_if["kind"] == "Deployment" {

			image_pull_secrets := make([]map[string]string, 0)

			value := map[string]string{
				"name": "docker-secret",
			}

			image_pull_secrets = append(image_pull_secrets, value)

			yaml_if["spec"].(map[string]interface{})["template"].(map[string]interface{})["spec"].(map[string]interface{})["imgaePullSecrets"] = image_pull_secrets

			c_count := len(yaml_if["spec"].(map[string]interface{})["template"].(map[string]interface{})["spec"].(map[string]interface{})["containers"].([]interface{}))

			for j := 0; j < c_count; j++ {

				prefix := "/target_"

				reg_fix := regaddr_effective + prefix

				reg_fix += yaml_if["spec"].(map[string]interface{})["template"].(map[string]interface{})["spec"].(map[string]interface{})["containers"].([]interface{})[j].(map[string]interface{})["image"].(string)

				yaml_if["spec"].(map[string]interface{})["template"].(map[string]interface{})["spec"].(map[string]interface{})["containers"].([]interface{})[j].(map[string]interface{})["image"] = reg_fix

				yaml_if["spec"].(map[string]interface{})["template"].(map[string]interface{})["spec"].(map[string]interface{})["containers"].([]interface{})[j].(map[string]interface{})["imagePullPolicy"] = "Always"
			}
		}

		result_b, err := goya.Marshal(yaml_if)

		if err != nil {
			return "", fmt.Errorf("failed to create ops src: %s", err.Error())
		}

		ops_src_list = append(ops_src_list, result_b)

	}

	for i := 0; i < len(ops_src_list); i++ {

		ops_src_file = append(ops_src_file, []byte("---\n")...)

		ops_src_file = append(ops_src_file, ops_src_list[i]...)

	}

	err = os.WriteFile(".usr/ops_src.yaml", ops_src_file, 0644)

	if err != nil {
		return "", fmt.Errorf("failed to create ops src: %s", err.Error())
	}

	return ".usr/ops_src.yaml", nil
}

func CreateUsrDelOperationSource(resourcenm string) (string, error) {

	file_byte, err := os.ReadFile(".usr/ops_src.yaml")

	if err != nil {
		return "", fmt.Errorf("failed to create del src: %s", err.Error())
	}

	var kill_doc [][]byte

	var kill_file []byte

	nm_found := 0

	file_str := string(file_byte)

	file_str_list := strings.Split(file_str, "---\n")

	for _, content := range file_str_list {

		if content == "\n" || content == "" {
			continue
		}

		yaml_if := make(map[interface{}]interface{})

		c_byte := []byte(content)

		err = goya.Unmarshal(c_byte, &yaml_if)

		if err != nil {
			return "", fmt.Errorf("failed to create del src: %s", err.Error())
		}

		if yaml_if["metadata"].(map[string]interface{})["name"] == resourcenm {

			nm_found = 1

			b_tmp, err := goya.Marshal(yaml_if)

			if err != nil {

				return "", fmt.Errorf("failed to create del srd: %s", err.Error())

			}

			kill_doc = append(kill_doc, b_tmp)
		}

	}

	if nm_found == 0 {
		return "", fmt.Errorf("failed to create del src: %s", "matching name not found")
	}

	for _, res_content := range kill_doc {

		kill_file = append(kill_file, []byte("---\n")...)
		kill_file = append(kill_file, res_content...)

	}

	err = os.WriteFile(".usr/del_ops_src.yaml", kill_file, 0644)

	if err != nil {
		return "", fmt.Errorf("failed to create del src: %s", err.Error())
	}

	return ".usr/del_ops_src.yaml", nil
}
