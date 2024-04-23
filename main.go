package main

import "github.com/STCraft/DFLoader/loader"

/*
 * This is an example main.go file that is used to start a basic Dragonfly server on top
 * of which the DFLoader adds various functionalities through in built libraries.
 */
func main() {
	loader.Init()
	loader.Start()
	loader.Deinit()
}
