package commands

import (
	"github.com/STCraft/DFLoader/dragonfly"
	"github.com/STCraft/dragonfly/server/cmd"
	"github.com/STCraft/dragonfly/server/player"
	"github.com/go-gl/mathgl/mgl64"
)

// Teleport is used to teleport a player to the provided coordinates
type Teleport struct {
	X float64 `cmd:"x"`
	Y float64 `cmd:"y"`
	Z float64 `cmd:"z"`
}

// Run ..
func (c Teleport) Run(src cmd.Source, o *cmd.Output) {
	p := src.(*player.Player)

	o.Printf(dragonfly.Translation("teleported", c.X, c.Y, c.Z))
	p.Teleport(mgl64.Vec3{c.X, c.Y, c.Z})
}

// Allow ...
func (c Teleport) Allow(src cmd.Source) bool {
	s, isPlayer := src.(*player.Player)

	if !isPlayer || !dragonfly.IsOP(s.XUID()) {
		return false
	}

	return true
}
