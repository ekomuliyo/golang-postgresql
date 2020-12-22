package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var dbGlobal *sql.DB
var err error

// Init ...
func Init() {

	// load env
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	var dbHost = os.Getenv("DB_HOST")
	var dbPort = os.Getenv("DB_PORT")
	var dbName = os.Getenv("DB_DATABASE")
	var dbUsername = os.Getenv("DB_USERNAME")
	var dbPassword = os.Getenv("DB_PASSWORD")

	var url string = "postgres://" + dbUsername + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=disable"
	db, err := sql.Open("postgres", url)

	if err != nil {
		panic("error connection")
	}

	err = db.Ping()

	if err != nil {
		panic("no connected")
	}

	// initial table
	_, err = db.Exec(`
			CREATE TABLE IF NOT EXISTS users (
				id_user serial PRIMARY KEY,
				username VARCHAR (50) UNIQUE NOT NULL,
				password VARCHAR (450) NOT NULL,
				email VARCHAR (255) UNIQUE NOT NULL,
				photo VARCHAR (50) UNIQUE NOT NULL,
				id_group INT DEFAULT 2
			);`)

	if err != nil {
		panic(err)
	}

	dbGlobal = db
}

// CreateCon ...
func CreateCon() *sql.DB {
	return dbGlobal
}
