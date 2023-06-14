package information

import (
	"github.com/gin-gonic/gin"
	"log"
	"my-app/dao"
	"my-app/logs"
	"my-app/utils"
	"net/http"
)

func AddComment(c *gin.Context) {
	username := utils.GetUsername(c)
	if username == "" {
		utils.RespFail(c, "你还没有登录不能参与评论")
		return
	}
	is := dao.IsExist(username)
	if !is {
		utils.RespFail(c, "cookie 错误")
		return
	}
	content := c.PostForm("content")
	beID := c.PostForm("beID")
	questionID := c.PostForm("questionID")
	beUsername := c.PostForm("beUsername")
	err := dao.AddComment(username, content, beID, questionID, beUsername)
	if err != nil {
		logs.LogError(err)
		utils.RespFail(c, "internal err")
		return
	}
	logs.LogSuccess(username + " add comment successfully")
	utils.RespSuccess(c, "添加评论成功")

}

func GetAllComment(c *gin.Context) {
	ID := c.Query("beID")

	comments, err := dao.GetAllComment(ID)

	if err != nil {
		logs.LogError(err)
		utils.RespFail(c, "internal error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   200,
		"comments": comments,
	})

}

func deleteComments(questionID string) error {
	err := dao.DeleteComments(questionID)
	if err != nil {
		logs.LogError(err)
		return err
	}
	logs.LogSuccess("the question of" + questionID + " was deleted")
	return nil
}

func UpdateComment(c *gin.Context) {
	username := utils.GetUsername(c)

	if username == "" {
		utils.RespFail(c, "你还未登录")
		return
	}
	ID := c.PostForm("ID")
	newComment := c.PostForm("content")
	err := dao.UpdateComment(ID, newComment)
	if err != nil {
		logs.LogError(err)
		utils.RespFail(c, "internal error")
		return
	}
	logs.LogSuccess("the comment of" + ID + " was updated")
	utils.RespSuccess(c, "修改成功")

}

func DeleteComment(c *gin.Context) {
	ID := c.PostForm("ID")

	err := dao.DeleteComment(ID)

	if err != nil {
		log.Println(err)
		utils.RespFail(c, "internal error")
		return
	}
	logs.LogSuccess("the comment of" + ID + " was deleted")
	utils.RespSuccess(c, "删除成功")

}
