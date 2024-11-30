package main

import (
	"blazethenet/game"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"path/filepath"
)

const (
	gameTitle           string = "Hivemind Invader Game"
	fps                 int32  = 60
	ButtonImageFilePath        = "resources/button.png"
)

var (
	screenWidth            int32 = 2000
	screenHeight           int32 = 1200
	userMonitorWidth       int
	userMonitorHeight      int
	userMonitorCount       int
	isFullScreen           bool = false
	isGameOn               bool = true
	currentFont            rl.Font
	fontSize               int32
	buttonWidth            int32
	buttonHeight           int32
	buttonTexture          rl.Texture2D
	oneButtonTextureHeight int32
	buttonImageRectangle   rl.Rectangle
	gameState              int
)

const (
	InMainMenu = iota
	InGame
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, gameTitle)
	rl.SetTargetFPS(fps)

	gameState = InMainMenu
	initButton()
	defer rl.UnloadTexture(buttonTexture)
	initFont()
	defer rl.UnloadFont(currentFont)

	for isGameOn {
		mousePoint := rl.GetMousePosition()
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		switch gameState {
		case InMainMenu:
			if rl.WindowShouldClose() {
				isGameOn = false
				break
			}

			startButtonX := (screenWidth - buttonWidth) / 2
			startButtonY := int32(100)
			startButtonRect := rl.Rectangle{
				X:      float32(startButtonX),
				Y:      float32(startButtonY),
				Width:  float32(buttonWidth),
				Height: float32(buttonHeight),
			}
			if buttonControl(
				mousePoint,
				startButtonRect,
				buttonImageRectangle,
				oneButtonTextureHeight,
				buttonTexture,
				"Start",
				currentFont,
				fontSize,
				rl.White,
			) {
				gameState = InGame
			}

			exitButtonX := (screenWidth - buttonWidth) / 2
			exitButtonY := startButtonY + buttonHeight
			exitButtonRect := rl.Rectangle{
				X:      float32(exitButtonX),
				Y:      float32(exitButtonY),
				Width:  float32(buttonWidth),
				Height: float32(buttonHeight),
			}
			if buttonControl(
				mousePoint,
				exitButtonRect,
				buttonImageRectangle,
				oneButtonTextureHeight,
				buttonTexture,
				"Exit",
				currentFont,
				fontSize,
				rl.White,
			) {
				rl.CloseWindow()
				return
			}
			if rl.IsKeyPressed(rl.KeyF3) {
				printMainMenuInfos()
			}

			if rl.IsKeyPressed(rl.KeyF4) {
				changeResolution(1366, 768)
			}

			if rl.IsKeyPressed(rl.KeyF5) {
				changeResolution(1600, 900)
			}

			if rl.IsKeyPressed(rl.KeyF6) {
				changeResolution(1280, 1024)
			}

			if rl.IsKeyPressed(rl.KeyF7) {
				changeResolution(2800, 1400)
			}

			if rl.IsKeyPressed(rl.KeyF8) {
				makeItFullScreen()
			}

			break

		case InGame:
			game.Logic(
				currentFont,
				screenWidth,
				screenHeight,
			)
			if rl.IsKeyPressed(rl.KeyF10) {
				gameState = InMainMenu
			}

			if rl.IsKeyPressed(rl.KeyEscape) {
				gameState = InMainMenu
			}

			if rl.WindowShouldClose() {
				gameState = InMainMenu
			}

			break
		default:
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func printMainMenuInfos() {
	fmt.Println("=== Main Menu Information ===")

	fmt.Println("Game Metadata:")
	fmt.Printf("  Title: %s\n", gameTitle)
	fmt.Printf("  FPS: %d\n", fps)

	fmt.Println("\nScreen and Monitor Information:")
	fmt.Printf("  Screen Resolution: %dx%d\n", screenWidth, screenHeight)
	fmt.Printf("  Monitor Count: %d\n", userMonitorCount)
	fmt.Printf("  Primary Monitor Resolution: %dx%d\n", userMonitorWidth, userMonitorHeight)

	fmt.Println("\nGame State and Settings:")
	fmt.Printf("  Fullscreen: %v\n", isFullScreen)
	fmt.Printf("  Game Running: %v\n", isGameOn)
	fmt.Printf("  Current Game State: %s\n", gameState)

	fmt.Println("\nUI Elements Information:")

	fmt.Println("  Font Details:")
	fmt.Printf("    Font Size: %d\n", fontSize)
	fmt.Printf("    Font Texture ID: %d\n", currentFont.Texture.ID)
	fmt.Printf("    Font Base Size: %d\n", currentFont.BaseSize)

	fmt.Println("  Button Details:")
	fmt.Printf("    Button Width: %d\n", buttonWidth)
	fmt.Printf("    Button Height: %d\n", buttonHeight)
	fmt.Printf("    Button Texture ID: %d\n", buttonTexture.ID)
	fmt.Printf("    Button Texture Dimensions: %dx%d\n", buttonTexture.Width, buttonTexture.Height)
	fmt.Printf("    One Button Texture Height: %d\n", oneButtonTextureHeight)
	fmt.Printf("    Button Image Rectangle: {X: %.2f, Y: %.2f, Width: %.2f, Height: %.2f}\n",
		buttonImageRectangle.X, buttonImageRectangle.Y, buttonImageRectangle.Width, buttonImageRectangle.Height)

	fmt.Println("==============================")
}

func changeResolution(width, height int) {
	if rl.IsWindowFullscreen() {
		rl.ToggleFullscreen()
	}
	screenWidth = int32(width)
	screenHeight = int32(height)
	rl.CloseWindow()
	rl.InitWindow(screenWidth, screenHeight, gameTitle)
	initButton()
	initFont()
}

func makeItFullScreen() {
	if rl.IsWindowFullscreen() {
		return
	}
	display := rl.GetCurrentMonitor()
	userMonitorWidth = rl.GetMonitorWidth(display)
	userMonitorHeight = rl.GetMonitorHeight(display)
	screenWidth = int32(userMonitorWidth)
	screenHeight = int32(userMonitorHeight)

	rl.CloseWindow()
	rl.InitWindow(screenWidth, screenHeight, gameTitle)
	rl.SetWindowPosition(0, 200)
	rl.ToggleFullscreen()
	initButton()
	initFont()
}

func initFont() {
	fontPath := filepath.Join("resources", "font", "Noto_Sans_KR", "static", "NotoSansKR-ExtraBold.ttf")
	fontSize = screenWidth / 24
	currentFont = rl.LoadFontEx(fontPath, min(fontSize, 48), nil, 65535)
}

func initButton() {
	loadButtonTexture()
	buttonWidth = screenWidth / 4
	buttonHeight = screenHeight / 4
}

func buttonControl(
	mousePoint rl.Vector2,
	buttonRect rl.Rectangle,
	buttonImageRectangle rl.Rectangle,
	oneButtonTextureHeight int32,
	buttonTexture rl.Texture2D,
	buttonText string,
	font rl.Font,
	fontSize int32,
	textColor rl.Color,
) bool {
	isMouseOver := rl.CheckCollisionPointRec(mousePoint, buttonRect)
	var buttonTextureIndex int32 = 0
	if isMouseOver {
		buttonTextureIndex = 1
	}
	if isMouseOver && rl.IsMouseButtonDown(rl.MouseLeftButton) {
		buttonTextureIndex = 2
	}
	if isMouseOver && rl.IsMouseButtonReleased(rl.MouseLeftButton) {
		return true
	}
	buttonImageRectangle.Y = float32(buttonTextureIndex * oneButtonTextureHeight)
	rl.DrawTexturePro(
		buttonTexture,
		buttonImageRectangle,
		buttonRect,
		rl.Vector2{X: 0, Y: 0},
		0,
		rl.White,
	)

	textWidth := rl.MeasureTextEx(font, buttonText, float32(fontSize), 1).X
	textHeight := rl.MeasureTextEx(font, buttonText, float32(fontSize), 1).Y
	textPosition := rl.Vector2{
		X: buttonRect.X + (buttonRect.Width-textWidth)/2,
		Y: buttonRect.Y + (buttonRect.Height-textHeight)/2,
	}

	rl.DrawTextEx(
		font,
		buttonText,
		textPosition,
		float32(fontSize),
		1,
		textColor,
	)
	return false
}

func loadButtonTexture() {
	buttonTexture = rl.LoadTexture(ButtonImageFilePath)
	oneButtonTextureHeight = buttonTexture.Height / 3
	buttonImageRectangle = rl.Rectangle{Width: float32(buttonTexture.Width), Height: float32(oneButtonTextureHeight)}
}
