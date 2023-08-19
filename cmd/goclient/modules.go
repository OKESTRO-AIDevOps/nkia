package goclient

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/OKESTRO-AIDevOps/nkia/pkg/apistandard"
	"github.com/OKESTRO-AIDevOps/nkia/src/controller"
	"github.com/OKESTRO-AIDevOps/nkia/src/modules"
	goya "github.com/goccy/go-yaml"
)

func ClientAuthChallenge(client *http.Client) error {

	var session_sym_key string

	var kube_config map[interface{}]interface{}

	var req_body controller.AuthChallenge
	var resp_body controller.AuthChallenge

	req_challenge_records := make(modules.ChallengRecord)

	client_context_map := make(map[string]string)

	kube_config_path, err := modules.GetKubeConfigPath()

	if err != nil {
		return fmt.Errorf("chal: %s", err.Error())
	}

	kube_config_file_byte, err := os.ReadFile(kube_config_path)

	err = goya.Unmarshal(kube_config_file_byte, &kube_config)

	if err != nil {
		return fmt.Errorf("chal: %s", err.Error())
	}

	contexts_len := len(kube_config["contexts"].([]interface{}))

	for i := 0; i < contexts_len; i++ {

		context_nm := kube_config["contexts"].([]interface{})[i].(map[string]interface{})["name"].(string)

		pubkey_b, err := modules.GetContextClusterPublicKeyBytes(context_nm)

		if err != nil {
			return fmt.Errorf("chal: %s", err.Error())
		}

		client_context_map[context_nm] = string(pubkey_b)

	}

	req_challenge_records["ask_challenge"] = client_context_map

	req_body.ChallengeID = "ASK"

	req_body.ChallengeData = req_challenge_records

	req_b, err := json.Marshal(req_body)

	if err != nil {
		return fmt.Errorf("chal: %s", err.Error())
	}

	resp, err := client.Post(COMM_URL_AUTH, "application/json", bytes.NewBuffer(req_b))

	if err != nil {
		return fmt.Errorf("chal: %s", err.Error())
	}

	defer resp.Body.Close()

	body_bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("chal: %s", err.Error())
	}

	err = json.Unmarshal(body_bytes, &resp_body)

	if err != nil {
		return fmt.Errorf("chal: %s", err.Error())
	}
	challenge_map := resp_body.ChallengeData

	for chal_id, content := range challenge_map {

		for context, enc := range content {

			priv_key_b, err := modules.GetContextUserPrivateKeyBytes(context)

			if err != nil {
				return fmt.Errorf("chal: %s", err.Error())
			}

			priv_key, err := modules.BytesToPrivateKey(priv_key_b)

			if err != nil {
				return fmt.Errorf("chal: %s", err.Error())
			}

			enc_b, err := hex.DecodeString(enc)

			if err != nil {
				return fmt.Errorf("chal: %s", err.Error())
			}

			dec_b, err := modules.DecryptWithPrivateKey(enc_b, priv_key)

			if err != nil {
				return fmt.Errorf("chal: %s", err.Error())
			}

			dec := string(dec_b)

			challenge_map[chal_id][context] = dec

		}

	}

	req_body.ChallengeID = "ANS"

	req_body.ChallengeData = challenge_map

	req_b, err = json.Marshal(req_body)

	if err != nil {
		return fmt.Errorf("chal: %s", err.Error())
	}

	resp, err = client.Post(COMM_URL_AUTH, "application/json", bytes.NewBuffer(req_b))

	if err != nil {
		return fmt.Errorf("chal: %s", err.Error())
	}

	defer resp.Body.Close()

	body_bytes, err = io.ReadAll(resp.Body)

	if err != nil {
		return fmt.Errorf("chal: %s", err.Error())
	}

	err = json.Unmarshal(body_bytes, &resp_body)

	if err != nil {
		return fmt.Errorf("chal: %s", err.Error())
	}

	key_records := resp_body.ChallengeKey

	for context, enc := range key_records {

		priv_b, err := modules.GetContextUserPrivateKeyBytes(context)

		if err != nil {
			return fmt.Errorf("chal: %s", err.Error())
		}

		priv_key, err := modules.BytesToPrivateKey(priv_b)

		if err != nil {
			return fmt.Errorf("chal: %s", err.Error())
		}

		enc_b, err := hex.DecodeString(enc)

		if err != nil {
			return fmt.Errorf("chal: %s", err.Error())
		}

		dec_b, err := modules.DecryptWithPrivateKey(enc_b, priv_key)

		if err != nil {
			return fmt.Errorf("chal: %s", err.Error())
		}

		dec := string(dec_b)

		session_sym_key = dec

	}

	SESSION_SYM_KEY = session_sym_key

	return nil
}

func CommunicationHandler_LinearInstruction_PrintOnly(client *http.Client, linear_insruction string) {

	var req_body controller.APIMessageRequest
	var resp_body controller.APIMessageResponse

	query_plain := linear_insruction

	query_enc, err := modules.EncryptWithSymmetricKey([]byte(SESSION_SYM_KEY), []byte(query_plain))

	if err != nil {

		fmt.Println(err)
		return
	}

	req_body.Query = hex.EncodeToString(query_enc)

	req_b, err := json.Marshal(req_body)

	if err != nil {

		fmt.Println(err)
		return
	}

	resp, err := client.Post(COMM_URL, "application/json", bytes.NewBuffer(req_b))

	if err != nil {

		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body_bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(string(body_bytes))
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(body_bytes, &resp_body)

	if err != nil {

		fmt.Println(err)
		return
	}

	if resp_body.ServerMessage != "SUCCESS" {

		fmt.Println(resp_body.ServerMessage)
		return
	}

	result_enc := resp_body.QueryResult

	result_enc_b, err := hex.DecodeString(result_enc)

	if err != nil {

		fmt.Println(err)
		return
	}

	result_b, err := modules.DecryptWithSymmetricKey([]byte(SESSION_SYM_KEY), result_enc_b)

	if err != nil {

		fmt.Println(err)
		return
	}

	var api_out apistandard.API_OUTPUT

	err = json.Unmarshal(result_b, &api_out)

	if err != nil {

		fmt.Println(err)
		return
	}

	// for now

	fmt.Println(api_out)

}
