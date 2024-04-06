package db

import (
	"time"
)

// User represents a player that may be offline. This is useful for retrieving
// data about a player even when they are offline and also useful data such as
// last online time, and their registration time, etc.
type User struct {
	Name       string
	Xuid       string
	LastOnline int
	Registered int
}

// GetUser gets the user with the specified xuid from the database. This function
// will panic if the user does not exist.
func GetUser(xuid string) User {
	user := User{}

	rows, err := DB.Query(`SELECT * FROM "Users" where "Xuid" = $1`, xuid)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if err := rows.Scan(&user.Name, &user.Xuid, &user.LastOnline, &user.Registered); err != nil {
		panic(err)
	}

	return user
}

// GetUserFromName gets the user with the specified name from the database. This function
// will panic if the user does not exist.
func GetUserFromName(name string) User {
	user := User{}

	rows, err := DB.Query(`SELECT * FROM "Users" where "Name" = $1`, name)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if err := rows.Scan(&user.Name, &user.Xuid, &user.LastOnline, &user.Registered); err != nil {
		panic(err)
	}

	return user
}

// IsUser returns whether the user with the specified xuid exists.
func IsUser(xuid string) bool {
	rows, err := DB.Query(`SELECT * FROM "Users" where "Xuid" = $1`, xuid)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	return rows.Next()
}

// CreateUser creates a new user with the specified xuid and name
func CreateUser(xuid string, name string) {
	time := time.Now().UnixMilli()

	if _, err := DB.Exec(`INSERT INTO "Users" ("Name", "Xuid", "LastOnline", "Registered") VALUES ($1, $2, $3, $4)`, xuid, name, time, time); err != nil {
		panic(err)
	}
}

// UpdateUser is called to update the user's display name if changed and their last online time.
func UpdateUser(xuid string, name string) {
	time := time.Now().UnixMilli()

	if _, err := DB.Exec(`UPDATE "Users" SET "Name" = $1, "LastOnline" = $2 WHERE "Xuid" = $3`, name, time, xuid); err != nil {
		panic(err)
	}
}
