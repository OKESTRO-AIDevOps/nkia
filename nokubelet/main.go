package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"syscall"

	"fmt"

	_ "github.com/gin-gonic/contrib/sessions"
	_ "github.com/gin-gonic/gin"

	"github.com/OKESTRO-AIDevOps/nkia/nokubelet/modules"
	sock "github.com/OKESTRO-AIDevOps/nkia/nokubelet/oagent"
	_ "github.com/OKESTRO-AIDevOps/nkia/nokubelet/router"

	"golang.org/x/term"
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

	MODE_DEBUG := 0

	MODE_TEST := 0

	MODE_UPDATE := 0

	_ = MODE_DEBUG

	_ = MODE_TEST

	current_user, err := user.Current()

	if err != nil {
		fmt.Println(err.Error())

		return

	}

	if current_user.Username != "root" {

		fmt.Println("Error: not running as root")
		return

	}

	if len(os.Args) <= 1 {
		fmt.Println("Error: arg not specified")
		return
	}

	for i := 1; i < len(os.Args); i++ {

		flag := os.Args[i]

		if flag == "-g" || flag == "--debug" {

			MODE_DEBUG = 1

		} else if flag == "-t" || flag == "--test" {

			MODE_TEST = 1

		} else if flag == "-u" || flag == "--update" {

			MODE_UPDATE = 1

		}

	}

	if _, err := os.Stat("srv"); err != nil {

		if err_init := InitNpiaServer(); err_init != nil {

			fmt.Println(err_init.Error())

			return

		}

	}

	/* MODE_DEBUG
	gin_srv := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	gin_srv.Use(sessions.Sessions("npia-session", store))

	gin_srv = router.Init(gin_srv)

	gin_srv.Run("0.0.0.0:13337")
	*/

	/* MODE_TEST

	if err := sock.DetachedServerCommunicator_Test(address, email); err != nil {
		fmt.Println(err.Error())
		return
	}

	*/

	var address string

	var email string

	var cluster_id string

	var update_token string

	fmt.Println("orch.io address: ")

	fmt.Scanln(&address)

	fmt.Println("orch.io user email: ")

	fmt.Scanln(&email)

	fmt.Println("orch.io cluster id: ")

	fmt.Scanln(&cluster_id)

	if MODE_UPDATE == 1 {

		fmt.Println("orch.io update token: ")

		byte_passwd, err := term.ReadPassword(int(syscall.Stdin))

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		token_str := string(byte_passwd)

		update_token = strings.TrimSpace(token_str)

		if err := sock.DetachedServerCommunicatorWithUpdate(address, email, cluster_id, update_token); err != nil {
			fmt.Println(err.Error())
			return
		}

	} else {

		if err := sock.DetachedServerCommunicator(address, email, cluster_id); err != nil {
			fmt.Println(err.Error())
			return
		}

	}

}
