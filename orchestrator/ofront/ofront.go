package main

import (
	"github.com/OKESTRO-AIDevOps/nkia/orchestrator/ofront/orouter"

	"github.com/OKESTRO-AIDevOps/nkia/orchestrator/ofront/omodels"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func main() {

	omodels.DbEstablish()

	gin_srv := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	gin_srv.Use(sessions.Sessions("npia-session", store))

	gin_srv.LoadHTMLGlob("oview/*")

	gin_srv = orouter.Init(gin_srv)

	gin_srv.Run("0.0.0.0:1337")

}
