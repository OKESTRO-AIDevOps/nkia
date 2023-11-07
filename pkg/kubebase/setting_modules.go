package kubebase

import (
	"fmt"
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
