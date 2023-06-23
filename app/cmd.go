package main

import (
	"my-app/api"
	"my-app/cache"
	"my-app/configs"
	"my-app/dao"
	"my-app/dao/redis"
	"my-app/logs"
)

func main() {
	logs.InitLogger()
	configs.InitViper()
	dao.InitDB()
	redis.InitLike()
	cache.InitCache()
	api.InitRouter()
}
