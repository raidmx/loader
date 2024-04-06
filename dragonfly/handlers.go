package dragonfly

import (
	"net"
	"time"

	"github.com/STCraft/DFLoader/db"
	"github.com/STCraft/dragonfly/server/block/cube"
	"github.com/STCraft/dragonfly/server/cmd"
	"github.com/STCraft/dragonfly/server/event"
	"github.com/STCraft/dragonfly/server/item"
	"github.com/STCraft/dragonfly/server/player"
	"github.com/STCraft/dragonfly/server/player/skin"
	"github.com/STCraft/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

// DefaultHandler implements the Handler interface to handle DFLoader's events.
type DefaultHandler struct{}

func (DefaultHandler) HandleJoin(ctx *event.Context, xuid string) {
	p, ok := Server.PlayerByXUID(xuid)
	if !ok {
		return
	}

	if !db.IsUser(xuid) {
		db.CreateUser(xuid, p.Name())
	} else {
		db.UpdateUser(xuid, p.Name())
	}
}

func (DefaultHandler) HandleItemDrop(*event.Context, world.Entity)                                {}
func (DefaultHandler) HandleMove(*event.Context, mgl64.Vec3, float64, float64)                    {}
func (DefaultHandler) HandleJump()                                                                {}
func (DefaultHandler) HandleTeleport(*event.Context, mgl64.Vec3)                                  {}
func (DefaultHandler) HandleChangeWorld(*world.World, *world.World)                               {}
func (DefaultHandler) HandleToggleSprint(*event.Context, bool)                                    {}
func (DefaultHandler) HandleToggleSneak(*event.Context, bool)                                     {}
func (DefaultHandler) HandleCommandExecution(*event.Context, cmd.Command, []string)               {}
func (DefaultHandler) HandleTransfer(*event.Context, *net.UDPAddr)                                {}
func (DefaultHandler) HandleChat(*event.Context, *string)                                         {}
func (DefaultHandler) HandleSkinChange(*event.Context, *skin.Skin)                                {}
func (DefaultHandler) HandleStartBreak(*event.Context, cube.Pos)                                  {}
func (DefaultHandler) HandleBlockBreak(*event.Context, cube.Pos, *[]item.Stack, *int)             {}
func (DefaultHandler) HandleBlockPlace(*event.Context, cube.Pos, world.Block)                     {}
func (DefaultHandler) HandleBlockPick(*event.Context, cube.Pos, world.Block)                      {}
func (DefaultHandler) HandleSignEdit(*event.Context, bool, string, string)                        {}
func (DefaultHandler) HandleLecternPageTurn(*event.Context, cube.Pos, int, *int)                  {}
func (DefaultHandler) HandleItemPickup(*event.Context, *item.Stack)                               {}
func (DefaultHandler) HandleItemUse(*event.Context)                                               {}
func (DefaultHandler) HandleItemUseOnBlock(*event.Context, cube.Pos, cube.Face, mgl64.Vec3)       {}
func (DefaultHandler) HandleItemUseOnEntity(*event.Context, world.Entity)                         {}
func (DefaultHandler) HandleItemConsume(*event.Context, item.Stack)                               {}
func (DefaultHandler) HandleItemDamage(*event.Context, item.Stack, int)                           {}
func (DefaultHandler) HandleAttackEntity(*event.Context, world.Entity, *float64, *float64, *bool) {}
func (DefaultHandler) HandleExperienceGain(*event.Context, *int)                                  {}
func (DefaultHandler) HandlePunchAir(*event.Context)                                              {}
func (DefaultHandler) HandleHurt(*event.Context, *float64, *time.Duration, world.DamageSource)    {}
func (DefaultHandler) HandleHeal(*event.Context, *float64, world.HealingSource)                   {}
func (DefaultHandler) HandleFoodLoss(*event.Context, int, *int)                                   {}
func (DefaultHandler) HandleDeath(world.DamageSource, *bool)                                      {}
func (DefaultHandler) HandleRespawn(*mgl64.Vec3, **world.World)                                   {}
func (DefaultHandler) HandleQuit()                                                                {}

// handlers is a map of player handlers registered by the plugins and libraries using DFLoader.
var handlers = map[string]player.Handler{
	"default": DefaultHandler{},
}

// RegisterHandler registers the specified handler with the provided name. This returns false if a handler
// with the specified name already exists.
func RegisterHandler(key string, h player.Handler) bool {
	if _, ok := handlers[key]; ok {
		return false
	}

	handlers[key] = h
	return true
}
