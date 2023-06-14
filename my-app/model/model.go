package model

type User struct {
	Username    string
	Password    string
	AvatarImage string
	ImageHash   string
	Telephone   string
	Address     string
}

type Question struct {
	Title    string
	Details  string
	Username string
	ID       string
	Date     string
}

type Comment struct {
	BeID       string
	ID         string
	Username   string
	Date       string
	Content    string
	QuestionID string
	BeUsername string
}

type SmResult struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		FileID    int    `json:"file_id"`
		Width     int    `json:"width"`
		Height    int    `json:"height"`
		Filename  string `json:"filename"`
		Storename string `json:"storename"`
		Size      int    `json:"size"`
		Path      string `json:"path"`
		Hash      string `json:"hash"`
		URL       string `json:"url"`
		Delete    string `json:"delete"`
		Page      string `json:"page"`
	} `json:"data"`
	RequestID string `json:"RequestId"`
}

type DeleteResult struct {
	Success   bool   `json:"success"`
	Code      string `json:"code"`
	Message   string `json:"message"`
	Data      []any  `json:"data"`
	RequestID string `json:"RequestId"`
}

type Proxy struct {
	Proxy    string
	UseProxy bool
}
