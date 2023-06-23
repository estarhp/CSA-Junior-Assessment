package user

import (
	"github.com/gin-gonic/gin"
	"my-app/RSA"
	"my-app/dao"
	"my-app/logs"
	"my-app/model"
	"my-app/utils"
)

func Register(c *gin.Context) {
	var user model.User
	var err error
	user.Username, err = RSA.DecryptByPrivateKey(c.PostForm("username"))
	user.Password, err = RSA.DecryptByPrivateKey(c.PostForm("password"))
	logs.LogSuccess(user.Username)

	is := dao.IsExist(user.Username)

	if is {
		utils.RespFail(c, "the user already exists")
		return
	}

	err = dao.AddUser(user)
	if err != nil {
		logs.LogError(err)
		utils.RespFail(c, "the internal error")
		return
	}

	logs.LogSuccess(user.Username + " register successfully")
	utils.RespSuccess(c, "register successfully")
}
