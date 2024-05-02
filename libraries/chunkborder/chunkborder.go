package chunkborder

import (
	"sync"

	"github.com/stcraft/dragonfly/server/block"
	"github.com/stcraft/dragonfly/server/block/cube"
	"github.com/stcraft/dragonfly/server/player"
	"github.com/stcraft/dragonfly/server/world"
	"github.com/stcraft/loader/libraries/fakeblock"
)

// ChunkBorderRegistry is the registry of all the chunk borders on
// the server
type ChunkBorderRegistry struct {
	mu   sync.RWMutex
	list []*ChunkBorder
}

// chunkborders is an instance of ChunkBorderRegistry that stores all
// the chunk borders spawned in the server.
var chunkborders = ChunkBorderRegistry{
	mu:   sync.RWMutex{},
	list: []*ChunkBorder{},
}

// ChunkBorder represents basically a structure block that shows
// the chunk borders of the chunk the player is in.
type ChunkBorder struct {
	chunkX int
	chunkZ int

	fakeblock *fakeblock.FakeBlock
}

// New creates and initialises a new Chunk Border and returns it
func New(p cube.Pos, w *world.World) *ChunkBorder {
	cx := p.X() >> 4
	cz := p.Z() >> 4

	cornerX := cx * 16
	cornerZ := cz * 16
	cornerY := w.HighestBlock(cx, cz) - 1

	pos := cube.Pos{cornerX, cornerY, cornerZ}
	b := block.StructureBlock{X: int32(cornerX), Y: int32(cornerY), Z: int32(cornerZ)}

	fb := fakeblock.New(pos, w, b)
	cb := &ChunkBorder{
		chunkX:    cx,
		chunkZ:    cz,
		fakeblock: fb,
	}

	defer chunkborders.mu.Unlock()
	chunkborders.mu.Lock()

	chunkborders.list = append(chunkborders.list, cb)
	return cb
}

// Get returns a chunk border spawned in the provided world at the
// chunk containing the specified position. If no chunk border was
// previously spawned this returns nil.
func Get(p cube.Pos, w *world.World) *ChunkBorder {
	defer chunkborders.mu.RUnlock()
	chunkborders.mu.RLock()

	cx := p.X() >> 4
	cz := p.Z() >> 4

	for _, cb := range chunkborders.list {
		if cb.chunkX == cx && cb.chunkZ == cz {
			return cb
		}
	}

	return nil
}

// World returns the world the chunk border is spawned in
func (c *ChunkBorder) World() *world.World {
	return c.fakeblock.World()
}

// ChunkX returns the x coordinate of the chunk
func (c *ChunkBorder) ChunkX() int {
	return c.chunkX
}

// ChunkZ returns the z coordinate of the chunk
func (c *ChunkBorder) ChunkZ() int {
	return c.chunkZ
}

// Send sends the chunk border to the specified player and returns
// whether the operation was successful.
func (c *ChunkBorder) Send(p *player.Player) bool {
	return c.fakeblock.AddViewer(p)
}

// Remove removes the chunk border from the specified player and returns
// whether the operation was successful.
func (c *ChunkBorder) Remove(p *player.Player) bool {
	return c.fakeblock.RemoveViewer(p)
}

// Destroy destroys the chunk border and restores the original block for
// all the viewers, and deregisters the chunk border.
func (c *ChunkBorder) Destroy() {
	defer chunkborders.mu.Unlock()
	chunkborders.mu.Lock()

	c.fakeblock.Destroy()

	for index, cb := range chunkborders.list {
		if cb.World() == c.World() && cb.chunkX == c.chunkX && cb.chunkZ == c.chunkZ {
			chunkborders.list[index] = chunkborders.list[len(chunkborders.list)-1]
			chunkborders.list = chunkborders.list[:len(chunkborders.list)-1]
		}
	}
}
