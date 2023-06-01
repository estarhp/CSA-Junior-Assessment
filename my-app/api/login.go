package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"my-app/dao"
	"my-app/model"
	"my-app/utils"
	"strings"
)

func Login(c *gin.Context) {
	var user model.User
	user.Username = c.Query("username")
	user.Password = c.Query("password")

	if isexist := dao.IsExist(user.Username); !isexist {
		utils.RespFail(c, "the user does not exist")
		return
	}

	password, err := dao.QueryPassword(user.Username)
	if err != nil {
		log.Fatalln(err)
		return
	}
	if password != user.Password {
		utils.RespFail(c, "the password is wrong")
	}

	c.SetCookie("isLogin", user.Username+"+"+user.Password, 3600, "/", "127.0.0.1", false, true)
	utils.RespSuccess(c, "login successfully")
}

func AlreadyLogin(c *gin.Context) {
	isLogin, err := c.Cookie("isLogin")
	if err != nil {
		utils.RespSuccess(c, "false")
		return
	}
	arr := strings.Split(isLogin, "+")

	username := arr[0]

	is := dao.IsExist(username)

	if !is {
		utils.RespFail(c, "wrong cookie")
		return
	}

	utils.RespSuccess(c, "true")
}
