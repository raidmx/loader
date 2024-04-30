package fakeblock

import (
	"sync"

	"github.com/stcraft/dragonfly/server/block/cube"
	"github.com/stcraft/dragonfly/server/player"
	"github.com/stcraft/dragonfly/server/world"
)

// FakeBlockRegistry is the registry of all the fake blocks spawned in the
// server.
type FakeBlockRegistry struct {
	mu   *sync.RWMutex
	list []*FakeBlock
}

// fakeblocks is an instance of FakeBlockRegistry storing all the fake blocks
// spawned in the server.
var fakeblocks = FakeBlockRegistry{
	mu:   &sync.RWMutex{},
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

	world *world.World
	pos   cube.Pos

	mu      *sync.Mutex
	viewers map[string]*player.Player
}

// Creates and returns a new FakeBlock
func New(pos cube.Pos, w *world.World, block world.Block) *FakeBlock {
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

// Get returns the fake block at the passed position in the provided world.
// Returns nil if no fake block exists at the provided position.
func Get(w *world.World, pos cube.Pos) *FakeBlock {
	defer fakeblocks.mu.RUnlock()
	fakeblocks.mu.RLock()

	for _, fb := range fakeblocks.list {
		if fb.world == w && fb.pos == pos {
			return fb
		}
	}

	return nil
}

// World returns the world the fakeblock is spawned in
func (fb *FakeBlock) World() *world.World {
	return fb.world
}

// Pos returns the position of the fake block
func (fb *FakeBlock) Pos() cube.Pos {
	return fb.pos
}

// AddViewer adds a viewer for the fake block and returns whether the operation
// was successful
func (fb *FakeBlock) AddViewer(p *player.Player) bool {
	defer fb.mu.Unlock()
	fb.mu.Lock()

	if p.World() != fb.world {
		return false
	}

	p.Session().ViewBlockUpdate(fb.pos, fb.block, 0)
	fb.viewers[p.XUID()] = p

	return true
}

// RemoveViewer removes the viewer for the fake block and returns whether
// the operation was successful
func (fb *FakeBlock) RemoveViewer(p *player.Player) bool {
	defer fb.mu.Unlock()
	fb.mu.Lock()

	if p.World() != fb.world {
		return false
	}

	p.Session().ViewBlockUpdate(fb.pos, fb.original, 0)
	delete(fb.viewers, p.XUID())

	return true
}

// Destroy destroys the fake block for all the players in the world
func (fb *FakeBlock) Destroy() {
	defer fakeblocks.mu.Unlock()
	fakeblocks.mu.Lock()

	for _, p := range fb.viewers {
		fb.RemoveViewer(p)
	}

	for index, fakeblock := range fakeblocks.list {
		if fakeblock.world == fb.world && fakeblock.pos == fb.pos {
			fakeblocks.list[index] = fakeblocks.list[len(fakeblocks.list)-1]
			fakeblocks.list = fakeblocks.list[:len(fakeblocks.list)-1]
		}
	}
}
