package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Physics struct {
	Pos rl.Vector2
	Vel rl.Vector2
}

func newPhysics(pos rl.Vector2, vel rl.Vector2) Physics{
	return Physics{
		Pos:	pos,
		Vel:	vel,
	}

}

func (p *Physics) PhysicsUpdate() {
	p.VelocityTick()
}

func (p *Physics) VelocityTick() {
	adjustedVel := rl.Vector2Scale(p.Vel, rl.GetFrameTime())
	p.Pos = rl.Vector2Add(p.Pos, adjustedVel)
}