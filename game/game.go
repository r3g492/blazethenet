package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Logic() bool {
	rl.ClearBackground(rl.Black)

	buttonBounds := rl.Rectangle{X: 50, Y: 800, Width: 200, Height: 50}
	rl.DrawRectangleRec(buttonBounds, rl.Gray)
	return true
}
