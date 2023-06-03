package utils

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func GetUsername(c *gin.Context) string {
	isLogin, err := c.Cookie("isLogin")
	if err != nil {
		return ""
	}
	username := strings.Split(isLogin, "+")[0]
	return username
}
