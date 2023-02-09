package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/vietthangc1/mini-web-golang/routes"
)

func main() {
	// Set up env variables
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}

	router := routes.GenerateRoutes()
	fmt.Printf("Running at http://%s\n", os.Getenv("port"))
	router.Run(os.Getenv("port"))
}
