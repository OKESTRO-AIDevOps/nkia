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

func LoadConfig() {

	CONFIG_JSON = GetConfigJSON()

	OAUTH_JSON = GetOauthJSON()

	GoogleOauthConfig = GenerateGoogleOauthConfig()

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
