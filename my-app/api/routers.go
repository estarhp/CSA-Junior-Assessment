package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"my-app/api/middleWare"
)

func InitRouter() {
	r := gin.Default()

	r.Use(middleWare.Cors())

	r.GET("login", Login)
	r.POST("register", Register)
	r.GET("isLogin", AlreadyLogin)
	r.POST("addQuestion", AddQuestion)
	r.GET("getAllQuestions", GetALlQuestions)
	r.POST("addComment", AddComment)
	r.GET("getAllComments", GetAllComment)
	r.GET("userDetails", getUserDetail)
	r.POST("deleteQuestion", deleteQuestion)
	r.POST("logoff", Logoff)

	err := r.Run(":8000")
	if err != nil {
		log.Panicln(err)
	}

}
