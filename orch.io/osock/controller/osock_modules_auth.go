package controller

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"

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
