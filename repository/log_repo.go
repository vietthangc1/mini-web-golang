package repository

import "github.com/vietthangc1/mini-web-golang/models"

type LogRepo interface {
	AddLog(newLog models.Log) (models.Log, error)
}