package repository

import (
	"github.com/vietthangc1/mini-web-golang/models"
)

type UserService interface {
	AddUser(newUser models.User) (models.User, error)
	DeleteUser(id uint) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
}


