package main

import (
	"my-app/api"
	"my-app/configs"
	"my-app/dao"
	"my-app/logs"
)

func main() {
	logs.InitLogger()
	configs.InitViper()
	dao.InitDB()
	api.InitRouter()
}
