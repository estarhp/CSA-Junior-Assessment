package model

type User struct {
	Username string
	Password string
}

type Question struct {
	Title    string
	Details  string
	Username string
	ID       string
	Date     string
}

type Comment struct {
	beID     string
	ID       string
	Username string
	Date     string
	Content  string
}
