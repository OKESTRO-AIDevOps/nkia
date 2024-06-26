package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"

	_ "log"
	"net/http"
	_ "net/http"
	_ "os"
	_ "time"

	"github.com/OKESTRO-AIDevOps/nkia/nokubelet/modules"
	"github.com/gorilla/websocket"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type InstallSession struct {
	mu sync.Mutex

	INST_SESSION map[string]*[]byte

	INST_RESULT map[string]string
}

var ADDR = flag.String("addr", "0.0.0.0:7331", "service address")

var UPGRADER = websocket.Upgrader{} // use default options

var FRONT_CONNECTION = make(map[string]*websocket.Conn)

var FRONT_CONNECTION_FRONT = make(map[*websocket.Conn]string)

var SERVER_CONNECTION = make(map[string]*websocket.Conn)

var SERVER_CONNECTION_KEY = make(map[*websocket.Conn]string)

var SERVER_CONNECTION_FRONT = make(map[*websocket.Conn]string)

var FI_SESSIONS = InstallSession{
	INST_SESSION: make(map[string]*[]byte),
	INST_RESULT:  make(map[string]string),
}

func O_Init() error {
	challenge_records := make(modules.ChallengRecord)

	key_records := make(modules.KeyRecord)

	challenge_records["_INIT"] = map[string]string{
		"_INIT": "_INIT",
	}

	key_records["_INIT"] = "_INIT"

	cmd := exec.Command("mkdir", "-p", ".npia")

	err := cmd.Run()

	if err != nil {

		return fmt.Errorf("failed init npia orchestrator: %s", err.Error())
	}

	cmd = exec.Command("mkdir", "-p", "session")

	err = cmd.Run()

	if err != nil {

		return fmt.Errorf("failed init npia orchestrator: %s", err.Error())
	}

	challenge_records_b, err := json.Marshal(challenge_records)

	if err != nil {

		return fmt.Errorf("failed init npia orchestrator: %s", err.Error())
	}

	key_records_b, err := json.Marshal(key_records)

	if err != nil {

		return fmt.Errorf("failed init npia orchestrator: %s", err.Error())
	}

	err = os.WriteFile(".npia/challenge.json", challenge_records_b, 0644)

	if err != nil {

		return fmt.Errorf("failed init npia orchestrator: %s", err.Error())
	}

	err = os.WriteFile(".npia/key.json", key_records_b, 0644)

	if err != nil {

		return fmt.Errorf("failed init npia orchestrator: %s", err.Error())
	}

	return nil
}

func main() {

	err := O_Init()

	if err != nil {
		panic(err.Error())
	}

	LoadConfig()

	var db_info string

	if !CONFIG_JSON.DEBUG {

		db_info = fmt.Sprintf("%s:%s@tcp(%s)/%s",
			CONFIG_JSON.DB_ID,
			CONFIG_JSON.DB_PW,
			CONFIG_JSON.DB_HOST,
			CONFIG_JSON.DB_NAME,
		)

	} else {

		db_info = fmt.Sprintf("%s:%s@tcp(%s)/%s",
			CONFIG_JSON.DB_ID,
			CONFIG_JSON.DB_PW,
			CONFIG_JSON.DB_HOST_DEV,
			CONFIG_JSON.DB_NAME,
		)

	}

	DB, _ = sql.Open("mysql", db_info)

	DB.SetConnMaxLifetime(time.Second * 10)
	DB.SetConnMaxIdleTime(time.Second * 5)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	EventLogger("DB Connected")

	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/osock/server/test", ServerHandler_Test)
	http.HandleFunc("/osock/server", ServerHandler)
	http.HandleFunc("/osock/front/test", FrontHandler_Test)
	http.HandleFunc("/osock/front", FrontHandler)
	log.Fatal(http.ListenAndServe(*ADDR, nil))

}
