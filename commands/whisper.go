package commands

import (
	"github.com/STCraft/DFLoader/dragonfly"
	"github.com/STCraft/dragonfly/server/block/cube"
	"github.com/STCraft/dragonfly/server/cmd"
	"github.com/STCraft/dragonfly/server/player"
	"github.com/STCraft/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

// Whisper command is used to whisper to nearby players
type Whisper struct {
	Message cmd.Varargs `cmd:"message"`
}

// Run ...
func (c Whisper) Run(src cmd.Source, o *cmd.Output) {
	p := src.(*player.Player)
	msg := string(c.Message)

	min := p.Position().Sub(mgl64.Vec3{5, 5, 5})
	max := p.Position().Add(mgl64.Vec3{5, 5, 5})

	bbox := cube.Box(min.X(), min.Y(), min.Z(), max.X(), max.Y(), max.Z())
	entities := p.World().EntitiesWithin(bbox, func(e world.Entity) bool {
		_, ok := e.(*player.Player)
		return !ok
	})

	for _, e := range entities {
		t := e.(*player.Player)
		t.Message(dragonfly.Translation("whispers", p.Name(), msg))
	}
}

// Allow ...
func (c Whisper) Allow(src cmd.Source) bool {
	_, isPlayer := src.(*player.Player)
	return isPlayer
}
