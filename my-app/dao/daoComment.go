package dao

import (
	"my-app/logs"
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
		logs.LogError(result.Error)
		return result.Error
	}
	return nil
}

func UpdateComment(ID string, newComment string) error {
	var comment model.Comment
	result := DB.Where("ID = ?", ID).Find(&comment)

	if result.Error != nil {
		return result.Error
	}
	comment.Content = newComment
	result = DB.Save(&comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteComment(ID string) error {
	var comment model.Comment
	result := DB.Where("ID = ?", ID).Delete(&comment)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetUserComments(username string) (map[string][]model.Comment, error) {
	var comments []model.Comment
	var commentsMap = make(map[string][]model.Comment)

	result := DB.Where("Username = ?", username).Find(&comments)
	if result.Error != nil {
		return commentsMap, result.Error
	}

	commentsMap = utils.SortComments(comments)

	return commentsMap, nil
}
