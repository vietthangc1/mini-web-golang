package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("mysqlLogin"))
	if err != nil {
		return nil, err
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if err != nil {
		fmt.Println("Ping Failed!!")
		return nil,err
	}		
	fmt.Println("Successful database connection")
	return db, nil
}

func ConnectDatabaseORM() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(os.Getenv("mysqlLogin")), &gorm.Config{
		QueryFields: true,
	})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(
		&User{}, 
		&Product{}, 
		&Propertises{},
	)
	log.Println(db.Migrator().HasTable(&Product{}))
	return db, nil
}