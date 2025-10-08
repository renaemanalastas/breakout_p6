package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Ball struct {
	Position rl.Vector2
	Radius float32
	Color rl.Color
	// Rotation float32
	offPaddle bool
	Direction int
	Physics
}

func newBall(p rl.Vector2, v rl.Vector2, r float32, c rl.Color) Ball{
	return Ball{
		Position:	p,
		Radius:		r,
		Color:		c,
		offPaddle:	false,
		Direction: 0, //-1 left, 0 middle, 1 right
		Physics:	newPhysics(p, v),
	}
}

func (b *Ball) DrawBall(){
	rl.DrawCircleV(b.Position, b.Radius, b.Color)
	// rl.DrawPoly(b.Position, 60, b.Radius, b.Rotation, b.Color)
}

// func (b *Ball) launchBall(){
// 	b.Position = rl.Vector2Subtract(b.Position, rl.Vector2Scale(b.Vel, rl.GetFrameTime()))
// }

func (b *Ball) resetBall(ballInitPos rl.Vector2, ballVel rl.Vector2){
	b.Position = ballInitPos
	b.Vel = ballVel
	b.offPaddle = false
	b.DrawBall()
}

//FIXME get the ball to bounce off walls correctly because it keeps bouncing back to the paddle?
func (b *Ball) Bounce(){
	b.Vel = rl.Vector2Scale(b.Vel, -1)
	b.Position = rl.Vector2Subtract(b.Position, rl.Vector2Scale(b.Vel, rl.GetFrameTime()))
}

func (b *Ball) CheckBoundaries() bool{
	if int32(b.Position.Y - b.Radius) > 0 {
		if int32(b.Position.X - b.Radius) > 0 && int32(b.Position.X + b.Radius) < ScreenWidth {
			return true
		}
	}
	return false
}
