package fakeblock

import (
	"github.com/STCraft/dragonfly/server/block/cube"
	"github.com/STCraft/dragonfly/server/world"
)

type FakeBlock struct {
	original world.Block
	block    world.Block

	pos   cube.Pos
	world *world.World

	viewers []world.Viewer
}

func New(w *world.World, pos cube.Pos, block world.Block) FakeBlock {
	return FakeBlock{
		original: w.Block(pos),
		block:    block,
		pos:      pos,
		world:    w,
		viewers:  make([]world.Viewer, 0),
	}
}

func (fb *FakeBlock) AddViewer(v world.Viewer) {
	fb.viewers = append(fb.viewers, v)
	v.ViewBlockUpdate(fb.pos, fb.block, 0)
}

func (fb *FakeBlock) RemoveViewer(v world.Viewer) {
	v.ViewBlockUpdate(fb.pos, fb.original, 0)
}
