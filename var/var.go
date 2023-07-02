package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	goya "github.com/goccy/go-yaml"
	//"github.com/OKESTRO-AIDevOps/npia-api/pkg/apistandard"
)

type ChallengRecord map[string]map[string]string

type KeyRecord map[string]string

func GetKubeConfigPathSimple() {

	cmd := exec.Command("../srv/get_kubeconfig_path")

	out, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return

	}

	fmt.Println(string(out))
}

func RandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func GetKubeConfigPath() (string, error) {

	var kube_config_path string

	cmd := exec.Command("../srv/get_kubeconfig_path")

	out, err := cmd.Output()

	if err != nil {

		return "", fmt.Errorf("failed to get kube config path: %s", err.Error())

	}

	strout := string(out)

	ret_strout := strings.ReplaceAll(strout, "\n", "")

	ret_strout = strings.ReplaceAll(ret_strout, " ", "")

	kube_config_path = ret_strout

	return kube_config_path, nil
}

func LoadTest() (ChallengRecord, error) {

	var kube_config map[interface{}]interface{}

	challenge_records := make(ChallengRecord)

	new_challenge_records := make(ChallengRecord)

	context_challenges := make(map[string]string)

	file_byte, err := os.ReadFile("../srv/challenge.json")

	if err != nil {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", err.Error())
	}

	err = json.Unmarshal(file_byte, &challenge_records)

	if err != nil {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", err.Error())
	}

	new_challenge_id, _ := RandomHex(32)

	_, okay := challenge_records[new_challenge_id]

	if okay {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", "duplicate challenge id")
	}

	kube_config_path, err := GetKubeConfigPath()

	if err != nil {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", err.Error())
	}

	kube_config_file_byte, err := os.ReadFile(kube_config_path)

	err = goya.Unmarshal(kube_config_file_byte, &kube_config)

	fmt.Println(kube_config)

	if err != nil {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", err.Error())
	}

	contexts_len := len(kube_config["contexts"].([]interface{}))

	for i := 0; i < contexts_len; i++ {

		context_nm := kube_config["contexts"].([]interface{})[i].(map[string]interface{})["name"].(string)

		new_challenge_val, _ := RandomHex(256)

		context_challenges[context_nm] = new_challenge_val

	}

	challenge_records[new_challenge_id] = context_challenges

	new_challenge_records[new_challenge_id] = context_challenges

	challenge_records_byte, err := json.Marshal(challenge_records)

	if err != nil {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", err.Error())
	}

	err = os.WriteFile("../srv/challenge.json", challenge_records_byte, 0644)

	if err != nil {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", err.Error())
	}

	return new_challenge_records, nil

}

func main() {

	//	ASgi := apistandard.ASgi

	//	ASgi.PrintPrettyDefinition()

	// GetKubeConfigPathSimple()

	if rec, err := LoadTest(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(rec)
	}

}
