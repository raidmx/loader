package loader

import (
	"github.com/STCraft/dragonfly-loader/dragonfly"
	"github.com/STCraft/dragonfly-loader/registry"
)

// Start starts the dragonfly-loader mod and registers all the commands, libraries, and various
// other features that this library provides.
func Start() {
	registry.VanillaCommands()
	dragonfly.New()
}
