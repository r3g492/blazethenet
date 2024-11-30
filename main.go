package main

import (
	"blazethenet/game"
	rl "github.com/gen2brain/raylib-go/raylib"
	"path/filepath"
)

const (
	gameTitle           string = "Hivemind Invader Game"
	fps                 int32  = 60
	ButtonImageFilePath        = "resources/button.png"
	InOneFile                  = 3
)

var screenWidth int32 = 1600
var screenHeight int32 = 900
var userMonitorWidth int
var userMonitorHeight int
var userMonitorCount int
var isFullScreen bool = false
var isGameOn bool = true
var currentFont rl.Font

const (
	InMainMenu = iota
	InGame
)

func main() {

	rl.InitWindow(screenWidth, screenHeight, gameTitle)

	userMonitorCount = rl.GetMonitorCount()
	userMonitorWidth = rl.GetMonitorWidth(0)
	userMonitorHeight = rl.GetMonitorHeight(0)

	rl.SetTargetFPS(fps)

	gameState := InMainMenu

	buttonTexture,
		oneButtonTextureHeight,
		buttonImageRectangle := loadButtonTexture()
	defer unloadButtonTexture(buttonTexture)
	var buttonWidth int32 = 400
	var buttonHeight int32 = 200

	fontPath := filepath.Join("resources", "font", "Noto_Sans_KR", "static", "NotoSansKR-ExtraBold.ttf")
	currentFont = rl.LoadFontEx(fontPath, 32, nil, 65535)
	defer rl.UnloadFont(currentFont)
	var fontSize int32 = 150

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

			break

		case InGame:
			game.Logic(currentFont)
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

func loadButtonTexture() (rl.Texture2D, int32, rl.Rectangle) {
	buttonTexture := rl.LoadTexture(ButtonImageFilePath)
	oneButtonTextureHeight := buttonTexture.Height / 3
	buttonImageRectangle := rl.Rectangle{Width: float32(buttonTexture.Width), Height: float32(oneButtonTextureHeight)}
	return buttonTexture, oneButtonTextureHeight, buttonImageRectangle
}

func unloadButtonTexture(buttonTexture rl.Texture2D) {
	rl.UnloadTexture(buttonTexture)
}
