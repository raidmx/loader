package npc

import (
	"github.com/stcraft/DFLoader/dragonfly"
)

// Plugin represents a NPC plugin
type Plugin struct {
}

// Name ...
func (Plugin) Name() string {
	return "NPC"
}

// Description ...
func (Plugin) Description() string {
	return "Adds support for spawning Non Playable Characters"
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
	dragonfly.Server.RegisterHandler("npc", handler{})
}

// OnUnload ...
func (Plugin) OnUnload() {

}
