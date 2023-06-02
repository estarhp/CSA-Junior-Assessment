package dao

import (
	"my-app/model"
	"my-app/utils"
)

func AddQuestion(title string, details string, username string) error {
	question := model.Question{
		Title:    title,
		Details:  details,
		Username: username,
		ID:       utils.GenerateID(),
		Date:     utils.FormatTime(),
	}

	result := DB.Create(&question)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetALlQuestions() ([]model.Question, error) {
	//var question model.Question
	var questions []model.Question
	result := DB.Order("Date desc").Find(&questions)

	if result.Error != nil {
		return questions, result.Error
	}

	return questions, nil
}
