package commands

import (
	"github.com/STCraft/dragonfly/server/cmd"
	"github.com/STCraft/dragonfly/server/player"
	"github.com/STCraft/dragonfly/server/world"
)

// GameMode command updates the gamemode for the provided target. If no target is provided
type GameMode struct {
	Mode   gamemode                   `cmd:"gamemode"`
	Target cmd.Optional[[]cmd.Target] `cmd:"player"`
}

// Run ...
func (c GameMode) Run(src cmd.Source, o *cmd.Output) {
	players, _ := c.Target.Load()

	s, isPlayer := src.(*player.Player)

	if len(players) == 0 {
		if isPlayer {
			_ = setGamemode(s, c.Mode)
		} else {
			o.Errorf("Usage: /gamemode <mode> <player>")
		}
	}

	for _, target := range players {
		t, ok := target.(*player.Player)
		if !ok {
			continue
		}

		mode := setGamemode(t, c.Mode)
		o.Printf("§7You have updated §f%v's §7Gamemode to §f%v", t.Name(), mode)
	}

}

// setGamemode sets the gamemode of the provided player to the specified new gamemode. It also dispatches
// a message to the player saying that their gamemode got updated.
func setGamemode(t *player.Player, gm gamemode) string {
	var mode = ""

	switch gm {
	case "s", "survival", "0":
		mode = "Survival"
		t.SetGameMode(world.GameModeSurvival)
	case "c", "creative", "1":
		mode = "Creative"
		t.SetGameMode(world.GameModeCreative)
	case "a", "adventure", "2":
		mode = "Adventure"
		t.SetGameMode(world.GameModeAdventure)
	case "sp", "spectator", "3":
		mode = "Spectator"
		t.SetGameMode(world.GameModeSpectator)
	}

	t.Messagef("§7Your Game Mode has been updated to §f%v", mode)
	return mode
}

// gamemode ...
type gamemode string

// Type ...
func (gamemode) Type() string {
	return "gamemode"
}

// Options ...
func (gamemode) Options(_ cmd.Source) []string {
	return []string{"s", "survival", "0", "c", "creative", "1", "a", "adventure", "2", "sp", "spectator", "3"}
}
