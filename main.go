package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	rl.SetTargetFPS(60)

	var name string = ""

	var nameCreated bool = false

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		var charPressed int32 = rl.GetCharPressed()

		if !nameCreated {
			rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)
			rl.DrawText("Input your name:", 190, 240, 20, rl.LightGray)

			if charPressed != 0 {
				name += string(charPressed)
			}

			if rl.IsKeyPressed(rl.KeyBackspace) && len(name) > 0 {
				name = name[:len(name)-1]
			}

			if rl.IsKeyPressed(rl.KeyEnter) {
				nameCreated = true
			}

			rl.DrawText(name, 190, 280, 20, rl.LightGray)
		} else {
			rl.DrawText("Welcome to the game, "+name, 190, 200, 20, rl.LightGray)
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
