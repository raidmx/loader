package dragonfly

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// InitDB initialises the postgreSQL database.
func InitDB() {
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

// DBExec wraps around the DB to provide a sql.Result. This function panics if
// the query failed.
func DBQuery(query string, args ...any) *sql.Rows {
	rows, err := DB.Query(query, args)

	if err != nil {
		panic(err)
	}

	return rows
}

// DBExec wraps around the DB to provide a sql.Result. This function panics if
// the query failed.
func DBExec(query string, args ...any) sql.Result {
	result, err := DB.Exec(query, args)

	if err != nil {
		panic(err)
	}

	return result
}

// This function initialises the database by creating the default tables if they do not exist
func createTables() {
	if _, err := DB.Exec(`CREATE TABLE IF NOT EXISTS "Users" ("Name" VARCHAR(20) NOT NULL, "Xuid" VARCHAR(30) NOT NULL, "LastOnline" VARCHAR(20) NOT NULL, "Registered" VARCHAR(20) NOT NULL, PRIMARY KEY ("Xuid"))`); err != nil {
		panic(err)
	}
}