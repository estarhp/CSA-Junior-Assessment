package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"my-app/dao"
	"my-app/utils"
	"net/http"
)

func AddQuestion(c *gin.Context) {

	title := c.PostForm("title")
	details := c.PostForm("details")

	username := utils.GetUsername(c)

	if username == "" {
		utils.RespFail(c, "you does not login")
		return
	}

	err := dao.AddQuestion(title, details, username)
	if err != nil {
		log.Println(err)
		utils.RespFail(c, "add the question error")
		return
	}

	utils.RespSuccess(c, "add the question successfully")

}

func GetALlQuestions(c *gin.Context) {

	questions, err := dao.GetALlQuestions()

	if err != nil {
		log.Println(err)
		utils.RespFail(c, "get all questions error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    200,
		"questions": questions,
	})
}

func deleteQuestion(c *gin.Context) {
	ID := c.PostForm("ID")
	username := utils.GetUsername(c)
	err := dao.DeleteQuestion(username, ID)

	if err != nil {
		log.Println(err)
		utils.RespFail(c, "wrongly delete")
		return
	}
	beID := ID
	err = deleteComments(beID)

	if err != nil {
		utils.RespFail(c, "internal err")
		return
	}

	if err != nil {
		utils.RespFail(c, "wrongly delete")
		return
	}

	utils.RespSuccess(c, "delete successfully")
}
func editQuestion(c *gin.Context) {
	title := c.PostForm("title")
	details := c.PostForm("details")
	ID := c.PostForm("ID")

	err := dao.EditQuestion(title, details, ID)
	if err != nil {
		log.Println(err)
		utils.RespFail(c, "internal error")
		return
	}

	utils.RespSuccess(c, "修改问题成功")

}

func getQuestion(c *gin.Context) {
	ID := c.Query("ID")

	question, err := dao.GetQuestion(ID)
	if err != nil {
		log.Println(err)
		utils.RespFail(c, "internal error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"question": question,
		"status":   200,
	})
}

func getUserQuestion(c *gin.Context) {
	username := utils.GetUsername(c)

	questions, err := dao.GetUserQuestion(username)

	if err != nil {
		log.Println(err)
		utils.RespFail(c, "internal err")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"questions": questions,
		"status":    200,
	})
}

func getUserComments(c *gin.Context) {
	username := utils.GetUsername(c)

	comments, err := dao.GetUserComments(username)

	if err != nil {
		log.Println(err)
		utils.RespFail(c, "internal error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"comments": comments,
	})

}
