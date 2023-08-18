package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	_ "log"
	"net/http"
	_ "net/http"
	_ "os"
	_ "time"

	"github.com/OKESTRO-AIDevOps/nkia/src/modules"
	"github.com/gorilla/websocket"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var ADDR = flag.String("addr", "0.0.0.0:7331", "service address")

var UPGRADER = websocket.Upgrader{} // use default options

var FRONT_CONNECTION = make(map[string]*websocket.Conn)

var FRONT_CONNECTION_FRONT = make(map[*websocket.Conn]string)

var SERVER_CONNECTION = make(map[string]*websocket.Conn)

var SERVER_CONNECTION_KEY = make(map[*websocket.Conn]string)

var SERVER_CONNECTION_FRONT = make(map[*websocket.Conn]string)

func O_Init() error {
	challenge_records := make(modules.ChallengRecord)

	key_records := make(modules.KeyRecord)

	challenge_records["_INIT"] = map[string]string{
		"_INIT": "_INIT",
	}

	key_records["_INIT"] = "_INIT"

	cmd := exec.Command("mkdir", "-p", "srv")

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

	err = os.WriteFile("srv/challenge.json", challenge_records_b, 0644)

	if err != nil {

		return fmt.Errorf("failed init npia orchestrator: %s", err.Error())
	}

	err = os.WriteFile("srv/key.json", key_records_b, 0644)

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

	DB, _ = sql.Open("mysql", "npiaorchestrator:youdonthavetoknow@tcp(npiaorchestratordb:3306)/orchestrator")

	DB.SetConnMaxLifetime(time.Second * 10)
	DB.SetConnMaxIdleTime(time.Second * 5)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	EventLogger("DB Connected")

	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/osock/server-test", ServerHandler_Test)
	http.HandleFunc("/osock/server", ServerHandler)
	http.HandleFunc("/osock/front-test", FrontHandler_Test)
	http.HandleFunc("/osock/front", FrontHandler)
	log.Fatal(http.ListenAndServe(*ADDR, nil))

}
