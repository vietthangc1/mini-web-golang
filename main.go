package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/vietthangc1/mini-web-golang/app"
)

func main() {
	// Set up env variables
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}

	var _app app.App

	_app.Initialize()
	fmt.Printf("Running at http://%s\n", os.Getenv("PORT"))
	_app.Run()
}
