package main

import (
	"github.com/OKESTRO-AIDevOps/nkia/orch.io/ofront/orouter"

	"github.com/OKESTRO-AIDevOps/nkia/orch.io/ofront/omodels"

	"github.com/OKESTRO-AIDevOps/nkia/orch.io/ofront/omodules"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func main() {

	omodules.LoadConfig()

	if !omodules.CONFIG_JSON.DEBUG {

		omodels.DbEstablish(
			omodules.CONFIG_JSON.DB_ID,
			omodules.CONFIG_JSON.DB_PW,
			omodules.CONFIG_JSON.DB_HOST,
			omodules.CONFIG_JSON.DB_NAME,
		)

	} else {
		omodels.DbEstablish(
			omodules.CONFIG_JSON.DB_ID,
			omodules.CONFIG_JSON.DB_PW,
			omodules.CONFIG_JSON.DB_HOST_DEV,
			omodules.CONFIG_JSON.DB_NAME,
		)
	}

	gin_srv := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	gin_srv.Use(sessions.Sessions("npia-session", store))

	gin_srv.LoadHTMLGlob("oview/*")

	gin_srv.Static("/static", "./ostatic")

	gin_srv = orouter.Init(gin_srv)

	gin_srv.Run("0.0.0.0:1337")

}
