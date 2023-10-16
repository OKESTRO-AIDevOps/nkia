package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"os"
	"os/exec"

	_ "github.com/OKESTRO-AIDevOps/nkia/nokubeadm/admin"
	_ "github.com/OKESTRO-AIDevOps/nkia/nokubeadm/debug"
	nkadmdebug "github.com/OKESTRO-AIDevOps/nkia/nokubeadm/debug"
)

func InitGoClient() error {

	cmd := exec.Command("mkdir", "-p", "srv")

	err := cmd.Run()

	if err != nil {

		return fmt.Errorf("failed init npia go client: %s", err.Error())
	}

	get_kubeconfig_path_command_string :=
		`#!/bin/bash
[[ ! -z "$KUBECONFIG" ]] && echo "$KUBECONFIG" || echo "$HOME/.kube/config"`

	get_kubeconfig_path_command_b := []byte(get_kubeconfig_path_command_string)

	err = os.WriteFile("srv/get_kubeconfig_path", get_kubeconfig_path_command_b, 0755)

	if err != nil {

		return fmt.Errorf("failed init npia go client: %s", err.Error())
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

			nkadmdebug.ASgi_CliRef.PrintPrettyDefinition()

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

	err := InitGoClient()

	if err != nil {

		fmt.Println(err.Error())
		return
	}

	if len(os.Args) <= 1 {
		fmt.Println("error: no args specified")
		return
	}

	MODE_TEST := 0

	MODE_DEBUG := 0

	MODE_INTERACTIVE := 0

	// MODE_ADMIN := 0

	for i := 0; i < len(os.Args); i++ {

		flag := os.Args[i]

		if flag == "-t" || flag == "--test" {

			MODE_TEST = 1

		} else if flag == "-d" || flag == "--debug" {

			MODE_DEBUG = 1

		} else if flag == "-i" || flag == "--interactive" {

			MODE_INTERACTIVE = 1
		}

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
