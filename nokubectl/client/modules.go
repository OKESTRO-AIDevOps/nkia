package client

import (
	"bytes"
	"crypto/rsa"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/OKESTRO-AIDevOps/nkia/nokubectl/config"
	ctrl "github.com/OKESTRO-AIDevOps/nkia/nokubelet/controller"
	"github.com/OKESTRO-AIDevOps/nkia/nokubelet/modules"
	"github.com/OKESTRO-AIDevOps/nkia/pkg/apistandard"
	"github.com/gorilla/websocket"
)

var PRINT_ONLY_BODY map[string]string

func KeyAuthConn(client *http.Client, email string) (*websocket.Conn, error) {

	var c *websocket.Conn

	var req_orchestrator ctrl.OrchestratorRequest

	var resp_orchestrator ctrl.OrchestratorResponse

	req_orchestrator.Query = email

	req_b, err := json.Marshal(req_orchestrator)

	if err != nil {
		return c, fmt.Errorf("auth: %s", err.Error())
	}

	resp, err := client.Post(config.COMM_URL_AUTH, "application/json", bytes.NewBuffer(req_b))

	if err != nil {
		return c, fmt.Errorf("auth: %s", err.Error())
	}

	body_bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return c, fmt.Errorf("auth: %s", err.Error())
	}

	resp.Body.Close()

	err = json.Unmarshal(body_bytes, &resp_orchestrator)

	if err != nil {
		return c, fmt.Errorf("auth: %s", err.Error())
	}

	chal_rec := modules.ChallengRecord{}

	chal_rec_b := resp_orchestrator.QueryResult

	err = json.Unmarshal(chal_rec_b, &chal_rec)

	if err != nil {
		return c, fmt.Errorf("auth: %s", err.Error())
	}

	// get privkey from srv/.priv

	priv_key, err := LoadKeyAuthCredential()

	if err != nil {
		return c, fmt.Errorf("auth: %s", err.Error())
	}

	// decrypt the challenge

	chal_id := ""

	for k := range chal_rec {

		chal_id = k

	}

	cipher_txt := chal_rec[chal_id][email]

	cipher_b, err := hex.DecodeString(cipher_txt)

	if err != nil {
		return c, fmt.Errorf("auth: %s", err.Error())
	}

	ans, err := modules.DecryptWithPrivateKey(cipher_b, priv_key)

	if err != nil {
		return c, fmt.Errorf("auth: %s", err.Error())
	}

	ans_str := string(ans)

	chal_rec[chal_id][email] = ans_str

	query_b, err := json.Marshal(chal_rec)

	query_base64 := base64.StdEncoding.EncodeToString(query_b)

	req_orchestrator = ctrl.OrchestratorRequest{
		Query: query_base64,
	}

	req_b, err = json.Marshal(req_orchestrator)

	resp, err = client.Post(config.COMM_URL_AUTH_CALLBACK, "application/json", bytes.NewBuffer(req_b))

	if err != nil {
		return c, fmt.Errorf("auth: %s", err.Error())
	}

	resp_orchestrator = ctrl.OrchestratorResponse{}

	body_bytes, err = io.ReadAll(resp.Body)
	if err != nil {
		return c, fmt.Errorf("auth: %s", err.Error())
	}

	resp.Body.Close()

	err = json.Unmarshal(body_bytes, &resp_orchestrator)

	if err != nil {
		return c, fmt.Errorf("auth: %s", err.Error())
	}

	if resp_orchestrator.ServerMessage != "SUCCESS" {
		return c, fmt.Errorf("auth: %s", resp_orchestrator.ServerMessage)
	}

	req_key := resp_orchestrator.QueryResult

	req_key_b64 := base64.StdEncoding.EncodeToString(req_key)

	req_orchestrator = ctrl.OrchestratorRequest{
		RequestOption: req_key_b64,
	}

	fmt.Println("connecting to the command channel...")

	c, _, err = websocket.DefaultDialer.Dial(config.COMM_URL, nil)

	if err != nil {
		return c, fmt.Errorf("auth conn: %s", err.Error())
	}

	err = c.WriteJSON(req_orchestrator)

	if err != nil {
		return c, fmt.Errorf("auth conn: %s", err.Error())
	}

	return c, nil
}

func RequestHandler_LinearInstruction_Persist_PrintOnly(c *websocket.Conn, target string, option string, linear_instruction string) {

	var req_orchestrator ctrl.OrchestratorRequest

	recv := make(chan ctrl.OrchestratorResponse)

	go RequestHandler_ReadChannel(c, recv)

	req_orchestrator.RequestTarget = target

	req_orchestrator.RequestOption = option

	req_orchestrator.Query = linear_instruction

	c.WriteJSON(req_orchestrator)

	counter := 0

	for {

		select {

		case result := <-recv:

			fmt.Printf("\n----------> print srv message: \n")

			fmt.Println(result.ServerMessage)

			fmt.Printf("\n----------> print q result: \n")

			fmt.Println(string(result.QueryResult))

			return

		default:

			counter += 1

			if counter%10 == 0 {
				fmt.Printf(". . .\n")
			} else {
				fmt.Printf(". . . ")
			}

			if counter > 100 {

				fmt.Printf("\nrequest timeout\n")

				return
			}

			time.Sleep(time.Millisecond * 100)

		}
	}

}

func RequestHandler_APIX_Once_PrintOnly(c *websocket.Conn, req_orchestrator ctrl.OrchestratorRequest) {

	recv := make(chan ctrl.OrchestratorResponse)

	var body_ret = PRINT_ONLY_BODY

	go RequestHandler_ReadChannel(c, recv)

	c.WriteJSON(req_orchestrator)

	counter := 0

	for {

		select {

		case result := <-recv:

			fmt.Printf("\n----------> print srv message: \n")

			fmt.Println(result.ServerMessage)

			fmt.Printf("\n----------> print q result: \n")

			_ = json.Unmarshal(result.QueryResult, &body_ret)

			fmt.Println(body_ret["BODY"])

			return

		default:

			counter += 1

			if counter > 100 {

				fmt.Println("request timeout")

				return
			}

			time.Sleep(time.Millisecond * 100)

		}
	}

}

func RequestHandler_APIX_Store_Override(c *websocket.Conn, req_orchestrator ctrl.OrchestratorRequest) {

	recv := make(chan ctrl.OrchestratorResponse)

	var MSG string

	var OUT apistandard.API_OUTPUT

	var HEAD apistandard.API_METADATA

	var BODY string

	req_b, err := json.Marshal(req_orchestrator)

	if err != nil {
		panic(err.Error())

	}

	_ = os.WriteFile(".npia/_apix_o/REQ", req_b, 0644)

	go RequestHandler_ReadChannel(c, recv)

	c.WriteJSON(req_orchestrator)

	counter := 0

	for {

		select {

		case result := <-recv:

			OUT = apistandard.API_OUTPUT{}

			MSG = result.ServerMessage

			err := json.Unmarshal(result.QueryResult, &OUT)

			if err != nil {
				panic(err.Error())
			}

			if MSG != "SUCCESS" {
				panic(err.Error())
			}

			_ = os.WriteFile(".npia/_apix_o/MSG", []byte(MSG), 0644)

			HEAD = OUT.HEAD

			BODY = OUT.BODY

			head_b, err := json.Marshal(HEAD)

			if err != nil {
				panic(err.Error())
			}

			_ = os.WriteFile(".npia/_apix_o/HEAD", head_b, 0644)

			_ = os.WriteFile(".npia/_apix_o/BODY", []byte(BODY), 0644)

			fmt.Println("SUCCESS")

			return

		default:

			counter += 1

			if counter > 100 {

				fmt.Println("request timeout")

				return
			}

			time.Sleep(time.Millisecond * 100)

		}
	}

}

func RequestHandler_ReadChannel(c *websocket.Conn, recv chan ctrl.OrchestratorResponse) {

	var resp_body ctrl.OrchestratorResponse

	err := c.ReadJSON(&resp_body)
	if err != nil {
		panic("auth reader1:" + err.Error())
	}

	recv <- resp_body

}

func LoadKeyAuthCredential() (*rsa.PrivateKey, error) {

	var ret_key *rsa.PrivateKey

	key_b, err := os.ReadFile(".npia/.priv")

	if err != nil {
		return ret_key, fmt.Errorf("failed to load cred: %s", err.Error())
	}

	ret_key, err = modules.BytesToPrivateKey(key_b)

	if err != nil {
		return ret_key, fmt.Errorf("failed to load cred: %s", err.Error())
	}

	return ret_key, nil
}
