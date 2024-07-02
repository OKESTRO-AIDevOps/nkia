package models

import (
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	modules "github.com/OKESTRO-AIDevOps/nkia/pkg/challenge"
	"github.com/OKESTRO-AIDevOps/nkia/pkg/utils"
)

var DB *sql.DB

type OrchRecord struct {
	Email      string `json:"email"`
	Pubkey     string `json:"pubkey"`
	Osid       string `json:"osid"`
	RequestKey string `json:"request_key"`
}

type OrchClusterRecord struct {
	Email        string `json:"email"`
	ClusterId    string `json:"cluster_id"`
	Config       string `json:"config"`
	ConfigStatus string `json:"config_status"`
}

func DbQuery(query string, args []any) (*sql.Rows, error) {

	var empty_row *sql.Rows

	results, err := DB.Query(query, args[0:]...)

	if err != nil {

		return empty_row, err

	}

	return results, err

}

type OrchestratorRecord_EmailConfig struct {
	email  string
	config string
}

type OrchestratorRecord_Email struct {
	email string
}

type OrchestratorRecord_Pubkey struct {
	pubkey string
}

type OrchestratorRecord_RequestKey struct {
	request_key string
}

func GetPubkeyByEmail(email string) (string, error) {

	var ret string

	var result_container_pubkey []OrchestratorRecord_Pubkey

	q := "SELECT pubkey FROM orchestrator_record WHERE email = ?"

	a := []any{email}

	res, err := DbQuery(q, a)

	if err != nil {
		return "", fmt.Errorf("failed to get pubkey: %s", err.Error())
	}

	for res.Next() {

		var or OrchestratorRecord_Pubkey

		err = res.Scan(&or.pubkey)

		if err != nil {

			return "", fmt.Errorf("failed to get pubkey: %s", err.Error())

		}

		result_container_pubkey = append(result_container_pubkey, or)

	}

	res.Close()

	if len(result_container_pubkey) != 1 {
		return "", fmt.Errorf("failed to get pubkey: %s", "length")
	}

	ret = result_container_pubkey[0].pubkey

	return ret, nil
}

func GetKubeconfigByEmailAndClusterID(email string, cluster_id string) ([]byte, error) {

	var config []byte

	var result_container []OrchestratorRecord_EmailConfig

	query := "SELECT email, config FROM orchestrator_cluster_record WHERE email = ? AND cluster_id = ?"

	params := []any{email, cluster_id}

	res, err := DbQuery(query, params)

	if err != nil {
		return config, fmt.Errorf("failed to get config: %s", "db query error")
	}

	for res.Next() {

		var or OrchestratorRecord_EmailConfig

		err = res.Scan(&or.email, &or.config)

		if err != nil {

			fmt.Println(err.Error())

			return config, fmt.Errorf("failed to get config: %s", "records retrieval failed")

		}

		result_container = append(result_container, or)

	}

	if len(result_container) != 1 {

		fmt.Println(err.Error())

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

func OkeyEncryptor(stream []byte) ([]byte, error) {

	var ret_byte []byte

	okey_b, err := os.ReadFile("okey")

	if err != nil {
		return ret_byte, fmt.Errorf("okey failed: %s", err.Error())
	}

	enc_b, err := modules.EncryptWithSymmetricKey(okey_b, stream)

	if err != nil {
		return ret_byte, fmt.Errorf("okey failed: %s", err.Error())
	}

	return enc_b, nil

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

	res.Close()

	return email, nil
}

func UpdatePubkeyByEmail(email string, pubkey string) error {

	q := "UPDATE orchestrator_record SET pubkey = ? WHERE email = ?"

	a := []any{pubkey, email}

	res, err := DbQuery(q, a)

	if err != nil {
		return fmt.Errorf("failed to update pubkey: %s", err.Error())
	}

	res.Close()

	return nil

}

func UpdatePubkeyByEmail2(email string, pubkey string) error {

	var records = make([]OrchRecord, 0)

	updated := 0

	file_b, err := os.ReadFile(".npia/data/record.json")

	if err != nil {

		return fmt.Errorf("failed to update pubkey2: %s", err.Error())
	}

	err = json.Unmarshal(file_b, &records)

	if err != nil {

		return fmt.Errorf("failed to update pubkey2: %s", err.Error())
	}

	reclen := len(records)

	for i := 0; i < reclen; i++ {

		if records[i].Email == email {

			records[i].Pubkey = pubkey

			updated = 1

			break
		}

	}

	if updated != 1 {

		return fmt.Errorf("failed to updated pubkey2: %s", "no such user")

	}

	new_file_b, err := json.Marshal(records)

	if err != nil {

		return fmt.Errorf("failed to update pubkey2: %s", err.Error())
	}

	err = os.WriteFile(".npia/data/record.json", new_file_b, 0644)

	if err != nil {

		return fmt.Errorf("failed to update pubkey2: %s", err.Error())
	}

	return nil

}

func CreateClusterByEmail(email string, cluster_id string) (string, error) {

	var token string

	var err error

	var result_container []OrchestratorRecord_Email

	q := "SELECT email FROM orchestrator_cluster_record WHERE config_status != 'Y' AND config_status != 'ROTATE'"

	a := []any{}

	res, err := DbQuery(q, a)

	if err != nil {

		return token, fmt.Errorf("failed to create cluster: %s", err.Error())

	}

	for res.Next() {

		var or OrchestratorRecord_Email

		err = res.Scan(&or.email)

		if err != nil {

			return "", fmt.Errorf("failed to create cluster: %s", err.Error())

		}

		result_container = append(result_container, or)

	}

	if len(result_container) != 0 {
		return "", fmt.Errorf("failed to create cluster: %s", "another add in process")
	}

	res.Close()

	token, err = modules.RandomHex(16)

	if err != nil {
		return token, fmt.Errorf("failed to create cluster: %s", err.Error())
	}

	c_time := time.Now()

	c_time_fmt := c_time.Format("2006-01-02-15-04-05")

	config_chal := c_time_fmt + ":" + token

	q =
		`
	INSERT INTO

		orchestrator_cluster_record (email, cluster_id, config, config_status)
	
		VALUES(?, ?, ?, 'N')
	
	`

	a = []any{email, cluster_id, config_chal}

	res, err = DbQuery(q, a)

	if err != nil {
		return token, fmt.Errorf("failed to create cluster: %s", err.Error())
	}

	res.Close()

	return token, nil
}

func CreateClusterByEmail2(email string, clusterid string) (string, error) {

	var token string

	var records = make([]OrchClusterRecord, 0)

	var found_emails []string

	file_b, err := os.ReadFile(".npia/data/cluster_record.json")

	if err != nil {

		return token, fmt.Errorf("failed to update pubkey2: %s", err.Error())
	}

	err = json.Unmarshal(file_b, &records)

	if err != nil {

		return token, fmt.Errorf("failed to update pubkey2: %s", err.Error())
	}

	reclen := len(records)

	for i := 0; i < reclen; i++ {

		if records[i].ConfigStatus != "Y" && records[i].ConfigStatus != "ROTATE" {

			found_emails = append(found_emails, records[i].Email)

			break

		}

	}

	if utils.CheckIfSliceContains[string](found_emails, email) {

		return token, fmt.Errorf("failed to create cluster2: %s", "creation on going")

	}

	token, err = modules.RandomHex(16)

	if err != nil {
		return token, fmt.Errorf("failed to create cluster: %s", err.Error())
	}

	c_time := time.Now()

	c_time_fmt := c_time.Format("2006-01-02-15-04-05")

	config_chal := c_time_fmt + ":" + token

	records = append(records, OrchClusterRecord{
		Email:        email,
		ClusterId:    clusterid,
		Config:       config_chal,
		ConfigStatus: "N",
	})

	new_file_b, err := json.Marshal(records)

	if err != nil {

		return token, fmt.Errorf("failed to update clutser record: %s", err.Error())
	}

	err = os.WriteFile(".npia/data/cluster_record.json", new_file_b, 0644)

	if err != nil {

		return token, fmt.Errorf("failed to update cluster record: %s", err.Error())
	}

	return token, nil
}

func GetConfigChallengeByEmailAndClusterID(email string, cluster_id string) (string, error) {

	var token string

	var result_container []OrchestratorRecord_EmailConfig

	q := "SELECT email, config FROM orchestrator_cluster_record WHERE email = ? AND cluster_id = ? AND config_status = 'N'"

	a := []any{email, cluster_id}

	res, err := DbQuery(q, a)

	if err != nil {

		return token, fmt.Errorf("failed to get config: %s", err.Error())
	}

	for res.Next() {

		var or OrchestratorRecord_EmailConfig

		err = res.Scan(&or.email, &or.config)

		if err != nil {

			return token, fmt.Errorf("failed to get config: %s", err.Error())

		}

		result_container = append(result_container, or)

	}

	if len(result_container) != 1 {

		return token, fmt.Errorf("failed to get config: %s", "length")

	}

	res.Close()

	config_chal := result_container[0].config

	tmstamp_token := strings.Split(config_chal, ":")

	t_now := time.Now()

	t, _ := time.Parse("2006-01-02-15-04-05", tmstamp_token[0])

	diff := t_now.Sub(t)

	if diff.Seconds() > 3000 {

		q =
			`
		UPDATE orchestrator_cluster_record 
		SET 
			config = 'N', config_status = 'N' 
		WHERE
			email = ?
			AND cluster_id = ? 
		`

		a = []any{email, cluster_id}

		res, err = DbQuery(q, a)

		if err != nil {
			return token, fmt.Errorf("failed to get config: %s", "reset failed")
		}

		res.Close()

		return token, fmt.Errorf("failed to get config: %s", "timeout")

	}

	token = tmstamp_token[1]

	return token, nil

}

func GetConfigChallengeByEmailAndClusterID2(email string, cluster_id string) (string, error) {

	var token string

	var records = make([]OrchClusterRecord, 0)

	var found_index int = -1

	var found_pairs []struct {
		Email  string
		Config string
	}

	file_b, err := os.ReadFile(".npia/data/cluster_record.json")

	if err != nil {

		return token, fmt.Errorf("failed to get config chal2: %s", err.Error())
	}

	err = json.Unmarshal(file_b, &records)

	if err != nil {

		return token, fmt.Errorf("failed to get config chal2: %s", err.Error())
	}

	reclen := len(records)

	for i := 0; i < reclen; i++ {

		if records[i].Email == email && records[i].ClusterId == cluster_id && records[i].ConfigStatus == "N" {

			pair := struct {
				Email  string
				Config string
			}{
				Email:  records[i].Email,
				Config: records[i].Config,
			}

			found_index = i
			found_pairs = append(found_pairs, pair)

			break
		}

	}

	if len(found_pairs) != 1 {

		return token, fmt.Errorf("failed to get config chal2 : %s", "length")
	}

	config_chal := found_pairs[0].Config

	tmstamp_token := strings.Split(config_chal, ":")

	t_now := time.Now()

	t, _ := time.Parse("2006-01-02-15-04-05", tmstamp_token[0])

	diff := t_now.Sub(t)

	if diff.Seconds() > 3000 {

		records[found_index].Config = "N"

		records[found_index].ConfigStatus = "N"

		new_file_b, err := json.Marshal(records)

		if err != nil {

			return token, fmt.Errorf("failed to get config chal2: %s", err.Error())

		}

		err = os.WriteFile(".npia/data/cluster_record.json", new_file_b, 0644)

		if err != nil {

			return token, fmt.Errorf("failed to get config chal2: %s", err.Error())
		}

		return token, fmt.Errorf("failed to get config chal2: %s", "timeout")

	}

	token = tmstamp_token[1]

	return token, nil

}

func AddClusterByEmailAndClusterID(email string, cluster_id string, config string) error {

	var result_container []OrchestratorRecord_Email

	q := "SELECT email FROM orchestrator_cluster_record WHERE email = ? AND cluster_id = ? AND config_status = 'N'"

	a := []any{email, cluster_id}

	res, err := DbQuery(q, a)

	if err != nil {

		return fmt.Errorf("failed to add cluster: %s", err.Error())
	}

	for res.Next() {

		var or OrchestratorRecord_Email

		err = res.Scan(&or.email)

		if err != nil {

			return fmt.Errorf("failed to add cluster: %s", err.Error())

		}

		result_container = append(result_container, or)

	}

	if len(result_container) != 1 {

		return fmt.Errorf("failed to add cluster: %s", "length")

	}

	res.Close()

	config_enc_b, err := OkeyEncryptor([]byte(config))

	if err != nil {
		return fmt.Errorf("failed to add cluster: %s", err.Error())
	}

	enc_hex := hex.EncodeToString(config_enc_b)

	q =
		`
	UPDATE orchestrator_cluster_record
	SET config = ?, config_status = 'Y'
	WHERE 
		email = ?
		AND cluster_id = ? 
	`

	a = []any{enc_hex, email, cluster_id}

	res, err = DbQuery(q, a)

	if err != nil {

		return fmt.Errorf("failed to add cluster: %s", err.Error())

	}

	res.Close()

	return nil
}

func AddClusterByEmailAndClusterID2(email string, cluster_id string, config string) error {

	var records = make([]OrchClusterRecord, 0)

	var found_index int = -1

	file_b, err := os.ReadFile(".npia/data/cluster_record.json")

	if err != nil {

		return fmt.Errorf("failed to add cluster2: %s", err.Error())
	}

	err = json.Unmarshal(file_b, &records)

	if err != nil {

		return fmt.Errorf("failed to add cluster2: %s", err.Error())
	}

	reclen := len(records)

	for i := 0; i < reclen; i++ {

		if records[i].Email == email && records[i].ClusterId == cluster_id && records[i].ConfigStatus == "N" {

			found_index = i

			break
		}

	}

	records[found_index].Config = config
	records[found_index].ConfigStatus = "Y"

	new_file_b, err := json.Marshal(records)

	if err != nil {

		return fmt.Errorf("failed to add cluster2: %s", err.Error())

	}

	err = os.WriteFile(".npia/data/cluster_record.json", new_file_b, 0644)

	if err != nil {

		return fmt.Errorf("failed to add cluster2: %s", err.Error())
	}

	return nil
}
