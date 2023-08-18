package omodels

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/OKESTRO-AIDevOps/nkia/orchestrator/ofront/omodules"
	"github.com/OKESTRO-AIDevOps/nkia/src/modules"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

type OrchestratorRecord_Email struct {
	email string
}

type OrchestratorRecord_RequestKey struct {
	request_key string
}

func DbEstablish() {

	DB, _ = sql.Open("mysql", "npiaorchestrator:youdonthavetoknow@tcp(npiaorchestratordb:3306)/orchestrator")

	DB.SetConnMaxLifetime(time.Second * 10)
	DB.SetConnMaxIdleTime(time.Second * 5)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	fmt.Println("DB Connected")

}

func DbQuery(query string, args []any) (*sql.Rows, error) {

	var empty_row *sql.Rows

	results, err := DB.Query(query, args[0:]...)

	if err != nil {

		return empty_row, err

	}

	return results, err

}

func FrontAccessAuth(session_id string) (string, error) {

	var request_key string

	var result_container_request_key []OrchestratorRecord_RequestKey

	q := "SELECT request_key FROM orchestrator_record WHERE osid = ?"

	a := []any{session_id}

	res, err := DbQuery(q, a)

	if err != nil {
		return "", fmt.Errorf("failed to get access: %s", err.Error())
	}

	for res.Next() {

		var or OrchestratorRecord_RequestKey

		err = res.Scan(&or.request_key)

		if err != nil {

			return "", fmt.Errorf("failed to register: %s", err.Error())

		}

		result_container_request_key = append(result_container_request_key, or)

	}

	if len(result_container_request_key) != 1 {
		return "", fmt.Errorf("failed to get access: %s", "duplicate")
	}

	request_key = result_container_request_key[0].request_key

	res.Close()

	return request_key, nil
}

func RegisterOsidAndRequestKey(session_id string, oauth_struct omodules.OAuthStruct) (string, error) {

	var request_key string

	var result_container_email []OrchestratorRecord_Email

	var result_container_request_key []OrchestratorRecord_RequestKey

	q := "SELECT email FROM orchestrator_record WHERE email = ?"

	a := []any{oauth_struct.EMAIL}

	res, err := DbQuery(q, a)

	if err != nil {
		return "", fmt.Errorf("failed to register: %s", err.Error())
	}

	for res.Next() {

		var or OrchestratorRecord_Email

		err = res.Scan(&or.email)

		if err != nil {

			return "", fmt.Errorf("failed to register: %s", err.Error())

		}

		result_container_email = append(result_container_email, or)

	}

	res.Close()

	if len(result_container_email) != 1 {
		return "", fmt.Errorf("failed to register: %s", "duplicate")
	}

	request_key, err = modules.RandomHex(16)

	if err != nil {
		return "", fmt.Errorf("failed to register: %s", err.Error())
	}

	q = "UPDATE orchestrator_record SET osid = ?, request_key =? WHERE email = ?"

	a = []any{session_id, request_key, oauth_struct.EMAIL}

	res, err = DbQuery(q, a)

	if err != nil {
		return "", fmt.Errorf("failed to register: %s", err.Error())
	}

	res.Close()

	q = "SELECT request_key FROM orchestrator_record WHERE email = ?"

	a = []any{oauth_struct.EMAIL}

	res, err = DbQuery(q, a)

	if err != nil {
		return "", fmt.Errorf("failed to register: %s", err.Error())
	}

	for res.Next() {

		var or OrchestratorRecord_RequestKey

		err = res.Scan(&or.request_key)

		if err != nil {

			return "", fmt.Errorf("failed to register: %s", err.Error())

		}

		result_container_request_key = append(result_container_request_key, or)

	}

	if len(result_container_request_key) != 1 {
		return "", fmt.Errorf("failed to register: %s", "duplicate")
	}

	request_key = result_container_request_key[0].request_key

	res.Close()

	return request_key, nil

}
