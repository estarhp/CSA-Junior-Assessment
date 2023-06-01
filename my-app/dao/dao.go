package dao

import (
	"log"
	"my-app/model"
)

func IsExist(username string) bool {
	var user model.User

	sqlStr := "select username,password from users where username = ?"

	err := DB.QueryRow(sqlStr, username).Scan(&user.Username, &user.Password)

	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func AddUser(user model.User) error {

	sqlStr := "insert into users (username, password) values (?,?)"

	_, err := DB.Exec(sqlStr, user.Username, user.Password)

	if err != nil {
		log.Fatalln(err)
		return err
	}

	log.Println("add user successfully")
	return nil

}

func QueryPassword(username string) (string, error) {

	var password string
	sqlStr := "select password from users where username = ?"

	err := DB.QueryRow(sqlStr, username).Scan(&password)
	if err != nil {
		log.Fatalln(err)

		return "error", err
	}

	return password, nil

}
