package ocontroller

import (
	"github.com/OKESTRO-AIDevOps/nkia/orchestrator/ofront/omodels"
	"github.com/OKESTRO-AIDevOps/nkia/orchestrator/ofront/omodules"

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
		fmt.Printf("access auth failed: %s", "session id not found")
		c.String(403, "forbidden")
		return
	} else {
		session_id = v.(string)
	}

	request_key, err := omodels.FrontAccessAuth(session_id)

	if err != nil {
		fmt.Printf("access auth failed: %s", "request key not found")
		c.String(403, "forbidden")
		return
	}

	req_code_b64 := base64.StdEncoding.EncodeToString([]byte(request_key))

	c.HTML(200, "orchestrator.tmpl", gin.H{
		"request_key": req_code_b64,
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

	_, err = omodels.RegisterOsidAndRequestKey(session_id, oauth_struct)

	if err != nil {
		fmt.Printf("access auth failed: %s", err.Error())
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
