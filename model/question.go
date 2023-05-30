package model

import (
	"cubizy/plugins/db"
	"cubizy/util"
	"errors"

	"gorm.io/gorm"
)

// Question will all informetion of Questions
type Question struct {
	gorm.Model
	Question  string
	UpdatedBy uint
	CreatedBy uint
}

func init() {
	// Migrate the schema for user obj
	db.Conn.AutoMigrate(&Question{})
}

// Update will update product by given post argumnets
func (question *Question) Update(QuestionMap map[string]interface{}, UpdatedBy *User) (bool, error) {
	var err error
	var flag bool
	if UpdatedBy.ID == 0 {
		return flag, errors.New("invalid updater")
	}

	if _, ok := QuestionMap["Question"]; ok {
		Value := util.GetString(QuestionMap["Question"])
		if Value != question.Question {
			question.Question = Value
			flag = true
		}
	}

	if flag {
		if question.Question == "" {
			err = errors.New(" Question can not be empty")
		} else {
			question.UpdatedBy = UpdatedBy.ID
			if question.ID == 0 {
				question.CreatedBy = UpdatedBy.ID
				var duplicates Question
				err = db.Conn.First(&duplicates, " `question` = ?", question.Question).Error
				if err == nil {
					err = errors.New("duplicate question")
				} else {
					err = db.Conn.Create(&question).Error
				}
			} else {
				err = db.Conn.Save(&question).Error
			}
		}
	}
	return flag, err
}
