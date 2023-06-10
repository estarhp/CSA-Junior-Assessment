package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"my-app/dao"
	"my-app/model"
	"my-app/utils"
)

func Register(c *gin.Context) {
	var user model.User
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	log.Println(user)

	is := dao.IsExist(user.Username)

	if is {
		utils.RespFail(c, "the user already exists")
		return
	}

	err := dao.AddUser(user)
	if err != nil {
		utils.RespFail(c, "the internal error")
		return
	}
	utils.RespSuccess(c, "register successfully")
}

func Logoff(c *gin.Context) {
	c.SetCookie("isLogin", "", -1, "/", "", false, true)
	utils.RespSuccess(c, "log off successfully")
}
