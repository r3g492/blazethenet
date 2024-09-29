package button

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	FilePath  = "resources/button.png"
	InOneFile = 3
)

var buttonTexture rl.Texture2D

type Info struct {
	ScreenWidth  int32
	ScreenHeight int32
	Buttons      []rl.Rectangle
	Texts        []string
	Count        int
}

func NewInfo(screenWidth int32, screenHeight int32, buttons []rl.Rectangle, texts []string) (*Info, error) {
	if len(buttons) != len(texts) {
		return nil, fmt.Errorf("buttons and texts slices must have the same length: got %d buttons and %d texts", len(buttons), len(texts))
	}

	return &Info{
		ScreenWidth:  screenWidth,
		ScreenHeight: screenHeight,
		Buttons:      buttons,
		Texts:        texts,
		Count:        len(buttons),
	}, nil
}

func InitButtonTexture() {
	buttonTexture = rl.LoadTexture(FilePath)
}

func UnloadButtonTexture() {
	rl.UnloadTexture(buttonTexture)
}

func DrawButtonAction(bounds rl.Rectangle, mousePoint rl.Vector2, text string, screenRatio float32, font rl.Font) bool {
	frameHeight := buttonTexture.Height / InOneFile
	sourceRec := rl.Rectangle{Width: float32(buttonTexture.Width), Height: float32(frameHeight)}

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

	destRec := rl.Rectangle{
		X:      bounds.X,
		Y:      bounds.Y,
		Width:  bounds.Width,
		Height: bounds.Height,
	}

	rl.DrawTexturePro(buttonTexture, sourceRec, destRec, rl.Vector2{X: 0, Y: 0}, 0, rl.White)

	textSize := int32(20 * screenRatio)
	textWidth := rl.MeasureText(text, textSize)
	textX := bounds.X + (bounds.Width/2 - float32(textWidth)/2)
	textY := bounds.Y + (bounds.Height/2 - float32(textSize)/2)

	spacing := float32(1)
	fontSize := 20 * screenRatio
	rl.DrawTextEx(font, text, rl.Vector2{X: textX, Y: textY}, fontSize, spacing, rl.White)
	return btnAction
}
func Plan(texts []string, screenWidth int32, screenHeight int32, screenRatio float32) (Info, error) {

	buttonHeight := float32(buttonTexture.Height/InOneFile) * screenRatio
	buttonWidth := float32(buttonTexture.Width) * screenRatio

	buttons := make([]rl.Rectangle, len(texts))

	for i := range texts {
		yPos := float32(screenHeight/2) - (buttonHeight * float32(len(texts)) / 2) + (buttonHeight * float32(i))

		xPos := float32(screenWidth/2) - (buttonWidth / 2)

		buttons[i] = rl.Rectangle{
			X:      xPos,
			Y:      yPos,
			Width:  buttonWidth,
			Height: buttonHeight,
		}
	}
	result, _ := NewInfo(screenWidth, screenHeight, buttons, texts)
	return *result, nil
}
