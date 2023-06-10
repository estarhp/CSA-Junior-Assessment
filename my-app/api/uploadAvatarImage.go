package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/carlmjohnson/requests"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"my-app/dao"
	"my-app/model"
	"my-app/utils"
	"net/http"
	"net/url"
	"path/filepath"
)

var Baseurl = "https://sm.ms/api/v2/"
var Token = "2QFEYPIYapEV0XW4MgCiysJBbahMsN5s"

func uploadAvatarImage(c *gin.Context) {
	req, err := setRequest(c)
	if err != nil {
		return
	}
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: func(req *http.Request) (*url.URL, error) {
				// 设置代理服务器地址和端口号
				proxyUrl, err := url.Parse("http://127.0.0.1:1080")
				if err != nil {
					log.Println(err)
					utils.RespFail(c, "代理错误")
					return nil, err
				}
				return proxyUrl, nil
			},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		utils.RespFail(c, "internal error")
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	fmt.Println(resp.Status)

	respbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		utils.RespFail(c, "internal error")
		return
	}
	var result model.SmResult
	err = json.Unmarshal(respbody, &result)
	if err != nil {
		log.Println(err)
		utils.RespFail(c, "internal error")
		return
	}
	log.Println(result)
	if result.Success == false {
		log.Println("重复的图片")
		utils.RespFail(c, "重复的图片")
		return
	}
	username := utils.GetUsername(c)
	if username == "" {
		utils.RespFail(c, "还未登录噢")
		return
	}
	hash, _ := dao.GetImageHash(username)
	if hash != "" {
		err := deleteImage(hash)
		if err != nil {
			log.Println(err)
		}
	}

	err = dao.SaveAvatar(username, result.Data.URL, result.Data.Hash)
	if err != nil {
		log.Println(err)
		utils.RespFail(c, "internal error")
		return
	}
	utils.RespSuccess(c, "成功上传头像")

}
func deleteImage(hash string) error {
	var s string
	err := requests.URL(Baseurl).
		Pathf("delete/%s", hash).
		Header("Authorization", Token).
		ToString(&s).
		Fetch(context.Background())

	if err != nil {
		return err
	}

	return nil
}

func setRequest(c *gin.Context) (*http.Request, error) {
	file, err := c.FormFile("file")
	if err != nil {
		utils.RespFail(c, "file error")
		return nil, err
	}

	image, err := file.Open()
	if err != nil {
		log.Println(err)
		utils.RespFail(c, "failed to open file")
		return nil, err
	}

	//设置参数

	var body = new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("smfile", filepath.Base(file.Filename))
	if err != nil {
		log.Println(err)
		utils.RespFail(c, "error")
		return nil, err
	}

	_, err = io.Copy(part, image)

	if err != nil {
		log.Println(err)
		utils.RespFail(c, "failed to add file content")
		return nil, err
	}

	err = writer.Close()

	if err != nil {
		log.Println(err)
		utils.RespFail(c, "failed to close form data")
		return nil, err
	}

	req, err := http.NewRequest("POST", Baseurl+"upload", body)

	if err != nil {
		log.Println(err)
		utils.RespFail(c, "internal error")
		return nil, err
	}

	req.Header.Set("Authorization", Token)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	return req, nil
}
