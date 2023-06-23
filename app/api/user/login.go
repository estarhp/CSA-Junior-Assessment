package user

import (
	"github.com/gin-gonic/gin"
	"my-app/RSA"
	"my-app/dao"
	"my-app/logs"
	"my-app/model"
	"my-app/utils"
)

func Login(c *gin.Context) {
	var user model.User
	var err error
	user.Username, err = RSA.DecryptByPrivateKey(c.Query("username"))
	user.Password, err = RSA.DecryptByPrivateKey(c.Query("password"))

	if err != nil {
		logs.LogError(err)
		utils.RespFail(c, "internal error")
		return
	}

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

func Logoff(c *gin.Context) {
	c.SetCookie("isLogin", "", -1, "/", "", false, true)
	utils.RespSuccess(c, "log off successfully")
}
