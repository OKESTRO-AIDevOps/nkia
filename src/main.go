package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"os/user"

	"fmt"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/OKESTRO-AIDevOps/npia-server/src/modules"
	"github.com/OKESTRO-AIDevOps/npia-server/src/router"
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

	if _, err := os.Stat("srv"); err != nil {

		if err_init := InitNpiaServer(); err_init != nil {

			fmt.Println(err_init.Error())

			return

		}

	}

	gin_srv := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	gin_srv.Use(sessions.Sessions("npia-session", store))

	gin_srv = router.Init(gin_srv)

	gin_srv.Run("0.0.0.0:13337")
}
