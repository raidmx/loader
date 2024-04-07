package dragonfly

import (
	"time"

	"github.com/STCraft/dragonfly/server/event"
	"github.com/STCraft/dragonfly/server/player"
)

// DFUser represents a player that may be offline. This is useful for retrieving
// data about a player even when they are offline and also useful data such as
// last online time, and their registration time, etc.
type DFUser struct {
	Name       string
	Xuid       string
	LastOnline int
	Registered int
}

// User gets the user with the specified xuid from the database. This function
// will return nil if the user does not exist.
func User(xuid string) *DFUser {
	user := DFUser{}

	rows, err := DB.Query(`SELECT * FROM "Users" where "Xuid" = $1`, xuid)
	if err != nil {
		return nil
	}

	defer rows.Close()

	if err := rows.Scan(&user.Name, &user.Xuid, &user.LastOnline, &user.Registered); err != nil {
		return nil
	}

	return &user
}

// UserFromName gets the user with the specified name from the database. This function
// will return nil if the user does not exist.
func UserFromName(name string) *DFUser {
	user := DFUser{}

	rows, err := DB.Query(`SELECT * FROM "Users" where "Name" = $1`, name)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if err := rows.Scan(&user.Name, &user.Xuid, &user.LastOnline, &user.Registered); err != nil {
		return nil
	}

	return &user
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

// createUser creates a new user with the specified xuid and name
func createUser(xuid string, name string) {
	time := time.Now().UnixMilli()

	if _, err := DB.Exec(`INSERT INTO "Users" ("Name", "Xuid", "LastOnline", "Registered") VALUES ($1, $2, $3, $4)`, name, xuid, time, time); err != nil {
		panic(err)
	}
}

// updateUser is called to update the user's display name if changed and their last online time.
func updateUser(xuid string, name string) {
	time := time.Now().UnixMilli()

	if _, err := DB.Exec(`UPDATE "Users" SET "Name" = $1, "LastOnline" = $2 WHERE "Xuid" = $3`, name, time, xuid); err != nil {
		panic(err)
	}
}

// UserHandler handles the various events such as User Creation, User Update, etc.
type UserHandler struct {
	player.NopHandler
}

// HandleJoin ...
func (UserHandler) HandleJoin(ctx *event.Context, p *player.Player) {
	if !IsUser(p.XUID()) {
		createUser(p.XUID(), p.Name())
	} else {
		updateUser(p.XUID(), p.Name())
	}
}
