package logs

import (

	"github.com/vietthangc1/mini-web-golang/models"
	"github.com/vietthangc1/mini-web-golang/repository"
	"gorm.io/gorm"
)

type LogRepoImpl struct {
	db *gorm.DB
}

func NewLogRepo(db *gorm.DB) repository.LogRepo {
	return &LogRepoImpl{db: db}
}

func (l *LogRepoImpl) AddLog(newLog models.Log) (models.Log, error) {
	err := l.db.Create(&newLog).Error
	if err != nil {
		return models.Log{}, err
	}
	return newLog, nil
}

