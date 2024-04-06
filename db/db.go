package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// Init initialises the postgreSQL database.
func Init() {
	var err error

	connStr := "postgres://postgres:stcraft123@localhost/stcraft?sslmode=disable"
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}

	createTables()
}

// This function initialises the database by creating the default tables if they do not exist
func createTables() {
	if _, err := DB.Exec(`CREATE TABLE IF NOT EXISTS Users ("Name" VARCHAR(12), "Xuid" VARCHAR(30), "LastOnline" VARCHAR(20), "Registered" VARCHAR(20))`); err != nil {
		panic(err)
	}
}
