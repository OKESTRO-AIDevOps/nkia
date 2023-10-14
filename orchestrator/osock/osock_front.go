package main

import (
	"crypto/x509"
	b64 "encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"net/http"

	"time"

	ctrl "github.com/OKESTRO-AIDevOps/nkia/src/controller"
	"github.com/OKESTRO-AIDevOps/nkia/src/modules"
	_ "github.com/gorilla/websocket"
)

func FrontHandler(w http.ResponseWriter, r *http.Request) {

	EventLogger("Front access")

	UPGRADER.CheckOrigin = func(r *http.Request) bool { return true }

	c, err := UPGRADER.Upgrade(w, r, nil)
	if err != nil {
		EventLogger("upgrade:" + err.Error())
		return
	}

	c.SetReadDeadline(time.Time{})

	var req_server ctrl.APIMessageRequest
	var req_orchestrator ctrl.OrchestratorRequest

	auth_flag := 0

	defer c.Close()

	for auth_flag == 0 {

		err := c.ReadJSON(&req_orchestrator)
		if err != nil {
			EventLogger("auth:" + err.Error())
			return
		}

		request_key_b64 := req_orchestrator.RequestOption

		request_key_b, err := b64.StdEncoding.DecodeString(request_key_b64)

		if err != nil {
			EventLogger("auth:" + err.Error())
			return
		}

		request_key := string(request_key_b)

		EventLogger("sess key: " + request_key)

		email, err := CheckSessionAndGetEmailByRequestKey(request_key)

		if err != nil {
			EventLogger("auth:" + err.Error())
			return
		}

		FRONT_CONNECTION[email] = c

		FRONT_CONNECTION_FRONT[c] = email

		break
	}

	EventLogger("front accepted")

	for {

		req_orchestrator = ctrl.OrchestratorRequest{}

		req_server = ctrl.APIMessageRequest{}

		res_orchestrator := ctrl.OrchestratorResponse{}

		err := c.ReadJSON(&req_orchestrator)

		if err != nil {
			EventLogger("read front:" + err.Error())
			return
		}

		target := req_orchestrator.RequestTarget

		email, okay := FRONT_CONNECTION_FRONT[c]

		if !okay {
			EventLogger("read front: no connected front name")
			return
		}

		email_context := email + ":" + target

		req_option := req_orchestrator.RequestOption

		query_str := req_orchestrator.Query

		if req_option == "ADM" {

			ret, err := AdminRequest(email, query_str)

			if err != nil {
				res_orchestrator.ServerMessage = err.Error()

				c.WriteJSON(&res_orchestrator)

				return
			}

			res_orchestrator.ServerMessage = "SUCCESS"

			res_orchestrator.QueryResult = ret

			c.WriteJSON(&res_orchestrator)

			continue

		}

		server_c, okay := SERVER_CONNECTION[email_context]

		if !okay {
			EventLogger("read front: no connected server context")
			return
		}

		key_id, okay := SERVER_CONNECTION_KEY[server_c]

		if !okay {
			EventLogger("read front: no server context key")
			return
		}

		session_sym_key, err := modules.AccessAuth_Detached(key_id)

		if err != nil {
			EventLogger("read front: " + err.Error())
			return
		}

		query_b := []byte(query_str)

		query_enc, err := modules.EncryptWithSymmetricKey([]byte(session_sym_key), query_b)

		if err != nil {
			EventLogger("read front: " + err.Error())
			return
		}

		query_hex := hex.EncodeToString(query_enc)

		req_server.Query = query_hex

		err = server_c.WriteJSON(&req_server)

		if err != nil {
			EventLogger("write to server: " + err.Error())
			return
		}

	}

}

func AdminRequest(email string, query string) ([]byte, error) {

	var ret []byte

	switch query {

	case "keygen":

		privkey, pubkey, err := modules.GenerateKeyPair(4096)

		if err != nil {
			return ret, fmt.Errorf("admin req: %s", err.Error())
		}

		priv_pem := pem.EncodeToMemory(
			&pem.Block{
				Type:  "RSA PRIVATE KEY",
				Bytes: x509.MarshalPKCS1PrivateKey(privkey),
			},
		)

		pub_pem := pem.EncodeToMemory(
			&pem.Block{
				Type:  "RSA PUBLIC KEY",
				Bytes: x509.MarshalPKCS1PublicKey(pubkey),
			},
		)

		pub_pem_str := string(pub_pem)

		err = UpdatePubkeyByEmail(email, pub_pem_str)

		if err != nil {
			return ret, fmt.Errorf("admin req: %s", err.Error())
		}

		ret = priv_pem

	// case "signout" :

	default:

		return ret, fmt.Errorf("admin req: %s", "no matching operand")

	}

	return ret, nil
}
