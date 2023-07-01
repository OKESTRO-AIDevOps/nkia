package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/OKESTRO-AIDevOps/npia-api/pkg/apistandard"
)

var COMM_URL = "http://localhost:13337/api/v0alpha-test"

type JSONMessageRequest struct {
	Query string `json:"query"`
}

type JSONMessageResponse struct {
	ServerMessage string `json:"server_message"`

	QueryResult apistandard.API_OUTPUT `json:"query_result"`
}

func request_test() {

	var req_body JSONMessageRequest
	var resp_body JSONMessageResponse

	req_body.Query = "APPLY-DIST:test,test-addr.com,test-addr.com"

	req_b, err := json.Marshal(req_body)

	if err != nil {

		fmt.Println(err)
		return
	}

	resp, err := http.Post(COMM_URL, "application/json", bytes.NewBuffer(req_b))

	if err != nil {

		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body_bytes, err := io.ReadAll(resp.Body)
	if err != nil {

		fmt.Println(err)
		return
	}

	err = json.Unmarshal(body_bytes, &resp_body)

	if err != nil {

		fmt.Println(err)
		return
	}

	fmt.Println(resp_body.ServerMessage)
	fmt.Println(resp_body.QueryResult.BODY)

}

func main() {

	request_test()

}
