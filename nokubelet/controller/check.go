package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var healthiness_record = map[string]string{
	"id": "_healthiness_probe", "health": "Healthy",
}

func GetHealth(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, healthiness_record)
}
