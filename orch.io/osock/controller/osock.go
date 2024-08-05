package controller

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	orchcmd "github.com/OKESTRO-AIDevOps/nkia/orch.io/osock/cmd"
	ctrl "github.com/OKESTRO-AIDevOps/nkia/pkg/apistandard/apix"
	"github.com/gorilla/websocket"
)

var DB *sql.DB

var CONFIG_JSON ConfigJSON

var FRONT_CONNECTION = make(map[string]*websocket.Conn)

var FRONT_CONNECTION_FRONT = make(map[*websocket.Conn]string)

var SERVER_CONNECTION = make(map[string]*websocket.Conn)

var SERVER_CONNECTION_KEY = make(map[*websocket.Conn]string)

var SERVER_CONNECTION_FRONT = make(map[*websocket.Conn]string)

var SERVER_CONNECTION_CMD = make(map[*websocket.Conn]string)

var UPGRADER = websocket.Upgrader{} // use default options

var L = log.New(os.Stdout, "", 0)

func FrontHandler2(w http.ResponseWriter, r *http.Request) {

	EventLogger("Front access")

	UPGRADER.CheckOrigin = func(r *http.Request) bool { return true }

	c, err := UPGRADER.Upgrade(w, r, nil)
	if err != nil {
		EventLogger("upgrade:" + err.Error())
		return
	}

	c.SetReadDeadline(time.Time{})

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

	for {

		req_orchestrator = ctrl.OrchestratorRequest{}

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

			res_orchestrator.QueryResult = result

			c.WriteJSON(&res_orchestrator)

			continue

		}

		Emit(c, email, target, query_str)

	}

}

func ServerHandler(w http.ResponseWriter, r *http.Request) {

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
