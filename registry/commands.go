package registry

import (
	"github.com/STCraft/DFLoader/commands"
	"github.com/STCraft/dragonfly/server/cmd"
)

// RegisterVanillaCommands registers all the default vanilla commands to the dragonfly server.
func RegisterVanillaCommands() {
	cmd.Register(cmd.New("gamemode", "Updates the gamemode of the provided target", []string{"gm"}, commands.GameMode{}))
	cmd.Register(cmd.New("say", "Broadcasts a message to all the players", []string{"say"}, commands.Say{}))
	cmd.Register(cmd.New("weather", "Updates the weather", []string{}, commands.Weather{}))
	cmd.Register(cmd.New("op", "Sets a player to the server operator", []string{}, commands.Op{}))
	cmd.Register(cmd.New("deop", "Removes a player from being a server operator", []string{}, commands.Deop{}))
	cmd.Register(cmd.New("multiworld", "Multiworld System", []string{"mw"}, commands.MWLoad{}, commands.MWUnload{}, commands.MWTeleport{}, commands.MWList{}))
	cmd.Register(cmd.New("tell", "Message a player", []string{"msg"}, commands.Tell{}))
	cmd.Register(cmd.New("time", "Time management", []string{}, commands.TimeAdd{}, commands.TimeQuery{}, commands.TimeSet{}, commands.TimeSetSpec{}))
	cmd.Register(cmd.New("kill", "Kills a player", []string{""}, commands.Kill{}))
	cmd.Register(cmd.New("teleport", "Teleports you the specified position", []string{"tp"}, commands.Teleport{}))
	cmd.Register(cmd.New("setworldspawn", "Sets the world spawn", []string{}, commands.SetWorldSpawn{}))
	cmd.Register(cmd.New("transfer", "Transfer a player to another server", []string{}, commands.TransferServer{}))
	cmd.Register(cmd.New("defaultgamemode", "Sets the default gamemode for the provided world", []string{}, commands.DefaultGamemode{}))
}
