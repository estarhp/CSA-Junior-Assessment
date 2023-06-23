package user

import (
	"github.com/gin-gonic/gin"
	"my-app/dao/redis"
	"my-app/logs"
	"my-app/utils"
	"net/http"
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

func IsFollowed(c *gin.Context) {
	username := utils.GetUsername(c)

	if username == "" {
		utils.RespFail(c, "未登录")
		return
	}

	BeUsername := c.Query("BeUsername")

	followed, err := redis.IsFollowed(username, BeUsername)
	if err != nil {
		logs.LogError(err)
		utils.RespFail(c, "internal error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"IsFollowed": followed,
	})

}
