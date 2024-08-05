package controller

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	models "github.com/OKESTRO-AIDevOps/nkia/orch.io/osock/models"
	ctrl "github.com/OKESTRO-AIDevOps/nkia/pkg/apistandard/apix"
	modules "github.com/OKESTRO-AIDevOps/nkia/pkg/challenge"
	"github.com/gorilla/websocket"
)

var CA_CERT *x509.Certificate

func CertAuthChallenge(c *websocket.Conn) (string, error) {

	var v_email string

	var req_orchestrator = ctrl.OrchestratorRequest{}
	var res_orchestrator = ctrl.OrchestratorResponse{}

	err := c.ReadJSON(&req_orchestrator)

	if err != nil {

		return "", fmt.Errorf("key auth: json: %s", err.Error())

	}

	certString := req_orchestrator.Query

	cert_b := []byte(certString)

	clientcrt, err := modules.BytesToCert(cert_b)

	if err != nil {

		return "", fmt.Errorf("key auth: cert: %s", err.Error())
	}

	hash_sha := sha256.New()

	hash_sha.Write(clientcrt.RawTBSCertificate)

	hash_data := hash_sha.Sum(nil)

	pub_key := CA_CERT.PublicKey.(*rsa.PublicKey)

	err = rsa.VerifyPKCS1v15(pub_key, crypto.SHA256, hash_data, clientcrt.Signature)

	if err != nil {

		return "", fmt.Errorf("key auth: verify: %s", err.Error())
	}

	v_email = clientcrt.Subject.CommonName

	res_orchestrator.ServerMessage = "SUCCESS"

	res_orchestrator.QueryResult = []byte(v_email)

	err = c.WriteJSON(res_orchestrator)

	if err != nil {

		return "", fmt.Errorf("key auth: send: %s", err.Error())

	}

	return v_email, nil
}

func RemoteAccessAuthChallenge(c *websocket.Conn) (string, error) {

	auth_flag := 0

	iter_count := 0

	this_cluster := ""

	var ANSWER string

	for auth_flag == 0 {

		var req = ctrl.AuthChallenge{}
		var resp = ctrl.AuthChallenge{}

		if iter_count > 10 {
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
			if err != nil {

				return "", fmt.Errorf("auth iter write close:" + err.Error())
			}

			return "", fmt.Errorf("auth iter: limit")
		}

		err := c.ReadJSON(&req)
		if err != nil {

			return "", fmt.Errorf("auth:" + err.Error())
		}

		chal_id := req.ChallengeID

		switch chal_id {

		case "UPDATE":

			ANSWER, _ = modules.RandomHex(128)

			email_context := req.ChallengeMessage

			email_context_list := strings.Split(email_context, ":")

			if len(email_context_list) != 2 {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				return "", fmt.Errorf("auth update: wrong format")
			}

			email := email_context_list[0]

			cluster_id := email_context_list[1]

			token, err := models.GetConfigChallengeByEmailAndClusterID2(email, cluster_id)

			if err != nil {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				return "", fmt.Errorf("auth update: get config: " + err.Error())
			}

			token_b := []byte(token)

			QUEST, err := modules.EncryptWithSymmetricKey(token_b, []byte(ANSWER))

			if err != nil {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				return "", fmt.Errorf("auth update: encrypt: " + err.Error())
			}

			quest_hex := hex.EncodeToString(QUEST)

			resp.ChallengeID = "UPDATE"

			resp.ChallengeMessage = quest_hex

			err = c.WriteJSON(resp)

			if err != nil {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				return "", fmt.Errorf("auth update: write: " + err.Error())
			}

		case "ROTATE":

			email_context := req.ChallengeMessage

			email_context_list := strings.Split(email_context, ":")

			if len(email_context_list) != 4 {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				return "", fmt.Errorf("auth rotate: wrong format")
			}

			email := email_context_list[0]

			cluster_id := email_context_list[1]

			answer := email_context_list[2]

			config := email_context_list[3]

			token, err := models.GetConfigChallengeByEmailAndClusterID2(email, cluster_id)

			if err != nil {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				return "", fmt.Errorf("auth rotate: get config: " + err.Error())
			}

			token_b := []byte(token)

			if ANSWER != answer {

				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				return "", fmt.Errorf("auth rotate: answer: " + err.Error())
			}

			config_hex, err := hex.DecodeString(config)

			if err != nil {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				return "", fmt.Errorf("auth rotate: decode config: " + err.Error())
			}

			config_dec, err := modules.DecryptWithSymmetricKey(token_b, config_hex)

			if err != nil {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				return "", fmt.Errorf("auth rotate: decrypt config: " + err.Error())
			}

			config_dec_string := string(config_dec)

			err = models.AddClusterByEmailAndClusterID2(email, cluster_id, config_dec_string)

			if err != nil {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				return "", fmt.Errorf("auth rotate: add cluster: " + err.Error())
			}

			resp.ChallengeID = "ROTATE"

			resp.ChallengeMessage = "SUCCESS"

			err = c.WriteJSON(resp)

			if err != nil {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				return "", fmt.Errorf("auth rotate: write: " + err.Error())
			}

		case "ASK":

			email_context := req.ChallengeMessage

			email_context_list := strings.Split(email_context, ":")

			if len(email_context_list) != 2 {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				return "", fmt.Errorf("auth ask: wrong format")
			}

			email := email_context_list[0]

			cluster_id := email_context_list[1]

			config_b, err := models.GetKubeconfigByEmailAndClusterID(email, cluster_id)

			if err != nil {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				return "", fmt.Errorf("auth ask: get config: " + err.Error())
			}

			client_ca_pub_key := req.ChallengeData

			chal_rec, err := modules.GenerateChallenge_Detached(config_b, client_ca_pub_key)

			if err != nil {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				return "", fmt.Errorf("auth ask: gen chal: " + err.Error())
			}

			resp.ChallengeID = "ASK"
			resp.ChallengeData = chal_rec

			err = c.WriteJSON(&resp)

			if err != nil {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				return "", fmt.Errorf("auth ask: write: " + err.Error())
			}

		case "ANS":

			email_context := req.ChallengeMessage

			email_context_list := strings.Split(email_context, ":")

			if len(email_context_list) != 2 {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				return "", fmt.Errorf("auth ask: wrong format")
			}

			email := email_context_list[0]

			cluster_id := email_context_list[1]

			config_b, err := models.GetKubeconfigByEmailAndClusterID(email, cluster_id)

			if err != nil {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				return "", fmt.Errorf("auth ans: get config: " + err.Error())
			}

			answer := req.ChallengeData

			gen_key, key_rec, err := modules.VerifyChallange_Detached(config_b, answer)

			if err != nil {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				return "", fmt.Errorf("auth ans: verify: " + err.Error())
			}

			server_c, okay := SERVER_CONNECTION[email_context]

			if okay {

				_ = server_c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				SERVER_CONNECTION[email_context] = c

			} else {
				SERVER_CONNECTION[email_context] = c
			}

			SERVER_CONNECTION_KEY[c] = gen_key

			SERVER_CONNECTION_FRONT[c] = email

			resp.ChallengeID = "ASK"
			resp.ChallengeKey = key_rec

			err = c.WriteJSON(&resp)

			if err != nil {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				return "", fmt.Errorf("auth ans: write: " + err.Error())
			}

			this_cluster = cluster_id

			auth_flag = 1

		default:
			_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

			return "", fmt.Errorf("auth blank: default")

		}

		iter_count += 1

	}

	return this_cluster, nil
}

func KeyAuthChallenge(c *websocket.Conn) (string, error) {

	var req_orchestrator = ctrl.OrchestratorRequest{}
	var res_orchestrator = ctrl.OrchestratorResponse{}

	err := c.ReadJSON(&req_orchestrator)

	if err != nil {

		return "", fmt.Errorf("key auth: json: %s", err.Error())

	}

	email := req_orchestrator.Query

	pubkey, err := models.GetPubkeyByEmail(email)

	if err != nil {

		return "", fmt.Errorf("key auth: pkey: %s", err.Error())

	}

	char_rec, err := modules.GenerateChallenge_Key(email, pubkey)

	if err != nil {

		return "", fmt.Errorf("key auth: gen chal: %s", err.Error())

	}

	char_rec_b, err := json.Marshal(char_rec)

	if err != nil {

		return "", fmt.Errorf("key auth: chal marshal: %s", err.Error())
	}

	res_orchestrator.ServerMessage = "SUCCESS"

	res_orchestrator.QueryResult = char_rec_b

	c.WriteJSON(res_orchestrator)

	req_orchestrator = ctrl.OrchestratorRequest{}

	res_orchestrator = ctrl.OrchestratorResponse{}

	var answer modules.ChallengRecord

	err = c.ReadJSON(&req_orchestrator)

	if err != nil {

		return "", fmt.Errorf("key auth: json2: %s", err.Error())

	}

	answer_json_b64 := req_orchestrator.Query

	answer_json_b, err := base64.StdEncoding.DecodeString(answer_json_b64)

	if err != nil {

		return "", fmt.Errorf("key auth: b64: %s", err.Error())

	}

	err = json.Unmarshal(answer_json_b, &answer)

	if err != nil {

		return "", fmt.Errorf("key auth: ans unmarshal: %s", err.Error())
	}

	v_email, err := modules.VerifyChallange_Key(answer)

	if err != nil {

		return "", fmt.Errorf("key auth: verify: %s", err.Error())
	}

	res_orchestrator.ServerMessage = "SUCCESS"

	res_orchestrator.QueryResult = []byte(v_email)

	err = c.WriteJSON(res_orchestrator)

	if err != nil {

		return "", fmt.Errorf("key auth: send: %s", err.Error())

	}

	return v_email, nil
}
