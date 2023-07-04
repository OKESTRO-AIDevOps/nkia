package main

import (
	"os"

	"fmt"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/OKESTRO-AIDevOps/npia-server/src/router"
)

func main() {

	if _, err := os.Stat("srv"); err != nil {

		fmt.Println(err.Error())

		return

	}

	gin_srv := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	gin_srv.Use(sessions.Sessions("npia-session", store))

	gin_srv = router.Init(gin_srv)

	gin_srv.Run("0.0.0.0:13337")
}
