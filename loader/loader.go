package loader

import (
	"github.com/STCraft/DFLoader/dragonfly"
	"github.com/STCraft/DFLoader/registry"
	"github.com/STCraft/dragonfly/server"
)

// Init initialises the Dragonfly server. You can initialise your library specific requirements
// after calling this function.
func Init() {
	dragonfly.InitDB()
	dragonfly.LoadLanguage()
	dragonfly.LoadOperators()

	registry.RegisterVanillaCommands()

	dragonfly.Server = server.New()
	dragonfly.Server.RegisterHandler("user", dragonfly.UserHandler{})
}

// Start starts the Dragonfly server. This is a blocking function.
func Start() {
	if dragonfly.DB == nil {
		panic("Call loader.Init() first before calling this function")
	}

	defer func() {
		dragonfly.SaveOperators() // We must save the list of operators in the end.
		dragonfly.DB.Close()      // We must close our connection to the database.
	}()

	dragonfly.Server.Start()
}
