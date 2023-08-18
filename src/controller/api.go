package controller

import (
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"

	"github.com/OKESTRO-AIDevOps/npia-api/pkg/apistandard"
	"github.com/gin-gonic/gin"

	"github.com/OKESTRO-AIDevOps/nkia/src/modules"
)

func QueryAPI_LinearInstruction(c *gin.Context) {

	ASgi := apistandard.ASgi
	var req APIMessageRequest
	var resp APIMessageResponse

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

	api_input, err := ASgi.StdCmdInputBuildFromLinearInstruction(linear_instruction)

	if err != nil {
		resp.ServerMessage = err.Error()
		c.IndentedJSON(http.StatusBadRequest, resp)
		return
	}

	api_out, err := ASgi.Run(api_input)

	if err != nil {
		resp.ServerMessage = err.Error()
		resp.QueryResult = ""
		c.IndentedJSON(http.StatusOK, resp)
		return
	}

	api_out_b, err := json.Marshal(api_out)

	if err != nil {
		resp.ServerMessage = err.Error()
		c.IndentedJSON(http.StatusInternalServerError, resp)
		return
	}

	ret_byte, err := modules.EncryptWithSymmetricKey(key_b, api_out_b)

	if err != nil {
		resp.ServerMessage = err.Error()
		c.IndentedJSON(http.StatusBadRequest, resp)
		return
	}

	ret_enc := hex.EncodeToString(ret_byte)

	resp.ServerMessage = "SUCCESS"
	resp.QueryResult = ret_enc

	c.IndentedJSON(http.StatusOK, resp)

	return
}
