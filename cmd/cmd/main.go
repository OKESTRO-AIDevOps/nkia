package main

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"os"
	"os/exec"

	_ "github.com/OKESTRO-AIDevOps/nkia/cmd/goadmin"
	"github.com/OKESTRO-AIDevOps/nkia/cmd/goclient"
	_ "github.com/OKESTRO-AIDevOps/nkia/cmd/goclient"
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

func RunClientInteractive() {

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

	err = goclient.ClientAuthChallenge(client)

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

			goclient.ASgi_CliRef.PrintPrettyDefinition()

		case "exit":

			return

		default:

			goclient.CommunicationHandler_LinearInstruction_PrintOnly(client, in_raw_query)

		}

	}

}

func RunAdminInteractive() {

}

func RunAdminFork() {

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

	mode := os.Args[1]

	if mode == "test" {
		goclient.BaseFlow_APIThenMultiMode_Test()
	} else if mode == "client" {
		RunClientInteractive()

	} else if mode == "admin" {

		RunAdminInteractive()

	} else if mode == "admin-fork" {

		RunAdminFork()

	} else {
		fmt.Println("wrong mode")
	}

}
