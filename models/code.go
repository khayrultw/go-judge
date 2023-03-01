package models

import (
	"gorm.io/gorm"
)

type Code struct {
	Id            uint   `json:"id"`
	SourceCode    string `json:"sourceCode"`
	Language      string `json:"language"`
	ProblemNumber int8   `json:"problemNumber"`
}

func PostCode(db *gorm.DB, code *Code) (err error) {
	err = db.Create(code).Error
	if err != nil {
		return err
	}

	return nil
}

func GetCode(db *gorm.DB, code *Code, id int) (err error) {
	err = db.Where("id = ?", id).First(code).Error
	if err != nil {
		return err
	}

	return nil
}
