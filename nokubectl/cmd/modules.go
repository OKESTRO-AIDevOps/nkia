package cmd

import (
	"fmt"

	cctrl "github.com/OKESTRO-AIDevOps/nkia/nokubectl/controller"
	"github.com/OKESTRO-AIDevOps/nkia/pkg/apistandard"
)

func RequestForwardHandler(api_input apistandard.API_INPUT) (bool, string, error) {

	var result string = ""

	ASgi := apistandard.ASgi

	if v_failed := ASgi.Verify(api_input); v_failed != nil {

		return false, "", fmt.Errorf("run failed: %s", v_failed.Error())

	}

	cmd_id := api_input["id"]

	switch cmd_id {

	case "NKCTL-INIT":

		priv_path := api_input["path"]

		err := cctrl.InitCtl(priv_path)

		if err != nil {

			return false, result, fmt.Errorf("run failed: %s", err.Error())

		}

		result = "successfully initiated nokubectl\n"

		return false, result, nil

	default:

		fmt.Println("forward")

	}

	return true, "", nil
}
