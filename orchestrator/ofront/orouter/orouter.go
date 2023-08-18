package orouter

import (
	octrl "github.com/OKESTRO-AIDevOps/nkia/orchestrator/ofront/ocontroller"
	"github.com/gin-gonic/gin"
)

func Init(gin_srv *gin.Engine) *gin.Engine {

	gin_srv.GET("/", octrl.IndexFeed)

	gin_srv.GET("/orchestrate", octrl.OrchestratorFeed)

	gin_srv.GET("/oauth2/google/login", octrl.OauthGoogleLogin)

	gin_srv.GET("/oauth2/google/callback", octrl.OauthGoogleCallback)

	//	gin_srv.GET("/auth2/google/callback", octrl.OauthGoogleCallback_Test)

	return gin_srv
}
