package smms

import (
	"bytes"
	"context"
	"github.com/carlmjohnson/requests"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"mime/multipart"
	"my-app/model"
	"my-app/utils"
	"path/filepath"
)

func setRequest(c *gin.Context) (model.SmResult, error) {

	var result model.SmResult
	file, err := c.FormFile("file")
	if err != nil {
		utils.RespFail(c, "file error")
		return result, err
	}

	image, err := file.Open()
	if err != nil {
		log.Println(err)
		utils.RespFail(c, "failed to open file")
		return result, err
	}

	//设置参数

	var body = new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("smfile", filepath.Base(file.Filename))
	if err != nil {
		log.Println(err)
		utils.RespFail(c, "error")
		return result, err
	}

	_, err = io.Copy(part, image)

	if err != nil {
		log.Println(err)
		utils.RespFail(c, "failed to add file content")
		return result, err
	}

	err = writer.Close()

	if err != nil {
		log.Println(err)
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
		return result, err
	}

	return result, nil
}
