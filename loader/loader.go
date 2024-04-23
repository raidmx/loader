package loader

import (
	"github.com/STCraft/DFLoader/dragonfly"
	"github.com/STCraft/DFLoader/libraries/npc"
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
	dragonfly.Server.RegisterHandler("npc", npc.Handler{})
}

// Start starts the Dragonfly server. This is a blocking function.
func Start() {
	if dragonfly.DB == nil {
		panic("Call loader.Init() first before calling this function")
	}

	dragonfly.Server.Start()
}

// Deinit deinitialises the Dragonfly server, saves configs, closes the database.
func Deinit() {
	dragonfly.SaveOperators() // We must save the list of operators in the end.
	dragonfly.DB.Close()      // We must close our connection to the database.
}
