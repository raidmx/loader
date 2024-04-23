package fakeinv

import (
	"sync"
	"time"

	"github.com/STCraft/DFLoader/libraries/fakeblock"
	"github.com/STCraft/dragonfly/server/block"
	"github.com/STCraft/dragonfly/server/block/cube"
	"github.com/STCraft/dragonfly/server/item/inventory"
	"github.com/STCraft/dragonfly/server/player"
	"github.com/STCraft/dragonfly/server/world"
)

const (
	InventoryTypeChest byte = iota
	InventoryTypeDoubleChest
	InventoryTypeHopper
	InventoryTypeDispenser
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
	inventoryType byte

	fakeblocks []*fakeblock.FakeBlock
	container  block.Container

	mu      sync.RWMutex
	viewers map[string]FakeViewer
}

// New creates and returns a new Fake Inventory at the specified position
func New(pos cube.Pos, w *world.World, inventoryType byte) *FakeInventory {
	var fakeblocks []*fakeblock.FakeBlock
	var container block.Container

	switch inventoryType {
	case InventoryTypeChest:
		c := block.NewChest(block.ChestTypeSingle)

		fakeblocks = append(fakeblocks, fakeblock.New(pos, w, c))
		container = c
	case InventoryTypeDoubleChest:
		// Calculate the neighbour position and spawn the first pair of the chest
		neighbour := pos.Add(cube.Pos{1, 0, 0})
		c := block.NewChest(block.ChestTypeDouble)
		c.Facing = cube.North

		// Create another chest of a certain inventory type and set the
		// facing and inventory same as previous chest
		pair := block.NewChest(block.ChestTypeDouble)
		pair.Facing = c.Facing
		pair.SetInventory(c.Inventory())

		// Set the pairing values for this chest to link to the second
		// chest
		c.Paired = true
		c.PairX = int32(neighbour[0])
		c.PairZ = int32(neighbour[2])

		// Set the pairing values for other chest to link to first
		// chest
		pair.Paired = true
		pair.PairX = int32(pos[0])
		pair.PairZ = int32(pos[2])

		// Spawn the fakeblocks for both the chest pairs and set the container to the first chest
		fakeblocks = append(fakeblocks, fakeblock.New(pos, w, c), fakeblock.New(neighbour, w, pair))
		container = c
	case InventoryTypeHopper:
		h := block.NewHopper()

		fakeblocks = append(fakeblocks, fakeblock.New(pos, w, h))
		container = h
	case InventoryTypeDispenser:
		d := block.NewDispenser()

		fakeblocks = append(fakeblocks, fakeblock.New(pos, w, d))
		container = d
	}

	fakeinventory := &FakeInventory{
		inventoryType: inventoryType,
		fakeblocks:    fakeblocks,
		container:     container,
		mu:            sync.RWMutex{},
		viewers:       map[string]FakeViewer{},
	}

	defer fakeinventories.mu.Unlock()
	fakeinventories.mu.Lock()

	fakeinventories.list = append(fakeinventories.list, fakeinventory)
	return fakeinventory
}

// World returns the world the fake inventory is spawned in
func (inv *FakeInventory) World() *world.World {
	return inv.fakeblocks[0].World()
}

// Pos returns the position of the fake inventory
func (inv *FakeInventory) Pos() cube.Pos {
	return inv.fakeblocks[0].Pos()
}

// Inventory returns the fake block container inventory
func (inv *FakeInventory) Inventory() *inventory.Inventory {
	return inv.container.Inventory()
}

// AddViewer adds a player to view the fake inventory and returns whether
// it was successful
func (inv *FakeInventory) AddViewer(p *player.Player) bool {
	defer inv.mu.Unlock()
	inv.mu.Lock()

	for _, fb := range inv.fakeblocks {
		if !fb.AddViewer(p) {
			return false
		}
	}

	/*
	 * If the inventory type is of kind Double Chest then we must open the container for the
	 * player after atleast one tick to allow the pairing to complete client side.
	 */
	var duration time.Duration

	switch inv.inventoryType {
	case InventoryTypeDoubleChest:
		duration = time.Millisecond * 55
	default:
		duration = time.Second * 0
	}

	time.AfterFunc(duration, func() {
		pos := inv.Pos()
		windowID := p.Session().OpenFakeContainer(pos, inv.container)

		inv.viewers[p.Name()] = FakeViewer{
			p:        p,
			windowID: windowID,
		}
	})

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

	for _, fb := range inv.fakeblocks {
		if !fb.RemoveViewer(p) {
			return false
		}
	}

	return false
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

	for _, fb := range inv.fakeblocks {
		fb.Destroy()
	}
}
