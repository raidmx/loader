package commands

import (
	"github.com/STCraft/DFLoader/dragonfly"
	"github.com/STCraft/dragonfly/server/cmd"
	"github.com/STCraft/dragonfly/server/player"
)

// Tell command is used to send a private message to an online player
type Tell struct {
	Target  []cmd.Target `cmd:"player"`
	Message cmd.Varargs  `cmd:"message"`
}

// Run ...
func (c Tell) Run(src cmd.Source, o *cmd.Output) {
	if len(c.Target) > 1 {
		o.Print(dragonfly.Translation("single_target_expected"))
		return
	}

	sender := "CONSOLE"
	msg := string(c.Message)

	if p, ok := src.(*player.Player); ok {
		sender = p.Name()
	}

	p, ok := c.Target[0].(*player.Player)
	if !ok {
		o.Print(dragonfly.Translation("target_must_be_player"))
		return
	}

	if sender == p.Name() {
		o.Printf(dragonfly.Translation("must_specify_target"))
		return
	}

	p.Message(dragonfly.Translation("tell", sender, p.Name(), msg))
	o.Printf(dragonfly.Translation("tell", sender, p.Name(), msg))
}
