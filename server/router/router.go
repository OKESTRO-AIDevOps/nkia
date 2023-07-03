package router

import (
	"github.com/gin-gonic/gin"

	ctrl "github.com/OKESTRO-AIDevOps/npia-server/server/controller"
)

func Init(gin_srv *gin.Engine) *gin.Engine {

	gin_srv.GET("/get-health", ctrl.GetHealth)

	gin_srv.POST("/auth-challenge", ctrl.AuthChallengeHandler)

	// gin_srv.POST("/api/v0alpha", ctrl.QueryAPI_LinearInstruction)

	// gin_srv.POST("/multimod/v0alpha",ctrl.HandleMultiCluster_LinearInstruction)

	gin_srv.POST("/auth-challenge/test", ctrl.AuthChallenge_Test)

	gin_srv.POST("/api/v0alpha/test", ctrl.QueryAPI_LinearInstruction_Test)

	// gin_srv.POST("/multimod/v0alpha/test",ctrl.HandleMultiCluster_LinearInstruction_Test)

	return gin_srv
}
