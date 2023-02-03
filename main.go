package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vietthangc1/mini-web-golang/environment"
	"github.com/vietthangc1/mini-web-golang/routes"
)

func main() {
	// Set up env variables
	environment.SetUpEnvironmentVariables()

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

	router := routes.GenerateRoutes()

	fmt.Printf("Running at http://%s\n", os.Getenv("port"))
	router.Run(os.Getenv("port"))
}
