package orouter

import (
	octrl "github.com/OKESTRO-AIDevOps/npia-server/orchestrator/ofront/ocontroller"
	"github.com/gin-gonic/gin"
)

func Init(gin_srv *gin.Engine) *gin.Engine {

	gin_srv.GET("/", octrl.IndexFeed_Test)

	gin_srv.GET("/oauth2/google/login", octrl.OauthGoogleLogin)

	gin_srv.GET("/oauth2/google/callback", octrl.OauthGoogleCallback_Test)

	// gin_srv.GET("/auth2/google/callback", octrl.OauthGoogleCallback_Test)

	return gin_srv
}
