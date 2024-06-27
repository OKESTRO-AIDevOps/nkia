package cmd

import (
	"fmt"
	"os"

	"github.com/OKESTRO-AIDevOps/nkia/nokubelet/config"
	sock "github.com/OKESTRO-AIDevOps/nkia/nokubelet/oagent"
	"github.com/OKESTRO-AIDevOps/nkia/pkg/apistandard"
)

func RequestHandler(api_input apistandard.API_INPUT) error {

	address := config.ADDRESS

	email := config.EMAIL

	ASgi := apistandard.ASgi

	if v_failed := ASgi.Verify(api_input); v_failed != nil {

		return fmt.Errorf("run failed: %s", v_failed.Error())

	}

	cmd_id := api_input["id"]

	switch cmd_id {

	case "NKLET-CONNUP":

		clusterid := api_input["clusterid"]
		updatetoken, ut_okay := api_input["updatetoken"]

		_ = os.WriteFile(".npia/cluster_id", []byte(clusterid), 0644)

		if config.MODE == "test" && config.DEBUG == "true" && ut_okay {

			if err := sock.DetachedServerCommunicatorWithUpdate_Test_Debug(address, email, clusterid, updatetoken); err != nil {

				return err
			}

		} else if config.MODE == "test" && config.DEBUG != "true" && ut_okay {

			if err := sock.DetachedServerCommunicatorWithUpdate_Test(address, email, clusterid, updatetoken); err != nil {

				return err
			}

		} else if config.MODE == "test" && config.DEBUG == "true" && !ut_okay {

			if err := sock.DetachedServerCommunicator_Test_Debug(address, email, clusterid); err != nil {

				return err
			}

		} else if config.MODE == "test" && config.DEBUG != "true" && !ut_okay {

			if err := sock.DetachedServerCommunicator_Test(address, email, clusterid); err != nil {

				return err
			}

		}

	case "NKLET-CONN":

		return fmt.Errorf("failed to run api: %s", "not implemented")

	default:

		return fmt.Errorf("failed to run api: %s", "invalid command id")
	}

	return nil
}
