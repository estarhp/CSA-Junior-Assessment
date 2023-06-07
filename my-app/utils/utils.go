package utils

import (
	"github.com/gin-gonic/gin"
	"my-app/model"
	"net/http"
)

func RespSuccess(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": message,
	})
}

func RespFail(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"status":  500,
		"message": message,
	})
}

func SortComments(comments []model.Comment) map[string][]model.Comment {
	var commentsMap = make(map[string][]model.Comment)
	for i := 0; i < len(comments); i++ {
		_, ok := commentsMap[comments[i].QuestionID]
		if ok {
			commentsMap[comments[i].QuestionID] = append(commentsMap[comments[i].QuestionID], comments[i])
		} else {
			commentsMap[comments[i].QuestionID] = []model.Comment{comments[i]}
		}
	}

	return commentsMap

}
