package staffmode

import (
	"sync"

	"github.com/STCraft/dragonfly/server/block"
	"github.com/STCraft/dragonfly/server/block/cube"
	"github.com/STCraft/dragonfly/server/item"
	"github.com/STCraft/dragonfly/server/player"
	"github.com/go-gl/mathgl/mgl64"
)

// StaffModeRegistry is the registry of all the Staff Members that have
// enabled Staff Mode on the server
type StaffModeRegistry struct {
	mu   sync.RWMutex
	list map[*player.Player]StaffMember
}

// staffMembers is an instance of StaffModeRegistry that contains a list of
// staff members in staff mode
var staffMembers = StaffModeRegistry{
	mu:   sync.RWMutex{},
	list: map[*player.Player]StaffMember{},
}

// StaffMember stores the data of a Staff Member that is currently in Staff Mode
// to be restored when they are back
type StaffMember struct {
	main   []item.Stack
	armour []item.Stack

	health float64
	food   int

	pos mgl64.Vec3
	rot cube.Rotation
}

// inventoryItems is a list of Staff Mode items that are given to a player
// when they enter Staff Mode.
var inventoryItems = map[int]item.Stack{
	0: item.NewStack(item.Dye{Colour: item.ColourRed()}, 1),
	1: item.NewStack(item.Spyglass{}, 1),
	2: item.NewStack(block.BlueIce{}, 1),
	3: item.NewStack(item.Clock{}, 1),
	4: item.NewStack(block.EnderChest{}, 1),
	5: item.NewStack(item.Sword{Tier: item.ToolTierNetherite}, 1),
	6: item.NewStack(item.EnderPearl{}, 1),
	7: item.NewStack(item.Axe{Tier: item.ToolTierGold}, 1),
	8: item.NewStack(item.EnchantedBook{}, 1),
}

// armourItems is a list of Staff Mode items that are added in the player's
// armour inventory when they enter Staff Mode.
var armourItems = map[int]item.Stack{
	0: item.NewStack(item.Helmet{Tier: item.ArmourTierNetherite{}}, 1),
	1: item.NewStack(item.Chestplate{Tier: item.ArmourTierNetherite{}}, 1),
	2: item.NewStack(item.Leggings{Tier: item.ArmourTierNetherite{}}, 1),
	3: item.NewStack(item.Boots{Tier: item.ArmourTierNetherite{}}, 1),
}

// enable enables the Staff Mode for the player.
func enable(p *player.Player) {
	defer staffMembers.mu.Unlock()
	staffMembers.mu.Lock()

	health := p.Health()
	food := p.Food()

	p.EnableFlight()
	p.SetHealth(p.MaxHealth())
	p.SetFood(20)

	staffMembers.list[p] = StaffMember{
		main:   p.Inventory().Items(),
		armour: p.Armour().Items(),
		health: health,
		food:   food,
		pos:    p.Position(),
		rot:    p.Rotation(),
	}

	p.Inventory().Clear()
	p.Armour().Inventory().Clear()

	for slot, it := range inventoryItems {
		if it.Empty() {
			continue
		}

		p.Inventory().SetItem(slot, it)
	}

	for slot, it := range armourItems {
		if it.Empty() {
			continue
		}

		p.Armour().Inventory().SetItem(slot, it)
	}
}

// enabled returns whether the provided player is in Staff Mode
func enabled(p *player.Player) bool {
	defer staffMembers.mu.RUnlock()
	staffMembers.mu.RLock()

	_, ok := staffMembers.list[p]
	return ok
}

// disable disables the Staff Mode for the provided player
func disable(p *player.Player) {
	defer staffMembers.mu.Unlock()
	staffMembers.mu.Lock()

	staff := staffMembers.list[p]
	delete(staffMembers.list, p)

	p.DisableFlight()
	p.SetHealth(staff.health)
	p.SetFood(staff.food)

	p.Inventory().Clear()
	p.Armour().Inventory().Clear()

	for slot, it := range staff.main {
		if it.Empty() {
			continue
		}

		p.Inventory().SetItem(slot, it)
	}

	for slot, it := range staff.armour {
		if it.Empty() {
			continue
		}

		p.Armour().Inventory().SetItem(slot, it)
	}
}
