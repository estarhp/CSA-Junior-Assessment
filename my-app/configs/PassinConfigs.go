package configs

import (
	"github.com/spf13/viper"
	"my-app/logs"
	"my-app/model"
)

var Proxy model.Proxy

func InitViper() {
	viper.SetConfigFile("./configs/proxy/proxy.json")
	//viper.SetConfigName("proxy")
	//viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil { // 处理读取配置文件的错误
		logs.LogError(err)
	}

	Proxy.Proxy = viper.GetString("proxy")
	Proxy.UseProxy = viper.GetBool("use_proxy")
}
