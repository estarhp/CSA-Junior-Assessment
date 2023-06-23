package redis

import "github.com/go-redis/redis"

var client *redis.Client

func InitLike() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // Redis数据库没有密码
		DB:       0,  // 默认数据库为0
	})
}
