package main

import (
	"net/http"
	"strings"
	"time"

	ctrl "github.com/OKESTRO-AIDevOps/npia-server/src/controller"
	"github.com/OKESTRO-AIDevOps/npia-server/src/modules"
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

	defer c.Close()

	for auth_flag == 0 {

		err := c.ReadJSON(&req)
		if err != nil {
			EventLogger("auth:" + err.Error())
			return
		}

		chal_id := req.ChallengeID

		switch chal_id {

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

			SERVER_KEY_ENTRY[email_context] = gen_key

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

	}

	EventLogger("server accepted")

}
