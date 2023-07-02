package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/OKESTRO-AIDevOps/npia-api/pkg/apistandard"
	"github.com/gin-gonic/gin"

	_ "github.com/OKESTRO-AIDevOps/npia-server/server/modules"
)

func QueryAPI_LinearInstruction(c *gin.Context) {

	ASgi := apistandard.ASgi

	var req APIMessageRequest
	var resp APIMessageResponse

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

	linear_instruction := req.Query

	api_input, err := ASgi.StdCmdInputBuildFromLinearInstruction(linear_instruction)

	if err != nil {
		resp.ServerMessage = err.Error()
		c.IndentedJSON(http.StatusBadRequest, resp)
		return
	}

	api_out, err := ASgi.Run(api_input)

	if err != nil {
		resp.ServerMessage = "ERROR"
		api_out.BODY = err.Error()
		resp.QueryResult = api_out
		c.IndentedJSON(http.StatusOK, resp)
		return
	}

	resp.ServerMessage = "SUCCESS"
	resp.QueryResult = api_out
	c.IndentedJSON(http.StatusOK, resp)

	return
}
