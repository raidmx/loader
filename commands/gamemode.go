package commands

import (
	"github.com/stcraft/DFLoader/dragonfly"
	"github.com/stcraft/dragonfly/server/cmd"
	"github.com/stcraft/dragonfly/server/player"
	"github.com/stcraft/dragonfly/server/world"
)

// GameMode command updates the gamemode for the provided target. If no target is provided
type GameMode struct {
	Mode   gamemode                   `cmd:"gamemode"`
	Target cmd.Optional[[]cmd.Target] `cmd:"player"`
}

// Run ...
func (c GameMode) Run(src cmd.Source, o *cmd.Output) {
	players, ok := c.Target.Load()

	if !ok || len(players) == 0 {
		// If targets are not specified then we add the console as self target
		if p, ok := src.(*player.Player); ok {
			players = append(players, p)
		}
	}

	if len(players) == 0 {
		o.Printf(dragonfly.Translation("must_specify_target"))
		return
	}

	for _, target := range players {
		t, ok := target.(*player.Player)
		if !ok {
			continue
		}

		mode := setGamemode(t, c.Mode)
		o.Printf(dragonfly.Translation("target_gamemode_updated", t.Name(), mode))
	}

}

// Allow ...
func (c GameMode) Allow(src cmd.Source) bool {
	s, isPlayer := src.(*player.Player)

	if isPlayer && !dragonfly.IsOP(s.XUID()) {
		return false
	}

	return true
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

	t.Messagef(dragonfly.Translation("self_gamemode_updated", mode))
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
