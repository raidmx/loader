package commands

import (
	"time"

	"github.com/STCraft/dragonfly/server/cmd"
	"github.com/STCraft/dragonfly/server/player"
)

// Weather command is used to change the weather of the world the player is in for an optional
// duration if specified.
type Weather struct {
	Weather  weatherType        `cmd:"weather"`
	Duration cmd.Optional[uint] `cmd:"duration"`
}

// Run ..
func (c Weather) Run(src cmd.Source, o *cmd.Output) {
	p, ok := src.(*player.Player)
	if !ok {
		o.Errorf("Please run this command in-game")
		return
	}

	w := p.World()
	t := time.Duration(c.Duration.LoadOr(60)) * time.Second

	switch c.Weather {
	case "clear":
		w.StopRaining()
		w.StopThundering()
		o.Printf("Changing to clear weather")
	case "rain":
		w.StartRaining(t)
		o.Printf("Changing to rainy weather")
	case "thunder":
		w.StartRaining(t)
		w.StartThundering(t)
		o.Printf("Changing to rain and thunder")
	default:
		o.Errorf("Unknown weather type: %s", c.Weather)
		return
	}
}

// weatherType ...
type weatherType string

// Type ...
func (w weatherType) Type() string {
	return "weather"
}

// Options ...
func (w weatherType) Options(source cmd.Source) []string {
	return []string{
		"clear", "rain", "thunder",
	}
}
