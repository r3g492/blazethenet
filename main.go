package main

import rl "github.com/gen2brain/raylib-go/raylib"

type GameState int

const (
	NameCreation = iota
	Playing
)

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	rl.SetTargetFPS(60)

	var name string = ""

	// var gameState GameState = NameCreation

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.DrawText("Input your name:", 190, 240, 20, rl.LightGray)

		var charPressed int32 = rl.GetCharPressed()
		if charPressed != 0 {
			name += string(charPressed)
		}

		if rl.IsKeyPressed(rl.KeyBackspace) && len(name) > 0 {
			name = name[:len(name)-1]
		}

		rl.DrawText(name, 190, 260, 20, rl.LightGray)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
