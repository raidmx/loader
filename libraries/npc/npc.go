package npc

import (
	"math"
	"time"

	"github.com/STCraft/dragonfly/server/block/cube"
	"github.com/STCraft/dragonfly/server/player"
	"github.com/STCraft/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
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
	world           *world.World
}

// Spawn spawns a new NPC with the provided configuration. It registers the NPC
// to the NPCRegistry and also returns an instance to it.
func Spawn(config Config) *NPC {
	p := player.New(config.Name, config.Skin, config.Pos)
	p.SetNameTag(config.Name)
	p.SetMobile()
	p.SetVisible()

	npc := &NPC{
		Player:          p,
		lookAtPlayer:    config.LookAtPlayer,
		interactionType: config.InteractionType,
		interaction:     config.Interaction,
		world:           config.World,
	}

	config.World.AddEntity(npc)

	if npc.lookAtPlayer {
		go npc.lookAroundTask()
	}

	return npc
}

// Type returns the world.EntityType for the NPC.
func (n *NPC) Type() world.EntityType {
	return NPCType{}
}

// lookAroundTask is used to make the entity look at the players nearby it
func (n *NPC) lookAroundTask() {
	for n.world != nil {
		nPos := n.Position()
		min := n.Position().Sub(mgl64.Vec3{5, 5, 5})
		max := n.Position().Add(mgl64.Vec3{5, 5, 5})

		bbox := cube.Box(min.X(), min.Y(), min.Z(), max.X(), max.Y(), max.Z())

		entities := n.world.EntitiesWithin(bbox, func(e world.Entity) bool {
			_, ok := e.(*player.Player)
			return !ok
		})

		for _, e := range entities {
			p := e.(*player.Player)
			pos := p.Position()

			angle := math.Atan2(pos.Z()-nPos.Z(), pos.X()-nPos.X())
			yaw := ((angle * 180) / math.Pi) - 90

			px := math.Pow((nPos.X() - pos.X()), 2)
			pz := math.Pow((nPos.Z() - pos.Z()), 2)

			angle = math.Atan2(px+pz, pos.Y()-nPos.Y())
			pitch := ((angle * 180) / math.Pi) - 90

			p.Session().ViewEntityMovement(n, n.Position(), cube.Rotation{yaw, pitch}, true)
		}

		time.Sleep(time.Millisecond * 50)
	}
}

// Despawn kills and removes the entity from the world.
func (n *NPC) Despawn() {
	n.world.RemoveEntity(n)
	n.world = nil
}
