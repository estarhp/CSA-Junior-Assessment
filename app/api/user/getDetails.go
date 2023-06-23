package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"my-app/dao"
	"my-app/dao/redis"
	"my-app/logs"
	"my-app/utils"
	"net/http"
)

func GetUserDetail(c *gin.Context) {
	username := utils.GetUsername(c)

	user, err := dao.GetUserDetails(username)
	if err != nil {
		logs.LogError(err)
		utils.RespFail(c, "internal error")
		return
	}
	user.Password = ""

	follows, err := redis.GetAllFollows(username)
	if err != nil {
		logs.LogError(err)
		return
	}
	JsonFollows, err := json.Marshal(follows)

	if err != nil {
		logs.LogError(err)
		utils.RespFail(c, "internal error")
		return
	}
	user.Follows = string(JsonFollows)

	c.JSON(http.StatusOK, gin.H{
		"userDetails": user,
	})
}

func GetOtherUserDetails(c *gin.Context) {
	username := c.Query("username")

	user, err := dao.GetUserDetails(username)
	if err != nil {
		logs.LogError(err)
		utils.RespFail(c, "internal error")
		return
	}
	user.Password = ""
	follows, err := redis.GetAllFollows(username)
	if err != nil {
		logs.LogError(err)
		return
	}
	JsonFollows, err := json.Marshal(follows)

	if err != nil {
		logs.LogError(err)
		utils.RespFail(c, "internal error")
		return
	}
	user.Follows = string(JsonFollows)

	c.JSON(http.StatusOK, gin.H{
		"userDetails": user,
	})
}
