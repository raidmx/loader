package dragonfly

import (
	"github.com/STCraft/dragonfly/server"
	"github.com/STCraft/dragonfly/server/player"
)

// Server is a global instance of dragonfly
var Server *server.Server

// New creates and returns a New instance of the dragonfly server and stores it in the Server variable.
// This is a blocking function as it runs the Start() method on the dragonfly instance which starts
// listening for connections on all the listeners configured.
func New() {
	defer func() {
		saveOperators()
	}()

	loadLanguage()
	loadOperators()

	Server = server.New()
	Server.CloseOnProgramEnd()
	Server.Start()

	for Server.Accept(func(p *player.Player) {
		for key, h := range handlers {
			p.AddHandler(key, h)
		}
	}) {
	}
}
