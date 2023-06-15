package utils

import (
	"github.com/gin-gonic/gin"
	"my-app/RSA"
	"my-app/logs"
	"strings"
)

func GetUsername(c *gin.Context) string {
	isLogin, err := c.Cookie("isLogin")
	if err != nil {
		return ""
	}
	unencrypted, err := RSA.DecryptByPublicKey(isLogin)
	if err != nil {
		logs.LogError(err)
		return ""
	}
	username := strings.Split(unencrypted, "+")[0]
	return username
}
