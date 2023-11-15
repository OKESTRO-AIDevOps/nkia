package main

import (
	"crypto/x509"
	b64 "encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"strings"

	"time"

	ctrl "github.com/OKESTRO-AIDevOps/nkia/nokubelet/controller"
	"github.com/OKESTRO-AIDevOps/nkia/nokubelet/modules"
	"github.com/OKESTRO-AIDevOps/nkia/pkg/apistandard"

	"github.com/gorilla/websocket"
	_ "github.com/gorilla/websocket"
)

func FrontDestructor(c *websocket.Conn) {

	EventLogger("Front destructor called")

	fc, _ := FRONT_CONNECTION_FRONT[c]

	delete(FRONT_CONNECTION_FRONT, c)

	delete(FRONT_CONNECTION, fc)

	EventLogger("Front destructor exit")
}

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

			FrontDestructor(c)
			c.Close()
			EventLogger("read front:" + err.Error())
			return
		}

		target := req_orchestrator.RequestTarget

		email, okay := FRONT_CONNECTION_FRONT[c]

		if !okay {
			FrontDestructor(c)
			_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

			EventLogger("read front: no connected front name")

			return
		}

		email_context := email + ":" + target

		req_option := req_orchestrator.RequestOption

		query_str := req_orchestrator.Query

		if req_option == "admin" {

			ret, err := AdminRequest(email, query_str)

			if err != nil {

				EventLogger("read front: " + err.Error())

				FrontDestructor(c)

				res_orchestrator.ServerMessage = err.Error()

				c.WriteJSON(&res_orchestrator)

				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				return
			}

			res_orchestrator.ServerMessage = "SUCCESS"

			res_orchestrator.QueryResult = ret

			c.WriteJSON(&res_orchestrator)

			continue

		}

		server_c, okay := SERVER_CONNECTION[email_context]

		if !okay {
			FrontDestructor(c)
			_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
			EventLogger("read front: no connected server context")

			return
		}

		key_id, okay := SERVER_CONNECTION_KEY[server_c]

		if !okay {

			FrontDestructor(c)
			_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
			EventLogger("read front: no server context key")

			return
		}

		session_sym_key, err := modules.AccessAuth_Detached(key_id)

		if err != nil {

			FrontDestructor(c)
			_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
			EventLogger("read front: " + err.Error())
			return

		}

		query_b := []byte(query_str)

		query_enc, err := modules.EncryptWithSymmetricKey([]byte(session_sym_key), query_b)

		if err != nil {
			FrontDestructor(c)
			_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
			EventLogger("read front: " + err.Error())

			return
		}

		query_hex := hex.EncodeToString(query_enc)

		req_server.Query = query_hex

		err = server_c.WriteJSON(&req_server)

		if err != nil {
			FrontDestructor(c)
			c.Close()
			EventLogger("write to server: " + err.Error())

			return
		}

	}

}

func AdminRequest(email string, query string) ([]byte, error) {

	var ret []byte

	var ret_apiout apistandard.API_OUTPUT

	OP, args, err := AdminRequestParser_Linear(query)

	if err != nil {
		return ret, fmt.Errorf("admin req: %s", err.Error())
	}

	_ = args

	switch OP {

	case "ORCH-CONNCHK":

		var talkback string = "talking back list: "

		for i := 0; i < len(args); i++ {

			talkback += args[i] + " "

		}

		talkback += "\n"

		ret_apiout.BODY = talkback

		ret, err = json.Marshal(ret_apiout)

		if err != nil {
			return ret, fmt.Errorf("admin req: %s", err.Error())
		}

	case "ORCH-KEYGEN":

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

		pub_b, err := x509.MarshalPKIXPublicKey(pubkey)

		if err != nil {
			return ret, fmt.Errorf("admin req: %s", err.Error())
		}

		pub_pem := pem.EncodeToMemory(
			&pem.Block{
				Type:  "PUBLIC KEY",
				Bytes: pub_b,
			},
		)

		pub_pem_str := string(pub_pem)

		err = UpdatePubkeyByEmail(email, pub_pem_str)

		if err != nil {
			return ret, fmt.Errorf("admin req: %s", err.Error())
		}

		ret_apiout.BODY = string(priv_pem)

		ret, err = json.Marshal(ret_apiout)

		if err != nil {
			return ret, fmt.Errorf("admin req: %s", err.Error())
		}

	case "ORCH-ADDCL":

		cluster_id := args[0]

		token, err := CreateClusterByEmail(email, cluster_id)

		if err != nil {
			return ret, fmt.Errorf("admin req: %s", err.Error())
		}

		ret_apiout.BODY = token

		ret, err = json.Marshal(ret_apiout)

		if err != nil {
			return ret, fmt.Errorf("admin req: %s", err.Error())
		}

	case "ORCH-INSTCL":

		if len(FI_SESSIONS.INST_SESSION) > 100 {
			return ret, fmt.Errorf("admin req: too many remote install sessions")
		}

		cluster_id := args[0]
		targetip := args[1]
		targetid := args[2]
		targetpw := args[3]
		localip := args[4]
		osnm := args[5]
		cv := args[6]
		update_token := args[7]

		session_key := email + ":" + cluster_id

		_, okay := FI_SESSIONS.INST_SESSION[session_key]

		if okay {
			return ret, fmt.Errorf("admin req: already an ongoing installation")
		}

		FI_SESSIONS.INST_SESSION[session_key] = &[]byte{}

		FI_SESSIONS.INST_RESULT[session_key] = "-"

		go InstallCluster(session_key, cluster_id, targetip, targetid, targetpw, localip, osnm, cv, update_token)

		ret_apiout.BODY = "remote cluster installation started\n"

		ret, err = json.Marshal(ret_apiout)

		if err != nil {
			return ret, fmt.Errorf("admin req: %s", err.Error())
		}

	case "ORCH-INSTCLLOG":

		cluster_id := args[0]
		targetip := args[1]
		targetid := args[2]
		targetpw := args[3]

		session_key := email + ":" + cluster_id

		log_b, err := InstallClusterLog(session_key, cluster_id, targetip, targetid, targetpw)

		if err != nil {
			return ret, fmt.Errorf("admin req: %s", err.Error())
		}

		ret_apiout.BODY = string(log_b)

		ret, err = json.Marshal(ret_apiout)

		if err != nil {
			return ret, fmt.Errorf("admin req: %s", err.Error())
		}

	// case "SIGNOUT" :

	default:

		return ret, fmt.Errorf("admin req: %s", "no matching operand")

	}

	return ret, nil
}

func AdminRequestParser_Linear(query string) (string, []string, error) {

	var operand string

	args := make([]string, 0)

	linear_list := strings.SplitN(query, ":", 2)

	operand = linear_list[0]

	if len(linear_list) != 2 {
		return operand, args, fmt.Errorf("parsing linear inst: %s", "length")
	}

	args = append(args, strings.Split(linear_list[1], ",")...)

	return operand, args, nil

}
