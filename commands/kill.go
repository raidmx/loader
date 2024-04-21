package commands

import (
	"fmt"

	"github.com/STCraft/DFLoader/dragonfly"
	"github.com/STCraft/dragonfly/server/cmd"
	"github.com/STCraft/dragonfly/server/entity/effect"
	"github.com/STCraft/dragonfly/server/player"
)

// Kill command is used to kill the specified player
type Kill struct {
	Target []cmd.Target `cmd:"player"`
}

// Run ...
func (c Kill) Run(src cmd.Source, o *cmd.Output) {
	sender := "CONSOLE"
	if p, ok := src.(*player.Player); ok {
		sender = p.Name()
	}

	for _, t := range c.Target {
		p, ok := t.(*player.Player)
		if !ok {
			continue
		}

		p.Kill(effect.InstantDamageSource{}, fmt.Sprintf("Killed by %s", sender))
		o.Printf(dragonfly.Translation("killed_player", p.Name()))
		p.Message(dragonfly.Translation("killed"))
	}
}

// Allow ...
func (c Kill) Allow(src cmd.Source) bool {
	s, isPlayer := src.(*player.Player)

	if isPlayer && !dragonfly.IsOP(s.XUID()) {
		return false
	}

	return true
}
