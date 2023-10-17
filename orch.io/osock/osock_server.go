package main

import (
	"encoding/hex"
	"net/http"
	"strings"
	"time"

	ctrl "github.com/OKESTRO-AIDevOps/nkia/nokubelet/controller"
	"github.com/OKESTRO-AIDevOps/nkia/nokubelet/modules"
	"github.com/gorilla/websocket"
)

func ServerHandler(w http.ResponseWriter, r *http.Request) {

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

	var ANSWER string

	defer c.Close()

	for auth_flag == 0 {

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

		case "UPDATE":

			ANSWER, _ = modules.RandomHex(128)

			email_context := req.ChallengeMessage

			email_context_list := strings.Split(email_context, ":")

			if len(email_context_list) != 2 {
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
				if err != nil {
					EventLogger("auth update write close:" + err.Error())
					return
				}
				EventLogger("auth update: wrong format")
				return
			}

			email := email_context_list[0]

			cluster_id := email_context_list[1]

			token, err := GetConfigChallengeByEmailAndClusterID(email, cluster_id)

			if err != nil {
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
				if err != nil {
					EventLogger("auth update write close:" + err.Error())
					return
				}
				EventLogger("auth update: wrong format")
				return
			}

			token_b := []byte(token)

			QUEST, err := modules.DecryptWithSymmetricKey(token_b, []byte(ANSWER))

			if err != nil {
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
				if err != nil {
					EventLogger("auth update write close:" + err.Error())
					return
				}
				EventLogger("auth update: wrong format")
				return
			}

			quest_hex := hex.EncodeToString(QUEST)

			resp.ChallengeID = "UPDATE"

			resp.ChallengeMessage = quest_hex

			err = c.WriteJSON(resp)

			if err != nil {
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
				if err != nil {
					EventLogger("auth update write close:" + err.Error())
					return
				}
				EventLogger("auth update: wrong format")
				return
			}

		case "ROTATE":

			email_context := req.ChallengeMessage

			email_context_list := strings.Split(email_context, ":")

			if len(email_context_list) != 4 {
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
				if err != nil {
					EventLogger("auth update write close:" + err.Error())
					return
				}
				EventLogger("auth update: wrong format")
				return
			}

			email := email_context_list[0]

			cluster_id := email_context_list[1]

			answer := email_context_list[2]

			config := email_context_list[3]

			token, err := GetConfigChallengeByEmailAndClusterID(email, cluster_id)

			if err != nil {
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
				if err != nil {
					EventLogger("auth rotate write close:" + err.Error())
					return
				}
				EventLogger("auth rotate: wrong format")
				return
			}

			token_b := []byte(token)

			answer_b, err := hex.DecodeString(answer)

			if err != nil {
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
				if err != nil {
					EventLogger("auth rotate write close:" + err.Error())
					return
				}
				EventLogger("auth rotate: wrong format")
				return
			}

			submit_ans, err := modules.DecryptWithSymmetricKey(token_b, answer_b)

			if err != nil {
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
				if err != nil {
					EventLogger("auth rotate write close:" + err.Error())
					return
				}
				EventLogger("auth rotate: wrong format")
				return
			}

			if ANSWER != string(submit_ans) {

				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
				if err != nil {
					EventLogger("auth rotate write close:" + err.Error())
					return
				}
				EventLogger("auth rotate: wrong format")
				return
			}

			config_hex, err := hex.DecodeString(config)

			if err != nil {
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
				if err != nil {
					EventLogger("auth rotate write close:" + err.Error())
					return
				}
				EventLogger("auth rotate: wrong format")
				return
			}

			config_dec, err := modules.DecryptWithSymmetricKey(token_b, config_hex)

			if err != nil {
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
				if err != nil {
					EventLogger("auth rotate write close:" + err.Error())
					return
				}
				EventLogger("auth rotate: wrong format")
				return
			}

			config_dec_string := string(config_dec)

			err = AddClusterByEmailAndClusterID(email, cluster_id, config_dec_string)

			if err != nil {
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
				if err != nil {
					EventLogger("auth rotate write close:" + err.Error())
					return
				}
				EventLogger("auth rotate: wrong format")
				return
			}

			resp.ChallengeID = "ROTATE"

			resp.ChallengeMessage = "SUCCESS"

			err = c.WriteJSON(resp)

			if err != nil {
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
				if err != nil {
					EventLogger("auth rotate write close:" + err.Error())
					return
				}
				EventLogger("auth rotate: wrong format")
				return
			}

		case "ASK":

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

			cluster_id := email_context_list[1]

			config_b, err := GetKubeconfigByEmailAndClusterID(email, cluster_id)

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

			cluster_id := email_context_list[1]

			config_b, err := GetKubeconfigByEmailAndClusterID(email, cluster_id)

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
