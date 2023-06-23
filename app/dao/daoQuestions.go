package dao

import (
	"my-app/dao/redis"
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

func DeleteQuestion(username, ID string) error {
	var question model.Question
	question.ID = ID

	result := DB.Where("Username = ?", username).Delete(&question)
	if result.Error != nil {
		return result.Error
	}
	err := redis.DeleteAllLike(question.ID)
	if err != nil {
		return err
	}
	return nil
}

func EditQuestion(title string, details string, ID string) error {
	var question model.Question

	result := DB.Where("ID = ?", ID).Find(&question)

	if result.Error != nil {
		return result.Error
	}

	question.Title = title
	question.Details = details

	result = DB.Save(&question)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetQuestion(ID string) (model.Question, error) {
	var question model.Question
	result := DB.Where("ID = ?", ID).Find(&question)
	if result.Error != nil {
		return question, result.Error
	}

	return question, nil
}

func GetUserQuestion(username string) ([]model.Question, error) {
	var questions []model.Question

	result := DB.Where("Username = ?", username).Find(&questions)

	if result.Error != nil {
		return questions, result.Error
	}

	return questions, nil
}
