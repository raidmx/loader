package commands

import (
	"github.com/STCraft/DFLoader/dragonfly"
	"github.com/STCraft/dragonfly/server/cmd"
	"github.com/STCraft/dragonfly/server/player"
)

// OP can only be executed from the console to set a player as the server operator.
type OP struct {
	Target []cmd.Target `cmd:"player"`
}

// Run ...
func (c OP) Run(src cmd.Source, o *cmd.Output) {
	if len(c.Target) > 1 {
		o.Print(dragonfly.Translation("single_target_expected"))
		return
	}

	t, ok := c.Target[0].(*player.Player)
	if !ok {
		o.Print(dragonfly.Translation("target_must_be_player"))
		return
	}

	if dragonfly.IsOP(t.XUID()) {
		o.Printf(dragonfly.Translation("already_operator", t.Name()))
		return
	}

	dragonfly.SetOP(t.XUID())

	o.Printf(dragonfly.Translation("operator_granted", t.Name()))
	dragonfly.SendToast(t, "operator_access_granted")
}

// Allow ...
func (c OP) Allow(src cmd.Source) bool {
	if _, isPlayer := src.(*player.Player); isPlayer {
		return false
	}

	return true
}