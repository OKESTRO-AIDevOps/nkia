package main

import (
	"github.com/gin-gonic/gin"

	"github.com/OKESTRO-AIDevOps/npia-server/server/router"
)

func main() {

	gin_srv := gin.Default()

	gin_srv = router.Init(gin_srv)

	gin_srv.Run("0.0.0.0:13337")
}
