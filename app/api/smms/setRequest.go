package smms

import (
	"bytes"
	"context"
	"github.com/carlmjohnson/requests"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"my-app/logs"
	"my-app/model"
	"my-app/utils"
	"path/filepath"
)

func setRequest(c *gin.Context) (model.SmResult, error) {

	var result model.SmResult
	file, err := c.FormFile("file")
	if err != nil {
		logs.LogError(err)
		utils.RespFail(c, "file error")
		return result, err
	}

	image, err := file.Open()
	if err != nil {
		logs.LogError(err)
		utils.RespFail(c, "failed to open file")
		return result, err
	}

	//设置参数

	var body = new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("smfile", filepath.Base(file.Filename))
	if err != nil {
		logs.LogError(err)
		utils.RespFail(c, "error")
		return result, err
	}

	_, err = io.Copy(part, image)

	if err != nil {
		logs.LogError(err)
		utils.RespFail(c, "failed to add file content")
		return result, err
	}

	err = writer.Close()

	if err != nil {
		logs.LogError(err)
		utils.RespFail(c, "failed to close form data")
		return result, err
	}

	transport := utils.GetTransport()
	err = requests.
		URL(Baseurl+"upload").
		Header("Authorization", Token).
		ContentType(writer.FormDataContentType()).
		BodyReader(body).
		ToJSON(&result).
		Transport(transport).
		Fetch(context.Background())
	if err != nil {
		logs.LogError(err)
		return result, err
	}
	//urlString := result.Data.URL
	//
	//// 解析URL
	//u, err := url.Parse(urlString)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// 获取文件名
	//fileName := path.Base(u.Path)
	//
	//// 去掉后缀名
	//nameWithoutExtension := strings.TrimSuffix(fileName, path.Ext(fileName))
	//
	//newURL := "https://smms.app/image/" + nameWithoutExtension
	//result.Data.URL = newURL
	return result, nil
}
