package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth  int32  = 1000
	screenHeight int32  = 720
	gameTitle    string = "Detective Game"
	fps          int32  = 60
	buttonFile   string = "resources/button.png"
)

const (
	StartMenu = iota
	SettingsMenu
	ExitMenu
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, gameTitle)
	rl.SetTargetFPS(fps)

	button := rl.LoadTexture(buttonFile)
	defer rl.UnloadTexture(button)

	buttons := []rl.Rectangle{
		{X: float32(screenWidth/2 - button.Width/2), Y: float32(screenHeight/2 - button.Height/3), Width: float32(button.Width), Height: float32(button.Height / 3)},
		{X: float32(screenWidth/2 - button.Width/2), Y: float32(screenHeight / 2), Width: float32(button.Width), Height: float32(button.Height / 3)},
		{X: float32(screenWidth/2 - button.Width/2), Y: float32(screenHeight/2 + button.Height/3), Width: float32(button.Width), Height: float32(button.Height / 3)},
	}

	texts := []string{
		"Start",
		"Settings",
		"Exit",
	}

	buttonActions := make([]bool, len(buttons))
GameLoop:
	for !rl.WindowShouldClose() {
		mousePoint := rl.GetMousePosition()

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		for i, btnBounds := range buttons {
			buttonActions[i] = DrawButton(button, btnBounds, mousePoint, texts[i])

			if buttonActions[StartMenu] {

			}

			if buttonActions[SettingsMenu] {

			}

			if buttonActions[ExitMenu] {
				break GameLoop
			}
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func DrawButton(button rl.Texture2D, bounds rl.Rectangle, mousePoint rl.Vector2, text string) bool {
	frameHeight := button.Height / 3
	sourceRec := rl.Rectangle{Width: float32(button.Width), Height: float32(frameHeight)}

	var btnState int32
	var btnAction bool

	if rl.CheckCollisionPointRec(mousePoint, bounds) {
		if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			btnState = 2
		} else {
			btnState = 1
		}

		if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
			btnAction = true
		}
	} else {
		btnState = 0
	}

	sourceRec.Y = float32(btnState * frameHeight)

	rl.DrawTextureRec(button, sourceRec, rl.Vector2{X: bounds.X, Y: bounds.Y}, rl.White)

	textWidth := rl.MeasureText(text, 20)
	textX := bounds.X + (bounds.Width/2 - float32(textWidth)/2)
	textY := bounds.Y + (bounds.Height/2 - float32(20)/2)
	rl.DrawText(text, int32(textX), int32(textY), 20, rl.White)

	return btnAction
}
