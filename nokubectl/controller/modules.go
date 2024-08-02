package controller

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	apistd "github.com/OKESTRO-AIDevOps/nkia/pkg/apistandard"
	apix "github.com/OKESTRO-AIDevOps/nkia/pkg/apistandard/apix"
	goya "github.com/goccy/go-yaml"
)

type HelpJSON struct {
	Command string `json:"command"`
	Comment string `json:"comment"`
	Flag    []struct {
		Name    string `json:"name"`
		Comment string `json:"comment"`
	} `json:"flag"`
}

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

func HelpCtl(out_format string) (string, error) {

	var retstr string

	var retarr []HelpJSON

	for k, v := range apix.AXgi {

		command := strings.ReplaceAll(k, "-", " ")

		command_comment := apix.AXcmd[k]

		hjson := HelpJSON{
			Command: command,
			Comment: command_comment,
			Flag: []struct {
				Name    string "json:\"name\""
				Comment string "json:\"comment\""
			}{},
		}

		flag_list := apistd.ASgi[v]

		flag_len := len(flag_list)

		for i := 0; i < flag_len; i++ {

			flag := flag_list[i]

			if flag == "id" {
				continue
			}

			flag_dash := "--" + flag

			flag_comment := apix.AXflag[flag]

			hjson.Flag = append(hjson.Flag, struct {
				Name    string "json:\"name\""
				Comment string "json:\"comment\""
			}{
				Name:    flag_dash,
				Comment: flag_comment,
			})

		}

		retarr = append(retarr, hjson)
	}

	if out_format == "pretty" {

		retarr_len := len(retarr)

		for i := 0; i < retarr_len; i++ {

			idx_str := fmt.Sprintf("%d", i+1)

			retstr += idx_str + ". " + retarr[i].Command + "\n"
			retstr += "  comment: " + retarr[i].Comment + "\n"
			retstr += "  flags  : " + "\n"

			flag_len := len(retarr[i].Flag)

			if flag_len == 0 {

				retstr += "    none\n"

			}

			for j := 0; j < flag_len; j++ {

				retstr += "    " + retarr[i].Flag[j].Name + "\n"
				retstr += "        comment: " + retarr[i].Flag[j].Comment + "\n"

			}

		}

	} else {

		jb, err := json.Marshal(retarr)

		if err != nil {

			return "", fmt.Errorf("failed to get help: %s", err.Error())

		}

		retstr = string(jb)
	}

	return retstr, nil
}
