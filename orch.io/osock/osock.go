package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	_ "log"
	"net/http"
	_ "net/http"
	_ "os"
	_ "time"

	sctrl "github.com/OKESTRO-AIDevOps/nkia/orch.io/osock/controller"
	models "github.com/OKESTRO-AIDevOps/nkia/orch.io/osock/models"
	modules "github.com/OKESTRO-AIDevOps/nkia/pkg/challenge"
	"github.com/gorilla/websocket"

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

	data_records := make([]models.OrchRecord, 0)

	cluster_data_records := make([]models.OrchClusterRecord, 0)

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

	cmd = exec.Command("mkdir", "-p", ".npia/certs")

	err = cmd.Run()

	if err != nil {

		return fmt.Errorf("failed init npia orchestrator: %s", err.Error())
	}

	cmd = exec.Command("mkdir", "-p", ".npia/data")

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

	data_records_b, err := json.Marshal(data_records)

	if err != nil {

		return fmt.Errorf("failed init npia orchestrator: %s", err.Error())
	}

	cluster_data_records_b, err := json.Marshal(cluster_data_records)

	if err != nil {

		return fmt.Errorf("failed init npia orchestrator: %s", err.Error())
	}

	err = os.WriteFile(".npia/key.json", key_records_b, 0644)

	if err != nil {

		return fmt.Errorf("failed init npia orchestrator: %s", err.Error())
	}

	err = os.WriteFile(".npia/challenge.json", challenge_records_b, 0644)

	if err != nil {

		return fmt.Errorf("failed init npia orchestrator: %s", err.Error())
	}

	err = os.WriteFile(".npia/data/record.json", data_records_b, 0644)

	if err != nil {

		return fmt.Errorf("failed init npia orchestrator: %s", err.Error())
	}

	err = os.WriteFile(".npia/data/cluster_record.json", cluster_data_records_b, 0644)

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

	sctrl.LoadConfig()

	/*

		var db_info string

		if !sctrl.CONFIG_JSON.DEBUG {

			db_info = fmt.Sprintf("%s:%s@tcp(%s)/%s",
				sctrl.CONFIG_JSON.DB_ID,
				sctrl.CONFIG_JSON.DB_PW,
				sctrl.CONFIG_JSON.DB_HOST,
				sctrl.CONFIG_JSON.DB_NAME,
			)

		} else {

			db_info = fmt.Sprintf("%s:%s@tcp(%s)/%s",
				sctrl.CONFIG_JSON.DB_ID,
				sctrl.CONFIG_JSON.DB_PW,
				sctrl.CONFIG_JSON.DB_HOST_DEV,
				sctrl.CONFIG_JSON.DB_NAME,
			)

		}

		sctrl.DB, _ = sql.Open("mysql", db_info)

		sctrl.DB.SetConnMaxLifetime(time.Second * 10)
		sctrl.DB.SetConnMaxIdleTime(time.Second * 5)
		sctrl.DB.SetMaxOpenConns(10)
		sctrl.DB.SetMaxIdleConns(10)


	*/

	file_b, err := os.ReadFile(".npia/certs/ca.crt")

	if err != nil {

		panic(err)
	}

	crt, err := modules.BytesToCert(file_b)

	if err != nil {
		panic(err)
	}

	sctrl.CA_CERT = crt

	sctrl.EventLogger(fmt.Sprintf("started at: %s", *ADDR))

	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/osock/server/test", ServerHandler_Test)
	http.HandleFunc("/osock/server", ServerHandler)
	http.HandleFunc("/osock/front/test", FrontHandler2_Test)
	http.HandleFunc("/osock/front", FrontHandler2)
	log.Fatal(http.ListenAndServeTLS(*ADDR, ".npia/certs/server.crt", ".npia/certs/server.priv", nil))

}
