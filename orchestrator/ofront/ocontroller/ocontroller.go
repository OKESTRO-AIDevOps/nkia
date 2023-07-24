package ocontroller

import (
	"github.com/OKESTRO-AIDevOps/npia-server/orchestrator/ofront/omodels"
	"github.com/OKESTRO-AIDevOps/npia-server/orchestrator/ofront/omodules"

	ctrl "github.com/OKESTRO-AIDevOps/npia-server/src/controller"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"encoding/json"
	"fmt"
)

func IndexFeed(c *gin.Context) {

	c.HTML(200, "index.html", gin.H{
		"title": "index test",
	})

}

func IndexFeed_Test(c *gin.Context) {

	c.HTML(200, "index.html", gin.H{
		"title": "index test",
	})

}

func OauthGoogleLogin(c *gin.Context) {

	oauth_state := omodules.GenerateStateOauthCookie(c)

	u := omodules.GoogleOauthConfig.AuthCodeURL(oauth_state)

	c.Redirect(302, u)

}

func OauthGoogleCallback(c *gin.Context) {

	session := sessions.Default(c)

	var session_id string

	v := session.Get("OSID")

	if v == nil {
		fmt.Printf("access auth failed: %s", "session id not found")
		return
	} else {
		session_id = v.(string)
	}

	state := c.Request.FormValue("state")

	if state == "" {
		fmt.Printf("access auth failed: %s", "form state not found")
		return
	}

	if state != session_id {
		fmt.Printf("access auth failed: %s", "value not matching")
		c.Redirect(302, "/")
		return
	}

	data, err := omodules.GetUserDataFromGoogle(c.Request.FormValue("code"))
	if err != nil {
		fmt.Printf("access auth failed: %s", err.Error())
		c.Redirect(302, "/")
		return
	}

	fmt.Println(string(data))

	var oauth_struct omodules.OAuthStruct

	err = json.Unmarshal(data, &oauth_struct)

	if err != nil {
		fmt.Printf("access auth failed: %s", err.Error())
		c.Redirect(302, "/")
		return
	}

	if !oauth_struct.VERIFIED_EMAIL {
		fmt.Printf("access auth failed: %s", err.Error())
		c.Redirect(302, "/")
		return
	}

	request_key, err := omodels.RegisterOsidAndRequestKey(session_id, oauth_struct)

	if err != nil {
		fmt.Printf("access auth failed: %s", err.Error())
		c.Redirect(302, "/")
		return
	}

	var server_response ctrl.OrchestratorResponse

	server_response.ServerMessage = "SUCCESS"

	server_response.QueryResult = []byte(request_key)

	fmt.Println(request_key)

	c.IndentedJSON(200, server_response)

	return
}

func OauthGoogleCallback_Test(c *gin.Context) {

	session := sessions.Default(c)

	var session_id string

	v := session.Get("OSID")

	if v == nil {
		fmt.Printf("access auth failed: %s", "session id not found")
		return
	} else {
		session_id = v.(string)
	}

	state := c.Request.FormValue("state")

	if state == "" {
		fmt.Printf("access auth failed: %s", "form state not found")
		return
	}

	if state != session_id {
		fmt.Printf("access auth failed: %s", "value not matching")
		c.Redirect(302, "/")
		return
	}

	data, err := omodules.GetUserDataFromGoogle(c.Request.FormValue("code"))
	if err != nil {
		fmt.Printf("access auth failed: %s", err.Error())
		c.Redirect(302, "/")
		return
	}

	fmt.Println(string(data))

	var oauth_struct omodules.OAuthStruct

	err = json.Unmarshal(data, &oauth_struct)

	if err != nil {
		fmt.Printf("access auth failed: %s", err.Error())
		c.Redirect(302, "/")
		return
	}

	fmt.Println(oauth_struct)

	return
}
