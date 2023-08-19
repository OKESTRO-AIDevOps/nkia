package runtimefs

import (
	"fmt"
	"os"
	"strings"

	goya "github.com/goccy/go-yaml"
)

func OpenFilePointerForUsrBuildLog() (*os.File, error) {

	var outfile *os.File
	var err error

	if _, err := os.Stat(".usr/build_log"); err == nil {
		return outfile, fmt.Errorf("failed to get file pointer: %s", "another build in process")
	}

	if _, err := os.Stat(".usr/build_done"); err == nil {
		e := os.Remove(".usr/build_done")
		if e != nil {
			return outfile, fmt.Errorf("failed to get file pointer: %s", err.Error())
		}
	}

	outfile, err = os.Create(".usr/build_log")

	if err != nil {
		return outfile, fmt.Errorf("failed to get file pointer: %s", err.Error())
	}

	return outfile, nil
}

func CloseFilePointerForUsrBuildLogAndMarkDone(fp *os.File) error {

	err := fp.Close()

	if err != nil {
		return fmt.Errorf("failed to close file pointer: %s", err.Error())
	}

	file_byte, err := os.ReadFile(".usr/build_log")

	if err != nil {
		return fmt.Errorf("failed to close file pointer: %s", err.Error())
	}

	err = os.Remove(".usr/build_log")

	if err != nil {
		return fmt.Errorf("failed to close file pointer: %s", err.Error())
	}

	err = os.WriteFile(".usr/build_done", file_byte, 0644)

	if err != nil {
		return fmt.Errorf("failed to close file pointer: %s", err.Error())
	}

	return nil
}

func GetUsrTargetDockerComposeYamlBuild() (string, error) {

	file_byte, err := os.ReadFile(".usr/target/docker-compose.yaml")

	if err != nil {
		return "", fmt.Errorf("failed to get dc yaml build: %s", err.Error())
	}

	yaml_if := make(map[interface{}]interface{})

	err = goya.Unmarshal(file_byte, &yaml_if)

	if err != nil {
		return "", fmt.Errorf("failed to get dc yaml build: %s", err.Error())
	}

	service_map := yaml_if["services"].(map[string]interface{})

	keys := make([]string, 0)
	for k := range service_map {

		keys = append(keys, k)
	}

	for i := 0; i < len(keys); i++ {

		delete(yaml_if["services"].(map[string]interface{})[keys[i]].(map[string]interface{}), "ports")
		delete(yaml_if["services"].(map[string]interface{})[keys[i]].(map[string]interface{}), "volumes")

	}

	out_b, err := goya.Marshal(yaml_if)

	if err != nil {
		return "", fmt.Errorf("failed to get dc yaml build: %s", err.Error())
	}

	err = os.WriteFile(".usr/target/docker-compose.yaml.build", out_b, 0644)

	if err != nil {
		return "", fmt.Errorf("failed to get dc yaml build: %s", err.Error())
	}

	return ".usr/target/docker-compose.yaml.build", nil

}

func GetUsrTargetPushList(regaddr string) ([][]string, error) {

	var target_push_list [][]string

	reg_prefix := "/target_"

	img_prefix := "target_"

	file_byte, err := os.ReadFile(".usr/target/docker-compose.yaml.build")

	if err != nil {
		return target_push_list, fmt.Errorf("failed to get push list: %s", err.Error())
	}

	yaml_if := make(map[interface{}]interface{})

	err = goya.Unmarshal(file_byte, &yaml_if)

	if err != nil {
		return target_push_list, fmt.Errorf("failed to get push list: %s", err.Error())
	}

	service_map := yaml_if["services"].(map[string]interface{})

	keys := make([]string, 0)
	for k := range service_map {

		keys = append(keys, k)
	}

	for i := 0; i < len(keys); i++ {

		conversion_type := ""

		prop_map := yaml_if["services"].(map[string]interface{})[keys[i]].(map[string]interface{})
		for k := range prop_map {
			if k == "image" {
				conversion_type = k
				break
			} else if k == "build" {
				conversion_type = k
				break
			}
		}

		container_name := ""

		source := ""

		destination := ""

		if conversion_type == "image" {

			container_name = yaml_if["services"].(map[string]interface{})[keys[i]].(map[string]interface{})["container_name"].(string)

			container_name = strings.ReplaceAll(container_name, "_", "-")

			destination = regaddr + reg_prefix + container_name

			source = yaml_if["services"].(map[string]interface{})[keys[i]].(map[string]interface{})["image"].(string)

		} else if conversion_type == "build" {

			container_name = yaml_if["services"].(map[string]interface{})[keys[i]].(map[string]interface{})["container_name"].(string)

			container_name = strings.ReplaceAll(container_name, "_", "-")

			destination = regaddr + reg_prefix + container_name

			source = img_prefix + keys[i]

		} else {
			return target_push_list, fmt.Errorf("failed to get push list: %s", "build or image key not found")
		}

		target_push_list = append(target_push_list, []string{source, destination})

	}

	return target_push_list, nil
}

func GetUsrBuildLog() ([]byte, error) {

	var ret_byte []byte

	var err error

	if _, err = os.Stat(".usr/build_log"); err == nil {

		ret_byte, err = os.ReadFile(".usr/build_log")

	} else if _, err = os.Stat(".usr/build_done"); err == nil {

		ret_byte, err = os.ReadFile(".usr/build_done")

	} else {
		err = fmt.Errorf("failed to get build log: %s", "none found")
	}

	return ret_byte, err
}
