package staffmode

import (
	"github.com/stcraft/dragonfly/server/cmd"
	"github.com/stcraft/dragonfly/server/player"
	"github.com/stcraft/loader/dragonfly"
)

// Toggle is a command to enable or disable staff mode
type Toggle struct{}

// Run ...
func (c Toggle) Run(src cmd.Source, o *cmd.Output) {
	p := src.(*player.Player)

	if enabled(p) {
		disable(p)
		p.Message(dragonfly.Translation("staff_mode_disabled"))
	} else {
		enable(p)
		p.Message(dragonfly.Translation("staff_mode_enabled"))
	}
}

// Allow ...
func (c Toggle) Allow(src cmd.Source) bool {
	s, isPlayer := src.(*player.Player)

	if !isPlayer || !dragonfly.IsOP(s.XUID()) {
		return false
	}

	return true
}
