package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"os/user"

	"fmt"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/OKESTRO-AIDevOps/nkia/src/modules"
	"github.com/OKESTRO-AIDevOps/nkia/src/router"
	"github.com/OKESTRO-AIDevOps/nkia/src/sock"
)

func InitNpiaServer() error {

	challenge_records := make(modules.ChallengRecord)

	key_records := make(modules.KeyRecord)

	challenge_records["_INIT"] = map[string]string{
		"_INIT": "_INIT",
	}

	key_records["_INIT"] = "_INIT"

	cmd := exec.Command("mkdir", "-p", "srv")

	err := cmd.Run()

	if err != nil {

		return fmt.Errorf("failed init npia server: %s", err.Error())
	}

	challenge_records_b, err := json.Marshal(challenge_records)

	if err != nil {

		return fmt.Errorf("failed init npia server: %s", err.Error())
	}

	key_records_b, err := json.Marshal(key_records)

	if err != nil {

		return fmt.Errorf("failed init npia server: %s", err.Error())
	}

	err = os.WriteFile("srv/challenge.json", challenge_records_b, 0644)

	if err != nil {

		return fmt.Errorf("failed init npia server: %s", err.Error())
	}

	err = os.WriteFile("srv/key.json", key_records_b, 0644)

	if err != nil {

		return fmt.Errorf("failed init npia server: %s", err.Error())
	}

	get_kubeconfig_path_command_string :=
		`#!/bin/bash
[[ ! -z "$KUBECONFIG" ]] && echo "$KUBECONFIG" || echo "$HOME/.kube/config"`

	get_kubeconfig_path_command_b := []byte(get_kubeconfig_path_command_string)

	err = os.WriteFile("srv/get_kubeconfig_path", get_kubeconfig_path_command_b, 0755)

	if err != nil {

		return fmt.Errorf("failed init npia server: %s", err.Error())
	}

	return nil
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

	if len(os.Args) < 2 {
		fmt.Println("Error: wrong arguments")
		return
	}

	if _, err := os.Stat("srv"); err != nil {

		if err_init := InitNpiaServer(); err_init != nil {

			fmt.Println(err_init.Error())

			return

		}

	}

	option := os.Args[1]

	if option == "attached" {
		gin_srv := gin.Default()
		store := sessions.NewCookieStore([]byte("secret"))
		gin_srv.Use(sessions.Sessions("npia-session", store))

		gin_srv = router.Init(gin_srv)

		gin_srv.Run("0.0.0.0:13337")

	} else if option == "detached" {

		if len(os.Args) != 4 {
			fmt.Println("Error: wrong arguments")
			return
		}

		address := os.Args[2]

		email := os.Args[3]

		if err := sock.DetachedServerCommunicator(address, email); err != nil {
			fmt.Println(err.Error())
			return
		}

	} else if option == "detached-test" {

		if len(os.Args) != 4 {
			fmt.Println("Error: wrong arguments")
			return
		}

		address := os.Args[2]

		email := os.Args[3]

		if err := sock.DetachedServerCommunicator_Test(address, email); err != nil {
			fmt.Println(err.Error())
			return
		}

	} else {
		fmt.Println("Error: wrong option")
		return
	}

}
