package ocontroller

import (
	"fmt"

	"github.com/OKESTRO-AIDevOps/npia-server/orchestrator/ofront/omodules"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

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

	c.String(200, "text/plain", string(data))

	return
}
