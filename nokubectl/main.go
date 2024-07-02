package main

import (
	"crypto/tls"
	"fmt"
	"os"
	"os/exec"

	nkctlclient "github.com/OKESTRO-AIDevOps/nkia/nokubectl/client"
	ccmd "github.com/OKESTRO-AIDevOps/nkia/nokubectl/cmd"
	"github.com/OKESTRO-AIDevOps/nkia/nokubectl/config"
	"github.com/OKESTRO-AIDevOps/nkia/pkg/apistandard/apix"
	goya "github.com/goccy/go-yaml"
	"github.com/gorilla/websocket"
)

func InitCtl() error {

	cmd := exec.Command("mkdir", "-p", ".npia")

	err := cmd.Run()

	if err != nil {

		return fmt.Errorf("failed to init: %s", err.Error())
	}

	if len(os.Args) < 3 {
		return fmt.Errorf("failed to init: %s", "too few arguments")
	}

	cmd = exec.Command("mkdir", "-p", ".npia")

	err = cmd.Run()

	if err != nil {

		return fmt.Errorf("failed to init: %s", err.Error())
	}

	CONFIG_YAML := make(map[string]string)

	if _, err := os.Stat(".npia/config.yaml"); err == nil {

		fmt.Println("existing configyaml has been detected")

		file_b, err := os.ReadFile("./.npia/config.yaml")

		if err == nil {

			fmt.Println("successfully read the existing configyaml")

			err = goya.Unmarshal(file_b, &CONFIG_YAML)

			if err == nil {

				fmt.Println("existing configuration: ")

				for k, v := range CONFIG_YAML {

					fmt.Printf("  %s: %s\n", k, v)

				}

			} else {

				return fmt.Errorf("failed to init: %s", err.Error())

			}

		}

	}

	priv_b, err := os.ReadFile(CONFIG_YAML["PRIV_PATH"])

	if err != nil {

		return fmt.Errorf("failed to init: %s", err.Error())
	}

	_ = os.WriteFile(".npia/.priv", priv_b, 0644)

	return nil

}

func RunClientCmd(args []string) {

	api_input, err := apix.AXgi.BuildAPIInputFromCommandLine(args)

	if err != nil {

		fmt.Printf("failed: %s\n", err.Error())

		return
	}

	forward, result, err := ccmd.RequestForwardHandler(api_input)

	if err != nil {

		fmt.Printf("failed: %s\n", err.Error())

		return
	}

	if !forward {

		fmt.Println(result)

		return

	}

	oreq, err := apix.AXgi.BuildOrchRequestFromCommandLine(args)

	if err != nil {

		fmt.Printf("failed: %s\n", err.Error())

		return
	}

	certpool, err := config.LoadCertAuthCredential()

	if err != nil {

		fmt.Printf("failed: %s\n", err.Error())

		return
	}

	websocket.DefaultDialer.TLSClientConfig = &tls.Config{
		RootCAs: certpool,
	}

	c, _, err := websocket.DefaultDialer.Dial(config.COMM_URL, nil)

	if err != nil {

		fmt.Println(err.Error())

		return
	}
	err = nkctlclient.CertAuthConn(c)

	if err != nil {

		fmt.Println(err.Error())

		return
	}

	// nkctlclient.RequestHandler_APIX_Once_PrintOnly(c, oreq)

	nkctlclient.Do(c, oreq)

}

func ReadClientCmdResult() {

	nkctlclient.Read_APIX_Store_Override()
	return

}

func main() {

	args := os.Args[1:]

	RunClientCmd(args)

}
