package db

import (
	"log"
	"os"

	"github.com/vietthangc1/mini-web-golang/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabaseORM() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(os.Getenv("MYSQLHOST")), &gorm.Config{
		QueryFields: true,
	})
	if err != nil {
		return nil, err
	}

	log.Println("Connect to dtb")

	db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Propertises{},
	)
	return db, nil
}
