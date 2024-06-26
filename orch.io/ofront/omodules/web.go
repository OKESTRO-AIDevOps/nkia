package omodules

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GenerateStateAuthCookie(c *gin.Context) string {

	b := make([]byte, 16)
	rand.Read(b)

	session := sessions.Default(c)

	state := base64.URLEncoding.EncodeToString(b)

	session.Set("OSID", state)

	session.Save()

	return state
}
