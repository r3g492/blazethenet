package main

import (
	"blazethenet/button"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	gameTitle string = "Detective Game"
	fps       int32  = 60
)

var screenWidth int32 = 800
var screenHeight int32 = 600
var userMonitorWidth int
var userMonitorHeight int
var userMonitorCount int
var isFullScreen bool = false

const (
	MainMenu = iota
	GameStart
	Settings
)

/*
 * Main menu buttons
 */
const (
	StartMenu = iota
	SettingsMenu
	ExitMenu
)

/*
 * Game settings
 */
const (
	FullScreen = iota
	Resolution800600
	ReturnToMain
)

func main() {

	rl.InitWindow(screenWidth, screenHeight, gameTitle)

	userMonitorCount = rl.GetMonitorCount()
	userMonitorWidth = rl.GetMonitorWidth(0)
	userMonitorHeight = rl.GetMonitorHeight(0)

	rl.SetTargetFPS(fps)

	button.InitButtonTexture()
	defer button.UnloadButtonTexture()

	gameState := MainMenu

GameLoop:
	for !rl.WindowShouldClose() {
		mousePoint := rl.GetMousePosition()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		switch gameState {
		case MainMenu:

			texts := []string{
				"Start",
				"Settings",
				"Exit",
			}
			buttonInfo, _ := button.Plan(texts, screenWidth, screenHeight)
			buttonActions := make([]bool, len(buttonInfo.Buttons))

			for i, btnBounds := range buttonInfo.Buttons {
				buttonActions[i] = button.DrawButtonAction(btnBounds, mousePoint, texts[i])

				if buttonActions[StartMenu] {
					// Start the game
				}

				if buttonActions[SettingsMenu] {
					gameState = Settings
				}

				if buttonActions[ExitMenu] {
					break GameLoop
				}
			}
		case Settings:

			texts := []string{
				"Fullscreen",
				"800x600 Resolution",
				"Main Menu",
			}
			buttonInfo, _ := button.Plan(texts, screenWidth, screenHeight)
			buttonActions := make([]bool, len(buttonInfo.Buttons))

			for i, btnBounds := range buttonInfo.Buttons {
				buttonActions[i] = button.DrawButtonAction(btnBounds, mousePoint, texts[i])

				if buttonActions[FullScreen] {
					if !isFullScreen {
						screenWidth = int32(userMonitorWidth)
						screenHeight = int32(userMonitorHeight)
						rl.SetWindowSize(int(screenWidth), int(screenHeight))
						rl.SetWindowPosition(0, 0)
						isFullScreen = true
					}
				}

				if buttonActions[Resolution800600] {
					screenWidth = 800
					screenHeight = 600
					rl.SetWindowSize(int(screenWidth), int(screenHeight))
					rl.SetWindowPosition(userMonitorWidth/4, userMonitorHeight/4)
					isFullScreen = false
				}

				if buttonActions[ReturnToMain] {
					gameState = MainMenu
				}
			}
		default:
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
