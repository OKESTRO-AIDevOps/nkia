package main

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/OKESTRO-AIDevOps/nkia/src/modules"
)

var DB *sql.DB

func DbQuery(query string, args []any) (*sql.Rows, error) {

	var empty_row *sql.Rows

	results, err := DB.Query(query, args[0:]...)

	if err != nil {

		return empty_row, err

	}

	return results, err

}

var L = log.New(os.Stdout, "", 0)

func EventLogger(msg string) {
	L.SetPrefix(time.Now().UTC().Format("2006-01-02 15:04:05.000") + " [INFO] ")
	L.Print(msg)
}

type OrchestratorRecord_EmailConfig struct {
	email  string
	config string
}

type OrchestratorRecord_Email struct {
	email string
}

func GetKubeconfigByEmail(email string) ([]byte, error) {

	var config []byte

	var result_container []OrchestratorRecord_EmailConfig

	query := "SELECT email, config FROM orchestrator_record WHERE email = ?"

	params := []any{email}

	res, err := DbQuery(query, params)

	if err != nil {
		return config, fmt.Errorf("failed to get config: %s", "db query error")
	}

	for res.Next() {

		var or OrchestratorRecord_EmailConfig

		err = res.Scan(&or.email, &or.config)

		if err != nil {

			EventLogger(err.Error())

			return config, fmt.Errorf("failed to get config: %s", "records retrieval failed")

		}

		result_container = append(result_container, or)

	}

	if len(result_container) != 1 {

		EventLogger("auth failed")

		return config, fmt.Errorf("failed to get config: %s", "false count")

	}

	config_enc := result_container[0].config

	config_enc_b, err := hex.DecodeString(config_enc)

	config, err = OkeyDecryptor(config_enc_b)

	if err != nil {
		return config, fmt.Errorf("failed to get config: %s", err.Error())
	}

	return config, nil
}

func OkeyDecryptor(stream []byte) ([]byte, error) {

	var ret_byte []byte

	okey_b, err := os.ReadFile("okey")

	if err != nil {
		return ret_byte, fmt.Errorf("okey failed: %s", err.Error())
	}

	dec_b, err := modules.DecryptWithSymmetricKey(okey_b, stream)

	if err != nil {
		return ret_byte, fmt.Errorf("okey failed: %s", err.Error())
	}

	return dec_b, nil

}

func CheckSessionAndGetEmailByRequestKey(request_key string) (string, error) {

	var email string

	var result_container []OrchestratorRecord_Email

	q := "SELECT email FROM orchestrator_record WHERE osid != 'N' AND request_key = ?"

	a := []any{request_key}

	res, err := DbQuery(q, a)

	if err != nil {

		return "", fmt.Errorf("failed to check session: %s", err.Error())

	}

	for res.Next() {

		var or OrchestratorRecord_Email

		err = res.Scan(&or.email)

		if err != nil {

			return "", fmt.Errorf("failed to check session: %s", err.Error())

		}

		result_container = append(result_container, or)

	}

	if len(result_container) != 1 {
		return "", fmt.Errorf("failed to check session: %s", "wrong length")
	}

	email = result_container[0].email

	return email, nil
}
