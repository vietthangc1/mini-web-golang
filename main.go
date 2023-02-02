package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	db, err := sql.Open("mysql", "root:Chaugn@rs2@/mini_golang_project")
	if err != nil {
		panic(err.Error())
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Ping Failed!!")
	} else {
		fmt.Println("Successful database connection")
	}

	router := GenerateRoutes()

	fmt.Println("Running at http://127.0.0.1:8080")
	router.Run("127.0.0.1:8080")
}
