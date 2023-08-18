package main

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	goya "github.com/goccy/go-yaml"

	"github.com/OKESTRO-AIDevOps/nkia/src/modules"
	"github.com/OKESTRO-AIDevOps/npia-api/pkg/runtimefs"
	kalfs "github.com/OKESTRO-AIDevOps/npia-multi-mode/pkg/kaleidofs"
)

type ChallengRecord map[string]map[string]string

type KeyRecord map[string]string

func GetKubeConfigPathSimple() {

	cmd := exec.Command("../srv/get_kubeconfig_path")

	out, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return

	}

	fmt.Println(string(out))
}

func RandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func GetKubeConfigPath() (string, error) {

	var kube_config_path string

	cmd := exec.Command("../srv/get_kubeconfig_path")

	out, err := cmd.Output()

	if err != nil {

		return "", fmt.Errorf("failed to get kube config path: %s", err.Error())

	}

	strout := string(out)

	ret_strout := strings.ReplaceAll(strout, "\n", "")

	ret_strout = strings.ReplaceAll(ret_strout, " ", "")

	kube_config_path = ret_strout

	return kube_config_path, nil
}

func LoadTest() (ChallengRecord, error) {

	var kube_config map[interface{}]interface{}

	challenge_records := make(ChallengRecord)

	new_challenge_records := make(ChallengRecord)

	context_challenges := make(map[string]string)

	file_byte, err := os.ReadFile("../srv/challenge.json")

	if err != nil {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", err.Error())
	}

	err = json.Unmarshal(file_byte, &challenge_records)

	if err != nil {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", err.Error())
	}

	new_challenge_id, _ := RandomHex(32)

	_, okay := challenge_records[new_challenge_id]

	if okay {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", "duplicate challenge id")
	}

	kube_config_path, err := GetKubeConfigPath()

	if err != nil {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", err.Error())
	}

	kube_config_file_byte, err := os.ReadFile(kube_config_path)

	err = goya.Unmarshal(kube_config_file_byte, &kube_config)

	fmt.Println(kube_config)

	if err != nil {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", err.Error())
	}

	contexts_len := len(kube_config["contexts"].([]interface{}))

	for i := 0; i < contexts_len; i++ {

		context_nm := kube_config["contexts"].([]interface{})[i].(map[string]interface{})["name"].(string)

		new_challenge_val, _ := RandomHex(256)

		context_challenges[context_nm] = new_challenge_val

	}

	challenge_records[new_challenge_id] = context_challenges

	new_challenge_records[new_challenge_id] = context_challenges

	challenge_records_byte, err := json.Marshal(challenge_records)

	if err != nil {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", err.Error())
	}

	err = os.WriteFile("../srv/challenge.json", challenge_records_byte, 0644)

	if err != nil {
		return challenge_records, fmt.Errorf("failed to generate challenge: %s", err.Error())
	}

	return new_challenge_records, nil

}

func GetContextUserPrivateKeyBytes(context_nm string) ([]byte, error) {

	var kube_config map[interface{}]interface{}

	var ret_byte []byte

	kube_config_path, err := GetKubeConfigPath()

	if err != nil {
		return ret_byte, fmt.Errorf("failed to get context user private key: %s", err.Error())
	}

	kube_config_file_byte, err := os.ReadFile(kube_config_path)

	err = goya.Unmarshal(kube_config_file_byte, &kube_config)

	if err != nil {
		return ret_byte, fmt.Errorf("failed to get context user private key: %s", err.Error())
	}

	contexts_len := len(kube_config["contexts"].([]interface{}))

	context_user_nm := ""

	for i := 0; i < contexts_len; i++ {

		if kube_config["contexts"].([]interface{})[i].(map[string]interface{})["name"].(string) == context_nm {

			context_user_nm = kube_config["contexts"].([]interface{})[i].(map[string]interface{})["context"].(map[string]interface{})["user"].(string)

			break
		}

	}

	if context_user_nm == "" {
		return ret_byte, fmt.Errorf("failed to get context user private key: %s", "matching user not found")
	}

	user_len := len(kube_config["users"].([]interface{}))

	var user_priv_key_data []byte

	for i := 0; i < user_len; i++ {

		if kube_config["users"].([]interface{})[i].(map[string]interface{})["name"].(string) == context_user_nm {

			tmp_base64, okay := kube_config["users"].([]interface{})[i].(map[string]interface{})["user"].(map[string]interface{})["client-key-data"].(string)

			if !okay {
				return ret_byte, fmt.Errorf("failed to get context user private key: %s", "no key data")
			}

			dec_base64, err := base64.StdEncoding.DecodeString(tmp_base64)

			user_priv_key_data = dec_base64

			if err != nil {
				return ret_byte, fmt.Errorf("failed to get context user private key: %s", err.Error())
			}

			break
		}

	}

	ret_byte = user_priv_key_data

	return ret_byte, nil
}

func GetContextUserPublicKeyBytes(context_nm string) ([]byte, error) {

	var kube_config map[interface{}]interface{}

	var ret_byte []byte

	kube_config_path, err := GetKubeConfigPath()

	if err != nil {
		return ret_byte, fmt.Errorf("failed to get context user public key: %s", err.Error())
	}

	kube_config_file_byte, err := os.ReadFile(kube_config_path)

	err = goya.Unmarshal(kube_config_file_byte, &kube_config)

	if err != nil {
		return ret_byte, fmt.Errorf("failed to get context user public key: %s", err.Error())
	}

	contexts_len := len(kube_config["contexts"].([]interface{}))

	context_user_nm := ""

	for i := 0; i < contexts_len; i++ {

		if kube_config["contexts"].([]interface{})[i].(map[string]interface{})["name"].(string) == context_nm {

			context_user_nm = kube_config["contexts"].([]interface{})[i].(map[string]interface{})["context"].(map[string]interface{})["user"].(string)

			break
		}

	}

	if context_user_nm == "" {
		return ret_byte, fmt.Errorf("failed to get context user public key: %s", "matching user not found")
	}

	user_len := len(kube_config["users"].([]interface{}))

	var user_pub_key_data []byte

	user_certificate_data := ""

	for i := 0; i < user_len; i++ {

		if kube_config["users"].([]interface{})[i].(map[string]interface{})["name"].(string) == context_user_nm {

			tmp_base64, okay := kube_config["users"].([]interface{})[i].(map[string]interface{})["user"].(map[string]interface{})["client-certificate-data"].(string)

			if !okay {
				return ret_byte, fmt.Errorf("failed to get context user public key: %s", "no key data")
			}

			dec_base64, err := base64.StdEncoding.DecodeString(tmp_base64)

			user_certificate_data = string(dec_base64)

			if err != nil {
				return ret_byte, fmt.Errorf("failed to get context user public key: %s", err.Error())
			}

			break
		}

	}

	if user_certificate_data == "" {
		return ret_byte, fmt.Errorf("failed to get context user public key: %s", "no matching user key")
	}

	block, _ := pem.Decode([]byte(user_certificate_data))
	var cert *x509.Certificate
	cert, err = x509.ParseCertificate(block.Bytes)

	if err != nil {
		return ret_byte, fmt.Errorf("failed to get context user public key: %s", err.Error())
	}

	rsaPublicKey := cert.PublicKey.(*rsa.PublicKey)

	user_pub_key_data, err = PublicKeyToBytes(rsaPublicKey)

	if err != nil {
		return ret_byte, fmt.Errorf("failed to get context user public key: %s", err.Error())
	}

	ret_byte = user_pub_key_data

	return ret_byte, nil

}

func GenerateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {

	var pubkey *rsa.PublicKey

	privkey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return privkey, pubkey, fmt.Errorf("failed to gen key pair: %s", err.Error())
	}

	pubkey = &privkey.PublicKey

	return privkey, pubkey, nil
}

func PrivateKeyToBytes(priv *rsa.PrivateKey) ([]byte, error) {
	var ret_byte []byte
	privBytes := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(priv),
		},
	)

	if privBytes == nil {
		return ret_byte, fmt.Errorf("failed to encode priv key to bytes: %s", "invalid")
	}

	return privBytes, nil
}

func PublicKeyToBytes(pub *rsa.PublicKey) ([]byte, error) {

	var ret_byte []byte

	pubASN1, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return pubASN1, fmt.Errorf("failed to encode pub key to bytes: %s", err.Error())
	}

	pubBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})

	if pubBytes == nil {
		return ret_byte, fmt.Errorf("failed to encode priv key to bytes: %s", "invalid")
	}

	return pubBytes, nil
}

func BytesToPrivateKey(priv []byte) (*rsa.PrivateKey, error) {

	var privkey *rsa.PrivateKey

	block, _ := pem.Decode(priv)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		fmt.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return privkey, fmt.Errorf("failed to decode bytes to priv key: %s", err.Error())
		}
	}
	key, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		return privkey, fmt.Errorf("failed to decode bytes to priv key: %s", err.Error())
	}
	return key, nil
}

func BytesToPublicKey(pub []byte) (*rsa.PublicKey, error) {

	var pubkey *rsa.PublicKey

	block, _ := pem.Decode(pub)
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	var err error
	if enc {
		fmt.Println("is encrypted pem block")
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return pubkey, fmt.Errorf("failed to decode bytes to pub key: %s", err.Error())
		}
	}
	ifc, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		return pubkey, fmt.Errorf("failed to decode bytes to pub key: %s", err.Error())
	}
	key, ok := ifc.(*rsa.PublicKey)
	if !ok {
		return pubkey, fmt.Errorf("failed to decode bytes to pub key: %s", err.Error())
	}
	return key, nil
}

func EncryptWithPublicKey(msg []byte, pub *rsa.PublicKey) ([]byte, error) {
	hash := sha512.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, pub, msg, nil)
	if err != nil {
		return ciphertext, fmt.Errorf("failed to encrypt with public key: %s", err.Error())
	}
	return ciphertext, nil
}

func DecryptWithPrivateKey(ciphertext []byte, priv *rsa.PrivateKey) ([]byte, error) {
	hash := sha512.New()
	plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, priv, ciphertext, nil)
	if err != nil {
		return plaintext, fmt.Errorf("failed to decrypt with private key: %s", err.Error())
	}
	return plaintext, nil
}

func EncryptWithSymmetricKey(key []byte, file_byte []byte) ([]byte, error) {

	var ret_byte []byte

	c, err := aes.NewCipher(key)

	if err != nil {
		return ret_byte, fmt.Errorf("failed to encrypt with symmetric key: %s", err.Error())

	}

	gcm, err := cipher.NewGCM(c)

	if err != nil {
		return ret_byte, fmt.Errorf("failed to encrypt with symmetric key: %s", err.Error())

	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return ret_byte, fmt.Errorf("failed to encrypt with symmetric key: %s", err.Error())
	}

	enc_file := gcm.Seal(nonce, nonce, file_byte, nil)

	ret_byte = enc_file

	return ret_byte, nil

}

func DecryptWithSymmetricKey(key []byte, file_byte []byte) ([]byte, error) {

	var ret_byte []byte

	c, err := aes.NewCipher(key)
	if err != nil {
		return ret_byte, fmt.Errorf("failed to decrypt with symmetric key: %s", err.Error())
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return ret_byte, fmt.Errorf("failed to decrypt with symmetric key: %s", err.Error())
	}

	nonceSize := gcm.NonceSize()
	if len(file_byte) < nonceSize {
		return ret_byte, fmt.Errorf("failed to decrypt with symmetric key: %s", err.Error())
	}

	nonce, ciphertext := file_byte[:nonceSize], file_byte[nonceSize:]
	plain_file, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return ret_byte, fmt.Errorf("failed to decrypt with symmetric key: %s", err.Error())
	}

	ret_byte = plain_file

	return ret_byte, nil
}

func GetContextClusterPublicKeyBytes(context_nm string) ([]byte, error) {

	var kube_config map[interface{}]interface{}

	var ret_byte []byte

	kube_config_path, err := GetKubeConfigPath()

	if err != nil {
		return ret_byte, fmt.Errorf("failed to get context cluster public key: %s", err.Error())
	}

	kube_config_file_byte, err := os.ReadFile(kube_config_path)

	err = goya.Unmarshal(kube_config_file_byte, &kube_config)

	if err != nil {
		return ret_byte, fmt.Errorf("failed to get context cluster public key: %s", err.Error())
	}

	contexts_len := len(kube_config["contexts"].([]interface{}))

	context_cluster_nm := ""

	for i := 0; i < contexts_len; i++ {

		if kube_config["contexts"].([]interface{})[i].(map[string]interface{})["name"].(string) == context_nm {

			context_cluster_nm = kube_config["contexts"].([]interface{})[i].(map[string]interface{})["context"].(map[string]interface{})["cluster"].(string)

			break
		}

	}

	if context_cluster_nm == "" {
		return ret_byte, fmt.Errorf("failed to get context user public key: %s", "matching cluster not found")
	}

	clusters_len := len(kube_config["clusters"].([]interface{}))

	var cluster_pub_key_data []byte

	cluster_certificate_data := ""

	for i := 0; i < clusters_len; i++ {

		if kube_config["clusters"].([]interface{})[i].(map[string]interface{})["name"].(string) == context_cluster_nm {

			tmp_base64, okay := kube_config["clusters"].([]interface{})[i].(map[string]interface{})["cluster"].(map[string]interface{})["certificate-authority-data"].(string)

			if !okay {
				return ret_byte, fmt.Errorf("failed to get context cluster public key: %s", "no key data")
			}

			dec_base64, err := base64.StdEncoding.DecodeString(tmp_base64)

			cluster_certificate_data = string(dec_base64)

			if err != nil {
				return ret_byte, fmt.Errorf("failed to get context cluster public key: %s", err.Error())
			}

			break
		}

	}

	if cluster_certificate_data == "" {
		return ret_byte, fmt.Errorf("failed to get context user public key: %s", "no matching user key")
	}

	block, _ := pem.Decode([]byte(cluster_certificate_data))
	var cert *x509.Certificate
	cert, err = x509.ParseCertificate(block.Bytes)

	if err != nil {
		return ret_byte, fmt.Errorf("failed to get context cluster public key: %s", err.Error())
	}

	rsaPublicKey := cert.PublicKey.(*rsa.PublicKey)

	cluster_pub_key_data, err = PublicKeyToBytes(rsaPublicKey)

	if err != nil {
		return ret_byte, fmt.Errorf("failed to get context cluster public key: %s", err.Error())
	}

	ret_byte = cluster_pub_key_data

	return ret_byte, nil

}

func GetContextUserCertificateBytes(context_nm string) ([]byte, error) {

	var kube_config map[interface{}]interface{}

	var ret_byte []byte

	kube_config_path, err := GetKubeConfigPath()

	if err != nil {
		return ret_byte, fmt.Errorf("failed to get context user certificate: %s", err.Error())
	}

	kube_config_file_byte, err := os.ReadFile(kube_config_path)

	err = goya.Unmarshal(kube_config_file_byte, &kube_config)

	if err != nil {
		return ret_byte, fmt.Errorf("failed to get context user certificate: %s", err.Error())
	}

	contexts_len := len(kube_config["contexts"].([]interface{}))

	context_user_nm := ""

	for i := 0; i < contexts_len; i++ {

		if kube_config["contexts"].([]interface{})[i].(map[string]interface{})["name"].(string) == context_nm {

			context_user_nm = kube_config["contexts"].([]interface{})[i].(map[string]interface{})["context"].(map[string]interface{})["user"].(string)

			break
		}

	}

	if context_user_nm == "" {
		return ret_byte, fmt.Errorf("failed to get context user certificate: %s", "matching user not found")
	}

	user_len := len(kube_config["users"].([]interface{}))

	user_certificate_data := ""

	for i := 0; i < user_len; i++ {

		if kube_config["users"].([]interface{})[i].(map[string]interface{})["name"].(string) == context_user_nm {

			tmp_base64, okay := kube_config["users"].([]interface{})[i].(map[string]interface{})["user"].(map[string]interface{})["client-certificate-data"].(string)

			if !okay {
				return ret_byte, fmt.Errorf("failed to get context user certificate: %s", "no key data")
			}

			dec_base64, err := base64.StdEncoding.DecodeString(tmp_base64)

			user_certificate_data = string(dec_base64)

			if err != nil {
				return ret_byte, fmt.Errorf("failed to get context user certificate: %s", err.Error())
			}

			break
		}

	}

	if user_certificate_data == "" {
		return ret_byte, fmt.Errorf("failed to get context user certificate: %s", "no matching user key")
	}

	ret_byte = []byte(user_certificate_data)

	return ret_byte, nil

}

func enc_dec_asym() {

	test_message, _ := RandomHex(32)

	test_message_b := []byte(test_message)

	priv_key_bytes, err := GetContextUserPrivateKeyBytes("kind-kindcluster1")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	priv_key, err := BytesToPrivateKey(priv_key_bytes)

	if err != nil {

		fmt.Println(err.Error())
		return
	}

	pub_key_bytes, err := GetContextUserPublicKeyBytes("kind-kindcluster1")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	pub_key, err := BytesToPublicKey(pub_key_bytes)

	if err != nil {

		fmt.Println(err.Error())
		return
	}

	enc_message, err := EncryptWithPublicKey(test_message_b, pub_key)

	if err != nil {

		fmt.Println(err.Error())
		return
	}

	dec_message, err := DecryptWithPrivateKey(enc_message, priv_key)

	dec_message_str := string(dec_message)

	if dec_message_str == test_message {
		fmt.Println("SUCCESS")
	} else {
		fmt.Println("FAIL")
	}

}

func enc_dec_sym() {

	test_key, _ := RandomHex(16)

	test_key_b := []byte(test_key)

	file_byte, err := os.ReadFile("loremipsum_short")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	enc_txt, err := EncryptWithSymmetricKey(test_key_b, file_byte)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = os.WriteFile("loremipsum_short_enc", enc_txt, 0644)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	file_byte_enc, err := os.ReadFile("loremipsum_short_enc")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	dec_txt, err := DecryptWithSymmetricKey(test_key_b, file_byte_enc)

	str_file_byte := string(file_byte)

	str_dec_txt := string(dec_txt)

	if str_file_byte == str_dec_txt {
		fmt.Println("SUCCESS")
	} else {
		fmt.Println("FAIL")
	}

}

func save_test() {

	get_kubeconfig_path_command_string :=
		`#!/bin/bash
[[ ! -z "$KUBECONFIG" ]] && echo "$KUBECONFIG" || echo "$HOME/.kube/config"`

	get_kubeconfig_path_command_b := []byte(get_kubeconfig_path_command_string)

	err := os.WriteFile("srv/get_kubeconfig_path", get_kubeconfig_path_command_b, 0755)

	if err != nil {

		fmt.Println(err.Error())
		return
	}

	cmd := exec.Command("srv/get_kubeconfig_path")

	t, err := cmd.Output()

	if err != nil {

		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(t))

}

func VerifyCertAgainstPub() {

	pub_b, err := GetContextClusterPublicKeyBytes("kind-kindcluster1")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	pub_key, err := BytesToPublicKey([]byte(pub_b))

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	cert_b, err := GetContextUserCertificateBytes("kind-kindcluster1")

	block, _ := pem.Decode(cert_b)

	var cert *x509.Certificate

	cert, err = x509.ParseCertificate(block.Bytes)

	hash_sha := sha256.New()

	hash_sha.Write(cert.RawTBSCertificate)

	hash_data := hash_sha.Sum(nil)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = rsa.VerifyPKCS1v15(pub_key, crypto.SHA256, hash_data, cert.Signature)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

}

func Save_Empty() {

	challenge_records := make(modules.ChallengRecord)

	key_records := make(modules.KeyRecord)

	multi_record := make(kalfs.MultiAppOrigin)

	var single_record runtimefs.AppOrigin

	var record runtimefs.RecordInfo

	var repo runtimefs.RepoInfo

	var reg runtimefs.RegInfo

	challenge_records["test"] = map[string]string{
		"test": "test",
	}

	key_records["test"] = "test"

	single_record.RECORDS = append(single_record.RECORDS, record)

	single_record.REPOS = append(single_record.REPOS, repo)

	single_record.REGS = append(single_record.REGS, reg)

	multi_record["_INIT"] = single_record

	challenge_records_b, err := json.Marshal(challenge_records)

	if err != nil {

		fmt.Println(err.Error())

		return
	}

	err = os.WriteFile("challenge.json", challenge_records_b, 0644)

	key_records_b, err := json.Marshal(key_records)

	if err != nil {

		fmt.Println(err.Error())
		return
	}

	err = os.WriteFile("key.json", key_records_b, 0644)

	multi_record_b, err := json.Marshal(multi_record)

	if err != nil {

		fmt.Println(err.Error())

		return
	}

	single_record_b, err := json.Marshal(single_record)

	if err != nil {

		fmt.Println(err.Error())

		return
	}

	err = os.WriteFile("multi.json", multi_record_b, 0644)

	err = os.WriteFile("single.json", single_record_b, 0644)
}

func main() {

	//	ASgi := apistandard.ASgi

	//	ASgi.PrintPrettyDefinition()

	// GetKubeConfigPathSimple()
	/*
		if rec, err := LoadTest(); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(rec)
		}
	*/

	/*
		if a, e := GetContextUserPrivateKeyBytes("kind-kindcluster1"); e != nil {
			fmt.Println(e.Error())
		} else {
			fmt.Println(string(a))
		}
	*/
	/*
		if a, e := GetContextUserPublicKeyBytes("kind-kindcluster1"); e != nil {
			fmt.Println(e.Error())
		} else {
			fmt.Println(string(a))
		}
	*/

	// enc_dec_asym()

	// enc_dec_sym()

	// VerifyCertAgainstPub()

	// save_test()

	Save_Empty()
}
