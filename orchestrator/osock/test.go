package main

import (
	b64 "encoding/base64"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
	"time"

	ctrl "github.com/OKESTRO-AIDevOps/nkia/src/controller"
	"github.com/OKESTRO-AIDevOps/nkia/src/modules"
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

	var req ctrl.AuthChallenge
	var resp ctrl.AuthChallenge

	auth_flag := 0

	iter_count := 0

	defer c.Close()

	for auth_flag == 0 {

		req = ctrl.AuthChallenge{}

		if iter_count > 5 {
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

		case "ASK":

			email_context := req.ChallengeMessage

			EventLogger(email_context)

			email_context_list := strings.Split(email_context, ":")

			if len(email_context_list) != 2 {
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
				if err != nil {
					EventLogger("auth ask write close:" + err.Error())
					return
				}
				EventLogger("auth ask: wrong format")
				return
			}

			email := email_context_list[0]

			config_b, err := GetKubeconfigByEmail(email)

			if err != nil {
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
				if err != nil {
					EventLogger("auth ask write close:" + err.Error())
					return
				}
				EventLogger("auth ask:" + err.Error())
				return
			}

			client_ca_pub_key := req.ChallengeData

			chal_rec, err := modules.GenerateChallenge_Detached(config_b, client_ca_pub_key)

			if err != nil {
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
				if err != nil {
					EventLogger("auth ask write close:" + err.Error())
					return
				}
				EventLogger("auth ask:" + err.Error())
				return
			}

			resp.ChallengeID = "ASK"
			resp.ChallengeData = chal_rec

			err = c.WriteJSON(&resp)

			if err != nil {
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
				if err != nil {
					EventLogger("auth ask write close:" + err.Error())
					return
				}
				EventLogger("auth ask:" + err.Error())
				return
			}

		case "ANS":

			email_context := req.ChallengeMessage

			email_context_list := strings.Split(email_context, ":")

			if len(email_context_list) != 2 {
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
				if err != nil {
					EventLogger("auth ask write close:" + err.Error())
					return
				}
				EventLogger("auth ask: wrong format")
				return
			}

			email := email_context_list[0]

			config_b, err := GetKubeconfigByEmail(email)

			if err != nil {
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
				if err != nil {
					EventLogger("auth ans write close:" + err.Error())
					return
				}
				EventLogger("auth ans:" + err.Error())
				return
			}

			answer := req.ChallengeData

			gen_key, key_rec, err := modules.VerifyChallange_Detached(config_b, answer)

			if err != nil {
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
				if err != nil {
					EventLogger("auth ans write close:" + err.Error())
					return
				}
				EventLogger("auth ans:" + err.Error())
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
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
				if err != nil {
					EventLogger("auth ans write close:" + err.Error())
					return
				}
				EventLogger("auth ans:" + err.Error())
				return
			}

			auth_flag = 1

		default:
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
			if err != nil {
				EventLogger("auth blank write close:" + err.Error())
				return
			}
			EventLogger("auth blank: default")
			return

		}

		iter_count += 1

	}

	EventLogger("server accepted")

	var res_server ctrl.APIMessageResponse
	var res_orchestrator ctrl.OrchestratorResponse

	fmt.Println("server connection ---------- ")
	fmt.Println(SERVER_CONNECTION)
	fmt.Println("server connection key ------ ")
	fmt.Println(SERVER_CONNECTION_KEY)
	fmt.Println("server connection front ---- ")
	fmt.Println(SERVER_CONNECTION_FRONT)

	for {

		err := c.ReadJSON(&res_server)
		if err != nil {
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
			if err != nil {
				EventLogger("response write close:" + err.Error())
				return
			}
			EventLogger("response:" + err.Error())
			return
		}

		key_id, okay := SERVER_CONNECTION_KEY[c]

		if !okay {

			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
			if err != nil {
				EventLogger("response key write close:" + err.Error())
				return
			}
			EventLogger("response key:" + err.Error())
			return

		}

		front_name, okay := SERVER_CONNECTION_FRONT[c]

		if !okay {

			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
			if err != nil {
				EventLogger("response front write close:" + err.Error())
				return
			}
			EventLogger("response front:" + err.Error())
			return

		}

		front_c, okay := FRONT_CONNECTION[front_name]

		if !okay {

			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
			if err != nil {
				EventLogger("response front conn write close:" + err.Error())
				return
			}
			EventLogger("response front conn:" + err.Error())
			return

		}

		session_sym_key, err := modules.AccessAuth_Detached(key_id)

		if err != nil {
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
			if err != nil {
				EventLogger("response access write close:" + err.Error())
				return
			}
			EventLogger("response access:" + err.Error())
			return

		}

		res_orchestrator.ServerMessage = res_server.ServerMessage

		resp_enc := res_server.QueryResult

		resp_enc_b, err := hex.DecodeString(resp_enc)

		if err != nil {
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
			if err != nil {
				EventLogger("response dec write close:" + err.Error())
				return
			}
			EventLogger("response dec:" + err.Error())
			return
		}

		resp_dec, err := modules.DecryptWithSymmetricKey([]byte(session_sym_key), resp_enc_b)

		if err != nil {
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
			if err != nil {
				EventLogger("response dec 2 write close:" + err.Error())
				return
			}
			EventLogger("response dec 2:" + err.Error())
			return
		}

		res_orchestrator.QueryResult = resp_dec

		err = front_c.WriteJSON(&res_orchestrator)

		if err != nil {
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

		query_str := req_orchestrator.Query

		query_b := []byte(query_str)

		query_enc, err := modules.EncryptWithSymmetricKey([]byte(session_sym_key), query_b)

		if err != nil {
			EventLogger("read front: " + err.Error())
			return
		}

		query_hex := hex.EncodeToString(query_enc)

		req_server.Query = query_hex

		fmt.Println("******* server side")
		fmt.Println("written")
		fmt.Println(req_server)

		err = server_c.WriteJSON(&req_server)

		if err != nil {
			EventLogger("write to server: " + err.Error())
			return
		}

	}

}
