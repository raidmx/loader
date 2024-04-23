package staffmode

import (
	"github.com/STCraft/DFLoader/dragonfly"
	"github.com/STCraft/dragonfly/server/cmd"
	"github.com/STCraft/dragonfly/server/player"
)

// Toggle is a command to enable or disable staff mode
type Toggle struct{}

// Run ...
func (c Toggle) Run(src cmd.Source, o *cmd.Output) {

}

// Allow ...
func (c Toggle) Allow(src cmd.Source) bool {
	s, isPlayer := src.(*player.Player)

	if isPlayer && !dragonfly.IsOP(s.XUID()) {
		return false
	}

	return true
}
