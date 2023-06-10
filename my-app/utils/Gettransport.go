package utils

import (
	"log"
	"net/http"
	"net/url"
)

func GetTransport() *http.Transport {

	transport := &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
		// 设置代理服务器地址和端口号
		proxyUrl, err := url.Parse("http://127.0.0.1:1080")
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return proxyUrl, nil
	}}

	return transport
}
