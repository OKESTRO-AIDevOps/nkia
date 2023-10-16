package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	ctrl "github.com/OKESTRO-AIDevOps/nkia/nokubelet/controller"
	"github.com/OKESTRO-AIDevOps/nkia/nokubelet/modules"
	"github.com/gorilla/websocket"
)

func KeyAuthConn(client *http.Client, email string) (*websocket.Conn, error) {

	var c *websocket.Conn

	var req_orchestrator ctrl.OrchestratorRequest

	var resp_orchestrator ctrl.OrchestratorResponse

	req_orchestrator.Query = email

	req_b, err := json.Marshal(req_orchestrator)

	if err != nil {
		return c, fmt.Errorf("auth: %s", err.Error())
	}

	resp, err := client.Post(COMM_URL_AUTH, "application/json", bytes.NewBuffer(req_b))

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

	// get privkey from srv/key.json

	// decrypt the challenge

	// send back

	// receive the token

	// establish web socket connection

	return c, nil
}

func RequestHandler_LinearInstruction_PrintOnly(c *websocket.Conn, target string, option string, linear_instruction string) {

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

			fmt.Println("print result: ")

			fmt.Println(result)

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
