package like

import (
	"github.com/gin-gonic/gin"
	"my-app/dao/like"
	"my-app/logs"
	"my-app/utils"
	"net/http"
)

func Like(c *gin.Context) {
	username := utils.GetUsername(c)
	if username == "" {
		utils.RespFail(c, "你还未登录，不能参与带点赞")
		return
	}
	questionID := c.Query("questionID")
	err := like.Like(username, questionID)
	if err != nil {
		logs.LogError(err)
		utils.RespFail(c, "internal error")
		return
	}
	utils.RespSuccess(c, "点赞成功")
}
func UnLike(c *gin.Context) {
	username := utils.GetUsername(c)
	if username == "" {
		utils.RespFail(c, "你还未登录，不能参与带点赞")
		return
	}
	questionID := c.Query("questionID")
	err := like.Unlike(username, questionID)
	if err != nil {
		logs.LogError(err)
		utils.RespFail(c, "internal error")
		return
	}
	utils.RespSuccess(c, "取消点赞成功")
}
func NumberLike(c *gin.Context) {
	questionID := c.Query("questionID")
	likes, err := like.GetLikes(questionID)
	if err != nil {
		logs.LogError(err)
		utils.RespFail(c, "internal error")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"likes": likes,
	})
}

func IsLike(c *gin.Context) {
	username := utils.GetUsername(c)
	if username == "" {
		utils.RespFail(c, "你还没有登录")
		return
	}
	questionID := c.Query("questionID")

	isLike, err := like.IsLike(username, questionID)
	if err != nil {
		logs.LogError(err)
		utils.RespFail(c, "internal error")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"isLike": isLike,
	})
}
