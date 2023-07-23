package sock

import (
	"fmt"
	"os"

	ctrl "github.com/OKESTRO-AIDevOps/npia-server/src/controller"
	"github.com/OKESTRO-AIDevOps/npia-server/src/modules"
	"github.com/gorilla/websocket"

	goya "github.com/goccy/go-yaml"
)

var READ = 0

func ServerAuth_ReaderChannel(c *websocket.Conn, ch_read chan ctrl.AuthChallenge) {

	var resp_body ctrl.AuthChallenge

	for READ == 0 {

		err := c.ReadJSON(&resp_body)
		if err != nil {
			panic("auth reader:" + err.Error())
		}

		ch_read <- resp_body

	}

}

func ServerAuthChallenge(c *websocket.Conn) error {

	ch_read := make(chan ctrl.AuthChallenge)

	ServerAuth_ReaderChannel(c, ch_read)

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

	req_body.ChallengeData = req_challenge_records

	err = c.WriteJSON(&req_body)

	if err != nil {
		return fmt.Errorf("chal: %s", err.Error())
	}

	return nil
}
