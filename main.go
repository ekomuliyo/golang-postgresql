package main

import (
	"golang-postgresql/db"
	"golang-postgresql/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	// load env
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// setting port dev / prod
	var port string
	if os.Getenv("PORT") == "" {
		port = os.Getenv("APP_PORT")
	} else {
		port = os.Getenv("PORT")
	}

	// fmt.Println(port)

	// initial db
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":" + port))
}
