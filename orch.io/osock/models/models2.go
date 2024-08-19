package models

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	modules "github.com/OKESTRO-AIDevOps/nkia/pkg/challenge"
	"github.com/OKESTRO-AIDevOps/nkia/pkg/utils"
)

func GetKubeconfigByEmailAndClusterID2(email string, cluster_id string) ([]byte, error) {

	var records = make([]OrchClusterRecord, 0)

	var found_pairs []struct {
		Email  string
		Config string
	}

	file_b, err := os.ReadFile(".npia/data/cluster_record.json")

	if err != nil {

		return nil, fmt.Errorf("failed to get config: %s", err.Error())
	}

	err = json.Unmarshal(file_b, &records)

	if err != nil {

		return nil, fmt.Errorf("failed to get config: %s", err.Error())
	}

	reclen := len(records)

	for i := 0; i < reclen; i++ {

		if records[i].Email == email && records[i].ClusterId == cluster_id && records[i].ConfigStatus == "Y" {

			pair := struct {
				Email  string
				Config string
			}{
				Email:  records[i].Email,
				Config: records[i].Config,
			}

			found_pairs = append(found_pairs, pair)

			break
		}

	}

	if len(found_pairs) != 1 {

		return nil, fmt.Errorf("failed to get config: %s", "length")
	}

	config_b := []byte(found_pairs[0].Config)

	return config_b, nil
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
