package main

import (
	b64 "encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
	"time"

	ctrl "github.com/OKESTRO-AIDevOps/nkia/nokubelet/controller"
	"github.com/OKESTRO-AIDevOps/nkia/nokubelet/modules"
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

	auth_flag := 0

	iter_count := 0

	var ANSWER string

	defer c.Close()

	for auth_flag == 0 {

		var req = ctrl.AuthChallenge{}
		var resp = ctrl.AuthChallenge{}

		if iter_count > 10 {
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
			if err != nil {
				EventLogger("auth iter write close:" + err.Error())
				return
			}
			EventLogger("auth iter: limit")
			return
		}

		err := c.ReadJSON(&req)
		if err != nil {
			EventLogger("auth:" + err.Error())
			return
		}

		chal_id := req.ChallengeID

		switch chal_id {

		case "UPDATE":

			ANSWER, _ = modules.RandomHex(128)

			email_context := req.ChallengeMessage

			email_context_list := strings.Split(email_context, ":")

			if len(email_context_list) != 2 {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
				EventLogger("auth update: wrong format")
				return
			}

			email := email_context_list[0]

			cluster_id := email_context_list[1]

			token, err := GetConfigChallengeByEmailAndClusterID(email, cluster_id)

			if err != nil {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
				EventLogger("auth update: get config: " + err.Error())
				return
			}

			token_b := []byte(token)

			QUEST, err := modules.EncryptWithSymmetricKey(token_b, []byte(ANSWER))

			if err != nil {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
				EventLogger("auth update: encrypt: " + err.Error())
				return
			}

			quest_hex := hex.EncodeToString(QUEST)

			resp.ChallengeID = "UPDATE"

			resp.ChallengeMessage = quest_hex

			err = c.WriteJSON(resp)

			if err != nil {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
				EventLogger("auth update: write: " + err.Error())
				return
			}

		case "ROTATE":

			email_context := req.ChallengeMessage

			email_context_list := strings.Split(email_context, ":")

			if len(email_context_list) != 4 {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
				EventLogger("auth rotate: wrong format")
				return
			}

			email := email_context_list[0]

			cluster_id := email_context_list[1]

			answer := email_context_list[2]

			config := email_context_list[3]

			token, err := GetConfigChallengeByEmailAndClusterID(email, cluster_id)

			if err != nil {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
				EventLogger("auth rotate: get config: " + err.Error())
				return
			}

			token_b := []byte(token)

			if ANSWER != answer {

				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				EventLogger("auth rotate: answer: " + err.Error())
				return
			}

			config_hex, err := hex.DecodeString(config)

			if err != nil {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				EventLogger("auth rotate: decode config: " + err.Error())
				return
			}

			config_dec, err := modules.DecryptWithSymmetricKey(token_b, config_hex)

			if err != nil {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				EventLogger("auth rotate: decrypt config: " + err.Error())
				return
			}

			config_dec_string := string(config_dec)

			err = AddClusterByEmailAndClusterID(email, cluster_id, config_dec_string)

			if err != nil {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
				EventLogger("auth rotate: add cluster: " + err.Error())
				return
			}

			resp.ChallengeID = "ROTATE"

			resp.ChallengeMessage = "SUCCESS"

			err = c.WriteJSON(resp)

			if err != nil {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				EventLogger("auth rotate: write: " + err.Error())
				return
			}

		case "ASK":

			email_context := req.ChallengeMessage

			email_context_list := strings.Split(email_context, ":")

			if len(email_context_list) != 2 {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				EventLogger("auth ask: wrong format")
				return
			}

			email := email_context_list[0]

			cluster_id := email_context_list[1]

			config_b, err := GetKubeconfigByEmailAndClusterID(email, cluster_id)

			if err != nil {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				EventLogger("auth ask: get config: " + err.Error())
				return
			}

			client_ca_pub_key := req.ChallengeData

			chal_rec, err := modules.GenerateChallenge_Detached(config_b, client_ca_pub_key)

			if err != nil {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				EventLogger("auth ask: gen chal: " + err.Error())
				return
			}

			resp.ChallengeID = "ASK"
			resp.ChallengeData = chal_rec

			err = c.WriteJSON(&resp)

			if err != nil {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				EventLogger("auth ask: write: " + err.Error())
				return
			}

		case "ANS":

			email_context := req.ChallengeMessage

			email_context_list := strings.Split(email_context, ":")

			if len(email_context_list) != 2 {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				EventLogger("auth ask: wrong format")
				return
			}

			email := email_context_list[0]

			cluster_id := email_context_list[1]

			config_b, err := GetKubeconfigByEmailAndClusterID(email, cluster_id)

			if err != nil {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				EventLogger("auth ans: get config: " + err.Error())
				return
			}

			answer := req.ChallengeData

			gen_key, key_rec, err := modules.VerifyChallange_Detached(config_b, answer)

			if err != nil {
				_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

				EventLogger("auth ans: verify: " + err.Error())
				return
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

				EventLogger("auth ans: write: " + err.Error())
				return
			}

			auth_flag = 1

		default:
			_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

			EventLogger("auth blank: default")
			return

		}

		iter_count += 1

	}

	EventLogger("server accepted")

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

	var res_server ctrl.APIMessageResponse
	var res_orchestrator ctrl.OrchestratorResponse

	for {

		err := c.ReadJSON(&res_server)
		if err != nil {
			ServerDestructor(c)
			c.Close()
			EventLogger("response:" + err.Error())
			return
		}

		fmt.Println("************************")
		fmt.Println("RECV SERVER")
		fmt.Println(res_server)

		key_id, okay := SERVER_CONNECTION_KEY[c]

		if !okay {

			ServerDestructor(c)
			_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

			EventLogger("response key:" + err.Error())
			return

		}

		front_name, okay := SERVER_CONNECTION_FRONT[c]

		if !okay {

			ServerDestructor(c)
			_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

			EventLogger("response front:" + err.Error())
			return

		}

		front_c, okay := FRONT_CONNECTION[front_name]

		if !okay {

			ServerDestructor(c)
			_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

			EventLogger("response front conn:" + err.Error())
			return

		}

		session_sym_key, err := modules.AccessAuth_Detached(key_id)

		if err != nil {
			ServerDestructor(c)
			_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

			EventLogger("response access:" + err.Error())
			return

		}

		res_orchestrator.ServerMessage = res_server.ServerMessage

		resp_enc := res_server.QueryResult

		resp_enc_b, err := hex.DecodeString(resp_enc)

		if err != nil {
			ServerDestructor(c)
			_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

			EventLogger("response dec:" + err.Error())
			return
		}

		resp_dec, err := modules.DecryptWithSymmetricKey([]byte(session_sym_key), resp_enc_b)

		if err != nil {
			ServerDestructor(c)
			_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

			EventLogger("response dec 2:" + err.Error())
			return
		}

		res_orchestrator.QueryResult = resp_dec

		err = front_c.WriteJSON(&res_orchestrator)

		fmt.Println("************************")
		fmt.Println("SENT TO FRONT")
		fmt.Println(res_orchestrator)

		if err != nil {
			ServerDestructor(c)
			c.Close()
			EventLogger("response send:" + err.Error())
			return
		}

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
