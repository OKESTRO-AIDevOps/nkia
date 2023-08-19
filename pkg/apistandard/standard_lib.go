package apistandard

import (
	"fmt"
	"strings"

	pkgutils "github.com/OKESTRO-AIDevOps/nkia/pkg/utils"
)

func (asgi API_STD) Verify(verifiable API_INPUT) error {

	cmd_id, okay := verifiable["id"]

	var duplicate_check []string

	if !okay {
		return fmt.Errorf("verification failed: %s", "missing command id")
	}

	v_list, okay := asgi[cmd_id]

	if !okay {

		return fmt.Errorf("verification failed: %s", "invalid command id")
	}

	if len(v_list) != len(verifiable) {
		return fmt.Errorf("verification failed: %s", "invalid command structure")
	}

	for i := range verifiable {

		hit := pkgutils.CheckIfSliceContains[string](duplicate_check, i)

		if hit {
			return fmt.Errorf("verification failed: %s", "invalid command structure: duplicate key")
		}

		hit = pkgutils.CheckIfSliceContains[string](v_list, i)

		if !hit {
			return fmt.Errorf("verification failed: %s", "invalid command structure: wrong key")
		}

		duplicate_check = append(duplicate_check, i)
	}

	return nil

}

func (asgi API_STD) StdCmdInputBuildFromLinearInstruction(linear_string string) (API_INPUT, error) {

	var ret_api_std API_INPUT

	linear_split := strings.SplitN(linear_string, ":", 2)

	linear_key := linear_split[0]

	std_keys, okay := asgi[linear_key]

	if !okay {
		return ret_api_std, fmt.Errorf("failed to interpret linear instruction: %s", "matching key not found")
	}

	linear_value_split := strings.Split(linear_split[1], ",")

	linear_value_split = pkgutils.InsertToSliceByIndex[string](linear_value_split, 0, linear_key)

	ret_api_std, err := asgi.StdCmdInputBuildHelper(std_keys, linear_value_split)

	if err != nil {
		return ret_api_std, fmt.Errorf("failed to interpret linear instruction: %s", err.Error())
	}

	return ret_api_std, nil
}

func (asgi API_STD) StdCmdInputBuildHelper(std_keys []string, v_list []string) (API_INPUT, error) {

	ret_api_std := make(API_INPUT)

	for i := 0; i < len(std_keys); i++ {

		ret_api_std[std_keys[i]] = v_list[i]

	}

	return ret_api_std, nil
}

func (asgi API_STD) PrintPrettyDefinition() {

	fmt.Println(API_DEFINITION)
}

func (asgi API_STD) PrintRawDefinition() {

	fmt.Println(asgi)

}
