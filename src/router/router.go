package router

import (
	"github.com/gin-gonic/gin"

	ctrl "github.com/OKESTRO-AIDevOps/nkia/src/controller"
)

func Init(gin_srv *gin.Engine) *gin.Engine {

	gin_srv.GET("/get-health", ctrl.GetHealth)

	gin_srv.POST("/auth-challenge", ctrl.AuthChallengeHandler)

	gin_srv.POST("/api/v0alpha", ctrl.QueryAPI_LinearInstruction)

	gin_srv.POST("/multimode/v0alpha", ctrl.Multimode_LinearInstruction)

	gin_srv.POST("/auth-challenge/test", ctrl.AuthChallenge_Test)

	gin_srv.POST("/api/v0alpha/test", ctrl.QueryAPI_LinearInstruction_Test)

	gin_srv.POST("/multimode/v0alpha/test", ctrl.Multimode_LinearInstruction_Test)

	return gin_srv
}
