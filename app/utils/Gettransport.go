package utils

import (
	"my-app/configs"
	"my-app/logs"
	"net/http"
	"net/url"
)

func GetTransport() *http.Transport {
	transport := &http.Transport{}
	if configs.Proxy.UseProxy {
		transport = &http.Transport{Proxy: func(req *http.Request) (*url.URL, error) {
			// 设置代理服务器地址和端口号
			proxyUrl, err := url.Parse(configs.Proxy.Proxy)
			if err != nil {
				logs.LogError(err)
				return nil, err
			}
			return proxyUrl, nil
		}}
	}

	return transport
}
