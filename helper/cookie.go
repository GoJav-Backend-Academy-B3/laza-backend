package helper

import (
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/gin-gonic/gin"
)

func GenerateStateOauthCookie(c *gin.Context) (state string) {
	var maxAge = 2 * time.Minute
	b := make([]byte, 16)
	rand.Read(b)
	state = base64.URLEncoding.EncodeToString(b)
	c.SetCookie("oauthstate", state, int(maxAge), "/", "", true, true)
	return
}
