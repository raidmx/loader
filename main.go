package main

import (
	"github.com/stcraft/DFLoader/loader"
	"github.com/stcraft/DFLoader/plugins/npc"
	"github.com/stcraft/DFLoader/plugins/staffmode"
)

/*
 * This is an example main.go file that is used to start a basic Dragonfly server on top
 * of which the DFLoader adds various functionalities through in built libraries.
 */
func main() {
	loader.LoadPlugin(staffmode.Plugin{})
	loader.LoadPlugin(npc.Plugin{})

	loader.Init()
	loader.Start()
	loader.Deinit()
}
