package models

import (
	"log"
	"os"

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
		&User{},
		&Product{},
		&Propertises{},
	)
	log.Println(db.Migrator().HasTable(&Product{}))
	return db, nil
}
