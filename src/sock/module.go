package sock

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"time"

	ctrl "github.com/OKESTRO-AIDevOps/nkia/src/controller"
	"github.com/OKESTRO-AIDevOps/nkia/src/modules"
	"github.com/OKESTRO-AIDevOps/npia-api/pkg/apistandard"
	"github.com/gorilla/websocket"

	goya "github.com/goccy/go-yaml"
)

var READ = 0

func ServerAuth_ReaderChannel(c *websocket.Conn, ch_read chan ctrl.AuthChallenge) {

	var resp_body ctrl.AuthChallenge

	err := c.ReadJSON(&resp_body)
	if err != nil {
		panic("auth reader1:" + err.Error())
	}

	ch_read <- resp_body

	resp_body = ctrl.AuthChallenge{}

	err = c.ReadJSON(&resp_body)
	if err != nil {
		panic("auth reader2:" + err.Error())
	}

	ch_read <- resp_body

}

func SockCommunicationHandler_ReaderChannel(c *websocket.Conn, ch_read chan ctrl.APIMessageRequest) {

	var req_body ctrl.APIMessageRequest

	for READ == 0 {

		err := c.ReadJSON(&req_body)

		if err != nil {
			panic("auth reader:" + err.Error())
		}

		ch_read <- req_body

		if READ != 0 {
			break
		}

	}

}

func ServerAuthChallenge(c *websocket.Conn, email string) error {

	ch_read := make(chan ctrl.AuthChallenge)

	go ServerAuth_ReaderChannel(c, ch_read)

	var session_sym_key string

	var kube_config map[interface{}]interface{}

	var req_body ctrl.AuthChallenge
	var resp_body ctrl.AuthChallenge

	req_challenge_records := make(modules.ChallengRecord)

	client_context_map := make(map[string]string)

	kube_config_path, err := modules.GetKubeConfigPath()

	if err != nil {
		return fmt.Errorf("chal: %s", err.Error())
	}

	kube_config_file_byte, err := os.ReadFile(kube_config_path)

	if err != nil {
		return fmt.Errorf("chal: %s", err.Error())
	}

	err = goya.Unmarshal(kube_config_file_byte, &kube_config)

	if err != nil {
		return fmt.Errorf("chal: %s", err.Error())
	}

	contexts_len := len(kube_config["contexts"].([]interface{}))

	if contexts_len != 1 {
		return fmt.Errorf("chal: %s", "too many contexts")
	}

	context_nm := kube_config["contexts"].([]interface{})[0].(map[string]interface{})["name"].(string)

	pubkey_b, err := modules.GetContextClusterPublicKeyBytes(context_nm)

	if err != nil {
		return fmt.Errorf("chal: %s", err.Error())
	}

	client_context_map[context_nm] = string(pubkey_b)

	req_challenge_records["ask_challenge"] = client_context_map

	req_body.ChallengeID = "ASK"

	req_body.ChallengeMessage = email + ":" + context_nm

	req_body.ChallengeData = req_challenge_records

	err = c.WriteJSON(&req_body)

	if err != nil {
		return fmt.Errorf("chal: %s", err.Error())
	}

	ask_loop := 0
	time_limit := 0
	ticker := time.NewTicker(time.Second)

	for ask_loop == 0 {
		select {
		case <-ticker.C:
			time_limit += 1
			if time_limit == 10 {
				READ = 1
				return fmt.Errorf("chal: %s", "ask time limit exceeded")
			}

		case read_body := <-ch_read:
			resp_body = read_body
			ticker.Stop()
			ask_loop = 1
			break
		}
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

	req_body.ChallengeMessage = email + ":" + context_nm

	req_body.ChallengeData = challenge_map

	err = c.WriteJSON(&req_body)

	if err != nil {
		return fmt.Errorf("chal: %s", err.Error())
	}

	ans_loop := 0
	time_limit = 0
	ticker = time.NewTicker(time.Second)

	for ans_loop == 0 {
		select {
		case <-ticker.C:
			time_limit += 1
			if time_limit == 10 {
				READ = 1
				return fmt.Errorf("chal: %s", "ans time limit exceeded")
			}

		case read_body := <-ch_read:
			resp_body = read_body
			ticker.Stop()
			ans_loop = 1
			break
		}
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

	READ = 1

	SESSION_SYM_KEY = session_sym_key

	return nil
}

func SockCommunicationHandler_LinearInstruction_PrintOnly(c *websocket.Conn) error {

	READ = 0

	ASgi := apistandard.ASgi

	var req_body ctrl.APIMessageRequest
	var resp_body ctrl.APIMessageResponse

	ch_read := make(chan ctrl.APIMessageRequest)

	interrupt := make(chan os.Signal, 1)

	go SockCommunicationHandler_ReaderChannel(c, ch_read)

	comm_loop := 0

	for comm_loop == 0 {

		select {

		case read_body := <-ch_read:

			req_body = read_body

			enc_query := req_body.Query

			enc_query_b, err := hex.DecodeString(enc_query)

			if err != nil {
				READ = 1
				resp_body.ServerMessage = err.Error()
				err = c.WriteJSON(&read_body)

				if err != nil {
					return fmt.Errorf("comm handler write: %s", err.Error())
				}

				return fmt.Errorf("comm handler: %s", err.Error())
			}

			key_b := []byte(SESSION_SYM_KEY)

			linear_instruction_b, err := modules.DecryptWithSymmetricKey(key_b, enc_query_b)

			if err != nil {
				READ = 1
				resp_body.ServerMessage = err.Error()
				err = c.WriteJSON(&resp_body)

				if err != nil {
					return fmt.Errorf("comm handler write: %s", err.Error())
				}

				return fmt.Errorf("comm handler: %s", err.Error())

			}

			linear_instruction := string(linear_instruction_b)

			api_input, err := ASgi.StdCmdInputBuildFromLinearInstruction(linear_instruction)

			if err != nil {
				READ = 1
				resp_body.ServerMessage = err.Error()
				err = c.WriteJSON(&resp_body)

				if err != nil {
					return fmt.Errorf("comm handler write: %s", err.Error())
				}

				return fmt.Errorf("comm handler: %s", err.Error())

			}

			api_out, err := ASgi.Run(api_input)

			if err != nil {
				READ = 1
				resp_body.ServerMessage = err.Error()
				err = c.WriteJSON(&resp_body)

				if err != nil {
					return fmt.Errorf("comm handler write: %s", err.Error())
				}

				return fmt.Errorf("comm handler: %s", err.Error())

			}

			api_out_b, err := json.Marshal(api_out)

			if err != nil {
				READ = 1
				resp_body.ServerMessage = err.Error()
				err = c.WriteJSON(&resp_body)

				if err != nil {
					return fmt.Errorf("comm handler write: %s", err.Error())
				}

				return fmt.Errorf("comm handler: %s", err.Error())

			}

			ret_byte, err := modules.EncryptWithSymmetricKey(key_b, api_out_b)

			if err != nil {
				READ = 1
				resp_body.ServerMessage = err.Error()
				err = c.WriteJSON(&resp_body)

				if err != nil {
					return fmt.Errorf("comm handler write: %s", err.Error())
				}

				return fmt.Errorf("comm handler: %s", err.Error())

			}

			ret_enc := hex.EncodeToString(ret_byte)

			resp_body.ServerMessage = "SUCCESS"
			resp_body.QueryResult = ret_enc

			err = c.WriteJSON(&resp_body)

			if err != nil {
				READ = 1
				return fmt.Errorf("comm handler: %s", err.Error())
			}

		case <-interrupt:
			READ = 1
			comm_loop = 1
			break
		}

	}

	return nil
}
