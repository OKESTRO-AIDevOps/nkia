package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"os"
	"os/exec"

	"github.com/OKESTRO-AIDevOps/nkia/nokubectl/apix"
	nkctlclient "github.com/OKESTRO-AIDevOps/nkia/nokubectl/client"
	"github.com/OKESTRO-AIDevOps/nkia/nokubectl/config"
	goya "github.com/goccy/go-yaml"
)

func InitCtl() error {

	var priv_loc string

	var file_b []byte

	cmd := exec.Command("mkdir", "-p", ".npia")

	err := cmd.Run()

	if err != nil {

		return fmt.Errorf("failed to init: %s", err.Error())
	}

	fmt.Println("enter privkey filepath: ")

	fmt.Scanln(&priv_loc)

	file_b, err = os.ReadFile(priv_loc)

	if err != nil {
		return fmt.Errorf("failed to init: %s", err.Error())
	}

	err = os.WriteFile(".npia/.priv", file_b, 0644)

	if err != nil {
		return fmt.Errorf("failed to init: %s", err.Error())
	}

	cmd = exec.Command("mkdir", "-p", ".npia")

	err = cmd.Run()

	if err != nil {

		return fmt.Errorf("failed to init: %s", err.Error())
	}

	cmd = exec.Command("mkdir", "-p", ".npia/_apix_o")

	err = cmd.Run()

	if err != nil {

		return fmt.Errorf("failed to init: %s", err.Error())
	}

	CONFIG_YAML := make(map[string]string)

	if _, err := os.Stat("./.npia/config.yaml"); err == nil {

		fmt.Println("existing configyaml has been detected")

		file_b, err := os.ReadFile("./.npia/config.yaml")

		if err == nil {

			fmt.Println("successfully read the existing configyaml")

			err = goya.Unmarshal(file_b, &CONFIG_YAML)

			if err == nil {

				fmt.Println("existing configuration: ")

				yn := "y"

				for k, v := range CONFIG_YAML {

					fmt.Printf("  %s: %s\n", k, v)

				}

				fmt.Println("use the existing conf ? : [ y | n ]")

				fmt.Scanln(&yn)

				if yn == "y" || yn == "Y" {

					return nil

				}

			} else {
				fmt.Println(err.Error())
			}

		}

	}

	var MODE string
	var BASE_URL string
	var BASE_URL_SOCK string
	var EMAIL string

	fmt.Println("MODE: ")

	fmt.Scanln(&MODE)

	fmt.Println("BASE_URL: ")

	fmt.Scanln(&BASE_URL)

	fmt.Println("BASE_URL_SOCK: ")

	fmt.Scanln(&BASE_URL_SOCK)

	fmt.Println("EMAIL: ")

	fmt.Scanln(&EMAIL)

	CONFIG_YAML["MODE"] = MODE

	CONFIG_YAML["BASE_URL"] = BASE_URL

	CONFIG_YAML["BASE_URL_SOCK"] = BASE_URL_SOCK

	CONFIG_YAML["EMAIL"] = EMAIL

	outconf, err := goya.Marshal(CONFIG_YAML)

	if err != nil {
		return fmt.Errorf("failed to init: %s", err.Error())
	}

	err = os.WriteFile(".npia/config.yaml", outconf, 0644)

	if err != nil {
		return fmt.Errorf("failed to init: %s", err.Error())
	}

	return nil
}

func RunClientInteractive() {

	var email string

	var in_raw_query string

	var err error

	jar, err := cookiejar.New(nil)

	if err != nil {

		fmt.Println(err.Error())

		return

	}

	email = config.EMAIL

	client := &http.Client{
		Jar: jar,
	}

	c, err := nkctlclient.KeyAuthConn(client, email)

	if err != nil {

		fmt.Println(err.Error())

		return
	}

	for {

		fmt.Println("query: ")

		fmt.Scanln(&in_raw_query)

		switch in_raw_query {

		case "exit":

			return

		default:

			var target string

			var option string

			fmt.Println("target: ")

			fmt.Scanln(&target)

			fmt.Println("option: ")

			fmt.Scanln(&option)

			nkctlclient.RequestHandler_LinearInstruction_Persist_PrintOnly(c, target, option, in_raw_query)

		}

	}

}

func RunClientCmd(args []string) {

	var email string

	var err error

	jar, err := cookiejar.New(nil)

	if err != nil {

		fmt.Println(err.Error())

		return

	}

	email = config.EMAIL

	client := &http.Client{
		Jar: jar,
	}

	c, err := nkctlclient.KeyAuthConn(client, email)

	if err != nil {

		fmt.Println(err.Error())

		return
	}

	oreq, err := apix.AXgi.BuildOrchRequestFromCommandLine(args)

	if err != nil {

		fmt.Printf("failed: %s\n", err.Error())

		return
	}

	// nkctlclient.RequestHandler_APIX_Once_PrintOnly(c, oreq)

	nkctlclient.RequestHandler_APIX_Store_Override(c, oreq)

}

func ReadClientCmdResult() {

	nkctlclient.Read_APIX_Store_Override()
	return

}

func main() {

	flag, args, err := apix.GetNKCTLFlagAndReduceArgs()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if flag == "help" {

	} else if flag == "apix-md" {

		apix.ExportMD()
		return

	} else if flag == "apix-js" {

	} else if flag == "apix-py" {

	} else if flag == "init" {

		err := InitCtl()

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("successfully initiated")
		return

	} else if flag == "interactive" {

		RunClientInteractive()
		return

	} else if flag == "read" {

		ReadClientCmdResult()
		return

	} else {

		RunClientCmd(args)
		return

	}

	//if (MODE_INTERACTIVE) > 1 {
	//	fmt.Println("error: more than one option used together")
	//	return
	//}

	return
}
