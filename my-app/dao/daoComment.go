package dao

import (
	"log"
	"my-app/model"
	"my-app/utils"
)

func AddComment(username string, content string, beID string, questionID string, beUsername string) error {
	comment := model.Comment{
		ID:         utils.GenerateID(),
		Username:   username,
		Date:       utils.FormatTime(),
		Content:    content,
		BeID:       beID,
		QuestionID: questionID,
		BeUsername: beUsername,
	}
	result := DB.Create(&comment)

	if result.Error != nil {
		return result.Error
	}

	return nil

}

func GetAllComment(ID string) ([]model.Comment, error) {
	var comments []model.Comment

	result := DB.Where("question_id = ?", ID).Order("Date desc").Find(&comments)

	if result.Error != nil {
		return comments, result.Error
	}

	return comments, nil

}

func DeleteComments(questionID string) error {
	var comment model.Comment
	result := DB.Where("question_id = ?", questionID).Delete(&comment)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}
