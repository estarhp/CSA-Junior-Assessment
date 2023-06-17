package information

import (
	"github.com/gin-gonic/gin"
	"my-app/cache"
	"my-app/dao"
	"my-app/logs"
	"my-app/model"
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
		logs.LogError(err)
		utils.RespFail(c, "add the question error")
		return
	}
	logs.LogSuccess(username + " add the question successfully")
	utils.RespSuccess(c, "add the question successfully")
	err = cache.DeleteCache("allQuestions")
	if err != nil {
		logs.LogError(err)
	}

}

func GetALlQuestions(c *gin.Context) {
	var questions []model.Question
	err := cache.GetCache("allQuestions", &questions)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":    200,
			"questions": questions,
		})
		return
	}

	questions, err = dao.GetALlQuestions()

	if err != nil {
		logs.LogError(err)
		utils.RespFail(c, "get all questions error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    200,
		"questions": questions,
	})
	err = cache.SetCache("allQuestions", &questions)
	if err != nil {
		logs.LogError(err)
		return
	}
}

func DeleteQuestion(c *gin.Context) {
	ID := c.PostForm("ID")
	username := utils.GetUsername(c)
	err := dao.DeleteQuestion(username, ID)

	if err != nil {
		logs.LogError(err)
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
	logs.LogSuccess(username + " delete the question successfully whose id is" + ID)
	utils.RespSuccess(c, "delete successfully")
	err = cache.DeleteCache("allQuestions")
	if err != nil {
		logs.LogError(err)
	}
}
func EditQuestion(c *gin.Context) {
	title := c.PostForm("title")
	details := c.PostForm("details")
	ID := c.PostForm("ID")

	err := dao.EditQuestion(title, details, ID)
	if err != nil {
		logs.LogError(err)
		utils.RespFail(c, "internal error")
		return
	}
	logs.LogSuccess(" edit the question successfully whose id is" + ID)
	utils.RespSuccess(c, "修改问题成功")
	err = cache.DeleteCache("allQuestions")
	if err != nil {
		logs.LogError(err)
	}

}

func GetQuestion(c *gin.Context) {
	ID := c.Query("ID")

	question, err := dao.GetQuestion(ID)
	if err != nil {
		logs.LogError(err)
		utils.RespFail(c, "internal error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"question": question,
		"status":   200,
	})
}

func GetUserQuestion(c *gin.Context) {
	username := utils.GetUsername(c)

	questions, err := dao.GetUserQuestion(username)

	if err != nil {
		logs.LogError(err)
		utils.RespFail(c, "internal err")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"questions": questions,
		"status":    200,
	})
}

func GetUserComments(c *gin.Context) {
	username := utils.GetUsername(c)

	comments, err := dao.GetUserComments(username)

	if err != nil {
		logs.LogError(err)
		utils.RespFail(c, "internal error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"comments": comments,
	})

}
