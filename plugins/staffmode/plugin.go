package staffmode

import (
	"github.com/stcraft/dragonfly/server/cmd"
	"github.com/stcraft/loader/dragonfly"
)

// Plugin is a Staff Mode Plugin
type Plugin struct {
}

// Name ...
func (Plugin) Name() string {
	return "Staff Mode"
}

// Description ...
func (Plugin) Description() string {
	return "Adds support for Staff Mode"
}

// Author ...
func (Plugin) Author() string {
	return "Crayder"
}

// Version ...
func (Plugin) Version() string {
	return "1.0.0"
}

// OnLoad ...
func (Plugin) OnLoad() {
	dragonfly.Server.RegisterHandler("staffmode", handler{})
	cmd.Register(cmd.New("staffmode", "Toggles Staff Mode", []string{"sm"}, Toggle{}))
}

// OnUnload ...
func (Plugin) OnUnload() {

}
