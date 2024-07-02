package client

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"time"

	cconf "github.com/OKESTRO-AIDevOps/nkia/nokubectl/config"
	"github.com/OKESTRO-AIDevOps/nkia/pkg/apistandard"
	ctrl "github.com/OKESTRO-AIDevOps/nkia/pkg/apistandard/apix"
	modules "github.com/OKESTRO-AIDevOps/nkia/pkg/challenge"
	"github.com/gorilla/websocket"
)

var PRINT_ONLY_BODY map[string]string

func CertAuthConn(client *websocket.Conn) error {

	var req_orchestrator ctrl.OrchestratorRequest

	var resp_orchestrator ctrl.OrchestratorResponse

	file_b, err := os.ReadFile(".npia/certs/client.crt")

	if err != nil {

		return fmt.Errorf("cert auth: %s", err.Error())

	}

	req_orchestrator.Query = string(file_b)

	err = client.WriteJSON(req_orchestrator)

	if err != nil {
		return fmt.Errorf("auth: %s", err.Error())
	}

	err = client.ReadJSON(&resp_orchestrator)

	if err != nil {
		return fmt.Errorf("auth: %s", err.Error())
	}

	if resp_orchestrator.ServerMessage != "SUCCESS" {
		return fmt.Errorf("auth: %s", err.Error())
	}

	return nil
}

func KeyAuthConn(client *websocket.Conn, email string) error {

	var req_orchestrator ctrl.OrchestratorRequest

	var resp_orchestrator ctrl.OrchestratorResponse

	req_orchestrator.Query = email

	err := client.WriteJSON(req_orchestrator)

	if err != nil {
		return fmt.Errorf("auth: %s", err.Error())
	}

	err = client.ReadJSON(&resp_orchestrator)

	if err != nil {
		return fmt.Errorf("auth: %s", err.Error())
	}

	if resp_orchestrator.ServerMessage != "SUCCESS" {
		return fmt.Errorf("auth: %s", err.Error())
	}

	chal_rec := modules.ChallengRecord{}

	chal_rec_b := resp_orchestrator.QueryResult

	err = json.Unmarshal(chal_rec_b, &chal_rec)

	if err != nil {
		return fmt.Errorf("auth: %s", err.Error())
	}

	// get privkey from srv/.priv

	priv_key, err := cconf.LoadKeyAuthCredential()

	if err != nil {
		return fmt.Errorf("auth: %s", err.Error())
	}

	// decrypt the challenge

	chal_id := ""

	for k := range chal_rec {

		chal_id = k

	}

	cipher_txt := chal_rec[chal_id][email]

	cipher_b, err := hex.DecodeString(cipher_txt)

	if err != nil {
		return fmt.Errorf("auth: %s", err.Error())
	}

	ans, err := modules.DecryptWithPrivateKey(cipher_b, priv_key)

	if err != nil {
		return fmt.Errorf("auth: %s", err.Error())
	}

	ans_str := string(ans)

	chal_rec[chal_id][email] = ans_str

	query_b, err := json.Marshal(chal_rec)

	query_base64 := base64.StdEncoding.EncodeToString(query_b)

	req_orchestrator = ctrl.OrchestratorRequest{
		Query: query_base64,
	}

	err = client.WriteJSON(req_orchestrator)

	if err != nil {
		return fmt.Errorf("auth: %s", err.Error())
	}

	resp_orchestrator = ctrl.OrchestratorResponse{}

	err = client.ReadJSON(&resp_orchestrator)

	if err != nil {
		return fmt.Errorf("auth: %s", err.Error())
	}

	if resp_orchestrator.ServerMessage != "SUCCESS" {
		return fmt.Errorf("auth: %s", resp_orchestrator.ServerMessage)
	}

	return nil
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

func Do(c *websocket.Conn, req_orchestrator ctrl.OrchestratorRequest) {

	recv := make(chan ctrl.OrchestratorResponse)

	var MSG string

	var OUT apistandard.API_OUTPUT

	//var HEAD apistandard.API_METADATA

	var BODY string

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
				fmt.Fprintf(os.Stderr, "err: %s", err.Error())

				return
			}

			if MSG != "SUCCESS" {

				fmt.Fprintf(os.Stderr, "failed: %s", MSG)

				return
			}

			//HEAD = OUT.HEAD

			BODY = OUT.BODY

			//head_b, err := json.Marshal(HEAD)

			if err != nil {

				fmt.Fprintf(os.Stderr, "err: %s", err.Error())
			}

			//_ = os.WriteFile(".npia/_output/HEAD", head_b, 0644)

			//_ = os.WriteFile(".npia/_output/BODY", []byte(BODY), 0644)

			fmt.Fprintf(os.Stdout, "%s", BODY)

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
