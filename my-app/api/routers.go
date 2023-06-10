package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"my-app/api/information"
	"my-app/api/middleWare"
	"my-app/api/smms"
	"my-app/api/user"
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
		p.GET("/login", user.Login)
		p.POST("/register", user.Register)
		p.GET("/isLogin", user.AlreadyLogin)
		p.POST("/addQuestion", information.AddQuestion)
		p.GET("/getAllQuestions", information.GetALlQuestions)
		p.POST("/addComment", information.AddComment)
		p.GET("/getAllComments", information.GetAllComment)
		p.GET("/userDetails", user.GetUserDetail)
		p.POST("/deleteQuestion", information.DeleteQuestion)
		p.POST("/logoff", user.Logoff)
		p.POST("/updateComment", information.UpdateComment)
		p.POST("/deleteComment", information.DeleteComment)
		p.POST("/editQuestion", information.EditQuestion)
		p.GET("/getQuestion", information.GetQuestion)
		p.GET("/getUserQuestion", information.GetUserQuestion)
		p.GET("/getUserComments", information.GetUserComments)
		p.POST("/upload", smms.UploadAvatarImage)
		p.POST("/saveUserDetails", user.SaveUserDetails)
	}

	err := r.Run(":8000")
	if err != nil {
		log.Panicln(err)
	}

}
