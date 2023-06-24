package api

import (
	"github.com/gin-gonic/gin"
	"my-app/api/information"
	"my-app/api/like"
	"my-app/api/middleWare"
	"my-app/api/smms"
	"my-app/api/user"
	"my-app/logs"
)

func InitRouter() {
	r := gin.Default()

	r.Use(middleWare.Cors())
	r.StaticFile("/", "./static/dist/index.html")
	r.Static("/assets", "./static/dist/assets")
	r.StaticFile("/vite.svg", "./static/dist/vite.svg")
	r.Static("/public", "./static/dist")

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
		p.GET("/like", like.Like)
		p.GET("/unlike", like.UnLike)
		p.GET("/likeNumber", like.NumberLike)
		p.GET("/isLike", like.IsLike)
		p.GET("/otherUserDetails", user.GetOtherUserDetails)
		p.GET("/follow", user.Follow)
		p.GET("/IsFollowed", user.IsFollowed)
		p.GET("/unfollow", user.Unfollow)
	}

	err := r.Run(":8000")
	if err != nil {
		logs.LogError(err)
	}

}
