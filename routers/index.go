package routers

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {

	c.Redirect(http.StatusTemporaryRedirect, "/edit/"+generateToken())
}

func generateToken() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}
