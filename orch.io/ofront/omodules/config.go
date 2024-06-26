package omodules

import (
	"encoding/json"
	"os"
)

type ConfigJSON struct {
	DEBUG       bool   `json:"DEBUG"`
	DB_HOST     string `json:"DB_HOST"`
	DB_ID       string `json:"DB_ID"`
	DB_PW       string `json:"DB_PW"`
	DB_NAME     string `json:"DB_NAME"`
	DB_HOST_DEV string `json:"DB_HOST_DEV"`
}

type OauthJSON struct {
	Web struct {
		ClientID                string   `json:"client_id"`
		ProjectID               string   `json:"project_id"`
		AuthURI                 string   `json:"auth_uri"`
		TokenURI                string   `json:"token_uri"`
		AuthProviderX509CertURL string   `json:"auth_provider_x509_cert_url"`
		ClientSecret            string   `json:"client_secret"`
		RedirectUris            []string `json:"redirect_uris"`
	} `json:"web"`
}

func LoadConfig() {

	CONFIG_JSON = GetConfigJSON()

	OAUTH_JSON = GetOauthJSON()

}

func GetConfigJSON() ConfigJSON {

	var cj ConfigJSON

	file_byte, err := os.ReadFile("config.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file_byte, &cj)

	if err != nil {
		panic(err)
	}

	return cj

}

func GetOauthJSON() OauthJSON {

	var oj OauthJSON

	file_byte, err := os.ReadFile("oauth.json")

	if err != nil {

		panic(err)
	}

	err = json.Unmarshal(file_byte, &oj)

	if err != nil {

		panic(err)

	}

	return oj

}
