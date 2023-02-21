package users

import (
	"github.com/vietthangc1/mini-web-golang/models"
	"github.com/vietthangc1/mini-web-golang/repository"
	"gorm.io/gorm"
)

type UserRepoImpl struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) repository.UserRepo {
	return &UserRepoImpl{db: db}
}

func (r *UserRepoImpl) AddUser(newUser models.User) (models.User, error) {
	err := r.db.Create(&newUser).Error
	if err != nil {
		return models.User{}, err
	}
	return newUser, nil
}

func (r *UserRepoImpl) DeleteUser(id uint) (models.User, error) {
	var userDelete models.User
	r.db.Where("id = ?", id).Delete(&userDelete)
	return userDelete, nil
}

func (r *UserRepoImpl) GetUserByEmail(email string) (models.User, error) {
	var userQuery models.User
	err := r.db.Preload("Products.Propertises").Where("email = ?", email).First(&userQuery).Error

	if err != nil {
		return models.User{}, err
	}
	return userQuery, nil
}
