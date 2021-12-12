package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	width  int32 = 800
	height int32 = 450
)

const (
	stepSpeed = 4
	ballSpeed = 10
)

type PongBall struct {
	X         int32
	Y         int32
	Direction rl.Vector2
	Speed     int32
}

func (p PongBall) Pos() rl.Vector2 {
	return rl.Vector2{float32(p.X), float32(p.Y)}
}

func (p PongBall) Draw() {
	rl.DrawCircle(p.X, p.Y, 10, rl.Red)
}

func (p *PongBall) Move() {
	p.X += int32(p.Direction.X)
	p.Y += int32(p.Direction.Y)
}

func (p *PongBall) Bounce() {
	p.Direction.X *= -1
	p.Direction.Y *= -1
	p.Direction.X += rl.Clamp(rand.Float32(), -1, 1)
	p.Direction.Y += rl.Clamp(rand.Float32(), -1, 1)
}

type Paddle struct {
	X int32
	Y int32
}

func (pd Paddle) Draw() {
	rl.DrawRectangle(pd.X, pd.Y, 15, height/8, rl.Black)
}

func (pd Paddle) Rect() rl.Rectangle {
	return rl.Rectangle{float32(pd.X), float32(pd.Y), 15, float32(height / 2)}
}

func main() {
	rl.InitWindow(width, height, "Pong")

	rl.SetTargetFPS(60)

	ball := PongBall{width / 2, height / 2, rl.Vector2{X: -1.0, Y: 0.0}, ballSpeed}

	leftPaddle := Paddle{width / 8, height / 2}
	rightPaddle := Paddle{width - (width / 8), height / 2}

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.White)

		rightPaddle.Draw()
		leftPaddle.Draw()

		ball.Draw()
		ball.Move()

		if ball.X < 0 || ball.X > width || ball.Y < 0 || ball.Y > height || rl.CheckCollisionCircleRec(ball.Pos(), 10, leftPaddle.Rect()) || rl.CheckCollisionCircleRec(ball.Pos(), 10, rightPaddle.Rect()) {
			ball.Bounce()
		}

		if rl.IsKeyDown(rl.KeyUp) && rightPaddle.Y > 0 {
			rightPaddle.Y -= stepSpeed
		}
		if rl.IsKeyDown(rl.KeyDown) && rightPaddle.Y < height-(height/8) {
			rightPaddle.Y += stepSpeed
		}

		if rl.IsKeyDown(rl.KeyW) && leftPaddle.Y > 0 {
			leftPaddle.Y -= stepSpeed
		}
		if rl.IsKeyDown(rl.KeyS) && leftPaddle.Y < height-(height/8) {
			leftPaddle.Y += stepSpeed
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
