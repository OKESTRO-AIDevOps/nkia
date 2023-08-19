package apistandard

import (
	"fmt"
	"strings"

	pkgutils "github.com/OKESTRO-AIDevOps/nkia/pkg/utils"
)

type LEGACY_EXCHANGE struct {
	ACD  string
	CMD  string
	DATA string
	CNT  string
	MSG  string
}

type LEGACY_FORM struct {
	FORM LEGACY_API_OUTPUT
}

type LEGACY_API_OUTPUT struct {
	HEAD string

	BODY string
}

func (asgi API_STD) LegacyInputTranslate(legacy_in string) (API_INPUT, error) {

	ret_api_input := make(API_INPUT)

	legacy_list := strings.SplitN(legacy_in, ":", 2)

	match_key := 0

	var err error

	new_key := legacy_list[0]

	c_list := strings.Split(legacy_list[1], ",")

	if new_key == "SUBMIT" {
		match_key = 1

	} else if new_key == "CALLME" {
		match_key = 1

	} else if new_key == "GITLOG" {
		match_key = 1

	} else if new_key == "PIPEHIST" {
		match_key = 1

	} else if new_key == "PIPE" {
		match_key = 1

	} else if new_key == "PIPELOG" {
		match_key = 1

	} else if new_key == "BUILD" {
		match_key = 1

	} else if new_key == "BUILDLOG" {
		match_key = 1

	} else if new_key == "DELND" {
		match_key = 1

	} else if new_key == "EXIT" {
		match_key = 1

	}

	if match_key != 0 {

		legacy_c_list := pkgutils.InsertToSliceByIndex[string](c_list, 0, new_key)

		std_keys, okay := asgi[new_key]

		if !okay {
			return ret_api_input, fmt.Errorf("failed to translate: %s", "key not found")
		}

		ret_api_input, err = asgi.StdCmdInputBuildHelper(std_keys, legacy_c_list)

		if err != nil {

			return ret_api_input, fmt.Errorf("failed to translate: %s", err.Error())

		}

	}

	if new_key == "ADMIN" && match_key != 0 {

		adm_obj, adm_resized_c_list := pkgutils.PopFromSliceByIndex[string](c_list, 0)

		if adm_obj == "ADMRMTCHK" {

		} else if adm_obj == "ADMRMTLDHA" {
			match_key = 1

		} else if adm_obj == "ADMRMTLDMV" {
			match_key = 1

		} else if adm_obj == "ADMRMTMSR" {
			match_key = 1

		} else if adm_obj == "ADMRMTLDWRK" {
			match_key = 1

		} else if adm_obj == "ADMRMTWRK" {
			match_key = 1

		} else if adm_obj == "ADMRMTSTR" {
			match_key = 1

		} else if adm_obj == "ADMRMTLOG" {
			match_key = 1

		} else if adm_obj == "ADMRMTSTATUS" {
			match_key = 1

		} else if adm_obj == "remote-up" {
			match_key = 1
			adm_obj = "UP"
		} else if adm_obj == "remote-down" {
			match_key = 1
			adm_obj = "DOWN"
		} else if adm_obj == "remote-lead" {
			match_key = 1
			adm_obj = "LEAD"
		} else if adm_obj == "remote-master" {
			match_key = 1
			adm_obj = "MSR"
		} else if adm_obj == "remote-lead-worker" {
			match_key = 1
			adm_obj = "LDVOL"
		} else if adm_obj == "remote-worker" {
			match_key = 1
			adm_obj = "WRK"
		} else if adm_obj == "remote-volume" {
			match_key = 1
			adm_obj = "STR"
		} else if adm_obj == "kwadm-log" {
			match_key = 1
			adm_obj = "LOG"
		} else if adm_obj == "kwadm-status" {
			match_key = 1
			adm_obj = "STATUS"
		}

		adm_std_key_obj := ""

		adm_std_key_obj = new_key + "-" + adm_obj

		legacy_c_list := pkgutils.InsertToSliceByIndex[string](adm_resized_c_list, 0, adm_std_key_obj)

		std_keys, okay := asgi[adm_std_key_obj]

		if !okay {
			return ret_api_input, fmt.Errorf("failed to translate: %s", "key not found")
		}

		ret_api_input, err = asgi.StdCmdInputBuildHelper(std_keys, legacy_c_list)

		if err != nil {

			return ret_api_input, fmt.Errorf("failed to translate: %s", err.Error())

		}

	}

	if match_key != 0 {
		return ret_api_input, nil
	}

	obj, resized_c_list := pkgutils.PopFromSliceByIndex[string](c_list, 1)

	std_key_obj := ""

	std_key_obj = new_key + "-" + obj

	legacy_c_list := pkgutils.InsertToSliceByIndex[string](resized_c_list, 0, std_key_obj)

	std_keys, okay := asgi[std_key_obj]

	if !okay {
		return ret_api_input, fmt.Errorf("failed to translate: %s", "key not found")
	}

	ret_api_input, err = asgi.StdCmdInputBuildHelper(std_keys, legacy_c_list)

	if err != nil {

		return ret_api_input, fmt.Errorf("failed to translate: %s", err.Error())

	}

	return ret_api_input, nil

}

func (asgi API_STD) LegacyOutputTranslate() {

}
