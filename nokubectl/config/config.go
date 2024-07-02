package config

import (
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"os"

	"github.com/OKESTRO-AIDevOps/nkia/pkg/apistandard"
	modules "github.com/OKESTRO-AIDevOps/nkia/pkg/challenge"
	goya "github.com/goccy/go-yaml"
)

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

		return CONFIG_YAML["BASE_URL_SOCK"] + url_path + "/test"

	} else if CONFIG_YAML["MODE"] == "release" {

		return CONFIG_YAML["BASE_URL_SOCK"] + url_path

	} else {

		panic("mode option unavailable: " + CONFIG_YAML["MODE"])

	}

}

func _ConstructURL_NoCalc(url_path string) string {

	return CONFIG_YAML["BASE_URL"] + url_path

}

func LoadCertAuthCredential() (*x509.CertPool, error) {

	certpool := x509.NewCertPool()

	file_b, err := os.ReadFile(".npia/certs/ca.crt")

	if err != nil {

		fmt.Printf("failed: %s\n", err.Error())

		return nil, fmt.Errorf("failed: %s\n", err.Error())
	}

	okay := certpool.AppendCertsFromPEM(file_b)

	if !okay {

		return nil, fmt.Errorf("failed to parse cert: ca.crt")
	}

	return certpool, nil

}

func LoadKeyAuthCredential() (*rsa.PrivateKey, error) {

	var ret_key *rsa.PrivateKey

	key_b, err := os.ReadFile(".npia/.priv")

	if err != nil {
		return ret_key, fmt.Errorf("failed to load cred: %s", err.Error())
	}

	ret_key, err = modules.BytesToPrivateKey(key_b)

	if err != nil {
		return ret_key, fmt.Errorf("failed to load cred: %s", err.Error())
	}

	return ret_key, nil
}

var CONFIG_YAML = _LoadConfigYaml()

var COMM_URL = _ConstructURL("/osock/front")

var COMM_URL_AUTH = _ConstructURL_NoCalc("/keyauth/login")

var COMM_URL_AUTH_CALLBACK = _ConstructURL_NoCalc("/keyauth/callback")

var EMAIL = CONFIG_YAML["EMAIL"]

var ASgi_CliRef = apistandard.ASgi
