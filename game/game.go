package game

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	IsGameInit = false
	Turn       = 0
)

func Init() {
	IsGameInit = true
}

func Logic(
	font rl.Font,
	fontSize float32,
	screenWidth int32,
	screenHeight int32,
) bool {
	rl.ClearBackground(rl.Black)

	// do control
	mousePoint := rl.GetMousePosition()
	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		processPrimaryClick(mousePoint)
	}

	if rl.IsMouseButtonReleased(rl.MouseRightButton) {
		processTurn()
		Turn++
	}

	// do rendering
	printTurn(
		font,
		fontSize,
		float32(screenWidth),
	)

	return true
}

func processPrimaryClick(mousePoint rl.Vector2) {
	rl.DrawCircle(
		int32(mousePoint.X),
		int32(mousePoint.Y),
		50,
		rl.Purple,
	)
}

func processTurn() {
	// TODO: do turn logics
}

func printTurn(
	font rl.Font,
	fontSize float32,
	screenWidth float32,
) {
	text := fmt.Sprintf("Turn %d", Turn)
	textWidth := rl.MeasureTextEx(font, text, fontSize, 10).X
	textX := (screenWidth - textWidth) / 2
	rl.DrawTextEx(font,
		text,
		rl.Vector2{X: textX, Y: 0},
		fontSize,
		0,
		rl.White,
	)
}
