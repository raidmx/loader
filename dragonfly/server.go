package dragonfly

import (
	"github.com/STCraft/dragonfly/server"
)

// Server is a global instance of dragonfly
var Server *server.Server

// New is a blocking function that creates and initialises a new Dragonfly Server
// instance.
func New() {
	defer func() {
		saveOperators() // We must save the list of operators in the end.
		DB.Close()      // We must close our connection to the database.
	}()

	loadLanguage()
	loadOperators()

	Server = server.New()
	Server.RegisterHandler("user", UserHandler{})

	Server.Start()
}
