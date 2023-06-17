package smms

import (
	"github.com/gin-gonic/gin"
	"my-app/dao"
	"my-app/logs"
	"my-app/utils"
)

var Baseurl = "https://sm.ms/api/v2/"
var Token = "2QFEYPIYapEV0XW4MgCiysJBbahMsN5s"

func UploadAvatarImage(c *gin.Context) {

	username := utils.GetUsername(c)
	if username == "" {
		utils.RespFail(c, "还未登录噢")
		return
	}

	result, err := setRequest(c)
	if err != nil {
		logs.LogError(err)
		utils.RespFail(c, "internal error")
		return
	}

	if result.Success == false {
		logs.LogSuccess(result.Message)
		utils.RespFail(c, result.Message)
		return
	}

	logs.LogSuccess(result.Message)

	hash, _ := dao.GetImageHash(username)
	if hash != "" {
		err := deleteImage(hash)
		if err != nil {
			logs.LogError(err)
		}
	}

	err = dao.SaveAvatar(username, result.Data.URL, result.Data.Hash)
	if err != nil {
		logs.LogError(err)
		utils.RespFail(c, "internal error")
		return
	}

	utils.RespSuccess(c, "成功上传头像")

}
