package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"os"
	"os/exec"
	"os/user"
	"time"

	"github.com/OKESTRO-AIDevOps/nkia/nokubectl/apix"

	nkadmcmd "github.com/OKESTRO-AIDevOps/nkia/nokubeadm/cmd"
	"github.com/OKESTRO-AIDevOps/nkia/nokubeadm/config"
	_ "github.com/OKESTRO-AIDevOps/nkia/nokubeadm/debug"
	nkadmdebug "github.com/OKESTRO-AIDevOps/nkia/nokubeadm/debug"
	"github.com/OKESTRO-AIDevOps/nkia/pkg/kubebase"
	goya "github.com/goccy/go-yaml"
)

func InitAdm() error {

	cmd := exec.Command("mkdir", "-p", ".npia")

	err := cmd.Run()

	if err != nil {

		return fmt.Errorf("failed init npia go client: %s", err.Error())
	}

	get_kubeconfig_path_command_string :=
		`#!/bin/bash
[[ ! -z "$KUBECONFIG" ]] && echo "$KUBECONFIG" || echo "$HOME/.kube/config"`

	get_kubeconfig_path_command_b := []byte(get_kubeconfig_path_command_string)

	err = os.WriteFile(".npia/get_kubeconfig_path", get_kubeconfig_path_command_b, 0755)

	if err != nil {

		return fmt.Errorf("failed init npia go client: %s", err.Error())
	}

	cmd = exec.Command("mkdir", "-p", ".npia")

	err = cmd.Run()

	if err != nil {

		return fmt.Errorf("failed to init: %s", err.Error())
	}

	CONFIG_YAML := make(map[string]string)

	if _, err := os.Stat(".npia/config.yaml"); err == nil {

		file_b, err := os.ReadFile(".npia/config.yaml")

		if err == nil {

			err = goya.Unmarshal(file_b, &CONFIG_YAML)

			if err == nil {

				fmt.Println("existing configuration: ")

				yn := "y"

				for k, v := range CONFIG_YAML {

					fmt.Printf("%s: %s\n", k, v)

				}

				fmt.Println("use the existing conf ? : [ y | n ]")

				fmt.Scanln(&yn)

				if yn == "y" || yn == "Y" {

					return nil

				}

			}

		}

	}

	var MODE string
	var BASE_URL string
	var EMAIL string

	fmt.Println("MODE: ")

	fmt.Scanln(&MODE)

	fmt.Println("BASE_URL: ")

	fmt.Scanln(&BASE_URL)

	fmt.Println("EMAIL: ")

	fmt.Scanln(&EMAIL)

	CONFIG_YAML["MODE"] = MODE

	CONFIG_YAML["BASE_URL"] = BASE_URL

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

func InitAdmDefault() error {

	if _, err := os.Stat(".npia/config.yaml"); err != nil {

		return fmt.Errorf("no default config yaml exists")

	} else {

		file_b, err := os.ReadFile(".npia/config.yaml")

		CONFIG_YAML := make(map[string]string)

		if err == nil {

			err = goya.Unmarshal(file_b, &CONFIG_YAML)

			if err == nil {

				fmt.Println("existing configuration: ")

				for k, v := range CONFIG_YAML {

					fmt.Printf("%s: %s\n", k, v)

				}

			} else {

				return fmt.Errorf("failed marshalling config yaml: %s", err.Error())
			}

		} else {

			return fmt.Errorf("failed reading config yaml: %s", err.Error())

		}

	}

	return nil
}

func RunDebugInteractive() {

	var in_raw_query string

	var err error

	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{
		Jar: jar,
	}

	err = nkadmdebug.ClientAuthChallenge(client)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("help: available queries")
	fmt.Println("exit: exit")

	for {

		fmt.Println("query: ")

		fmt.Scanln(&in_raw_query)

		switch in_raw_query {

		case "help":

			config.ASgi_CliRef.PrintPrettyDefinition()

		case "exit":

			return

		default:

			nkadmdebug.CommunicationHandler_LinearInstruction_PrintOnly(client, in_raw_query)

		}

	}

}

func RunAdminInteractive() {

}

func RunAdminCmd(args []string) {

	apistd_in, err := apix.AXgi.BuildAPIInputFromCommandLine(args)

	if err != nil {
		fmt.Printf("failed: %s\n", err.Error())

		return
	}

	err = nkadmcmd.RequestHandler(apistd_in)

	if err != nil {
		fmt.Printf("failed: %s\n", err.Error())

		return
	}

}

func main() {

	current_user, err := user.Current()

	if err != nil {
		fmt.Println(err.Error())

		return

	}

	if current_user.Username != "root" {

		fmt.Println("Error: not running as root")
		return

	}

	flag, args, err := apix.GetNKADMFlagAndReduceArgs()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if flag == "help" {

	} else if flag == "init" {

		err := InitAdm()

		if err != nil {

			fmt.Println(err.Error())

		}
		fmt.Println("successfully initiated")
		return

	} else if flag == "init-npia" {

		err_init := InitAdm()

		if err_init != nil {
			fmt.Println(err_init.Error())
		}

		go kubebase.AdminInitNPIA()

		t_start := time.Now()

		done := 0

		for time.Now().Sub(t_start).Seconds() < 30 {

			if _, err := os.Stat("npia_init_done"); err == nil {

				done = 1
				break

			}

		}

		if done == 1 {
			b, _ := kubebase.AdminGetInitLog()

			fmt.Println("-----INITLOG-----")

			fmt.Println(string(b))

			fmt.Println("-----------------")

			fmt.Println("successfully initiated")
		} else {

			b, _ := kubebase.AdminGetInitLog()

			fmt.Println("-----FAILLOG-----")

			fmt.Println(string(b))

			fmt.Println("-----------------")

			fmt.Println("initiation timeout")
		}
		return

	} else if flag == "init-npia-default" {

		err_init := InitAdmDefault()

		if err_init != nil {
			fmt.Println(err_init.Error())
		}

		go kubebase.AdminInitNPIA()

		t_start := time.Now()

		done := 0

		for time.Now().Sub(t_start).Seconds() < 30 {

			if _, err := os.Stat("npia_init_done"); err == nil {

				done = 1
				break

			}

		}

		if done == 1 {
			b, _ := kubebase.AdminGetInitLog()

			fmt.Println("-----INITLOG-----")

			fmt.Println(string(b))

			fmt.Println("-----------------")

			fmt.Println("successfully initiated")
		} else {

			b, _ := kubebase.AdminGetInitLog()

			fmt.Println("-----FAILLOG-----")

			fmt.Println(string(b))

			fmt.Println("-----------------")

			fmt.Println("initiation timeout")
		}
		return

	} else if flag == "interactive" {

		RunAdminInteractive()

	} else if flag == "debug" {

		RunDebugInteractive()

	} else {
		RunAdminCmd(args)
	}

	// nkadmdebug.BaseFlow_APIThenMultiMode_Test()

}
