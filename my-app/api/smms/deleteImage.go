package smms

import (
	"context"
	"github.com/carlmjohnson/requests"
	"my-app/logs"
	"my-app/model"
	"my-app/utils"
)

func deleteImage(hash string) error {
	var s model.DeleteResult
	transport := utils.GetTransport()

	err := requests.URL(Baseurl).
		Pathf("delete/%s", hash).
		Header("Authorization", Token).
		Transport(transport).
		ToJSON(&s).
		Fetch(context.Background())

	logs.LogSuccess(s.Message)

	if err != nil {
		return err
	}

	return nil
}
