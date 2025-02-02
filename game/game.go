package game

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ButtonImageFilePath = "resources/button.png"
)

var (
	IsGameInit = false
	Turn       = 0

	// button settings
	buttonWidth            int32
	buttonHeight           int32
	buttonTexture          rl.Texture2D
	oneButtonTextureHeight int32
	buttonImageRectangle   rl.Rectangle
	mergeMap               MergeMap
	linkMap                LinkMap
)

func Init(
	font rl.Font,
	fontSize float32,
	screenWidth int32,
	screenHeight int32,
	mergeWidth int32,
	mergeHeight int32,
) {
	IsGameInit = true
	initButton(
		screenWidth,
		screenHeight,
	)
	mergeMap = CreateMerge(
		mergeWidth,
		mergeHeight,
		screenWidth,
		screenHeight,
	)
	// TODO: delete this test codes
	mergeMap.AddFireGen(0)
	linkMap = CreateLinkMap()
}

func Unload() {
	IsGameInit = false
	unloadButton()
}

func ReInit() {
	IsGameInit = false
}

func Game(
	font rl.Font,
	fontSize float32,
	screenWidth int32,
	screenHeight int32,
) bool {
	rl.ClearBackground(rl.Black)

	mousePoint := rl.GetMousePosition()
	mergeMap.Control(mousePoint)
	linkMap.Control(mousePoint)

	endTurnButtonX := screenWidth - buttonWidth
	endTurnButtonY := screenHeight - buttonHeight
	endTurnButtonRect := rl.Rectangle{
		X:      float32(endTurnButtonX),
		Y:      float32(endTurnButtonY),
		Width:  float32(buttonWidth),
		Height: float32(buttonHeight),
	}
	if buttonControl(
		mousePoint,
		endTurnButtonRect,
		buttonImageRectangle,
		oneButtonTextureHeight,
		buttonTexture,
		"End Turn",
		font,
		int32(fontSize),
		rl.White,
	) {
		processTurn()
		Turn++
	}

	mergeMap.Render(
		font,
		fontSize,
		screenWidth,
		screenHeight,
		mousePoint,
	)
	mergeMap.Control(
		mousePoint,
	)
	linkMap.Render()

	// do rendering
	printTurn(
		font,
		fontSize,
		float32(screenWidth),
	)

	return true
}

func processTurn() {
	// TODO: do turn logics
	// refill merge
	mergeMap.ProcessTurn()
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

	// textWidth := rl.MeasureTextEx(font, buttonText, float32(fontSize), 1).X
	textHeight := rl.MeasureTextEx(font, buttonText, float32(fontSize), 1).Y
	textPosition := rl.Vector2{
		X: buttonRect.X + 80,
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

func initButton(
	screenWidth int32,
	screenHeight int32,
) {
	loadButtonTexture()
	buttonWidth = screenWidth / 10
	buttonHeight = screenHeight / 10
}

func unloadButton() {
	rl.UnloadTexture(buttonTexture)
}
