package main

import (
	"blazethenet/game"
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"path/filepath"
)

const (
	gameTitle           string = "Merge Rogue"
	fps                 int32  = 60
	ButtonImageFilePath        = "resources/button.png"
)

var (
	screenWidth            int32 = 1600
	screenHeight           int32 = 900
	userMonitorWidth       int
	userMonitorHeight      int
	userMonitorCount       int
	isGameOn               bool = true
	currentFont            rl.Font
	fontSize               int32
	buttonWidth            int32
	buttonHeight           int32
	buttonTexture          rl.Texture2D
	oneButtonTextureHeight int32
	buttonImageRectangle   rl.Rectangle
	gameState              int
	loadingText            string
	backgroundMusic        rl.Music
)

const (
	InMainMenu = iota
	InGame
	ResolutionSettings
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, gameTitle)

	initLoading()
	setFps()
	setGameStateAsMain()
	initButton()
	defer unloadButton()
	initFont()
	defer unloadFont()
	initAudio()
	defer unloadAudio()
	initMainMusic()
	defer unloadMainMusic()
	defer game.Unload()

	for isGameOn {
		rl.ClearBackground(rl.Black)
		rl.BeginDrawing()
		rl.UpdateMusicStream(backgroundMusic)
		if rl.WindowShouldClose() && gameState != InGame {
			isGameOn = false
			break
		}
		if rl.WindowShouldClose() && gameState == InGame {
			gameState = InMainMenu
		}

		mousePoint := rl.GetMousePosition()
		switch gameState {
		case InMainMenu:
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

			resolutionOptionButtonX := (screenWidth - buttonWidth) / 2
			resolutionOptionButtonY := startButtonY + buttonHeight
			resolutionOptionButtonRect := rl.Rectangle{
				X:      float32(resolutionOptionButtonX),
				Y:      float32(resolutionOptionButtonY),
				Width:  float32(buttonWidth),
				Height: float32(buttonHeight),
			}
			if buttonControl(
				mousePoint,
				resolutionOptionButtonRect,
				buttonImageRectangle,
				oneButtonTextureHeight,
				buttonTexture,
				"ResolutionSettings",
				currentFont,
				fontSize,
				rl.White,
			) {
				gameState = ResolutionSettings
			}

			exitButtonX := (screenWidth - buttonWidth) / 2
			exitButtonY := resolutionOptionButtonY + buttonHeight
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
		case ResolutionSettings:
			ButtonX1 := (screenWidth - buttonWidth) / 2
			ButtonY1 := int32(100)
			ButtonRect1 := rl.Rectangle{
				X:      float32(ButtonX1),
				Y:      float32(ButtonY1),
				Width:  float32(buttonWidth),
				Height: float32(buttonHeight),
			}
			if buttonControl(
				mousePoint,
				ButtonRect1,
				buttonImageRectangle,
				oneButtonTextureHeight,
				buttonTexture,
				"FullScreen",
				currentFont,
				fontSize,
				rl.White,
			) {
				makeItFullScreen()
			}

			ButtonX2 := (screenWidth - buttonWidth) / 2
			ButtonY2 := ButtonY1 + buttonHeight
			ButtonRect2 := rl.Rectangle{
				X:      float32(ButtonX2),
				Y:      float32(ButtonY2),
				Width:  float32(buttonWidth),
				Height: float32(buttonHeight),
			}
			if buttonControl(
				mousePoint,
				ButtonRect2,
				buttonImageRectangle,
				oneButtonTextureHeight,
				buttonTexture,
				"1600x900",
				currentFont,
				fontSize,
				rl.White,
			) {
				changeResolution(1600, 900)
			}

			ButtonX3 := (screenWidth - buttonWidth) / 2
			ButtonY3 := ButtonY2 + buttonHeight
			ButtonRect3 := rl.Rectangle{
				X:      float32(ButtonX3),
				Y:      float32(ButtonY3),
				Width:  float32(buttonWidth),
				Height: float32(buttonHeight),
			}
			if buttonControl(
				mousePoint,
				ButtonRect3,
				buttonImageRectangle,
				oneButtonTextureHeight,
				buttonTexture,
				"1280x1024",
				currentFont,
				fontSize,
				rl.White,
			) {
				changeResolution(1280, 1024)
			}

			exitButtonX := (screenWidth - buttonWidth) / 2
			exitButtonY := ButtonY3 + buttonHeight
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
				"ToMainMenu",
				currentFont,
				fontSize,
				rl.White,
			) {
				gameState = InMainMenu
			}
			break

		case InGame:
			if !game.IsGameInit {
				game.Init(
					currentFont,
					float32(fontSize),
					screenWidth,
					screenHeight,
				)
			} else {
				game.Game(
					currentFont,
					float32(fontSize),
					screenWidth,
					screenHeight,
				)
			}

			if rl.IsKeyPressed(rl.KeyF10) {
				gameState = InMainMenu
			}

			break
		default:
		}

		rl.EndDrawing()
	}
	rl.CloseWindow()
}

func setGameStateAsMain() {
	logLoadingLn("setting game state...")
	gameState = InMainMenu
}

func setFps() {
	logLoadingLn("setting FPS...")
	rl.SetTargetFPS(fps)
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
	fmt.Printf("  Fullscreen: %v\n", rl.IsWindowFullscreen())
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

func initLoading() {
	loadingText = ""
	rl.ClearBackground(rl.Black)
	logLoadingLn("loading initiated...")
}

func logLoadingLn(text string) {
	loadingText += text + "\n"
	rl.ClearBackground(rl.Black)
	rl.BeginDrawing()
	rl.DrawText(loadingText, 50, 80, 24, rl.White)
	rl.EndDrawing()
}

func changeResolution(width, height int) {
	if rl.IsWindowFullscreen() {
		rl.ToggleFullscreen()
	}
	screenWidth = int32(width)
	screenHeight = int32(height)
	rl.CloseWindow()
	rl.InitWindow(screenWidth, screenHeight, gameTitle)
	initLoading()
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
	initLoading()
	rl.SetWindowPosition(0, 0)
	rl.ToggleFullscreen()
	initButton()
	initFont()
}

func initFont() {
	logLoadingLn("setting fonts...")
	fontPath := filepath.Join("resources", "font", "Noto_Sans_KR", "static", "NotoSansKR-ExtraBold.ttf")
	fontSize = screenWidth / 48
	currentFont = rl.LoadFontEx(fontPath, min(fontSize, 48), nil, 65535)
}

func unloadFont() {
	logLoadingLn("unsetting font...")
	rl.UnloadFont(currentFont)
}

func initButton() {
	logLoadingLn("setting buttons...")
	loadButtonTexture()
	buttonWidth = screenWidth / 5
	buttonHeight = screenHeight / 5
}

func unloadButton() {
	logLoadingLn("unsetting buttons...")
	rl.UnloadTexture(buttonTexture)
}

func initAudio() {
	logLoadingLn("setting audios...")
	rl.InitAudioDevice()
}

func unloadAudio() {
	logLoadingLn("unsetting audio...")
	rl.CloseAudioDevice()
}

func initMainMusic() {
	logLoadingLn("setting main music...")
	backgroundMusic = rl.LoadMusicStream("resources/audio/background_music.ogg")
	rl.PlayMusicStream(backgroundMusic)
}

func unloadMainMusic() {
	logLoadingLn("unsetting main music...")
	rl.UnloadMusicStream(backgroundMusic)
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
