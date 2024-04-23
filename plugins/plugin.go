package plugins

// Plugin is an interface that must be implemented by all statically compiled
// optional plugins. Plugins are used to extend the functionalities of the server
// software beyond the scope of what the DFLoader can do.
type Plugin interface {
	// Name returns the name of the plugin
	Name() string

	// Description returns the description of the plugin
	Description() string

	// Author returns the plugin author name
	Author() string

	// Version returns the plugin version
	Version() string

	// OnLoad is called when the plugin is loaded into the server. At this
	// stage you can fetch plugin data, initialise if required.
	OnLoad()

	// OnUnload is called when the plugin is unloaded from the server
	// At this stage, you can update plugin data, deinitialise if required.
	OnUnload()
}
