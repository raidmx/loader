package staffmode

import (
	"time"

	"github.com/go-gl/mathgl/mgl64"
	"github.com/stcraft/dragonfly/server/block/cube"
	"github.com/stcraft/dragonfly/server/event"
	"github.com/stcraft/dragonfly/server/item"
	"github.com/stcraft/dragonfly/server/player"
	"github.com/stcraft/dragonfly/server/player/skin"
	"github.com/stcraft/dragonfly/server/world"
)

// handler is a Staff Mode Handler that prevents Staff from abusing by breaking
// and placing blocks, dropping items, etc.
type handler struct {
	player.NopHandler
	p *player.Player
}

// New ...
func (handler) New(p *player.Player) player.Handler {
	return handler{p: p}
}

// HandleItemDrop ...
func (h handler) HandleItemDrop(ctx *event.Context, e world.Entity) {
	if enabled(h.p) {
		ctx.Cancel()
	}
}

// HandleStartBreak ...
func (h handler) HandleStartBreak(ctx *event.Context, pos cube.Pos) {
	if enabled(h.p) {
		ctx.Cancel()
	}
}

// HandleBlockBreak ...
func (h handler) HandleBlockBreak(ctx *event.Context, pos cube.Pos, drops *[]item.Stack, xp *int) {
	if enabled(h.p) {
		ctx.Cancel()
	}
}

// HandleBlockPlace
func (h handler) HandleBlockPlace(ctx *event.Context, pos cube.Pos, b world.Block) {
	if enabled(h.p) {
		ctx.Cancel()
	}
}

// HandleBlockPick ...
func (h handler) HandleBlockPick(ctx *event.Context, pos cube.Pos, b world.Block) {
	if enabled(h.p) {
		ctx.Cancel()
	}
}

// HandleSignEdit ...
func (h handler) HandleSignEdit(ctx *event.Context, frontSide bool, oldText string, newText string) {
	if enabled(h.p) {
		ctx.Cancel()
	}
}

// HandleFoodLoss ...
func (h handler) HandleFoodLoss(ctx *event.Context, from int, to *int) {
	if enabled(h.p) {
		ctx.Cancel()
	}
}

// HandleHeal ...
func (h handler) HandleHeal(ctx *event.Context, health *float64, src world.HealingSource) {
	if enabled(h.p) {
		ctx.Cancel()
	}
}

// HandleHurt ...
func (h handler) HandleHurt(ctx *event.Context, damage *float64, attackImmunity *time.Duration, src world.DamageSource) {
	if enabled(h.p) {
		ctx.Cancel()
	}
}

// HandleSkinChange ...
func (h handler) HandleSkinChange(ctx *event.Context, skin *skin.Skin) {
	if enabled(h.p) {
		ctx.Cancel()
	}
}

// HandleItemUse ...
func (h handler) HandleItemUse(ctx *event.Context) {
	if enabled(h.p) {
		ctx.Cancel()
	}
}

// HandleItemUseOnBlock ...
func (h handler) HandleItemUseOnBlock(ctx *event.Context, pos cube.Pos, face cube.Face, clickPos mgl64.Vec3) {
	if enabled(h.p) {
		ctx.Cancel()
	}
}

// HandleItemUseOnEntity ...
func (h handler) HandleItemUseOnEntity(ctx *event.Context, e world.Entity) {
	if enabled(h.p) {
		ctx.Cancel()
	}
}

// HandleItemConsume ...
func (h handler) HandleItemConsume(ctx *event.Context, item item.Stack) {
	if enabled(h.p) {
		ctx.Cancel()
	}
}

// HandleAttackEntity ...
func (h handler) HandleAttackEntity(ctx *event.Context, e world.Entity, force, height *float64, critical *bool) {
	if enabled(h.p) {
		ctx.Cancel()
	}
}

// HandleExperienceGain ...
func (h handler) HandleExperienceGain(ctx *event.Context, amount *int) {
	if enabled(h.p) {
		ctx.Cancel()
	}
}

// HandlePunchAir ...
func (h handler) HandlePunchAir(ctx *event.Context) {
	if enabled(h.p) {
		ctx.Cancel()
	}
}

// HandleItemDamage ...
func (h handler) HandleItemDamage(ctx *event.Context, i item.Stack, damage int) {
	if enabled(h.p) {
		ctx.Cancel()
	}
}

// HandleItemPickup ...
func (h handler) HandleItemPickup(ctx *event.Context, i *item.Stack) {
	if enabled(h.p) {
		ctx.Cancel()
	}
}

// HandleQuit ...
func (h handler) HandleQuit() {
	if enabled(h.p) {
		disable(h.p)
	}
}
