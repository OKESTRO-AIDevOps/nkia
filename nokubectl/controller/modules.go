package controller

import (
	"fmt"
	"os"
	"os/exec"

	goya "github.com/goccy/go-yaml"
)

func InitCtl() error {

	if _, err := os.Stat(".npia/.init"); err == nil {

		return fmt.Errorf("failed to init: %s", "already initiated")
	}

	cmd := exec.Command("mkdir", "-p", ".npia")

	err := cmd.Run()

	if err != nil {

		return fmt.Errorf("failed to init: %s", err.Error())
	}

	if len(os.Args) < 3 {
		return fmt.Errorf("failed to init: %s", "too few arguments")
	}

	cmd = exec.Command("mkdir", "-p", ".npia")

	err = cmd.Run()

	if err != nil {

		return fmt.Errorf("failed to init: %s", err.Error())
	}

	CONFIG_YAML := make(map[string]string)

	if _, err := os.Stat(".npia/config.yaml"); err == nil {

		fmt.Println("existing configyaml has been detected")

		file_b, err := os.ReadFile("./.npia/config.yaml")

		if err == nil {

			fmt.Println("successfully read the existing configyaml")

			err = goya.Unmarshal(file_b, &CONFIG_YAML)

			if err == nil {

				fmt.Println("existing configuration: ")

				for k, v := range CONFIG_YAML {

					fmt.Printf("  %s: %s\n", k, v)

				}

			} else {

				return fmt.Errorf("failed to init: %s", err.Error())

			}

		}

	}

	return nil

}

func ToCtl(to string) error {

	err := os.WriteFile(".npia/.to", []byte(to), 0644)

	if err != nil {

		return fmt.Errorf("failed to write .to: %s", err.Error())

	}

	return nil
}

func AsCtl(as string) error {

	err := os.WriteFile(".npia/.as", []byte(as), 0644)

	if err != nil {

		return fmt.Errorf("failed to write .as: %s", err.Error())

	}

	return nil

}
