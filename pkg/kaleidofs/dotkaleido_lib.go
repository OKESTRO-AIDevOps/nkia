package kaleidofs

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func CheckAvailable() bool {

	if !CheckKaleidoRoot() {
		return false
	}

	if !CheckNpiaRoot() {
		return false
	}

	return true
}

func CheckKaleidoRoot() bool {

	if _, err := os.Stat(".kaleido"); err != nil {

		return false

	}

	return true

}

func CheckNpiaRoot() bool {

	if _, err := os.Stat(".etc"); err != nil {

		return false

	}

	if _, err := os.Stat(".usr"); err != nil {

		return false

	}

	if _, err := os.Stat("lib"); err != nil {

		return false

	}

	return true
}

func CheckMode() string {

	if _, err := os.Stat(".npia"); err == nil {
		return ".npia"
	}

	return "none"
}

func GetKubeConfigPath() (string, error) {

	var kube_config_path string

	cmd := exec.Command(".npia/get_kubeconfig_path")

	out, err := cmd.Output()

	if err != nil {

		return "", fmt.Errorf("failed to get kube config path: %s", err.Error())

	}

	strout := string(out)

	ret_strout := strings.ReplaceAll(strout, "\n", "")

	ret_strout = strings.ReplaceAll(ret_strout, " ", "")

	kube_config_path = ret_strout

	return kube_config_path, nil
}

func LoadMultiOrigin() (MultiAppOrigin, error) {

	kalku := make(MultiAppOrigin)

	file_byte, err := os.ReadFile(".kaleido/kube.json")

	if err != nil {

		return kalku, fmt.Errorf("failed to load kaleido kube: %s", err.Error())

	}

	err = json.Unmarshal(file_byte, &kalku)

	if err != nil {
		return kalku, fmt.Errorf("failed to load kaleido kube: %s", err.Error())
	}

	return kalku, nil

}

func UnloadMultiOrigin(kalku MultiAppOrigin) error {

	new_kalku, err := json.Marshal(kalku)

	if err != nil {

		return fmt.Errorf("failed to unload kaleido kube: %s", err.Error())

	}

	err = os.WriteFile(".kaleido/kube.json", new_kalku, 0644)

	if err != nil {

		return fmt.Errorf("failed to unload kaleido kube: %s", err.Error())

	}

	return nil

}
