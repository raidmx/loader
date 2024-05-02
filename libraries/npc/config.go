package npc

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/stcraft/dragonfly/server/player"
	"github.com/stcraft/dragonfly/server/player/skin"
	"github.com/stcraft/dragonfly/server/world"
)

const (
	InteractionTypeLeftClick byte = iota
	InteractionTypeRightClick
	InteractionTypeBoth
)

// OnInteract is the function that should be called once a player interacts with the NPC.
// It is called by the NPC library.
type OnInteract = func(p *player.Player)

// Config represents a NPC config that can be used to enable and disable
// various settings related to the NPC before spawning the NPC
type Config struct {
	// Name is the name that should be displayed on top of the Entity
	Name string

	// Skin is the skin that the NPC should be spawned with
	Skin skin.Skin

	// Pos is the position to spawn the NPC in the world
	Pos mgl64.Vec3

	// World is the world in which the NPC should be spawned
	World *world.World

	// LookAtPlayer is whether the NPC should look at the Player if the player
	// is in a certain range
	LookAtPlayer bool

	// InteractionType is the type of interaction for which the NPC's actions should
	// be triggered. It is one of the constants which can be found above.
	InteractionType byte

	// Interaction is the logic that should be executed after an interaction is successful.
	// The player that interacted with the NPC is passed in the parameter of this function
	Interaction OnInteract
}
