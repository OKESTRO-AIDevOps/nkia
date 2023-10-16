package client

import (
	"fmt"
	"net/http"
	"time"

	ctrl "github.com/OKESTRO-AIDevOps/nkia/nokubelet/controller"
	"github.com/gorilla/websocket"
)

func KeyAuthConn(client *http.Client) (*websocket.Conn, error) {

	var c *websocket.Conn

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
