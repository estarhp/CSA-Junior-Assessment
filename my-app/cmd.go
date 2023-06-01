package main

import (
	"my-app/api"
	"my-app/dao"
)

func main() {
	dao.InitDB()
	api.InitRouter()
}
