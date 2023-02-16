package modules

import (
	"github.com/vietthangc1/mini-web-golang/models"
	"gorm.io/gorm"
)

func AddUser(db *gorm.DB, newUser *models.User) (error) {
	err := db.Create(newUser).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(db *gorm.DB, userDelete *models.User, id uint) (error) {
	db.Where("id = ?", id).Delete(userDelete)
	return nil
}

func GetUserByEmail (db *gorm.DB, userQuery *models.User , email string) (error) {
	err := db.Preload("Products.Propertises").Where("email = ?", email).First(userQuery).Error

	if err != nil {
		return err
	}
	return nil
}