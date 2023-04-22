package db

import (
	"fmt"
	"log"
	"os"

	"github.com/vietthangc1/mini-web-golang/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabaseORM() (*gorm.DB, error) {
	uri := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DB"),
	)

	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{
		QueryFields: true,
	})
	if err != nil {
		return nil, err
	}

	log.Println("Connect to dtb")

	err = db.Migrator().CreateConstraint(&models.Product{}, "Price > 0")
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Propertises{},
		&models.Log{},
	)
	if err != nil {
		return nil, err
	}
	
	return db, nil
}
