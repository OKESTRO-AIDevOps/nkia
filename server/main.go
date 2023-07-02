package main

import (
	"os"

	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/OKESTRO-AIDevOps/npia-server/server/router"
)

func main() {

	if _, err := os.Stat("srv"); err != nil {

		fmt.Println(err.Error())

		return

	}

	gin_srv := gin.Default()

	gin_srv = router.Init(gin_srv)

	gin_srv.Run("0.0.0.0:13337")
}
