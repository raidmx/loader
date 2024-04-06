package registry

import (
	"github.com/STCraft/DFLoader/commands"
	"github.com/STCraft/dragonfly/server/cmd"
)

// VanillaCommands registers all the default vanilla commands to the dragonfly server.
func VanillaCommands() {
	cmd.Register(cmd.New("gamemode", "Updates the gamemode of the provided target", []string{"gm"}, commands.GameMode{}))
	cmd.Register(cmd.New("say", "Broadcasts a message to all the players", []string{"say"}, commands.Say{}))
	cmd.Register(cmd.New("weather", "Updates the weather", []string{}, commands.Weather{}))
	cmd.Register(cmd.New("op", "Sets a player to the server operator", []string{}, commands.OP{}))
	cmd.Register(cmd.New("deop", "Removes a player from being a server operator", []string{}, commands.DEOP{}))
}
