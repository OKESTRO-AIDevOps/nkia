package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"os"
	"os/exec"
	"time"

	_ "github.com/OKESTRO-AIDevOps/nkia/nokubeadm/admin"
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

			err = goya.Unmarshal(file_b, CONFIG_YAML)

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

func RunAdmin() {

}

func main() {

	INIT := 0

	INIT_NPIA := 0

	MODE_TEST := 0

	MODE_DEBUG := 0

	MODE_INTERACTIVE := 0

	// MODE_ADMIN := 0

	for i := 0; i < len(os.Args); i++ {

		flag := os.Args[i]

		if flag == "init" {

			INIT = 1

			break
		}

		if flag == "init-npia" {

			INIT_NPIA = 1

			break

		}

		if flag == "-t" || flag == "--test" {

			MODE_TEST = 1

		} else if flag == "-d" || flag == "--debug" {

			MODE_DEBUG = 1

		} else if flag == "-i" || flag == "--interactive" {

			MODE_INTERACTIVE = 1
		}

	}

	if INIT == 1 {

		err := InitAdm()

		if err != nil {

			fmt.Println(err.Error())

		}
		fmt.Println("successfully initiated")
		return
	}

	if INIT_NPIA == 1 {
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

	}

	if (MODE_TEST + MODE_DEBUG + MODE_INTERACTIVE) > 1 {
		fmt.Println("error: more than one option used together")
		return
	}

	if MODE_TEST == 1 {

		nkadmdebug.BaseFlow_APIThenMultiMode_Test()

	} else if MODE_DEBUG == 1 {

		RunDebugInteractive()

	} else if MODE_INTERACTIVE == 1 {

		RunAdminInteractive()

	} else {

		RunAdmin()

	}
}
