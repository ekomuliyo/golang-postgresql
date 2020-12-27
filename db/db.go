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
			);
			CREATE TABLE IF NOT EXISTS rooms (
				id_room serial PRIMARY KEY,
				name_room VARCHAR (255) NOT NULL,
				price INT NOT NULL,
				max_capacity INT NOT NULL
			);
			CREATE TABLE IF NOT EXISTS status_booking_room (
				id_status_booking serial PRIMARY KEY,
				status_booking VARCHAR (155) NOT NULL,
				description VARCHAR (255) NULL
			);
			CREATE TABLE IF NOT EXISTS booking_room (
				id_booking_room serial PRIMARY KEY,
				id_room INT NOT NULL,
				id_user INT NOT NULL,
				date_booking DATE NOT NULL,
				amount_capacity INT NOT NULL,
				id_status_booking INT NOT NULL,
				FOREIGN KEY(id_room) REFERENCES rooms (id_room),
				FOREIGN KEY(id_user) REFERENCES users (id_user),
				FOREIGN KEY(id_status_booking) REFERENCES status_booking_room(id_status_booking)
			);
			`)

	if err != nil {
		panic(err)
	}

	dbGlobal = db
}

// CreateCon ...
func CreateCon() *sql.DB {
	return dbGlobal
}
