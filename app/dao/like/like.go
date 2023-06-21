package like

import (
	"fmt"
	"github.com/go-redis/redis"
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
	_, err := client.SAdd(key, userId).Result()
	if err != nil {
		return err
	}

	return nil

}

// 取消点赞
func Unlike(userId string, questionID string) error {
	key := fmt.Sprintf("article:%s:likes", questionID)
	_, err := client.SRem(key, userId).Result()
	if err != nil {
		return err
	}

	return nil

}

// 获取文章点赞数
func GetLikes(questionID string) (int64, error) {
	key := fmt.Sprintf("article:%s:likes", questionID)
	return client.SCard(key).Result()
}

func IsLike(username string, questionID string) (bool, error) {
	key := fmt.Sprintf("article:%s:likes", questionID)
	result, err := client.SIsMember(key, username).Result()
	if err != nil {
		return false, err
	}

	return result, err
}

func DeleteAllLike(questionID string) error {
	key := fmt.Sprintf("article:%s:likes", questionID)

	// 获取当前集合中的所有成员
	members, err := client.SMembers(key).Result()
	if err != nil {
		return err
	}

	// 逐个删除集合中的成员
	for _, member := range members {
		err := client.SRem(key, member).Err()
		if err != nil {
			return err
		}
	}
	return nil
}
