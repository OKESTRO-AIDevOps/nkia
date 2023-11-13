package kubebase

import (
	"fmt"
	"os/exec"
	"strings"
)

func ConstructBaseName(def_id string, os_nm string) (string, error) {

	var ret_str string

	def_id_no_hypen := strings.ReplaceAll(def_id, "-", "")

	def_sanitized := strings.ToLower(def_id_no_hypen)

	os_checked, okay := OS_CHECKER[os_nm]

	if !okay {

		return ret_str, fmt.Errorf("base name construct: %s", "no such os: "+os_nm)

	}

	ret_str = def_sanitized + "-" + os_checked

	return ret_str, nil
}

func GetJoinToken() (string, error) {

	var ret_command string

	cmd := exec.Command("kubeadm", "token", "create", "--print-join-command")

	token_b, err := cmd.Output()

	if err != nil {
		return ret_command, fmt.Errorf("failed to get join token: %s", err.Error())
	}

	ret_command = string(token_b)

	return ret_command, nil

}
