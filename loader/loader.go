package loader

import (
	"github.com/STCraft/DFLoader/dragonfly"
	"github.com/STCraft/DFLoader/registry"
)

// Init initialises the various libraries and utilities that may be used in common for both
// the DFLoader and Plugins / Libraries using this loader. For example: DB. This is put in a
// separate function so that plugins can initialise their own specific stuff after initialising
// the loader. For example: plugins may want to initialise their tables after initialising of the
// database.
func Init() {
	dragonfly.InitDB()
}

// Start starts the DFLoader mod and registers all the commands, libraries, and various
// other features that this library provides.
func Start() {
	if dragonfly.DB == nil {
		panic("Call loader.Init() first before calling this function")
	}

	registry.RegisterVanillaCommands()
	dragonfly.New()
}
