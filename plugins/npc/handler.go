package npc

import (
	"github.com/STCraft/dragonfly/server/event"
	"github.com/STCraft/dragonfly/server/player"
	"github.com/STCraft/dragonfly/server/world"
)

// handler is a NPC handler that handles the Player <-> NPC interaction events and triggers
// the OnInteract function that was registered when the NPC was created.
type handler struct {
	player.NopHandler
	p *player.Player
}

// New ...
func (handler) New(p *player.Player) player.Handler {
	return handler{p: p}
}

// HandleItemUseOnEntity ...
func (h handler) HandleItemUseOnEntity(ctx *event.Context, e world.Entity) {
	npc, ok := e.(*NPC)
	if !ok {
		return
	}

	ctx.Cancel()

	if npc.interactionType == InteractionTypeLeftClick {
		return
	}

	npc.interaction(h.p)
}

// HandleAttackEntity ...
func (h handler) HandleAttackEntity(ctx *event.Context, e world.Entity, force *float64, height *float64, critical *bool) {
	npc, ok := e.(*NPC)
	if !ok {
		return
	}

	ctx.Cancel()

	if npc.interactionType == InteractionTypeRightClick {
		return
	}

	npc.interaction(h.p)
}
