package user

import (
	"github.com/gin-gonic/gin"
	"my-app/RSA"
	"my-app/dao"
	"my-app/logs"
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
		logs.LogError(err)
		return
	}
	if password != user.Password {
		utils.RespFail(c, "the password is wrong")
		return
	}
	encrypted, err := RSA.EncryptByPrivateKey(user.Username + "+" + user.Password)
	if err != nil {
		logs.LogError(err)
		utils.RespFail(c, "internal error")
		return
	}
	c.SetCookie("isLogin", encrypted, 3600, "/", "127.0.0.1", false, true)

	logs.LogSuccess(user.Username + " login successfully")
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
		logs.LogError(err)
		utils.RespFail(c, "internal error")
		return
	}
	user.Password = ""
	c.JSON(http.StatusOK, gin.H{
		"userDetails": user,
	})
}
