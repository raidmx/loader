package fakeblock

import (
	"sync"

	"github.com/STCraft/dragonfly/server/block/cube"
	"github.com/STCraft/dragonfly/server/player"
	"github.com/STCraft/dragonfly/server/world"
)

// FakeBlockRegistry is the registry of all the fake blocks spawned in the
// server.
type FakeBlockRegistry struct {
	mu   *sync.Mutex
	list []*FakeBlock
}

// fakeblocks is an instance of FakeBlockRegistry storing all the fake blocks
// spawned in the server.
var fakeblocks = FakeBlockRegistry{
	mu:   &sync.Mutex{},
	list: []*FakeBlock{},
}

// FakeBlock represents a block that is viewable by one or more specified clients.
// Usually the use cases include virtual inventories and they usually replace air blocks.
// Before viewing entities may go out of the scope of the fake block, you must remove it yourself
// by calling the RemoveViewer method. It is also assumed that the original block is same for all
// the viewing entities so that when a fakeblock is despawned, it sets it back to the original
// block for them.
type FakeBlock struct {
	original world.Block
	block    world.Block

	pos   cube.Pos
	world *world.World

	mu      *sync.Mutex
	viewers map[string]*player.Player
}

// Creates and returns a new FakeBlock
func New(w *world.World, pos cube.Pos, block world.Block) *FakeBlock {
	fb := &FakeBlock{
		original: w.Block(pos),
		block:    block,
		pos:      pos,
		world:    w,
		mu:       &sync.Mutex{},
		viewers:  map[string]*player.Player{},
	}

	defer fakeblocks.mu.Unlock()
	fakeblocks.mu.Lock()

	fakeblocks.list = append(fakeblocks.list, fb)
	return fb
}

// AddViewer adds a viewer for the fake block
func (fb *FakeBlock) AddViewer(p *player.Player) {
	defer fb.mu.Unlock()
	fb.mu.Lock()

	p.Session().ViewBlockUpdate(fb.pos, fb.block, 0)
	fb.viewers[p.XUID()] = p
}

// RemoveViewer removes the viewer for the fake block
func (fb *FakeBlock) RemoveViewer(p *player.Player) {
	defer fb.mu.Unlock()
	fb.mu.Lock()

	p.Session().ViewBlockUpdate(fb.pos, fb.original, 0)
	delete(fb.viewers, p.XUID())
}

// Despawn removes the fake block for all the players in the world
func (fb *FakeBlock) Despawn() {
	defer fb.mu.Unlock()
	fb.mu.Lock()

	for _, p := range fb.viewers {
		p.Session().ViewBlockUpdate(fb.pos, fb.original, 0)
	}
}
