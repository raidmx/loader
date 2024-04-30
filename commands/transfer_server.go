package commands

import (
	"fmt"
	"time"

	"github.com/stcraft/DFLoader/dragonfly"
	"github.com/stcraft/dragonfly/server/cmd"
	"github.com/stcraft/dragonfly/server/player"
)

// TransferServer command is used to transfer a player to the specified
// address and port
type TransferServer struct {
	Target []cmd.Target `cmd:"player"`
	IP     string       `cmd:"address"`
	Port   int          `cmd:"port"`
}

// Run ..
func (c TransferServer) Run(src cmd.Source, o *cmd.Output) {
	addr := fmt.Sprintf("%s:%d", c.IP, c.Port)
	count := 0

	for _, t := range c.Target {
		p, ok := t.(*player.Player)
		if !ok {
			continue
		}

		p.Message(dragonfly.Translation("server_transfer_target", addr))
		count += 1
	}

	o.Printf(dragonfly.Translation("server_transfer_sender", count, addr))

	go func() {
		time.Sleep(time.Second * 2)

		for _, t := range c.Target {
			p, ok := t.(*player.Player)
			if !ok {
				continue
			}

			p.Transfer(addr)
		}
	}()
}

// Allow ...
func (c TransferServer) Allow(src cmd.Source) bool {
	s, isPlayer := src.(*player.Player)

	if isPlayer && !dragonfly.IsOP(s.XUID()) {
		return false
	}

	return true
}
