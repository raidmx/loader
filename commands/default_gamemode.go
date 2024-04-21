package commands

import (
	"github.com/STCraft/DFLoader/dragonfly"
	"github.com/STCraft/dragonfly/server/cmd"
	"github.com/STCraft/dragonfly/server/player"
	"github.com/STCraft/dragonfly/server/world"
)

// DefaultGamemode is used to set the default gamemode for the provided world
// to the specified mode. It is persistent and will apply for new players that join.
type DefaultGamemode struct {
	World    worldName `cmd:"world"`
	GameMode gamemode  `cmd:"gamemode"`
}

// Run ..
func (c DefaultGamemode) Run(src cmd.Source, o *cmd.Output) {
	w := dragonfly.Server.World(string(c.World))
	mode := ""

	switch c.GameMode {
	case "s", "survival", "0":
		mode = "Survival"
		w.SetDefaultGameMode(world.GameModeSurvival)
	case "c", "creative", "1":
		mode = "Creative"
		w.SetDefaultGameMode(world.GameModeCreative)
	case "a", "adventure", "2":
		mode = "Adventure"
		w.SetDefaultGameMode(world.GameModeAdventure)
	case "sp", "spectator", "3":
		mode = "Spectator"
		w.SetDefaultGameMode(world.GameModeSpectator)
	}

	o.Printf(dragonfly.Translation("set_default_gamemode", mode, w.Name()))
}

// Allow ...
func (c DefaultGamemode) Allow(src cmd.Source) bool {
	s, isPlayer := src.(*player.Player)

	if isPlayer && !dragonfly.IsOP(s.XUID()) {
		return false
	}

	return true
}
