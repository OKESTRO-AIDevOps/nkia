package modules

import (
	"os"

	goya "github.com/goccy/go-yaml"
)

type ChallengRecord map[string]map[string]string

type KeyRecord map[string]string

func _LoadConfigYaml() map[string]string {

	var config_yaml map[string]string

	file_byte, err := os.ReadFile(".npia/config.yaml")

	if err != nil {
		panic(err.Error())
	}

	err = goya.Unmarshal(file_byte, &config_yaml)

	if err != nil {
		panic(err.Error())
	}

	return config_yaml
}

func _ConstructURL(url_path string) string {

	if CONFIG_YAML["MODE"] == "test" {

		return CONFIG_YAML["BASE_URL"] + url_path + "/test"

	} else if CONFIG_YAML["MODE"] == "release" {

		return CONFIG_YAML["BASE_URL"] + url_path

	} else {
		panic("mode option unavailable: " + CONFIG_YAML["MODE"])
	}

}

var CONFIG_YAML = _LoadConfigYaml()

var ADDRESS = _ConstructURL("/osock/server")

var EMAIL = CONFIG_YAML["EMAIL"]
