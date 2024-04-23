package npc

import (
	"github.com/STCraft/dragonfly/server/block/cube"
	"github.com/STCraft/dragonfly/server/player"
	"github.com/STCraft/dragonfly/server/world"
)

// NPCType is a world.EntityType implementation for NPC.
type NPCType struct{}

func (NPCType) EncodeEntity() string   { return "minecraft:player" }
func (NPCType) NetworkOffset() float64 { return 1.62 }
func (NPCType) BBox(e world.Entity) cube.BBox {
	p := e.(*NPC)
	s := p.Scale()

	switch {
	case p.Sneaking():
		return cube.Box(-0.3*s, 0, -0.3*s, 0.3*s, 1.5*s, 0.3*s)
	case p.Gliding(), p.Swimming():
		return cube.Box(-0.3*s, 0, -0.3*s, 0.3*s, 0.6*s, 0.3*s)
	default:
		return cube.Box(-0.3*s, 0, -0.3*s, 0.3*s, 1.8*s, 0.3*s)
	}
}

// NPC is a non playable character that can be spawned in a world and can be highly configured
// to display dialogues or to add custom logic to be triggered when a player interacts with
// the entity.
type NPC struct {
	*player.Player

	lookAtPlayer    bool
	interactionType byte
	interaction     OnInteract
}

// Spawn spawns a new NPC with the provided configuration. It registers the NPC
// to the NPCRegistry and also returns an instance to it.
func Spawn(config Config) *NPC {
	entity := player.New(config.Name, config.Skin, config.Pos)
	entity.SetNameTag(config.Name)

	npc := &NPC{
		Player:          entity,
		lookAtPlayer:    config.LookAtPlayer,
		interactionType: config.InteractionType,
		interaction:     config.Interaction,
	}

	config.World.AddEntity(npc)
	return npc
}

// Type returns the world.EntityType for the NPC.
func (e *NPC) Type() world.EntityType {
	return NPCType{}
}
