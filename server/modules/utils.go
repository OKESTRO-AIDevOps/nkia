package modules

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os/exec"
	"strings"
)

func RandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func GetKubeConfigPath() (string, error) {

	var kube_config_path string

	cmd := exec.Command("srv/get_kubeconfig_path")

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
