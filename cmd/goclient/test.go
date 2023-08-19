package goclient

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"

	"github.com/OKESTRO-AIDevOps/nkia/pkg/apistandard"
	"github.com/OKESTRO-AIDevOps/nkia/src/controller"
	"github.com/OKESTRO-AIDevOps/nkia/src/modules"
)

func BaseFlow_API_Test() {

	var err error

	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{
		Jar: jar,
	}

	err = ClientAuthChallenge(client)

	if err != nil {
		fmt.Println(err)
		return
	}

	var req_body controller.APIMessageRequest
	var resp_body controller.APIMessageResponse

	query_plain := "hello npia"

	query_enc, err := modules.EncryptWithSymmetricKey([]byte(SESSION_SYM_KEY), []byte(query_plain))

	if err != nil {

		fmt.Println(err)
		return
	}

	req_body.Query = hex.EncodeToString(query_enc)

	req_b, err := json.Marshal(req_body)

	if err != nil {

		fmt.Println(err)
		return
	}

	resp, err := client.Post(COMM_URL, "application/json", bytes.NewBuffer(req_b))

	if err != nil {

		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body_bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(string(body_bytes))
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(body_bytes, &resp_body)

	if err != nil {

		fmt.Println(err)
		return
	}

	result_enc := resp_body.QueryResult

	result_enc_b, err := hex.DecodeString(result_enc)

	if err != nil {

		fmt.Println(err)
		return
	}

	result_b, err := modules.DecryptWithSymmetricKey([]byte(SESSION_SYM_KEY), result_enc_b)

	if err != nil {

		fmt.Println(err)
		return
	}

	var api_out apistandard.API_OUTPUT

	err = json.Unmarshal(result_b, &api_out)

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println(resp_body.ServerMessage)
	fmt.Println(api_out)

}

func BaseFlow_APIThenMultiMode_Test() {

	var err error

	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{
		Jar: jar,
	}

	err = ClientAuthChallenge(client)

	if err != nil {
		fmt.Println(err)
		return
	}

	var req_body controller.APIMessageRequest
	var resp_body controller.APIMessageResponse

	query_plain := "hello npia"

	query_enc, err := modules.EncryptWithSymmetricKey([]byte(SESSION_SYM_KEY), []byte(query_plain))

	if err != nil {

		fmt.Println(err)
		return
	}

	req_body.Query = hex.EncodeToString(query_enc)

	req_b, err := json.Marshal(req_body)

	if err != nil {

		fmt.Println(err)
		return
	}

	resp, err := client.Post(COMM_URL, "application/json", bytes.NewBuffer(req_b))

	if err != nil {

		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body_bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(string(body_bytes))
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(body_bytes, &resp_body)

	if err != nil {

		fmt.Println(err)
		return
	}

	result_enc := resp_body.QueryResult

	result_enc_b, err := hex.DecodeString(result_enc)

	if err != nil {

		fmt.Println(err)
		return
	}

	result_b, err := modules.DecryptWithSymmetricKey([]byte(SESSION_SYM_KEY), result_enc_b)

	if err != nil {

		fmt.Println(err)
		return
	}

	var api_out apistandard.API_OUTPUT

	err = json.Unmarshal(result_b, &api_out)

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println("----------API----------")
	fmt.Println(resp_body.ServerMessage)
	fmt.Println(api_out)
	fmt.Println("-----------------------")
	fmt.Println(" ")

	query_plain = "INIT:"

	query_enc, err = modules.EncryptWithSymmetricKey([]byte(SESSION_SYM_KEY), []byte(query_plain))

	if err != nil {

		fmt.Println(err)
		return
	}

	req_body.Query = hex.EncodeToString(query_enc)

	req_b, err = json.Marshal(req_body)

	if err != nil {

		fmt.Println(err)
		return
	}

	resp, err = client.Post(COMM_URL_MULTIMODE, "application/json", bytes.NewBuffer(req_b))

	if err != nil {

		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body_bytes, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(string(body_bytes))
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(body_bytes, &resp_body)

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println(resp_body.ServerMessage)

	result_enc = resp_body.QueryResult

	result_enc_b, err = hex.DecodeString(result_enc)

	if err != nil {

		fmt.Println(err)
		return
	}

	result_b, err = modules.DecryptWithSymmetricKey([]byte(SESSION_SYM_KEY), result_enc_b)

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println("----------MULTIMODE----------")
	fmt.Println(string(result_b))

	query_plain = "SWITCH:kind-kindcluster2"

	query_enc, err = modules.EncryptWithSymmetricKey([]byte(SESSION_SYM_KEY), []byte(query_plain))

	if err != nil {

		fmt.Println(err)
		return
	}

	req_body.Query = hex.EncodeToString(query_enc)

	req_b, err = json.Marshal(req_body)

	if err != nil {

		fmt.Println(err)
		return
	}

	resp, err = client.Post(COMM_URL_MULTIMODE, "application/json", bytes.NewBuffer(req_b))

	if err != nil {

		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body_bytes, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(string(body_bytes))
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(body_bytes, &resp_body)

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println(resp_body.ServerMessage)

	result_enc = resp_body.QueryResult

	result_enc_b, err = hex.DecodeString(result_enc)

	if err != nil {

		fmt.Println(err)
		return
	}

	result_b, err = modules.DecryptWithSymmetricKey([]byte(SESSION_SYM_KEY), result_enc_b)

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println(string(result_b))

}
