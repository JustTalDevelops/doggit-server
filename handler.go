package main

import (
	"github.com/JustTalDevelops/doggit"
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/entity"
	"github.com/df-mc/dragonfly/dragonfly/entity/damage"
	"github.com/df-mc/dragonfly/dragonfly/entity/healing"
	"github.com/df-mc/dragonfly/dragonfly/event"
	"github.com/df-mc/dragonfly/dragonfly/item"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/df-mc/dragonfly/dragonfly/world"
	"github.com/go-gl/mathgl/mgl64"
	"net"
)

type PlayerHandler struct {
	handlers []doggit.Handler
	player   *player.Player
}

func StartHandling(player *player.Player, handlers []doggit.Handler) *PlayerHandler {
	for _, h := range handlers {
		h.HandleJoin(player)
	}
	return &PlayerHandler{player: player, handlers: handlers}
}

func (p PlayerHandler) HandleMove(ctx *event.Context, newPos mgl64.Vec3, newYaw, newPitch float64) {
	for _, h := range p.handlers {
		h.HandleMove(p.player, ctx, newPos, newYaw, newPitch)
	}
}

func (p PlayerHandler) HandleTeleport(ctx *event.Context, pos mgl64.Vec3) {
	for _, h := range p.handlers {
		h.HandleTeleport(p.player, ctx, pos)
	}
}

func (p PlayerHandler) HandleChat(ctx *event.Context, message *string) {
	for _, h := range p.handlers {
		h.HandleChat(p.player, ctx, message)
	}
}

func (p PlayerHandler) HandleFoodLoss(ctx *event.Context, from, to int) {
	for _, h := range p.handlers {
		h.HandleFoodLoss(p.player, ctx, from, to)
	}
}

func (p PlayerHandler) HandleHeal(ctx *event.Context, health *float64, src healing.Source) {
	for _, h := range p.handlers {
		h.HandleHeal(p.player, ctx, health, src)
	}
}

func (p PlayerHandler) HandleHurt(ctx *event.Context, damage *float64, src damage.Source) {
	for _, h := range p.handlers {
		h.HandleHurt(p.player, ctx, damage, src)
	}
}

func (p PlayerHandler) HandleDeath(src damage.Source) {
	for _, h := range p.handlers {
		h.HandleDeath(p.player, src)
	}
}

func (p PlayerHandler) HandleRespawn(pos *mgl64.Vec3) {
	for _, h := range p.handlers {
		h.HandleRespawn(p.player, pos)
	}
}

func (p PlayerHandler) HandleStartBreak(ctx *event.Context, pos world.BlockPos) {
	for _, h := range p.handlers {
		h.HandleStartBreak(p.player, ctx, pos)
	}
}

func (p PlayerHandler) HandleBlockBreak(ctx *event.Context, pos world.BlockPos) {
	for _, h := range p.handlers {
		h.HandleBlockBreak(p.player, ctx, pos)
	}
}

func (p PlayerHandler) HandleBlockPlace(ctx *event.Context, pos world.BlockPos, b world.Block) {
	for _, h := range p.handlers {
		h.HandleBlockPlace(p.player, ctx, pos, b)
	}
}

func (p PlayerHandler) HandleBlockPick(ctx *event.Context, pos world.BlockPos, b world.Block) {
	for _, h := range p.handlers {
		h.HandleBlockPick(p.player, ctx, pos, b)
	}
}

func (p PlayerHandler) HandleItemUse(ctx *event.Context) {
	for _, h := range p.handlers {
		h.HandleItemUse(p.player, ctx)
	}
}

func (p PlayerHandler) HandleItemUseOnBlock(ctx *event.Context, pos world.BlockPos, face world.Face, clickPos mgl64.Vec3) {
	for _, h := range p.handlers {
		h.HandleItemUseOnBlock(p.player, ctx, pos, face, clickPos)
	}
}

func (p PlayerHandler) HandleItemUseOnEntity(ctx *event.Context, e world.Entity) {
	for _, h := range p.handlers {
		h.HandleItemUseOnEntity(p.player, ctx, e)
	}
}

func (p PlayerHandler) HandleAttackEntity(ctx *event.Context, e world.Entity) {
	for _, h := range p.handlers {
		h.HandleAttackEntity(p.player, ctx, e)
	}
}

func (p PlayerHandler) HandleItemDamage(ctx *event.Context, i item.Stack, damage int) {
	for _, h := range p.handlers {
		h.HandleItemDamage(p.player, ctx, i, damage)
	}
}

func (p PlayerHandler) HandleItemPickup(ctx *event.Context, i item.Stack) {
	for _, h := range p.handlers {
		h.HandleItemPickup(p.player, ctx, i)
	}
}

func (p PlayerHandler) HandleItemDrop(ctx *event.Context, e *entity.Item) {
	for _, h := range p.handlers {
		h.HandleItemDrop(p.player, ctx, e)
	}
}

func (p PlayerHandler) HandleTransfer(ctx *event.Context, addr *net.UDPAddr) {
	for _, h := range p.handlers {
		h.HandleTransfer(p.player, ctx, addr)
	}
}

func (p PlayerHandler) HandleCommandExecution(ctx *event.Context, command cmd.Command, args []string) {
	for _, h := range p.handlers {
		h.HandleCommandExecution(p.player, ctx, command, args)
	}
}

func (p PlayerHandler) HandleQuit() {
	for _, h := range p.handlers {
		h.HandleQuit(p.player)
	}
}
