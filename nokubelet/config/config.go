package config

import (
	"os"

	goya "github.com/goccy/go-yaml"
)

func _CheckCfg() int {

	if _, err := os.Stat(".npia/config.yaml"); err != nil {
		return 0
	}
	return 1
}

func _LoadConfigYaml() map[string]string {

	var config_yaml map[string]string

	if CFG_EXIST == 1 {

		file_byte, err := os.ReadFile(".npia/config.yaml")

		if err != nil {
			panic(err.Error())
		}

		err = goya.Unmarshal(file_byte, &config_yaml)

		if err != nil {
			panic(err.Error())
		}

	}

	return config_yaml
}

func _ConstructURL(url_path string) string {

	if CFG_EXIST == 1 {

		if CONFIG_YAML["MODE"] == "test" {

			return CONFIG_YAML["BASE_URL"] + url_path + "/test"

		} else if CONFIG_YAML["MODE"] == "release" {

			return CONFIG_YAML["BASE_URL"] + url_path

		} else {
			panic("mode option unavailable: " + CONFIG_YAML["MODE"])
		}

	} else {
		return ""
	}

}

func _GetEmail() string {

	if CFG_EXIST == 1 {
		return CONFIG_YAML["EMAIL"]
	} else {
		return ""
	}

}

func _GetDebug() string {

	if CFG_EXIST == 1 {
		return CONFIG_YAML["DEBUG"]
	} else {
		return ""
	}

}

func _GetMode() string {

	if CFG_EXIST == 1 {
		return CONFIG_YAML["MODE"]
	} else {
		return ""
	}

}

var CFG_EXIST = _CheckCfg()

var CONFIG_YAML = _LoadConfigYaml()

var ADDRESS = _ConstructURL("/osock/server")

var EMAIL = _GetEmail()

var DEBUG = _GetDebug()

var MODE = _GetMode()
