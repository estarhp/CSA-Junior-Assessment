package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"my-app/api/middleWare"
)

func InitRouter() {
	r := gin.Default()

	r.Use(middleWare.Cors())
	r.StaticFile("/", "./static/dist/index.html")
	r.Static("/assets", "./static/dist/assets")
	r.StaticFile("/favicon.ico", "./static/dist/favicon.ico")
	// 其他静态资源
	r.Static("/public", "./api/static")

	p := r.Group("api")
	{
		p.GET("/login", Login)
		p.POST("/register", Register)
		p.GET("/isLogin", AlreadyLogin)
		p.POST("/addQuestion", AddQuestion)
		p.GET("/getAllQuestions", GetALlQuestions)
		p.POST("/addComment", AddComment)
		p.GET("/getAllComments", GetAllComment)
		p.GET("/userDetails", getUserDetail)
		p.POST("/deleteQuestion", deleteQuestion)
		p.POST("/logoff", Logoff)
		p.POST("/updateComment", updateComment)
		p.POST("/deleteComment", deleteComment)
		p.POST("/editQuestion", editQuestion)
		p.GET("/getQuestion", getQuestion)
		p.GET("/getUserQuestion", getUserQuestion)
		p.GET("/getUserComments", getUserComments)
		p.POST("/upload", uploadAvatarImage)
	}

	err := r.Run(":8000")
	if err != nil {
		log.Panicln(err)
	}

}
