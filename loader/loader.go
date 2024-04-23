package loader

import (
	"github.com/STCraft/DFLoader/dragonfly"
	"github.com/STCraft/DFLoader/plugins"
	"github.com/STCraft/DFLoader/registry"
	"github.com/STCraft/dragonfly/server"
)

// loadedPlugins contain a list of plugins that are registered and loaded
// by the Loader
var loadedPlugins = map[plugins.Plugin]struct{}{}

// LoadPlugin loads the provided plugin in the Loader
func LoadPlugin(plugin plugins.Plugin) {
	loadedPlugins[plugin] = struct{}{}
}

// UnloadPlugin unloads the provided plugin if found
func UnloadPlugin(plugin plugins.Plugin) {
	delete(loadedPlugins, plugin)
}

// Init initialises the Dragonfly server. You can initialise your library specific requirements
// after calling this function.
func Init() {
	dragonfly.InitDB()
	dragonfly.LoadLanguage()
	dragonfly.LoadOperators()

	registry.RegisterVanillaCommands()

	dragonfly.Server, dragonfly.Logger = server.New()
	dragonfly.Server.RegisterHandler("user", dragonfly.UserHandler{})

	for p := range loadedPlugins {
		dragonfly.Logger.Printf("Loading plugin %s (v%s) by %s\n", p.Name(), p.Version(), p.Author())
		p.OnLoad()
	}
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
	for p := range loadedPlugins {
		dragonfly.Logger.Printf("Unloading plugin %s\n", p.Name())
		p.OnUnload()
	}

	dragonfly.SaveOperators() // We must save the list of operators in the end.
	dragonfly.DB.Close()      // We must close our connection to the database.
}
