package user

import (
	"github.com/gin-gonic/gin"
	"my-app/dao/redis"
	"my-app/logs"
	"my-app/utils"
)

func Follow(c *gin.Context) {
	followed := c.Query("followed")
	username := utils.GetUsername(c)
	if username == "" {
		utils.RespFail(c, "你还没有登录")
		return
	}
	err := redis.Follow(username, followed)
	if err != nil {
		logs.LogError(err)
		utils.RespFail(c, "internal error")
		return
	}
	utils.RespSuccess(c, "关注成功")
}

func Unfollow(c *gin.Context) {
	followed := c.Query("followed")
	username := utils.GetUsername(c)
	if username == "" {
		utils.RespFail(c, "你还没有登录")
		return
	}
	err := redis.UnFollow(username, followed)
	if err != nil {
		logs.LogError(err)
		utils.RespFail(c, "internal error")
		return
	}
	utils.RespFail(c, "取消关注成功")
}

func GetAllFollows(c *gin.Context) []string {
	username := utils.GetUsername(c)
	if username == "" {
		return []string{}
	}
	follows, err := redis.GetAllFollows(username)
	if err != nil {
		logs.LogError(err)
		return nil
	}

	return follows
}
