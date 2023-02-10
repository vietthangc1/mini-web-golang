package modules

import (
	"database/sql"
	"fmt"

	"github.com/vietthangc1/mini-web-golang/models"
	"gorm.io/gorm"
)

func QueryAddUser(db *sql.DB, q string, u models.User) (models.User, error) {
	stmt, err := db.Prepare(q)
	if err != nil {
		return models.User{}, err
	}

	res, err := stmt.Exec(
		u.ID,
		u.Email,
		u.Password,
	)
	if err != nil {
		return models.User{}, err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return models.User{}, err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return models.User{}, err
	}
	fmt.Printf("ID = %d, affected = %d\n", lastId, rowCnt)

	return u, nil

}

func QueryDeleteUser(db *sql.DB, q, id string) (error) {
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(
		id,
	)
	if err != nil {
		return err
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
	return nil
}

func QueryGetUserByEmail (db *sql.DB, q string, email string) (models.User, error) {
	var userQuery models.User 

	err := db.QueryRow(q, email).Scan(
		&userQuery.ID,
		&userQuery.Email,
		&userQuery.Password,
	)

	if err != nil {
		return models.User{}, err
	}
	return userQuery, nil
}

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