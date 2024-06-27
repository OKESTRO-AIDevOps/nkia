package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"os/user"
	"time"

	"fmt"

	"github.com/OKESTRO-AIDevOps/nkia/pkg/apistandard/apix"
	_ "github.com/gin-gonic/contrib/sessions"
	_ "github.com/gin-gonic/gin"

	nkletcmd "github.com/OKESTRO-AIDevOps/nkia/nokubelet/cmd"
	"github.com/OKESTRO-AIDevOps/nkia/nokubelet/config"
	sock "github.com/OKESTRO-AIDevOps/nkia/nokubelet/oagent"
	modules "github.com/OKESTRO-AIDevOps/nkia/pkg/challenge"
	"github.com/OKESTRO-AIDevOps/nkia/pkg/kubebase"
	goya "github.com/goccy/go-yaml"
)

var DEBUG = 0

func InitNpiaServer() error {

	challenge_records := make(modules.ChallengRecord)

	key_records := make(modules.KeyRecord)

	challenge_records["_INIT"] = map[string]string{
		"_INIT": "_INIT",
	}

	key_records["_INIT"] = "_INIT"

	cmd := exec.Command("mkdir", "-p", ".npia")

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

	err = os.WriteFile(".npia/challenge.json", challenge_records_b, 0644)

	if err != nil {

		return fmt.Errorf("failed init npia server: %s", err.Error())
	}

	err = os.WriteFile(".npia/key.json", key_records_b, 0644)

	if err != nil {

		return fmt.Errorf("failed init npia server: %s", err.Error())
	}

	get_kubeconfig_path_command_string :=
		`#!/bin/bash
[[ ! -z "$KUBECONFIG" ]] && echo "$KUBECONFIG" || echo "$HOME/.kube/config"`

	get_kubeconfig_path_command_b := []byte(get_kubeconfig_path_command_string)

	err = os.WriteFile(".npia/get_kubeconfig_path", get_kubeconfig_path_command_b, 0755)

	if err != nil {

		return fmt.Errorf("failed init npia server: %s", err.Error())
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

				for k, v := range CONFIG_YAML {

					fmt.Printf("  %s: %s\n", k, v)

				}

				return nil

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

func InitNpiaServerDefault() error {

	challenge_records := make(modules.ChallengRecord)

	key_records := make(modules.KeyRecord)

	challenge_records["_INIT"] = map[string]string{
		"_INIT": "_INIT",
	}

	key_records["_INIT"] = "_INIT"

	cmd := exec.Command("mkdir", "-p", ".npia")

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

	err = os.WriteFile(".npia/challenge.json", challenge_records_b, 0644)

	if err != nil {

		return fmt.Errorf("failed init npia server: %s", err.Error())
	}

	err = os.WriteFile(".npia/key.json", key_records_b, 0644)

	if err != nil {

		return fmt.Errorf("failed init npia server: %s", err.Error())
	}

	get_kubeconfig_path_command_string :=
		`#!/bin/bash
[[ ! -z "$KUBECONFIG" ]] && echo "$KUBECONFIG" || echo "$HOME/.kube/config"`

	get_kubeconfig_path_command_b := []byte(get_kubeconfig_path_command_string)

	err = os.WriteFile(".npia/get_kubeconfig_path", get_kubeconfig_path_command_b, 0755)

	if err != nil {

		return fmt.Errorf("failed init npia server: %s", err.Error())
	}

	cmd = exec.Command("mkdir", "-p", ".npia")

	err = cmd.Run()

	if err != nil {

		return fmt.Errorf("failed to init: %s", err.Error())
	}

	CONFIG_YAML := make(map[string]string)

	if _, err := os.Stat(".npia/config.yaml"); err != nil {

		return fmt.Errorf("no default config yaml exists")

	} else {

		file_b, err := os.ReadFile(".npia/config.yaml")

		if err == nil {

			err = goya.Unmarshal(file_b, &CONFIG_YAML)

			if err == nil {

				fmt.Println("existing configuration: ")

				for k, v := range CONFIG_YAML {

					fmt.Printf("  %s: %s\n", k, v)

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

func RunLetCmd(args []string) {

	var address string

	var email string

	address = config.ADDRESS

	email = config.EMAIL

	apistd_in, err := apix.AXgi.BuildAPIInputFromCommandLine(args)

	if err != nil {
		fmt.Printf("failed: %s\n", err.Error())

		return
	}

	err = nkletcmd.RequestHandler(apistd_in)

	if err != nil {
		fmt.Printf("failed: %s\n", err.Error())

		return
	}

	cluster_id := apistd_in["clusterid"]

	update_token, ut_okay := apistd_in["updatetoken"]

	_ = os.WriteFile(".npia/cluster_id", []byte(cluster_id), 0644)

	if config.MODE == "test" && config.DEBUG == "true" && ut_okay {

		if err := sock.DetachedServerCommunicatorWithUpdate_Test_Debug(address, email, cluster_id, update_token); err != nil {
			fmt.Println(err.Error())
			return
		}

	} else if config.MODE == "test" && config.DEBUG != "true" && ut_okay {

		if err := sock.DetachedServerCommunicatorWithUpdate_Test(address, email, cluster_id, update_token); err != nil {
			fmt.Println(err.Error())
			return
		}

	} else if config.MODE == "test" && config.DEBUG == "true" && !ut_okay {

		if err := sock.DetachedServerCommunicator_Test_Debug(address, email, cluster_id); err != nil {
			fmt.Println(err.Error())
			return
		}

	} else if config.MODE == "test" && config.DEBUG != "true" && !ut_okay {

		if err := sock.DetachedServerCommunicator_Test(address, email, cluster_id); err != nil {
			fmt.Println(err.Error())
			return
		}

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

	args := os.Args[1:]

	if _, err := os.Stat(".npia/.init"); err != nil {

		err_init := InitNpiaServerDefault()

		if err_init != nil {
			fmt.Println(err_init.Error())
		}

		go kubebase.AdminInitNPIA()

		t_start := time.Now()

		done := 0

		for time.Now().Sub(t_start).Seconds() < 30 {

			if _, err := os.Stat(".npia/.init"); err == nil {

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

			return
		}

	}

	RunLetCmd(args)

}
