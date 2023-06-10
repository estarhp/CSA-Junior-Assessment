package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"my-app/dao"
	"my-app/model"
	"my-app/utils"
	"net/http"
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
		return
	}

	c.SetCookie("isLogin", user.Username+"+"+user.Password, 3600, "/", "127.0.0.1", false, true)
	utils.RespSuccess(c, "login successfully")
}

func AlreadyLogin(c *gin.Context) {
	username := utils.GetUsername(c)
	if username == "" {
		utils.RespSuccess(c, "false")
		return
	}
	is := dao.IsExist(username)

	if !is {
		utils.RespFail(c, "wrong cookie")
		return
	}

	utils.RespSuccess(c, "true")
}

func GetUserDetail(c *gin.Context) {
	username := utils.GetUsername(c)

	user, err := dao.GetUserDetails(username)
	if err != nil {
		utils.RespFail(c, "internal error")
		return
	}
	user.Password = ""
	c.JSON(http.StatusOK, gin.H{
		"userDetails": user,
	})
}
