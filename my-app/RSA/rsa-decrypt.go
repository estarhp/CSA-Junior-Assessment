package RSA

import (
	"github.com/wenzhenxi/gorsa"
	"my-app/logs"
)

func DecryptByPublicKey(encrypted string) (string, error) {
	result, err := gorsa.PublicDecrypt(encrypted, PublicKey)
	if err != nil {
		logs.LogError(err)
		return "", err
	}
	return result, nil
}

func DecryptByPrivateKey(encrypted string) (string, error) {
	result, err := gorsa.PriKeyEncrypt(encrypted, PrivateKey)
	if err != nil {
		logs.LogError(err)
		return "", err
	}
	return result, nil
}
