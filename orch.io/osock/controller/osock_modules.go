package controller

import (
	"encoding/hex"
	"encoding/json"
	"os"
	"strings"
	"sync"
	"time"

	ctrl "github.com/OKESTRO-AIDevOps/nkia/pkg/apistandard/apix"
	modules "github.com/OKESTRO-AIDevOps/nkia/pkg/challenge"
	"github.com/gorilla/websocket"
)

func EventLogger(msg string) {
	L.SetPrefix("[" + time.Now().UTC().Format("2006-01-02 15:04:05.000") + "]")
	L.Print(msg)
}

type ConfigJSON struct {
	DEBUG       bool   `json:"DEBUG"`
	DB_HOST     string `json:"DB_HOST"`
	DB_ID       string `json:"DB_ID"`
	DB_PW       string `json:"DB_PW"`
	DB_NAME     string `json:"DB_NAME"`
	DB_HOST_DEV string `json:"DB_HOST_DEV"`
}

type InstallSession struct {
	mu sync.Mutex

	INST_SESSION map[string]*[]byte

	INST_RESULT map[string]string
}

var FI_SESSIONS = InstallSession{
	INST_SESSION: make(map[string]*[]byte),
	INST_RESULT:  make(map[string]string),
}

func LoadConfig() {

	CONFIG_JSON = GetConfigJSON()

}

func GetConfigJSON() ConfigJSON {

	var cj ConfigJSON

	file_byte, err := os.ReadFile("config.json")

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file_byte, &cj)

	if err != nil {
		panic(err)
	}

	return cj

}

func Emit(c *websocket.Conn, email string, context string, query_str string) {

	var req_server ctrl.APIMessageRequest

	email_context := email + ":" + context

	server_c, okay := SERVER_CONNECTION[email_context]

	if !okay {
		FrontDestructor(c)
		_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

		msg := "emit: no connected server context"

		EventLogger(msg)

		return
	}

	key_id, okay := SERVER_CONNECTION_KEY[server_c]

	if !okay {

		FrontDestructor(c)
		_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

		msg := "emit: no server context key"
		EventLogger(msg)

		return
	}

	session_sym_key, err := modules.AccessAuth_Detached(key_id)

	if err != nil {

		FrontDestructor(c)
		_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

		msg := "emit: " + err.Error()

		EventLogger(msg)

		return
	}

	query_b := []byte(query_str)

	query_enc, err := modules.EncryptWithSymmetricKey([]byte(session_sym_key), query_b)

	if err != nil {
		FrontDestructor(c)
		_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))
		msg := "emit: " + err.Error()

		EventLogger(msg)

		return
	}

	query_hex := hex.EncodeToString(query_enc)

	req_server.Query = query_hex

	SERVER_CONNECTION_CMD[server_c] = query_str

	err = server_c.WriteJSON(&req_server)

	if err != nil {
		FrontDestructor(c)
		c.Close()
		msg := "emit: " + err.Error()

		EventLogger(msg)

		return
	}

	return
}

func OnData(c *websocket.Conn, cluster_id string) {

	var res_server ctrl.APIMessageResponse
	var res_orchestrator ctrl.OrchestratorResponse

	err := c.ReadJSON(&res_server)
	if err != nil {
		ServerDestructor(c)
		c.Close()
		EventLogger("on data:" + err.Error())

		return
	}

	key_id, okay := SERVER_CONNECTION_KEY[c]

	if !okay {

		ServerDestructor(c)
		_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

		EventLogger("on data:" + err.Error())

		return

	}

	front_name, okay := SERVER_CONNECTION_FRONT[c]

	if !okay {
		ServerDestructor(c)
		_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

		EventLogger("on data:" + err.Error())

		return

	}

	front_c, okay := FRONT_CONNECTION[front_name]

	if !okay {
		ServerDestructor(c)
		_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

		EventLogger("on data:" + err.Error())

		return

	}

	session_sym_key, err := modules.AccessAuth_Detached(key_id)

	if err != nil {
		ServerDestructor(c)
		_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

		EventLogger("on data:" + err.Error())

		return

	}

	res_orchestrator.ServerMessage = res_server.ServerMessage

	resp_enc := res_server.QueryResult

	resp_enc_b, err := hex.DecodeString(resp_enc)

	if err != nil {
		ServerDestructor(c)
		_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

		EventLogger("on data:" + err.Error())

		return
	}

	resp_dec, err := modules.DecryptWithSymmetricKey([]byte(session_sym_key), resp_enc_b)

	if err != nil {
		ServerDestructor(c)
		_ = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Connection Close"))

		EventLogger("on data:" + err.Error())

		return
	}

	res_orchestrator.QueryResult = resp_dec

	if res_orchestrator.ServerMessage != "SUCCESS" {

		err = front_c.WriteJSON(&res_orchestrator)

		if err != nil {
			ServerDestructor(c)
			c.Close()
			EventLogger("on data:" + err.Error())

			return
		}

		return
	}

}

func FrontDestructor(c *websocket.Conn) {

	EventLogger("front destructor called")

	fc, _ := FRONT_CONNECTION_FRONT[c]

	delete(FRONT_CONNECTION_FRONT, c)

	delete(FRONT_CONNECTION, fc)

	for k, _ := range SERVER_CONNECTION {

		if strings.Contains(k, fc) {

			delete(SERVER_CONNECTION, k)

		}

	}

	EventLogger("front destructor exit")
}

func ServerDestructor(c *websocket.Conn) {

	EventLogger("server destructor called")

	sc, _ := SERVER_CONNECTION_FRONT[c]

	delete(SERVER_CONNECTION_FRONT, c)

	delete(FRONT_CONNECTION, sc)

	delete(SERVER_CONNECTION_KEY, c)

	EventLogger("server destructor exit")
}
