package redis

import (
	"fmt"
)

func Follow(username string, followed string) error {
	key := fmt.Sprintf("follow:%s:follows", username)
	_, err := client.SAdd(key, followed).Result()
	if err != nil {
		return err
	}
	return nil
}

func UnFollow(username string, followed string) error {
	key := fmt.Sprintf("follow:%s:follows", username)
	_, err := client.SRem(key, followed).Result()
	if err != nil {
		return err
	}
	return nil
}

func GetAllFollows(username string) ([]string, error) {
	key := fmt.Sprintf("follow:%s:follows", username)
	members, err := client.SMembers(key).Result()
	if err != nil {
		return members, err
	}
	return members, nil
}
