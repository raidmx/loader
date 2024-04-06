package loader

import (
	"github.com/STCraft/DFLoader/db"
	"github.com/STCraft/DFLoader/dragonfly"
	"github.com/STCraft/DFLoader/registry"
)

// Start starts the DFLoader mod and registers all the commands, libraries, and various
// other features that this library provides.
func Start() {
	registry.VanillaCommands()
	db.Init()
	dragonfly.New()
}
