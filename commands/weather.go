package commands

import (
	"time"

	"github.com/STCraft/DFLoader/dragonfly"
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
	p := src.(*player.Player)

	w := p.World()
	t := time.Duration(c.Duration.LoadOr(60)) * time.Second

	switch c.Weather {
	case "clear":
		w.StopRaining()
		w.StopThundering()
		o.Print(dragonfly.Translation("weather_changed", "clear"))
	case "rain":
		w.StartRaining(t)
		o.Print(dragonfly.Translation("weather_changed", "rainy"))
	case "thunder":
		w.StartRaining(t)
		w.StartThundering(t)
		o.Print(dragonfly.Translation("weather_changed", "rainy and thunder"))
	default:
		o.Print(dragonfly.Translation("unknown_weather", string(c.Weather)))
		return
	}
}

// Allow ...
func (c Weather) Allow(src cmd.Source) bool {
	if _, ok := src.(*player.Player); ok {
		return false
	}

	return true
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
