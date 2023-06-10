package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"my-app/dao"
	"my-app/model"
	"my-app/utils"
)

func SaveUserDetails(c *gin.Context) {
	username := utils.GetUsername(c)
	if username == "" {
		log.Println("没登陆")
		utils.RespFail(c, "你都没登录")
		return
	}
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		utils.RespFail(c, "internal error")
		return
	}
	var user model.User

	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		log.Println(err)
		utils.RespFail(c, "internal error")
		return
	}

	err = dao.SaveUserDetails(username, user.Telephone, user.Address)
	if err != nil {
		log.Println(err)
		utils.RespFail(c, "internal error")
		return
	}

	utils.RespSuccess(c, "save successfully")

}
