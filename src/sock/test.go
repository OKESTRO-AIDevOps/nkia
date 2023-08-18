package sock

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	ctrl "github.com/OKESTRO-AIDevOps/nkia/src/controller"
	"github.com/OKESTRO-AIDevOps/nkia/src/modules"
	"github.com/OKESTRO-AIDevOps/npia-api/pkg/apistandard"
	"github.com/gorilla/websocket"
)

func DetachedServerCommunicator_Test(address string, email string) error {
	c, _, err := websocket.DefaultDialer.Dial(address, nil)
	if err != nil {
		return fmt.Errorf("comm failed: %s", err.Error())
	}
	defer c.Close()

	err = ServerAuthChallenge(c, email)

	if err != nil {
		return fmt.Errorf("comm failed: %s", err.Error())
	}

	if err := SockCommunicationHandler_LinearInstruction_PrintOnly_Test(c); err != nil {
		return fmt.Errorf("comm failed: %s", err.Error())
	}

	return nil
}

func SockCommunicationHandler_LinearInstruction_PrintOnly_Test(c *websocket.Conn) error {

	READ = 0

	// ASgi := apistandard.ASgi

	var req_body ctrl.APIMessageRequest
	var resp_body ctrl.APIMessageResponse

	ch_read := make(chan ctrl.APIMessageRequest)

	interrupt := make(chan os.Signal, 1)

	go SockCommunicationHandler_ReaderChannel(c, ch_read)

	comm_loop := 0

	for comm_loop == 0 {

		select {

		case read_body := <-ch_read:

			req_body = read_body

			enc_query := req_body.Query

			enc_query_b, err := hex.DecodeString(enc_query)

			if err != nil {
				READ = 1
				resp_body.ServerMessage = err.Error()
				err = c.WriteJSON(&read_body)

				if err != nil {
					return fmt.Errorf("comm handler write: %s", err.Error())
				}

				return fmt.Errorf("comm handler: %s", err.Error())
			}

			key_b := []byte(SESSION_SYM_KEY)

			linear_instruction_b, err := modules.DecryptWithSymmetricKey(key_b, enc_query_b)

			if err != nil {
				READ = 1
				resp_body.ServerMessage = err.Error()
				err = c.WriteJSON(&resp_body)

				if err != nil {
					return fmt.Errorf("comm handler write: %s", err.Error())
				}

				return fmt.Errorf("comm handler: %s", err.Error())

			}

			linear_instruction := string(linear_instruction_b)
			/*
				api_input, err := ASgi.StdCmdInputBuildFromLinearInstruction(linear_instruction)

				if err != nil {
					READ = 1
					resp_body.ServerMessage = err.Error()
					err = c.WriteJSON(&resp_body)

					if err != nil {
						return fmt.Errorf("comm handler write: %s", err.Error())
					}

					return fmt.Errorf("comm handler: %s", err.Error())

				}

				api_out, err := ASgi.Run(api_input)

				if err != nil {
					READ = 1
					resp_body.ServerMessage = err.Error()
					err = c.WriteJSON(&resp_body)

					if err != nil {
						return fmt.Errorf("comm handler write: %s", err.Error())
					}

					return fmt.Errorf("comm handler: %s", err.Error())

				}

			*/

			fmt.Println("******client side")
			fmt.Println(linear_instruction)

			var api_out apistandard.API_OUTPUT

			api_out.BODY = linear_instruction

			api_out_b, err := json.Marshal(api_out)

			if err != nil {
				READ = 1
				resp_body.ServerMessage = err.Error()
				err = c.WriteJSON(&resp_body)

				if err != nil {
					return fmt.Errorf("comm handler write: %s", err.Error())
				}

				return fmt.Errorf("comm handler: %s", err.Error())

			}

			ret_byte, err := modules.EncryptWithSymmetricKey(key_b, api_out_b)

			if err != nil {
				READ = 1
				resp_body.ServerMessage = err.Error()
				err = c.WriteJSON(&resp_body)

				if err != nil {
					return fmt.Errorf("comm handler write: %s", err.Error())
				}

				return fmt.Errorf("comm handler: %s", err.Error())

			}

			ret_enc := hex.EncodeToString(ret_byte)

			resp_body.ServerMessage = "SUCCESS"
			resp_body.QueryResult = ret_enc

			err = c.WriteJSON(&resp_body)

			if err != nil {
				READ = 1
				return fmt.Errorf("comm handler: %s", err.Error())
			}

		case <-interrupt:
			READ = 1
			comm_loop = 1
			break
		}

	}

	return nil
}
