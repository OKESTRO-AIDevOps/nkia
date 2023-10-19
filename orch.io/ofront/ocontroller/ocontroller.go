package ocontroller

import (
	ctrl "github.com/OKESTRO-AIDevOps/nkia/nokubelet/controller"
	"github.com/OKESTRO-AIDevOps/nkia/nokubelet/modules"
	"github.com/OKESTRO-AIDevOps/nkia/orch.io/ofront/omodels"
	"github.com/OKESTRO-AIDevOps/nkia/orch.io/ofront/omodules"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"encoding/base64"
	"encoding/json"
	"fmt"
)

func IndexFeed(c *gin.Context) {

	session := sessions.Default(c)

	var session_id string

	v := session.Get("OSID")

	if v != nil {
		session_id = v.(string)

		_, err := omodels.FrontAccessAuth(session_id)

		if err == nil {
			c.Redirect(302, "/orchestrate")
			return
		}

	}

	c.HTML(200, "index.html", gin.H{
		"title": "Index",
	})

}

func OrchestratorFeed(c *gin.Context) {

	session := sessions.Default(c)

	var session_id string

	v := session.Get("OSID")

	if v == nil {
		fmt.Printf("access auth failed: %s\n", "session id not found")
		c.String(403, "forbidden")
		return
	} else {
		session_id = v.(string)
	}

	request_key, err := omodels.FrontAccessAuth(session_id)

	if err != nil {
		fmt.Printf("access auth failed: %s\n", "request key not found")
		c.String(403, "forbidden")
		return
	}

	req_code_b64 := base64.StdEncoding.EncodeToString([]byte(request_key))

	c.HTML(200, "orchestrator.html", gin.H{
		"request_key": req_code_b64,
	})

}

func IndexFeed_Test(c *gin.Context) {

	c.HTML(200, "index.html", gin.H{
		"title": "index test",
	})

}

func OauthGoogleLogin(c *gin.Context) {

	oauth_state := omodules.GenerateStateAuthCookie(c)

	u := omodules.GoogleOauthConfig.AuthCodeURL(oauth_state)

	c.Redirect(302, u)

}

func OauthGoogleCallback(c *gin.Context) {

	session := sessions.Default(c)

	var session_id string

	v := session.Get("OSID")

	if v == nil {
		fmt.Printf("access auth failed: %s\n", "session id not found")
		return
	} else {
		session_id = v.(string)
	}

	state := c.Request.FormValue("state")

	if state == "" {
		fmt.Printf("access auth failed: %s\n", "form state not found")
		return
	}

	if state != session_id {
		fmt.Printf("access auth failed: %s\n", "value not matching")
		c.Redirect(302, "/")
		return
	}

	data, err := omodules.GetUserDataFromGoogle(c.Request.FormValue("code"))
	if err != nil {
		fmt.Printf("access auth failed: %s\n", err.Error())
		c.Redirect(302, "/")
		return
	}

	var oauth_struct omodules.OAuthStruct

	err = json.Unmarshal(data, &oauth_struct)

	if err != nil {
		fmt.Printf("access auth failed: %s\n", err.Error())
		c.Redirect(302, "/")
		return
	}

	if !oauth_struct.VERIFIED_EMAIL {
		fmt.Printf("access auth failed: %s\n", err.Error())
		c.Redirect(302, "/")
		return
	}

	_, err = omodels.RegisterOsidAndRequestKeyByOAuth(session_id, oauth_struct)

	if err != nil {
		fmt.Printf("access auth failed: %s\n", err.Error())
		c.Redirect(302, "/")
		return
	}

	c.Redirect(302, "/orchestrate")

	return
}

func OauthGoogleCallback_Test(c *gin.Context) {

	session := sessions.Default(c)

	var session_id string

	v := session.Get("OSID")

	if v == nil {
		fmt.Printf("access auth failed: %s\n", "session id not found")
		return
	} else {
		session_id = v.(string)
	}

	state := c.Request.FormValue("state")

	if state == "" {
		fmt.Printf("access auth failed: %s\n", "form state not found")
		return
	}

	if state != session_id {
		fmt.Printf("access auth failed: %s\n", "value not matching")
		c.Redirect(302, "/")
		return
	}

	data, err := omodules.GetUserDataFromGoogle(c.Request.FormValue("code"))
	if err != nil {
		fmt.Printf("access auth failed: %s\n", err.Error())
		c.Redirect(302, "/")
		return
	}

	fmt.Println(string(data))

	var oauth_struct omodules.OAuthStruct

	err = json.Unmarshal(data, &oauth_struct)

	if err != nil {
		fmt.Printf("access auth failed: %s\n", err.Error())
		c.Redirect(302, "/")
		return
	}

	fmt.Println(oauth_struct)

	return
}

func KeyAuthLogin(c *gin.Context) {

	var req_orchestrator = ctrl.OrchestratorRequest{}
	var res_orchestrator = ctrl.OrchestratorResponse{}

	err := c.BindJSON(&req_orchestrator)

	if err != nil {
		fmt.Printf("login failed: %s\n", err.Error())

		res_orchestrator.ServerMessage = err.Error()

		c.JSON(403, res_orchestrator)

		return
	}

	email := req_orchestrator.Query

	pubkey, err := omodels.GetPubkeyByEmail(email)

	if err != nil {

		fmt.Printf("login failed: %s\n", err.Error())

		res_orchestrator.ServerMessage = err.Error()

		c.JSON(403, res_orchestrator)

		return

	}

	chal_rec, err := modules.GenerateChallenge_Key(email, pubkey)

	if err != nil {

		fmt.Printf("login failed: %s\n", err.Error())

		res_orchestrator.ServerMessage = err.Error()

		c.JSON(500, res_orchestrator)

		return

	}

	chal_rec_b, err := json.Marshal(chal_rec)

	if err != nil {

		fmt.Printf("login failed: %s\n", err.Error())

		res_orchestrator.ServerMessage = err.Error()

		c.JSON(500, res_orchestrator)

		return

	}

	_ = omodules.GenerateStateAuthCookie(c)

	res_orchestrator.ServerMessage = "SUCCESS"

	res_orchestrator.QueryResult = chal_rec_b

	c.JSON(200, res_orchestrator)

	return

}

func KeyAuthCallback(c *gin.Context) {

	var req_orchestrator = ctrl.OrchestratorRequest{}
	var res_orchestrator = ctrl.OrchestratorResponse{}

	var answer modules.ChallengRecord

	session := sessions.Default(c)

	var session_id string

	v := session.Get("OSID")

	if v == nil {
		fmt.Printf("login callback failed: %s\n", "no session")

		res_orchestrator.ServerMessage = "no session"

		c.JSON(403, res_orchestrator)

		return
	} else {
		session_id = v.(string)
	}

	err := c.BindJSON(&req_orchestrator)

	if err != nil {
		fmt.Printf("login callback failed: %s\n", err.Error())

		res_orchestrator.ServerMessage = err.Error()

		c.JSON(403, res_orchestrator)

		return
	}

	answer_json_b64 := req_orchestrator.Query

	answer_json_b, err := base64.StdEncoding.DecodeString(answer_json_b64)

	if err != nil {
		fmt.Printf("login callback failed: %s\n", err.Error())

		res_orchestrator.ServerMessage = err.Error()

		c.JSON(403, res_orchestrator)

		return
	}

	err = json.Unmarshal(answer_json_b, &answer)

	if err != nil {
		fmt.Printf("login callback failed: %s\n", err.Error())

		res_orchestrator.ServerMessage = err.Error()

		c.JSON(403, res_orchestrator)

		return
	}

	email, err := modules.VerifyChallange_Key(answer)

	if err != nil {
		fmt.Printf("login callback failed: %s\n", err.Error())

		res_orchestrator.ServerMessage = err.Error()

		c.JSON(403, res_orchestrator)

		return
	}

	req_key, err := omodels.RegisterOsidAndRequestKeyByEmail(email, session_id)

	if err != nil {
		fmt.Printf("login callback failed: %s\n", err.Error())

		res_orchestrator.ServerMessage = err.Error()

		c.JSON(500, res_orchestrator)

		return
	}

	res_orchestrator.ServerMessage = "SUCCESS"

	res_orchestrator.QueryResult = []byte(req_key)

	c.JSON(200, res_orchestrator)

	return

}
