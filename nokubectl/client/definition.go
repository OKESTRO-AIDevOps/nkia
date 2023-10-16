package client

import (
	"os"

	"github.com/OKESTRO-AIDevOps/nkia/pkg/apistandard"
	goya "github.com/goccy/go-yaml"
)

func _LoadConfigYaml() map[string]string {

	var config_yaml map[string]string

	file_byte, err := os.ReadFile("config.yaml")

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

var COMM_URL = _ConstructURL("/osock/front")

var COMM_URL_AUTH = _ConstructURL("/keyauth/login")

var COMM_URL_AUTH_CALLBACK = _ConstructURL("/keyauth/callback")

var ASgi_CliRef = apistandard.ASgi
