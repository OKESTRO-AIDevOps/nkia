package controller

import (
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/OKESTRO-AIDevOps/nkia/pkg/apistandard"
	kalfs "github.com/OKESTRO-AIDevOps/nkia/pkg/kaleidofs"
	"github.com/OKESTRO-AIDevOps/nkia/src/modules"
)

func QueryAPI_LinearInstruction_Test(c *gin.Context) {

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

	linear_instruction := req.Query

	enc_query_b, err := hex.DecodeString(linear_instruction)

	if err != nil {
		resp.ServerMessage = err.Error()
		c.IndentedJSON(http.StatusBadRequest, resp)
		return
	}

	key_b := []byte(session_sym_key)

	message, err := modules.DecryptWithSymmetricKey(key_b, enc_query_b)

	if err != nil {
		resp.ServerMessage = err.Error()
		c.IndentedJSON(http.StatusBadRequest, resp)
		return
	}

	var api_out apistandard.API_OUTPUT

	api_out.BODY = "test_success"

	test_b, _ := json.Marshal(api_out)

	ret_byte, err := modules.EncryptWithSymmetricKey(key_b, test_b)

	if err != nil {
		resp.ServerMessage = err.Error()
		c.IndentedJSON(http.StatusBadRequest, resp)
		return
	}

	ret_enc := hex.EncodeToString(ret_byte)

	resp.ServerMessage = string(message)
	resp.QueryResult = ret_enc

	c.IndentedJSON(http.StatusOK, resp)

	return
}

func AuthChallenge_Test(c *gin.Context) {

	var req AuthChallenge

	var resp AuthChallenge

	body_byte, err := io.ReadAll(c.Request.Body)

	if err != nil {
		resp.ChallengeID = "NOPE"
		resp.ChallengeMessage = err.Error()
		c.IndentedJSON(http.StatusInternalServerError, resp)
		return
	}

	err = json.Unmarshal(body_byte, &req)

	if err != nil {
		resp.ChallengeID = "NOPE"
		resp.ChallengeMessage = err.Error()
		c.IndentedJSON(http.StatusBadRequest, resp)
		return
	}

	chal_id := req.ChallengeID

	switch chal_id {

	case "ASK":

		client_ca_pub_key := req.ChallengeData

		chal_rec, err := modules.GenerateChallenge(client_ca_pub_key)

		if err != nil {
			resp.ChallengeID = "NOPE"
			resp.ChallengeMessage = err.Error()
			c.IndentedJSON(http.StatusOK, resp)
			return
		}

		resp.ChallengeID = "ASK"
		resp.ChallengeData = chal_rec
		c.IndentedJSON(http.StatusOK, resp)
		return

	case "ANS":

		answer := req.ChallengeData

		gen_key, key_rec, err := modules.VerifyChallange(answer)

		if err != nil {
			resp.ChallengeID = "NOPE"
			resp.ChallengeMessage = err.Error()
			c.IndentedJSON(http.StatusOK, resp)
			return
		}

		session := sessions.Default(c)

		session.Set("SID", gen_key)
		session.Save()

		resp.ChallengeID = "ASK"
		resp.ChallengeKey = key_rec
		c.IndentedJSON(http.StatusOK, resp)
		return

	default:
		resp.ChallengeID = "NOPE"
		resp.ChallengeMessage = "invalid challenge id"
		c.IndentedJSON(http.StatusOK, resp)
		return

	}

}

func Multimode_LinearInstruction_Test(c *gin.Context) {

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
