package models

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

var DB *sql.DB

func ConnectDatabase() {
	DB, err := sql.Open("mysql", os.Getenv("mysqlLogin"))
	if err != nil {
		panic(err.Error())
	}

	// See "Important settings" section.
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	if err != nil {
		fmt.Println("Ping Failed!!")
	} else {
		fmt.Println("Successful database connection")
	}
}