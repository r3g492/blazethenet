package game

import rl "github.com/gen2brain/raylib-go/raylib"

func GameLogic() bool {
	rl.ClearBackground(rl.Black)
	rl.DrawText("Hello!!!!", 400, 400, 100, rl.Blue)
	return true
}
