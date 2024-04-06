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
	p, ok := src.(*player.Player)

	sender := ""

	if ok {
		sender = p.Name()
	} else {
		sender = "Console"
	}

	dragonfly.Server.Broadcast("§d%s says: %s", sender, c.Message)
}
