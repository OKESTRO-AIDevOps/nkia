package main

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"

	ci "github.com/OKESTRO-AIDevOps/nkia/test/ci"
	"gopkg.in/yaml.v3"

	"golang.org/x/term"
)

func GetRepo_Test() error {

	var repo_addr string

	var repo_id string

	var repo_pw string

	fmt.Printf("type repo addr: ")
	fmt.Scanln(&repo_addr)

	fmt.Printf("\ntype repo id: ")
	fmt.Scanln(&repo_id)
	fmt.Printf("\ntype repo pw: ")

	pw_b, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return fmt.Errorf("get repo test error: %s", err.Error())
	}

	repo_pw = string(pw_b)

	fmt.Printf("\n")

	if strings.HasPrefix(repo_addr, "https://") {

		repo_addr = strings.Replace(repo_addr, "https://", "", 1)

	}

	insert := "%s:%s@"

	repo_addr = insert + repo_addr

	repo_addr = fmt.Sprintf(repo_addr, repo_id, repo_pw)

	repo_addr = "https://" + repo_addr

	cmd := exec.Command("git", "clone", repo_addr)

	if err := cmd.Run(); err != nil {

		return fmt.Errorf("cmd run failed: %s", err.Error())

	}

	return nil

}

func ReadCIFile() {

	ci_targets, err := ci.LoadTargetsFromFile()

	if err != nil {
		fmt.Println(err.Error())
	}

	for k, v := range ci_targets {

		if k == "target.v1" {

			var target_arr []interface{}

			err := yaml.Unmarshal(v, &target_arr)

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			for i := 0; i < len(target_arr); i++ {

				var target_ci ci.TargetV1

				yaml_b, err := yaml.Marshal(target_arr[i])

				if err != nil {
					fmt.Println(err.Error())
					return
				}

				err = yaml.Unmarshal(yaml_b, &target_ci)

				if err != nil {
					fmt.Println(err.Error())
					return
				}

				fmt.Println(target_ci.GitPackage.Address)

			}

		}

	}

}

func main() {
	/*
		if err := GetRepo_Test(); err != nil {

			fmt.Println(err.Error())

		} else {

			fmt.Println("success!")

		}
	*/

	ReadCIFile()

}
