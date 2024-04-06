package commands

import (
	"github.com/STCraft/DFLoader/dragonfly"
	"github.com/STCraft/dragonfly/server/cmd"
	"github.com/STCraft/dragonfly/server/player"
)

// Say broadcasts a message on behalf of the player or the Console if the message was sent
// from the console.
type Say struct {
	Message cmd.Varargs `cmd:"message"`
}

// Run ...
func (c Say) Run(src cmd.Source, o *cmd.Output) {
	sender := "CONSOLE"
	msg := string(c.Message)

	if p, ok := src.(*player.Player); ok {
		sender = p.Name()
	}

	dragonfly.Server.Broadcast(dragonfly.Translation("broadcast_say", sender, msg))
}

// Allow ...
func (c Say) Allow(src cmd.Source) bool {
	s, isPlayer := src.(*player.Player)

	if isPlayer && !dragonfly.IsOP(s.XUID()) {
		return false
	}

	return true
}
