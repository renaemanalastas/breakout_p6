package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	// "fmt"
	"slices"
)

const (
	TileSize = 64
	ScreenWidth = int32(800)
	ScreenHeight = int32(450)
)

/* TODO

*/

func main() {
	paddle := rl.NewRectangle(float32(ScreenWidth/2 - 125/2), float32(ScreenHeight - 50), 125, 5)
	paddleSpeed := float32(200)

	ballVel := rl.NewVector2(0, 300)
	ballColor := rl.Red
	ballRad := float32(10)
	ballInitPos := rl.NewVector2(float32(paddle.X + 100/2), float32(paddle.Y - ballRad))
	ball := newBall(ballInitPos, ballVel, ballRad, ballColor)

	blockGrid := initGrid()

	rl.InitWindow(ScreenWidth, ScreenHeight, "Breakout")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose(){
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// block grid
		drawGrid(blockGrid)

		// paddle
		rl.DrawRectangleRec(paddle, rl.White)

		// paddle portions (...)
		paddle1 := rl.NewRectangle(paddle.X, paddle.Y, paddle.Width/7, paddle.Height)
		paddle2 := rl.NewRectangle((paddle.X + paddle.Width/7), paddle.Y, paddle.Width/7, paddle.Height)
		paddle3 := rl.NewRectangle((paddle.X + paddle.Width/7*2), paddle.Y, paddle.Width/7, paddle.Height)
		paddle4 := rl.NewRectangle((paddle.X + paddle.Width/7*3), paddle.Y, paddle.Width/7, paddle.Height)
		paddle5 := rl.NewRectangle((paddle.X + paddle.Width/7*4), paddle.Y, paddle.Width/7, paddle.Height)
		paddle6 := rl.NewRectangle((paddle.X + paddle.Width/7*5), paddle.Y, paddle.Width/7, paddle.Height)
		paddle7 := rl.NewRectangle((paddle.X + paddle.Width/7*6), paddle.Y, paddle.Width/7, paddle.Height)

		// paddle portions update according to paddle pos
		rl.DrawRectangleRec(paddle1, rl.Blank)
		rl.DrawRectangleRec(paddle2, rl.Blank)
		rl.DrawRectangleRec(paddle3, rl.Blank)
		rl.DrawRectangleRec(paddle4, rl.Blank)
		rl.DrawRectangleRec(paddle5, rl.Blank)
		rl.DrawRectangleRec(paddle6, rl.Blank)
		rl.DrawRectangleRec(paddle7, rl.Blank)

		// reset ball if it flies off screen
		if int32(ball.Position.Y - ball.Radius*1.5) > ScreenHeight{
			blockGrid = initGrid()
			ball.resetBall(rl.NewVector2(float32(paddle.X + 125/2), float32(paddle.Y - ballRad)), ballVel)
		}

		// keep ball on paddle until it is off/launched
		if !ball.offPaddle {
			ball.Position = rl.NewVector2(float32(paddle.X + 125/2), float32(paddle.Y - ballRad))
		}

		ball.Direction = 0
		//paddle movement, trap within bounds of screen
		if rl.IsKeyDown(rl.KeyA) && paddle.X > 0{
			ball.Direction = -1
			paddle.X -= paddleSpeed * rl.GetFrameTime()
		}
		if rl.IsKeyDown(rl.KeyD) && int32(paddle.X + paddle.Width) < ScreenWidth{
			ball.Direction = 1
			paddle.X += paddleSpeed * rl.GetFrameTime()
		}

		if rl.IsKeyPressed(rl.KeySpace) && !ball.offPaddle{
			ball.offPaddle = true
			ball.Vel = rl.Vector2Subtract(ball.Vel, rl.NewVector2(float32(ball.Direction)*100, 100))
			// fmt.Printf("%f, %f", ball.Vel.X, ball.Vel.Y)
		}

		// ball movement, launch ball off boundary as long as its off the paddle
		if ball.offPaddle && ball.CheckBoundaries() {
			ball.Position = rl.Vector2Subtract(ball.Position, rl.Vector2Scale(ball.Vel, rl.GetFrameTime()))
		}

		// check for collision between ball and paddle
		if rl.CheckCollisionCircleRec(ball.Position, ball.Radius, paddle) && ball.offPaddle{
			// really horrible way to do it.
			if rl.CheckCollisionCircleRec(ball.Position, ball.Radius, paddle1){
				ball.Vel = rl.Vector2Subtract(ball.Vel, rl.NewVector2(ball.Vel.X + 270, 0))
			} else if rl.CheckCollisionCircleRec(ball.Position, ball.Radius, paddle2){
				ball.Vel = rl.Vector2Subtract(ball.Vel, rl.NewVector2(ball.Vel.X + 180, 0))
			} else if rl.CheckCollisionCircleRec(ball.Position, ball.Radius, paddle3){
				ball.Vel = rl.Vector2Subtract(ball.Vel, rl.NewVector2(ball.Vel.X + 90, 0))
			} else if rl.CheckCollisionCircleRec(ball.Position, ball.Radius, paddle4){
				ball.Vel = rl.Vector2Subtract(ball.Vel, rl.NewVector2(ball.Vel.X + 0, 0))
			} else if rl.CheckCollisionCircleRec(ball.Position, ball.Radius, paddle5){
				ball.Vel = rl.Vector2Subtract(ball.Vel, rl.NewVector2(ball.Vel.X + -90, 0))
			} else if rl.CheckCollisionCircleRec(ball.Position, ball.Radius, paddle6){
				ball.Vel = rl.Vector2Subtract(ball.Vel, rl.NewVector2(ball.Vel.X + -180, 0))
			} else if rl.CheckCollisionCircleRec(ball.Position, ball.Radius, paddle7){
				ball.Vel = rl.Vector2Subtract(ball.Vel, rl.NewVector2(ball.Vel.X + -270, 0))
			}

			ball.Bounce()
		}

		// check collision with blocks in grid
		for i, block := range blockGrid{
			if rl.CheckCollisionCircleRec(ball.Position, ball.Radius, block.BlockTile){
				//remove from slice or refactor into a map idk
				//increase y velocity to speed up
				slices.Delete(blockGrid, i, i+1)
				ball.Vel.Y *= 1.075
				// ball.Vel = rl.Vector2Add(ball.Vel, rl.Vector2Scale(ball.Vel, 1.001))
				ball.Bounce()
			}
		}

		// check collision with wall and change direction
		if !ball.CheckBoundaries(){
			if ball.Position.X - ball.Radius <= 0{
				ball.Position.X = ball.Radius
				ball.Vel.X *= -1
			}

			if ball.Position.X + ball.Radius >= float32(ScreenWidth) {
				ball.Position.X = float32(ScreenWidth) - ball.Radius
				ball.Vel.X *= -1
			}

			if ball.Position.Y - ball.Radius <= 0 {
				ball.Position.Y = ball.Radius
				ball.Vel.Y *= -1
			}

			ball.Position = rl.Vector2Subtract(ball.Position, rl.Vector2Scale(ball.Vel, rl.GetFrameTime()))
		}

		// fmt.Println(ball.Vel)
		ball.PhysicsUpdate()
		ball.DrawBall()

		rl.EndDrawing()
	}
}