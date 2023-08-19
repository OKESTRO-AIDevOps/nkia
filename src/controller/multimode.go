package controller

import (
	"encoding/hex"
	"encoding/json"

	"io"
	"net/http"
	"strings"

	kalfs "github.com/OKESTRO-AIDevOps/nkia/pkg/kaleidofs"
	"github.com/OKESTRO-AIDevOps/nkia/src/modules"
	"github.com/gin-gonic/gin"
)

func Multimode_LinearInstruction(c *gin.Context) {

	var req APIMessageRequest
	var resp APIMessageResponse

	var server_message string

	session_sym_key, err := modules.AccessAuth(c)

	if err != nil {
		resp.ServerMessage = err.Error()
		c.IndentedJSON(http.StatusForbidden, resp)
		return
	}

	body_byte, err := io.ReadAll(c.Request.Body)

	if err != nil {
		resp.ServerMessage = err.Error()
		c.IndentedJSON(http.StatusInternalServerError, resp)
		return
	}

	err = json.Unmarshal(body_byte, &req)

	if err != nil {
		resp.ServerMessage = err.Error()
		c.IndentedJSON(http.StatusBadRequest, resp)
		return
	}

	enc_query := req.Query

	enc_query_b, err := hex.DecodeString(enc_query)

	if err != nil {
		resp.ServerMessage = err.Error()
		c.IndentedJSON(http.StatusBadRequest, resp)
		return
	}

	key_b := []byte(session_sym_key)

	linear_instruction_b, err := modules.DecryptWithSymmetricKey(key_b, enc_query_b)

	if err != nil {
		resp.ServerMessage = err.Error()
		c.IndentedJSON(http.StatusBadRequest, resp)
		return
	}

	linear_instruction := string(linear_instruction_b)

	linear_operand_variables := strings.SplitN(linear_instruction, ":", 2)

	operand := linear_operand_variables[0]

	variables := strings.Split(linear_operand_variables[1], ",")

	switch operand {
	case "INIT":

		err := kalfs.InitKaleidoRoot()

		if err != nil {
			resp.ServerMessage = err.Error()
			c.IndentedJSON(http.StatusOK, resp)
			return
		}

		server_message = "multimode initiation completed"

	case "SWITCH":

		switch_to := variables[0]

		err := kalfs.SaveAndSwitch(switch_to)

		if err != nil {
			resp.ServerMessage = err.Error()
			c.IndentedJSON(http.StatusOK, resp)
			return
		}

		server_message = "multimode switched to " + switch_to

	default:
		resp.ServerMessage = "invalid operand"
		c.IndentedJSON(http.StatusBadRequest, resp)
		return
	}

	server_message_b := []byte(server_message)

	ret_byte, err := modules.EncryptWithSymmetricKey(key_b, server_message_b)

	if err != nil {
		resp.ServerMessage = err.Error()
		c.IndentedJSON(http.StatusBadRequest, resp)
		return
	}

	ret_enc := hex.EncodeToString(ret_byte)

	resp.ServerMessage = "SUCCESS"
	resp.QueryResult = ret_enc

	c.IndentedJSON(http.StatusOK, resp)

}
