package omodules

import (
	"encoding/json"
	"os"
)

type ConfigJSON struct {
	DEBUG bool `json:"DEBUG"`
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
