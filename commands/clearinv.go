package commands

import (
	"github.com/stcraft/DFLoader/dragonfly"
	"github.com/stcraft/dragonfly/server/cmd"
	"github.com/stcraft/dragonfly/server/player"
)

// ClearInv command is used to clear the inventory of a player. You can specify
// the type of inventory that must be cleared.
type ClearInv struct {
	Target    []cmd.Target  `cmd:"player"`
	Inventory inventoryType `cmd:"type"`
}

// Run ...
func (c ClearInv) Run(src cmd.Source, o *cmd.Output) {
	sender := "CONSOLE"
	if p, ok := src.(*player.Player); ok {
		sender = p.Name()
	}

	for _, t := range c.Target {
		p, ok := t.(*player.Player)
		if !ok {
			continue
		}

		count := 0

		switch c.Inventory {
		case "main":
			count = len(p.Inventory().Clear())
		case "enderchest":
			count = len(p.EnderChestInventory().Clear())
		case "armour":
			count = len(p.Armour().Clear())
		}

		p.Message(dragonfly.Translation("inventory_cleared_target", count, string(c.Inventory), sender))
		o.Printf(dragonfly.Translation("inventory_cleared_sender", count, p.Name(), string(c.Inventory)))
	}
}

// Allow ...
func (c ClearInv) Allow(src cmd.Source) bool {
	s, isPlayer := src.(*player.Player)

	if isPlayer && !dragonfly.IsOP(s.XUID()) {
		return false
	}

	return true
}

// inventoryType is type of inventory to be cleared such as the main inventory
// or the enderchest inventory, etc.
type inventoryType string

// Type ...
func (inventoryType) Type() string {
	return "inventoryType"
}

// Options ...
func (inventoryType) Options(_ cmd.Source) []string {
	return []string{"main", "enderchest", "armour"}
}
