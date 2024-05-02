package dragonfly

import (
	"database/sql"
	_ "embed"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/stcraft/engine/config"
)

// DatabaseConfig represents the configuration file format that
// contains the database connection details.
type DatabaseConfig struct {
	Address  string `json:"address"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

// DB is the global instance of the PostgreSQL database
var DB *sql.DB

//go:embed postgres.json
var defaultDBCfg []byte

// InitDB initialises the postgreSQL database.
func InitDB() {
	var cfg = DatabaseConfig{}
	var err error

	if err := config.Load("", "./postgres.json", &cfg, defaultDBCfg); err != nil {
		panic(err)
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", cfg.Username, cfg.Password, cfg.Address, cfg.Database)
	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}

	if _, err := DB.Exec(`CREATE TABLE IF NOT EXISTS "Users" ("Name" VARCHAR(20) NOT NULL, "Xuid" VARCHAR(30) NOT NULL, "LastOnline" VARCHAR(20) NOT NULL, "Registered" VARCHAR(20) NOT NULL, PRIMARY KEY ("Xuid"))`); err != nil {
		panic(err)
	}
}

// QueryDB wraps around the DB to provide sql.Rows. This function panics if
// the query failed.
func QueryDB(query string, args ...any) *sql.Rows {
	rows, err := DB.Query(query, args...)

	if err != nil {
		panic(err)
	}

	return rows
}

// ExecDB wraps around the DB to provide a sql.Result. This function panics if
// the exec failed.
func ExecDB(query string, args ...any) sql.Result {
	result, err := DB.Exec(query, args...)

	if err != nil {
		panic(err)
	}

	return result
}
