package commands

import (
	"github.com/STCraft/DFLoader/dragonfly"
	"github.com/STCraft/dragonfly/server/cmd"
	"github.com/STCraft/dragonfly/server/player"
)

// spec is an enum of the various time options to choose from when
// setting the time of a world
type spec string

// Type ...
func (spec) Type() string {
	return "TimeSpec"
}

// Options ...
func (spec) Options(source cmd.Source) []string {
	return []string{"day", "night", "noon", "midnight", "sunrise", "sunset"}
}

// timeSpecToAmount contains a mapping of the time spec values to the integral amount values
var timeSpecToAmount = map[spec]int{
	"day": 1000, "night": 13000, "noon": 6000, "midnight": 18000, "sunrise": 23000, "sunset": 12000,
}

// TimeAdd is used to add time to the world of the command executor
type TimeAdd struct {
	World  worldName      `cmd:"world"`
	Add    cmd.SubCommand `cmd:"add"`
	Amount int            `name:"amount"`
}

// Run ...
func (c TimeAdd) Run(src cmd.Source, o *cmd.Output) {
	w := dragonfly.Server.World(string(c.World))
	w.SetTime(w.Time() + c.Amount)
	o.Printf(dragonfly.Translation("added_to_time", c.Amount, w.Name()))
}

// Allow ...
func (c TimeAdd) Allow(src cmd.Source) bool {
	s, isPlayer := src.(*player.Player)

	if isPlayer && !dragonfly.IsOP(s.XUID()) {
		return false
	}

	return true
}

// TimeQuery command is used to query the current time in the world of
// the command executor
type TimeQuery struct {
	World worldName      `cmd:"world"`
	Query cmd.SubCommand `cmd:"query"`
}

// Run ...
func (c TimeQuery) Run(src cmd.Source, o *cmd.Output) {
	w := dragonfly.Server.World(string(c.World))
	o.Printf(dragonfly.Translation("time_query", w.Name(), w.Time()))
}

// Allow ...
func (c TimeQuery) Allow(src cmd.Source) bool {
	s, isPlayer := src.(*player.Player)

	if isPlayer && !dragonfly.IsOP(s.XUID()) {
		return false
	}

	return true
}

// TimeSet command is used to set the time in the world of the command
// executor to the provided time
type TimeSet struct {
	World worldName      `cmd:"world"`
	Set   cmd.SubCommand `cmd:"set"`
	Time  int            `name:"time"`
}

// Run ...
func (c TimeSet) Run(src cmd.Source, o *cmd.Output) {
	w := dragonfly.Server.World(string(c.World))
	time := c.Time

	if time > 24000 {
		time -= 24000
	}

	w.SetTime(time)
	o.Printf(dragonfly.Translation("time_set", w.Name(), time))
}

// Allow ...
func (c TimeSet) Allow(src cmd.Source) bool {
	s, isPlayer := src.(*player.Player)

	if isPlayer && !dragonfly.IsOP(s.XUID()) {
		return false
	}

	return true
}

// TimeSetSpec command is used to set the time in the world of the command
// executor from one of the enum options
type TimeSetSpec struct {
	World worldName      `cmd:"world"`
	Set   cmd.SubCommand `cmd:"set"`
	Time  spec           `name:"time"`
}

// Run ...
func (c TimeSetSpec) Run(src cmd.Source, o *cmd.Output) {
	w := dragonfly.Server.World(string(c.World))
	t := timeSpecToAmount[c.Time]

	w.SetTime(t)
	o.Printf(dragonfly.Translation("time_set", w.Name(), t))
}

// Allow ...
func (c TimeSetSpec) Allow(src cmd.Source) bool {
	s, isPlayer := src.(*player.Player)

	if isPlayer && !dragonfly.IsOP(s.XUID()) {
		return false
	}

	return true
}
