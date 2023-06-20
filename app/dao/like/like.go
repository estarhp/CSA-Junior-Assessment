package like

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var client *redis.Client

func InitLike() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // Redis数据库没有密码
		DB:       0,  // 默认数据库为0
	})
}

func Like(userId string, questionID string) error {
	key := fmt.Sprintf("article:%s:likes", questionID)
	_, err := client.SAdd(context.TODO(), key, userId).Result()
	if err != nil {
		return err
	}

	return nil

}

// 取消点赞
func Unlike(userId string, questionID string) error {
	key := fmt.Sprintf("article:%s:likes", questionID)
	_, err := client.SRem(context.TODO(), key, userId).Result()
	if err != nil {
		return err
	}

	return nil

}

// 获取文章点赞数
func GetLikes(questionID string) (int64, error) {
	key := fmt.Sprintf("article:%s:likes", questionID)
	return client.SCard(context.TODO(), key).Result()
}

func IsLike(username string, questionID string) (bool, error) {
	key := fmt.Sprintf("article:%s:likes", questionID)
	result, err := client.SIsMember(context.TODO(), key, username).Result()
	if err != nil {
		return false, err
	}

	return result, err
}
