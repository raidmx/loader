package commands

import (
	"github.com/STCraft/DFLoader/dragonfly"
	"github.com/STCraft/dragonfly/server/cmd"
	"github.com/STCraft/dragonfly/server/player"
)

// Deop can be executed from the console or by another operator to remove a player
// from their operator status
type Deop struct {
	Target []cmd.Target `cmd:"player"`
}

// Run ...
func (c Deop) Run(src cmd.Source, o *cmd.Output) {
	if len(c.Target) > 1 {
		o.Print(dragonfly.Translation("single_target_expected"))
		return
	}

	t, ok := c.Target[0].(*player.Player)
	if !ok {
		o.Print(dragonfly.Translation("target_must_be_player"))
		return
	}

	if !dragonfly.IsOP(t.XUID()) {
		o.Printf(dragonfly.Translation("already_not_operator", t.Name()))
		return
	}

	dragonfly.RemoveOP(t.XUID())
	o.Printf(dragonfly.Translation("operator_withdrawn", t.Name()))
	dragonfly.SendToast(t, "operator_access_withdrawn")
}

// Allow ...
func (c Deop) Allow(src cmd.Source) bool {
	s, isPlayer := src.(*player.Player)

	if isPlayer && !dragonfly.IsOP(s.XUID()) {
		return false
	}

	return true
}
