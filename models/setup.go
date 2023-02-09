package models

import (
	"database/sql"
	"fmt"
	"os"
	"time"
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