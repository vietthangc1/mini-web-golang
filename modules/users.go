package modules

import (
	"github.com/vietthangc1/mini-web-golang/models"
	"gorm.io/gorm"
)

type UserRepository struct{
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) AddUser(newUser models.User) (models.User, error) {
	err := r.db.Create(&newUser).Error
	if err != nil {
		return models.User{}, err
	}
	return newUser, nil
}

func (r *UserRepository) DeleteUser(id uint) (models.User, error) {
	var userDelete models.User
	r.db.Where("id = ?", id).Delete(&userDelete)
	return userDelete, nil
}

func (r *UserRepository) GetUserByEmail (email string) (models.User, error) {
	var userQuery models.User
	err := r.db.Preload("Products.Propertises").Where("email = ?", email).First(&userQuery).Error

	if err != nil {
		return models.User{}, err
	}
	return userQuery, nil
}