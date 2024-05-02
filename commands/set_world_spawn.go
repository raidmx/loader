package commands

import (
	"github.com/stcraft/dragonfly/server/block/cube"
	"github.com/stcraft/dragonfly/server/cmd"
	"github.com/stcraft/dragonfly/server/player"
	"github.com/stcraft/loader/dragonfly"
)

// SetWorldSpawn is used to set the world spawn of the world the command
// executor is in at the position of the command executor
type SetWorldSpawn struct {
}

// Run ..
func (c SetWorldSpawn) Run(src cmd.Source, o *cmd.Output) {
	w := src.World()
	pos := cube.PosFromVec3(src.Position())

	w.SetSpawn(pos)
	o.Printf(dragonfly.Translation("set_world_spawn", pos.X(), pos.Y(), pos.Z()))
}

// Allow ...
func (c SetWorldSpawn) Allow(src cmd.Source) bool {
	s, isPlayer := src.(*player.Player)

	if !isPlayer || !dragonfly.IsOP(s.XUID()) {
		return false
	}

	return true
}
