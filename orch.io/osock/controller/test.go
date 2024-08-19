package controller

import (
	b64 "encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	orchcmd "github.com/OKESTRO-AIDevOps/nkia/orch.io/osock/cmd"
	"github.com/OKESTRO-AIDevOps/nkia/orch.io/osock/models"
	ctrl "github.com/OKESTRO-AIDevOps/nkia/pkg/apistandard/apix"
	modules "github.com/OKESTRO-AIDevOps/nkia/pkg/challenge"
	"github.com/gorilla/websocket"
)

func ServerHandler_Test(w http.ResponseWriter, r *http.Request) {
	EventLogger("Server access")

	UPGRADER.CheckOrigin = func(r *http.Request) bool { return true }

	c, err := UPGRADER.Upgrade(w, r, nil)
	if err != nil {
		EventLogger("upgrade:" + err.Error())
		return
	}

	c.SetReadDeadline(time.Time{})

	defer c.Close()

	cluster_id, err := RemoteAccessAuthChallenge(c)

	if err != nil {

		EventLogger(err.Error())

		return
	}

	EventLogger("server accepted")

	for {

		OnData(c, cluster_id)

	}
}

func FrontHandler_Test(w http.ResponseWriter, r *http.Request) {

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

		email, err := models.CheckSessionAndGetEmailByRequestKey(request_key)

		if err != nil {
			EventLogger("auth:" + err.Error())
			return
		}

		FRONT_CONNECTION[email] = c

		FRONT_CONNECTION_FRONT[c] = email

		break
	}

	EventLogger("front accepted")

	fmt.Println("server connection ---------- ")
	fmt.Println(SERVER_CONNECTION)
	fmt.Println("server connection key ------ ")
	fmt.Println(SERVER_CONNECTION_KEY)
	fmt.Println("server connection front ---- ")
	fmt.Println(SERVER_CONNECTION_FRONT)
	fmt.Println("front connection ------------ ")
	fmt.Println(FRONT_CONNECTION)
	fmt.Println("front conntction front ------ ")
	fmt.Println(FRONT_CONNECTION_FRONT)

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

		fmt.Println("************************")
		fmt.Println("RECV FRONT")
		fmt.Println(req_orchestrator)

		target := req_orchestrator.RequestTarget

		email, okay := FRONT_CONNECTION_FRONT[c]

		if !okay {
			FrontDestructor(c)
			_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

			EventLogger("read front: no connected front name")
			fmt.Println("front conntction front ------ ")
			fmt.Println(FRONT_CONNECTION_FRONT)
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
			EventLogger("read front: no server context key")
			return
		}

		query_b := []byte(query_str)

		query_enc, err := modules.EncryptWithSymmetricKey([]byte(session_sym_key), query_b)

		if err != nil {
			FrontDestructor(c)
			_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
			EventLogger("read front: no server context key")
			return
		}

		query_hex := hex.EncodeToString(query_enc)

		req_server.Query = query_hex

		err = server_c.WriteJSON(&req_server)

		fmt.Println("************************")
		fmt.Println("SENT TO SOCK")
		fmt.Println(req_server)

		if err != nil {
			FrontDestructor(c)
			c.Close()
			EventLogger("write to server: " + err.Error())
			return
		}

	}

}

func FrontHandler2_Test(w http.ResponseWriter, r *http.Request) {

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

	defer c.Close()

	user_id, err := CertAuthChallenge(c)

	if err != nil {

		EventLogger(fmt.Sprintf("auth: %s", err.Error()))
		return
	}

	FRONT_CONNECTION[user_id] = c

	FRONT_CONNECTION_FRONT[c] = user_id

	EventLogger("front accepted")

	fmt.Println("server connection ---------- ")
	fmt.Println(SERVER_CONNECTION)
	fmt.Println("server connection key ------ ")
	fmt.Println(SERVER_CONNECTION_KEY)
	fmt.Println("server connection front ---- ")
	fmt.Println(SERVER_CONNECTION_FRONT)
	fmt.Println("front connection ------------ ")
	fmt.Println(FRONT_CONNECTION)
	fmt.Println("front conntction front ------ ")
	fmt.Println(FRONT_CONNECTION_FRONT)

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

		fmt.Println("************************")
		fmt.Println("RECV FRONT")
		fmt.Println(req_orchestrator)

		target := req_orchestrator.RequestTarget

		email, okay := FRONT_CONNECTION_FRONT[c]

		if !okay {
			FrontDestructor(c)
			_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

			EventLogger("read front: no connected front name")

			return
		}

		email_context := email + ":" + target

		query_str := req_orchestrator.Query

		forward, result, err := orchcmd.RequestForwardHandler(email, query_str)

		if err != nil {

			EventLogger("read front: " + err.Error())

			FrontDestructor(c)

			res_orchestrator.ServerMessage = err.Error()

			c.WriteJSON(&res_orchestrator)

			_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

			return

		}

		if !forward {

			res_orchestrator.ServerMessage = "SUCCESS"

			res_orchestrator.QueryResult = []byte(result)

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

		fmt.Println("************************")
		fmt.Println("SENT TO SOCK")
		fmt.Println(req_server)

		if err != nil {
			FrontDestructor(c)
			c.Close()
			EventLogger("write to server: " + err.Error())

			return
		}

	}

}
