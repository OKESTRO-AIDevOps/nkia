package controller

import (
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/OKESTRO-AIDevOps/npia-api/pkg/apistandard"
	"github.com/OKESTRO-AIDevOps/npia-server/server/modules"
	_ "github.com/OKESTRO-AIDevOps/npia-server/server/modules"
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

		chal_rec, err := modules.GenerateChallenge()

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
