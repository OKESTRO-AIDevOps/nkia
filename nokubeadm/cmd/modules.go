package admin

import (
	"fmt"

	"github.com/OKESTRO-AIDevOps/nkia/pkg/apistandard"
	"github.com/OKESTRO-AIDevOps/nkia/pkg/kubebase"
)

func RequestHandler(api_input apistandard.API_INPUT) error {

	var ret_api_out apistandard.API_OUTPUT

	ASgi := apistandard.ASgi

	if v_failed := ASgi.Verify(api_input); v_failed != nil {

		return fmt.Errorf("run failed: %s", v_failed.Error())

	}

	cmd_id := api_input["id"]

	switch cmd_id {

	case "NKADM-INSTCTRL":

		localip := api_input["localip"]
		osnm := api_input["osnm"]
		cv := api_input["cv"]

		cmd_err := kubebase.InstallControlPlane(localip, osnm, cv)

		b_out := []byte("npia install control plane success\n")

		if cmd_err != nil {

			return fmt.Errorf("run failed: %s", cmd_err.Error())

		}

		ret_api_out.BODY = string(b_out)

	//	case "NKADM-INSTANCTRLCRT":

	//	case "NKADM-INSTANCTRLOL":

	//	case "NKADM-INSTANCTRLOR":

	case "NKADM-INSTWKOL":

		localip := api_input["localip"]
		osnm := api_input["osnm"]
		cv := api_input["cv"]
		token := api_input["token"]

		cmd_err := kubebase.InstallWorkerOnLocal(localip, osnm, cv, token)

		b_out := []byte("npia install worker successful\n")

		if cmd_err != nil {

			return fmt.Errorf("run failed: %s", cmd_err.Error())

		}

		ret_api_out.BODY = string(b_out)

	case "NKADM-INSTVOLOL":

		localip := api_input["localip"]

		cmd_err := kubebase.InstallVolumeOnLocal(localip)

		b_out := []byte("npia install volume successful\n")

		if cmd_err != nil {

			return fmt.Errorf("run failed: %s", cmd_err.Error())

		}

		ret_api_out.BODY = string(b_out)

	case "NKADM-INSTTKOL":

		cmd_err := kubebase.InstallToolKitOnLocal()

		b_out := []byte("npia install toolkit successful\n")

		if cmd_err != nil {

			return fmt.Errorf("run failed: %s", cmd_err.Error())

		}

		ret_api_out.BODY = string(b_out)

	case "NKADM-INSTLOGOL":

		b_out, cmd_err := kubebase.InstallLogOnLocal()

		if cmd_err != nil {

			return fmt.Errorf("run failed: %s", cmd_err.Error())

		}

		ret_api_out.BODY = string(b_out)

	default:

		return fmt.Errorf("failed to run api: %s", "invalid command id")

	}

	body := ret_api_out.BODY

	fmt.Println("----------MESSAGE----------")
	fmt.Println(body)

	return nil

}
