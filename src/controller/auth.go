package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/OKESTRO-AIDevOps/nkia/src/modules"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthChallengeHandler(c *gin.Context) {

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
