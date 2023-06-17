package RSA

import (
	"github.com/wenzhenxi/gorsa"
	"my-app/logs"
)

const PublicKey string = "-----BEGIN PUBLIC KEY-----\nMFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAMz4u0O0AI4TfvilTlJxQedoZgA8NuEC\nM/N16k3Pb2mObSpPGDIYdkqm1w3mjUGqK+Sld0Za2JEGAHrN1w+VvVECAwEAAQ==\n-----END PUBLIC KEY-----\n"
const PrivateKey string = "-----BEGIN PRIVATE KEY-----\nMIIBVAIBADANBgkqhkiG9w0BAQEFAASCAT4wggE6AgEAAkEAzPi7Q7QAjhN++KVO\nUnFB52hmADw24QIz83XqTc9vaY5tKk8YMhh2SqbXDeaNQaor5KV3RlrYkQYAes3X\nD5W9UQIDAQABAkAxCYTLssG7O+DRncK6KIxqz2gvwDgk5sEFCv3ONcVizBlBSXnF\nz5xc616oHLJNHslpel2CzZmdHcNJeMOglkQBAiEA+ZfjFYMK7Vfu3K4kSQfKx7nz\nCkwVql7SdvHGHSOm6cECIQDSO6ECQS2qLrG3FIxcySVH6m9vEJaVOkssSODZQ8sX\nkQIhALn1TB9u7uk+ppyMskQnJhIAnO+DGHFDDJPKNsznDykBAiAGJJ3ozfTCo0io\nG96aG3qOZmhJK4fq5mAp9Bs13ghCMQIgMM2CNIuYzOFByxXlw5q3pd1tDDFFG+E0\nXm0/thtRlFk=\n-----END PRIVATE KEY-----"

func EncryptByPublicKey(unencrypted string) (string, error) {
	result, err := gorsa.PublicEncrypt(unencrypted, PublicKey)
	if err != nil {
		logs.LogError(err)
		return "", err
	}
	return result, nil
}

func EncryptByPrivateKey(unencrypted string) (string, error) {
	result, err := gorsa.PriKeyEncrypt(unencrypted, PrivateKey)
	if err != nil {
		logs.LogError(err)
		return "", err
	}
	return result, nil
}
