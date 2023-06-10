package smms

import (
	"context"
	"github.com/carlmjohnson/requests"
	"log"
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

	log.Println(s.Message)

	if err != nil {
		return err
	}

	return nil
}
