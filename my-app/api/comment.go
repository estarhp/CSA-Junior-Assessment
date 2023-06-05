package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"my-app/dao"
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
	err := dao.AddComment(username, content, beID, questionID)
	if err != nil {
		log.Println(err)
		utils.RespFail(c, "internal err")
		return
	}

	utils.RespSuccess(c, "添加评论成功")

}

func GetAllComment(c *gin.Context) {
	ID := c.Query("beID")

	comments, err := dao.GetAllComment(ID)

	if err != nil {
		utils.RespFail(c, "internal error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   200,
		"comments": comments,
	})

}

func deleteComments(beID string) error {
	err := dao.DeleteComments(beID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
