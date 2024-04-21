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
	Add    cmd.SubCommand `cmd:"add"`
	Amount int            `name:"amount"`
}

// Run ...
func (c TimeAdd) Run(src cmd.Source, o *cmd.Output) {
	w := src.World()
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
	Query cmd.SubCommand `cmd:"query"`
}

// Run ...
func (TimeQuery) Run(src cmd.Source, o *cmd.Output) {
	w := src.World()
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
	Set  cmd.SubCommand `cmd:"set"`
	Time int            `name:"time"`
}

// TimeSetSpec command is used to set the time in the world of the command
// executor from one of the enum options
type TimeSetSpec struct {
	Set  cmd.SubCommand `cmd:"set"`
	Time spec           `name:"time"`
}

// Run ...
func (c TimeSet) Run(source cmd.Source, output *cmd.Output) {
	setTime(source, output, c.Time)
}

// Allow ...
func (c TimeSet) Allow(src cmd.Source) bool {
	s, isPlayer := src.(*player.Player)

	if isPlayer && !dragonfly.IsOP(s.XUID()) {
		return false
	}

	return true
}

// Run ...
func (c TimeSetSpec) Run(source cmd.Source, output *cmd.Output) {
	setTime(source, output, timeSpecToAmount[c.Time])
}

// Allow ...
func (c TimeSetSpec) Allow(src cmd.Source) bool {
	s, isPlayer := src.(*player.Player)

	if isPlayer && !dragonfly.IsOP(s.XUID()) {
		return false
	}

	return true
}

// setTime is called to set the time of the world of the command source to
// the provided value
func setTime(source cmd.Source, output *cmd.Output, t int) {
	w := source.World()
	w.SetTime(t)

	output.Printf(dragonfly.Translation("time_set", w.Name(), t))
}
