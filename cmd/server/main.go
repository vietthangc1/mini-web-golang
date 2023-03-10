package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	// "github.com/vietthangc1/mini-web-golang/app"
)

func main() {
	// Set up env variables
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}

	// dependency not using google wire
	// var _app app.App
	// _app.Initialize()
	// fmt.Printf("Running at http://%s\n", os.Getenv("PORT"))
	// _app.Run()

	// using google wire
	_app, err := InitializeApp()
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Running at http://%s\n", os.Getenv("PORT"))
	_app.Router.Run(os.Getenv("PORT"))
}
