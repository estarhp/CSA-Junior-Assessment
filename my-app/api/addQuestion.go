package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"my-app/dao"
	"my-app/utils"
	"net/http"
	"strings"
)

func AddQuestion(c *gin.Context) {

	title := c.PostForm("title")
	details := c.PostForm("details")

	isLogin, err := c.Cookie("isLogin")

	if err != nil {
		log.Println(err)
		utils.RespFail(c, "you does not login")
		return
	}

	username := strings.Split(isLogin, "+")[0]

	err = dao.AddQuestion(title, details, username)
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
