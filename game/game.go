package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Logic(
	font rl.Font,
	screenWidth int32,
	screenHeight int32,
) bool {
	rl.ClearBackground(rl.SkyBlue)

	rl.DrawTextEx(font, "this is in game!", rl.Vector2{X: 150, Y: 150}, 32, 10, rl.White)
	return true
}
