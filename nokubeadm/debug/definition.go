package goclient

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

var SESSION_SYM_KEY = ""

var COMM_URL = _ConstructURL("/api/v0alpha")

var COMM_URL_AUTH = _ConstructURL("/auth-challenge")

var COMM_URL_MULTIMODE = _ConstructURL("/multimode/v0alpha")

var ASgi_CliRef = apistandard.ASgi

// challenge_id : ASK, ANS
// response     : NOPE
