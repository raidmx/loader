package fakeinv

import (
	"sync"

	"github.com/STCraft/DFLoader/libraries/fakeblock"
	"github.com/STCraft/dragonfly/server/block"
	"github.com/STCraft/dragonfly/server/block/cube"
	"github.com/STCraft/dragonfly/server/player"
	"github.com/STCraft/dragonfly/server/world"
)

const (
	FakeInventoryTypeChest byte = iota
	FakeInventoryTypeDoubleChest
	FakeInventoryTypeHopper
	FakeInventoryTypeDispenser
)

// FakeInventoryRegistry is the registry of all the FakeInventories on the server
type FakeInventoryRegistry struct {
	mu   *sync.RWMutex
	list []*FakeInventory
}

// fakeinventories is an instance of FakeInventoryRegistry storing all the fake inventories
// spawned in the server.
var fakeinventories = FakeInventoryRegistry{
	mu:   &sync.RWMutex{},
	list: []*FakeInventory{},
}

// FakeViewer is a player entity viewing the FakeInventory with the window ID
// specific to them.
type FakeViewer struct {
	p        *player.Player
	windowID uint32
}

// FakeInventory represents a fake block container that is used to open
// client side inventory menu over a container that is a fake block which
// is only visible to the specified players.
type FakeInventory struct {
	fakeblock *fakeblock.FakeBlock
	container block.Container

	mu      sync.RWMutex
	viewers map[string]FakeViewer
}

// New creates and returns a new Fake Inventory at the specified position
func New(w *world.World, pos cube.Pos, inventoryType byte) *FakeInventory {
	container := block.NewChest()
	fakeblock := fakeblock.New(w, pos, container)

	fakeinventory := &FakeInventory{
		container: container,
		fakeblock: fakeblock,
	}

	defer fakeinventories.mu.Unlock()
	fakeinventories.mu.Lock()

	fakeinventories.list = append(fakeinventories.list, fakeinventory)
	return fakeinventory
}

// World returns the world the fake inventory is spawned in
func (inv *FakeInventory) World() *world.World {
	return inv.fakeblock.World()
}

// Pos returns the position of the fake inventory
func (inv *FakeInventory) Pos() cube.Pos {
	return inv.fakeblock.Pos()
}

// AddViewer adds a player to view the fake inventory and returns whether
// it was successful
func (inv *FakeInventory) AddViewer(p *player.Player) bool {
	defer inv.mu.Unlock()
	inv.mu.Lock()

	if !inv.fakeblock.AddViewer(p) {
		return false
	}

	pos := inv.Pos()
	windowID := p.Session().OpenFakeContainer(pos, inv.container)

	inv.viewers[p.Name()] = FakeViewer{
		p:        p,
		windowID: windowID,
	}
	return true
}

// RemoveViewer removes a viewing player from the fake inventory and
// returns whether it was successful
func (inv *FakeInventory) RemoveViewer(p *player.Player) bool {
	defer inv.mu.Unlock()
	inv.mu.Lock()

	fviewer, ok := inv.viewers[p.Name()]

	if !ok {
		return false
	}

	delete(inv.viewers, p.Name())

	if p.World() != inv.World() {
		return false
	}

	if p.Session().OpenedWindowID() != fviewer.windowID {
		return false
	}

	inv.container.RemoveViewer(p.Session(), p.World(), inv.Pos())
	p.Session().CloseFakeContainer()

	return !inv.fakeblock.RemoveViewer(p)
}

// Destroy tries to destroy the fake inventory by removing all the opened containers
// if any opened and also destorys the fake block for all viewers
func (inv *FakeInventory) Destroy() {
	defer fakeinventories.mu.Unlock()
	fakeinventories.mu.Lock()

	for _, p := range inv.viewers {
		inv.RemoveViewer(p.p)
	}

	for index, fakeinventory := range fakeinventories.list {
		if fakeinventory.World() == inv.World() && fakeinventory.Pos() == inv.Pos() {
			fakeinventories.list[index] = fakeinventories.list[len(fakeinventories.list)-1]
			fakeinventories.list = fakeinventories.list[:len(fakeinventories.list)-1]
		}
	}

	inv.fakeblock.Destroy()
}
