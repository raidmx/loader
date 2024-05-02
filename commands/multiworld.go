package commands

import (
	"github.com/stcraft/dragonfly/server"
	"github.com/stcraft/dragonfly/server/cmd"
	"github.com/stcraft/dragonfly/server/player"
	"github.com/stcraft/dragonfly/server/world"
	"github.com/stcraft/loader/dragonfly"
)

// dimension is the world.Dimension
type dimension string

// Type ...
func (dimension) Type() string {
	return "dimension"
}

// Options ...
func (dimension) Options(_ cmd.Source) []string {
	return []string{"overworld", "nether", "end"}
}

// generator is the world.Generator
type generator string

// Type ...
func (generator) Type() string {
	return "generator"
}

// Options ...
func (generator) Options(_ cmd.Source) []string {
	return []string{"flat", "void"}
}

// MWLoad is used to load a world with the specified name and settings. It
// creates a new world if it doesn't exist.
type MWLoad struct {
	Load      cmd.SubCommand `cmd:"load"`
	Name      string         `cmd:"name"`
	Dimension dimension      `cmd:"dimension"`
	Generator generator      `cmd:"generator"`
	ReadOnly  bool           `cmd:"readOnly"`
}

// Run ...
func (c MWLoad) Run(src cmd.Source, o *cmd.Output) {
	var dim world.Dimension
	var gen world.Generator

	switch c.Dimension {
	case "overworld":
		dim = world.Overworld
	case "nether":
		dim = world.Nether
	case "end":
		dim = world.End
	}

	switch c.Generator {
	case "flat":
		gen = server.FlatGenerator(dim)
	case "void":
		gen = server.VoidGenerator(dim)
	}

	dragonfly.Server.LoadWorld(c.Name, dim, gen, c.ReadOnly)
	o.Printf(dragonfly.Translation("world_loaded", c.Name, c.Dimension, c.Generator))
}

// Allow ...
func (c MWLoad) Allow(src cmd.Source) bool {
	s, isPlayer := src.(*player.Player)

	if isPlayer && !dragonfly.IsOP(s.XUID()) {
		return false
	}

	return true
}

// MWUnload is used to unload the world with the specified name
type MWUnload struct {
	Unload cmd.SubCommand `cmd:"unload"`
	Name   string         `cmd:"name"`
}

// Run ...
func (c MWUnload) Run(src cmd.Source, o *cmd.Output) {
	if c.Name == "overworld" {
		o.Printf(dragonfly.Translation("cannot_unload_overworld"))
		return
	}

	if !dragonfly.Server.WorldExists(c.Name) {
		o.Printf(dragonfly.Translation("world_not_found", c.Name))
		return
	}

	dragonfly.Server.UnloadWorld(c.Name)
	o.Printf(dragonfly.Translation("world_unloaded", c.Name))
}

// Allow ...
func (c MWUnload) Allow(src cmd.Source) bool {
	s, isPlayer := src.(*player.Player)

	if isPlayer && !dragonfly.IsOP(s.XUID()) {
		return false
	}

	return true
}

// MWList is used to list the loaded worlds
type MWList struct {
	List cmd.SubCommand `cmd:"list"`
}

// Run ...
func (c MWList) Run(src cmd.Source, o *cmd.Output) {
	loaded := dragonfly.Server.LoadedWorlds()
	o.Printf(dragonfly.Translation("loaded_worlds_count", len(loaded)))

	for _, name := range loaded {
		w := dragonfly.Server.World(name)
		o.Printf(dragonfly.Translation("loaded_world_entry", w.Name(), w.Dimension()))
	}
}

// Allow ...
func (c MWList) Allow(src cmd.Source) bool {
	s, isPlayer := src.(*player.Player)

	if isPlayer && !dragonfly.IsOP(s.XUID()) {
		return false
	}

	return true
}

// MWTeleport is used to teleport to the world with the provided name
type MWTeleport struct {
	Teleport cmd.SubCommand `cmd:"teleport"`
	Name     string         `cmd:"name"`
}

// Run ...
func (c MWTeleport) Run(src cmd.Source, o *cmd.Output) {
	w := dragonfly.Server.World(c.Name)

	if w == nil {
		o.Printf(dragonfly.Translation("world_not_found", c.Name))
		return
	}

	p := src.(*player.Player)
	p.World().RemoveEntity(p)
	w.AddEntity(p)

	o.Printf(dragonfly.Translation("teleported_to_world", c.Name))
}

// Allow ...
func (c MWTeleport) Allow(src cmd.Source) bool {
	s, isPlayer := src.(*player.Player)

	if !isPlayer || !dragonfly.IsOP(s.XUID()) {
		return false
	}

	return true
}
