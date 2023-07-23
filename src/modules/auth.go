package modules

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	goya "github.com/goccy/go-yaml"
)

func AccessAuth(c *gin.Context) (string, error) {

	var session_sym_key string

	var key_records KeyRecord

	session := sessions.Default(c)

	var session_id string

	v := session.Get("SID")

	if v == nil {
		return "", fmt.Errorf("access auth failed: %s", "session id not found")
	} else {
		session_id = v.(string)
	}

	file_byte, err := os.ReadFile("srv/key.json")

	if err != nil {
		return "", fmt.Errorf("access auth failed: %s", err.Error())
	}

	err = json.Unmarshal(file_byte, &key_records)

	if err != nil {
		return "", fmt.Errorf("access auth failed: %s", err.Error())
	}

	ssk, okay := key_records[session_id]

	if !okay {
		return "", fmt.Errorf("access auth failed: %s", "session not found")
	}

	session_sym_key = ssk

	return session_sym_key, nil

}

func AccessAuth_Detached(key_id string) (string, error) {

	var session_sym_key string

	var key_records KeyRecord

	file_byte, err := os.ReadFile("srv/key.json")

	if err != nil {
		return "", fmt.Errorf("access auth failed: %s", err.Error())
	}

	err = json.Unmarshal(file_byte, &key_records)

	if err != nil {
		return "", fmt.Errorf("access auth failed: %s", err.Error())
	}

	ssk, okay := key_records[key_id]

	if !okay {
		return "", fmt.Errorf("access auth failed: %s", "session not found")
	}

	session_sym_key = ssk

	return session_sym_key, nil

}

func GenerateChallenge(client_ca_pub_key ChallengRecord) (ChallengRecord, error) {

	var kube_config map[interface{}]interface{}

	server_contexts := make([]string, 0)

	challenge_records := make(ChallengRecord)

	new_challenge_records := make(ChallengRecord)

	context_challenges := make(map[string]string)

	context_challenges_cipher := make(map[string]string)

	client_context_map, okay := client_ca_pub_key["ask_challenge"]

	client_contexts := make([]string, 0)

	if !okay {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", "no ask_challenge key")
	}

	kube_config_path, err := GetKubeConfigPath()

	if err != nil {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", err.Error())
	}

	kube_config_file_byte, err := os.ReadFile(kube_config_path)

	err = goya.Unmarshal(kube_config_file_byte, &kube_config)

	if err != nil {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", err.Error())
	}

	client_context_length := len(client_context_map)

	contexts_len := len(kube_config["contexts"].([]interface{}))

	if contexts_len != client_context_length {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", "invalid formate: 01")
	}

	for i := 0; i < contexts_len; i++ {

		context_nm := kube_config["contexts"].([]interface{})[i].(map[string]interface{})["name"].(string)

		server_contexts = append(server_contexts, context_nm)

	}

	for cc := range client_context_map {

		exists := CheckIfSliceContains[string](client_contexts, cc)

		if exists {
			return challenge_records, fmt.Errorf("failed to generate challenge: %s", "invalid format: 02")
		}

		client_contexts = append(client_contexts, cc)

	}

	for i := 0; i < len(client_contexts); i++ {

		exists := CheckIfSliceContains[string](server_contexts, client_contexts[i])

		if !exists {
			return challenge_records, fmt.Errorf("failed to generate challenge: %s", "invalid format: 03")
		}

	}

	for context, pub_str := range client_context_map {

		context_user_cert_b, err := GetContextUserCertificateBytes(context)

		if err != nil {
			return challenge_records, fmt.Errorf("failed to generate challenge: %s", err.Error())
		}

		block, _ := pem.Decode(context_user_cert_b)

		var cert *x509.Certificate

		cert, err = x509.ParseCertificate(block.Bytes)

		hash_sha := sha256.New()

		hash_sha.Write(cert.RawTBSCertificate)

		hash_data := hash_sha.Sum(nil)

		pub_key, err := BytesToPublicKey([]byte(pub_str))

		if err != nil {
			return challenge_records, fmt.Errorf("failed to generate challenge: %s", "invalid format: 04")
		}

		err = rsa.VerifyPKCS1v15(pub_key, crypto.SHA256, hash_data, cert.Signature)

		if err != nil {
			return challenge_records, fmt.Errorf("failed to generate challenge: %s", "invalid pub key")
		}
	}

	file_byte, err := os.ReadFile("srv/challenge.json")

	if err != nil {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", err.Error())
	}

	err = json.Unmarshal(file_byte, &challenge_records)

	if err != nil {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", err.Error())
	}

	new_challenge_id, _ := RandomHex(8)

	_, okay = challenge_records[new_challenge_id]

	if okay {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", "duplicate challenge id")
	}

	for i := 0; i < contexts_len; i++ {

		context_nm := kube_config["contexts"].([]interface{})[i].(map[string]interface{})["name"].(string)

		pubkey_b, err := GetContextUserPublicKeyBytes(context_nm)

		if err != nil {
			return challenge_records, fmt.Errorf("failed to generate challenge: %s", err.Error())
		}

		pubkey, err := BytesToPublicKey(pubkey_b)

		if err != nil {
			return challenge_records, fmt.Errorf("failed to generate challenge: %s", err.Error())
		}

		new_challenge_val, _ := RandomHex(32)

		context_challenges[context_nm] = new_challenge_val

		chal_val_enc, err := EncryptWithPublicKey([]byte(new_challenge_val), pubkey)

		context_challenges_cipher[context_nm] = hex.EncodeToString(chal_val_enc)

	}

	challenge_records[new_challenge_id] = context_challenges

	new_challenge_records[new_challenge_id] = context_challenges_cipher

	challenge_records_byte, err := json.Marshal(challenge_records)

	if err != nil {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", err.Error())
	}

	err = os.WriteFile("srv/challenge.json", challenge_records_byte, 0644)

	return new_challenge_records, nil

}

func VerifyChallange(answer ChallengRecord) (string, KeyRecord, error) {

	answer_key := ""

	answer_map := make(map[string]string)

	answer_contexts := make([]string, 0)

	challenge_records := make(ChallengRecord)

	key_records := make(KeyRecord)

	new_key_record := make(KeyRecord)

	challenge_file_byte, err := os.ReadFile("srv/challenge.json")

	if err != nil {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", err.Error())
	}

	err = json.Unmarshal(challenge_file_byte, &challenge_records)

	if err != nil {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", err.Error())
	}

	key_file_byte, err := os.ReadFile("srv/key.json")

	if err != nil {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", err.Error())
	}

	err = json.Unmarshal(key_file_byte, &key_records)

	if err != nil {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", err.Error())
	}

	if len(answer) != 1 {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", "invalid format: 01")
	}

	for ak := range answer {
		answer_key = ak
	}

	challenge_map, okay := challenge_records[answer_key]

	if !okay {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", "invalid key")
	}

	answer_map, okay = answer[answer_key]

	if !okay {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", "invalid format: 02")
	}

	for am_k := range answer_map {

		exists := CheckIfSliceContains[string](answer_contexts, am_k)

		if exists {
			return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", "invalid format: 03")
		}

		answer_contexts = append(answer_contexts, am_k)

	}

	challenge_map_length := len(challenge_map)

	answer_contexts_length := len(answer_contexts)

	if challenge_map_length != answer_contexts_length {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", "invalid format: 04")
	}

	for i := 0; i < answer_contexts_length; i++ {

		ans, okay := challenge_map[answer_contexts[i]]

		if !okay {
			return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", "invalid format: 05")
		}

		if ans != answer_map[answer_contexts[i]] {
			return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", "wrong value")
		}

	}

	gen_key, err := RandomHex(32)

	if err != nil {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", "gen key failed")
	}

	new_sym_key, err := RandomHex(16)

	if err != nil {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", "gen key failed")
	}

	key_records[gen_key] = new_sym_key

	key_records_byte, err := json.Marshal(key_records)

	if err != nil {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", "key write failed")
	}

	err = os.WriteFile("srv/key.json", key_records_byte, 0644)

	if err != nil {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", "key write failed")
	}

	cidx := GetRandIntInRange(0, answer_contexts_length-1)

	rand_context := answer_contexts[cidx]

	pub_b, err := GetContextUserPublicKeyBytes(rand_context)

	if err != nil {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", "key write failed")
	}

	pubkey, err := BytesToPublicKey(pub_b)

	if err != nil {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", "key write failed")
	}

	new_sym_key_enc, err := EncryptWithPublicKey([]byte(new_sym_key), pubkey)

	if err != nil {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", "key write failed")
	}

	new_key_record[rand_context] = hex.EncodeToString(new_sym_key_enc)

	return gen_key, new_key_record, nil

}

func GenerateChallenge_Detached(config_b []byte, client_ca_pub_key ChallengRecord) (ChallengRecord, error) {

	var kube_config map[interface{}]interface{}

	detached_contexts := make([]string, 0)

	challenge_records := make(ChallengRecord)

	new_challenge_records := make(ChallengRecord)

	context_challenges := make(map[string]string)

	context_challenges_cipher := make(map[string]string)

	client_context_map, okay := client_ca_pub_key["ask_challenge"]

	client_context := ""

	if !okay {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", "no ask_challenge key")
	}

	err := goya.Unmarshal(config_b, &kube_config)

	if err != nil {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", err.Error())
	}

	client_context_length := len(client_context_map)

	o_contexts_len := len(kube_config["contexts"].([]interface{}))

	if client_context_length != 1 {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", "invalid formate: 01")
	}

	for i := 0; i < o_contexts_len; i++ {

		context_nm := kube_config["contexts"].([]interface{})[i].(map[string]interface{})["name"].(string)

		detached_contexts = append(detached_contexts, context_nm)

	}

	for cc := range client_context_map {

		exists := CheckIfSliceContains[string](detached_contexts, cc)

		if !exists {
			return challenge_records, fmt.Errorf("failed to generate challenge: %s", "invalid formate: 02")
		}

		client_context = cc

	}

	pub_str := client_context_map[client_context]

	context_user_cert_b, err := GetContextUserCertificateBytes_Detached(config_b, client_context)

	if err != nil {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", err.Error())
	}

	block, _ := pem.Decode(context_user_cert_b)

	var cert *x509.Certificate

	cert, err = x509.ParseCertificate(block.Bytes)

	hash_sha := sha256.New()

	hash_sha.Write(cert.RawTBSCertificate)

	hash_data := hash_sha.Sum(nil)

	pub_key, err := BytesToPublicKey([]byte(pub_str))

	if err != nil {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", "invalid format: 04")
	}

	err = rsa.VerifyPKCS1v15(pub_key, crypto.SHA256, hash_data, cert.Signature)

	if err != nil {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", "invalid pub key")
	}

	file_byte, err := os.ReadFile("srv/challenge.json")

	if err != nil {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", err.Error())
	}

	err = json.Unmarshal(file_byte, &challenge_records)

	if err != nil {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", err.Error())
	}

	new_challenge_id, _ := RandomHex(8)

	_, okay = challenge_records[new_challenge_id]

	if okay {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", "duplicate challenge id")
	}

	pubkey_b, err := GetContextUserPublicKeyBytes_Detached(config_b, client_context)

	if err != nil {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", err.Error())
	}

	pubkey, err := BytesToPublicKey(pubkey_b)

	if err != nil {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", err.Error())
	}

	new_challenge_val, _ := RandomHex(32)

	context_challenges[client_context] = new_challenge_val

	chal_val_enc, err := EncryptWithPublicKey([]byte(new_challenge_val), pubkey)

	context_challenges_cipher[client_context] = hex.EncodeToString(chal_val_enc)

	challenge_records[new_challenge_id] = context_challenges

	new_challenge_records[new_challenge_id] = context_challenges_cipher

	challenge_records_byte, err := json.Marshal(challenge_records)

	if err != nil {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", err.Error())
	}

	err = os.WriteFile("srv/challenge.json", challenge_records_byte, 0644)

	return new_challenge_records, nil
}

func VerifyChallange_Detached(config_b []byte, answer ChallengRecord) (string, KeyRecord, error) {

	answer_key := ""

	answer_map := make(map[string]string)

	answer_contexts := make([]string, 0)

	challenge_records := make(ChallengRecord)

	key_records := make(KeyRecord)

	new_key_record := make(KeyRecord)

	challenge_file_byte, err := os.ReadFile("srv/challenge.json")

	if err != nil {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", err.Error())
	}

	err = json.Unmarshal(challenge_file_byte, &challenge_records)

	if err != nil {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", err.Error())
	}

	key_file_byte, err := os.ReadFile("srv/key.json")

	if err != nil {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", err.Error())
	}

	err = json.Unmarshal(key_file_byte, &key_records)

	if err != nil {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", err.Error())
	}

	if len(answer) != 1 {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", "invalid format: 01")
	}

	for ak := range answer {
		answer_key = ak
	}

	challenge_map, okay := challenge_records[answer_key]

	if !okay {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", "invalid key")
	}

	answer_map, okay = answer[answer_key]

	if !okay {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", "invalid format: 02")
	}

	for am_k := range answer_map {

		exists := CheckIfSliceContains[string](answer_contexts, am_k)

		if exists {
			return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", "invalid format: 03")
		}

		answer_contexts = append(answer_contexts, am_k)

	}

	// challenge_map_length := len(challenge_map)

	answer_contexts_length := len(answer_contexts)

	if answer_contexts_length != 1 {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", "invalid format: 04")
	}

	ans, okay := challenge_map[answer_contexts[0]]

	if !okay {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", "invalid format: 05")
	}

	if ans != answer_map[answer_contexts[0]] {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", "wrong value")
	}

	gen_key, err := RandomHex(32)

	if err != nil {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", "gen key failed")
	}

	new_sym_key, err := RandomHex(16)

	if err != nil {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", "gen key failed")
	}

	key_records[gen_key] = new_sym_key

	key_records_byte, err := json.Marshal(key_records)

	if err != nil {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", "key write failed")
	}

	err = os.WriteFile("srv/key.json", key_records_byte, 0644)

	if err != nil {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", "key write failed")
	}

	matching_context := answer_contexts[0]

	pub_b, err := GetContextUserPublicKeyBytes_Detached(config_b, matching_context)

	if err != nil {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", "key write failed")
	}

	pubkey, err := BytesToPublicKey(pub_b)

	if err != nil {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", "key write failed")
	}

	new_sym_key_enc, err := EncryptWithPublicKey([]byte(new_sym_key), pubkey)

	if err != nil {
		return "", new_key_record, fmt.Errorf("failed to verify challenge: %s", "key write failed")
	}

	new_key_record[matching_context] = hex.EncodeToString(new_sym_key_enc)

	return gen_key, new_key_record, nil

}
