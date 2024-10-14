package main

import (
	"blazethenet/button"
	"blazethenet/game"
	"blazethenet/text"
	rl "github.com/gen2brain/raylib-go/raylib"
	"path/filepath"
)

const (
	gameTitle string = "Detective Game"
	fps       int32  = 60
)

var screenWidth int32 = 1024
var screenHeight int32 = 800
var userMonitorWidth int
var userMonitorHeight int
var userMonitorCount int
var isFullScreen bool = false
var isGameOn bool = true
var screenRatio float32 = 1
var koreanFont rl.Font
var currentFont rl.Font

const (
	InMainMenu = iota
	InGame
	InSettings
	InResolutionSettings
	InLanguageSettings
	ScenarioOver
)

/*
 * Main menu buttons
 */
const (
	StartMenu = iota
	ReturnToMyGame
	SettingsMenu
	ExitMenu
)

/*
 * settings
 */
const (
	ResolutionSettings = iota
	LanguageSettings
	BackFromSettings
)

/*
 * Resolution settings
 */
const (
	FullScreen = iota
	Resolution800600
	Resolution1024768
	Resolution19201080
	BackFromResolution
)

/*
 * Language settings
 */
const (
	English = iota
	Korean
	BackFromLanguage
)

func main() {

	rl.InitWindow(screenWidth, screenHeight, gameTitle)

	userMonitorCount = rl.GetMonitorCount()
	userMonitorWidth = rl.GetMonitorWidth(0)
	userMonitorHeight = rl.GetMonitorHeight(0)

	rl.SetTargetFPS(fps)

	button.InitButtonTexture()
	defer button.UnloadButtonTexture()

	gameState := InMainMenu

	fontPath := filepath.Join("resources", "font", "Noto_Sans_KR", "static", "NotoSansKR-ExtraBold.ttf")
	koreanFont = rl.LoadFontEx(fontPath, 32, nil, 65535)
	defer rl.UnloadFont(koreanFont)

	for isGameOn {
		if screenWidth > 1200 {
			screenRatio = 1.3
		} else if screenWidth < 600 {
			screenRatio = 0.8
		} else {
			screenRatio = 1
		}

		mousePoint := rl.GetMousePosition()
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		switch gameState {
		case InMainMenu:
			if rl.WindowShouldClose() {
				isGameOn = false
				break
			}
			texts := []string{
				text.NewText(
					"StartNew",
					"새로 시작",
				).Get(),
				text.NewText(
					"ReturnToGame",
					"돌아가기",
				).Get(),
				text.NewText(
					"InSettings",
					"설정",
				).Get(),
				text.NewText(
					"Exit",
					"종료",
				).Get(),
			}
			buttonInfo, _ := button.Plan(texts, screenWidth, screenHeight, screenRatio)
			buttonActions := make([]bool, len(buttonInfo.Buttons))

			for i, btnBounds := range buttonInfo.Buttons {
				buttonActions[i] = button.DrawButtonAction(btnBounds, mousePoint, texts[i], screenRatio, currentFont)

				if buttonActions[StartMenu] {
					gameState = InGame
				}

				if buttonActions[ReturnToMyGame] {
					gameState = InGame
				}

				if buttonActions[SettingsMenu] {
					gameState = InSettings
				}

				if buttonActions[ExitMenu] {
					isGameOn = false
				}
			}
			break
		case InSettings:
			texts := []string{
				text.NewText(
					"ResolutionSettings",
					"해상도 세팅",
				).Get(),
				text.NewText(
					"LanguageSettings",
					"언어 설정",
				).Get(),
				text.NewText(
					"InMainMenu",
					"주 메뉴",
				).Get(),
			}
			buttonInfo, _ := button.Plan(texts, screenWidth, screenHeight, screenRatio)
			buttonActions := make([]bool, len(buttonInfo.Buttons))

			for i, btnBounds := range buttonInfo.Buttons {
				buttonActions[i] = button.DrawButtonAction(btnBounds, mousePoint, texts[i], screenRatio, currentFont)

				if buttonActions[ResolutionSettings] {
					gameState = InResolutionSettings
				}

				if buttonActions[LanguageSettings] {
					gameState = InLanguageSettings
				}

				if buttonActions[BackFromSettings] {
					gameState = InMainMenu
				}

				if rl.IsKeyPressed(rl.KeyEscape) {
					gameState = InMainMenu
				}
			}
			break
		case InResolutionSettings:
			texts := []string{
				text.NewText(
					"FullScreen",
					"전체 화면",
				).Get(),
				text.NewText(
					"800x600Resolution",
					"해상도800x600",
				).Get(),
				text.NewText(
					"1024x768Resolution",
					"해상도1024x768",
				).Get(),
				text.NewText(
					"1920x1080Resolution",
					"해상도1920x1080",
				).Get(),
				text.NewText(
					"Back",
					"돌아가기",
				).Get(),
			}
			buttonInfo, _ := button.Plan(texts, screenWidth, screenHeight, screenRatio)
			buttonActions := make([]bool, len(buttonInfo.Buttons))

			for i, btnBounds := range buttonInfo.Buttons {
				buttonActions[i] = button.DrawButtonAction(btnBounds, mousePoint, texts[i], screenRatio, currentFont)

				if buttonActions[FullScreen] {
					if !isFullScreen {
						screenWidth = int32(userMonitorWidth)
						screenHeight = int32(userMonitorHeight)
						rl.SetWindowSize(int(screenWidth), int(screenHeight))
						rl.SetWindowPosition(0, 0)
						rl.ToggleFullscreen()
						isFullScreen = true

					}
				}

				if buttonActions[Resolution800600] {
					screenWidth = 800
					screenHeight = 600
					rl.SetWindowSize(int(screenWidth), int(screenHeight))
					rl.SetWindowPosition(0, 30)
					isFullScreen = false
					if rl.IsWindowFullscreen() {
						rl.ToggleFullscreen()
					}
				}

				if buttonActions[Resolution1024768] {
					screenWidth = 1024
					screenHeight = 768
					rl.SetWindowSize(int(screenWidth), int(screenHeight))
					rl.SetWindowPosition(0, 30)
					isFullScreen = false
					if rl.IsWindowFullscreen() {
						rl.ToggleFullscreen()
					}
				}

				if buttonActions[Resolution19201080] {
					screenWidth = 1920
					screenHeight = 1080
					rl.SetWindowSize(int(screenWidth), int(screenHeight))
					rl.SetWindowPosition(0, 30)
					isFullScreen = false
					if rl.IsWindowFullscreen() {
						rl.ToggleFullscreen()
					}
				}

				if buttonActions[BackFromResolution] {
					gameState = InSettings
				}

				if rl.IsKeyPressed(rl.KeyEscape) {
					gameState = InMainMenu
				}
			}
			break
		case InLanguageSettings:
			texts := []string{
				text.NewText(
					"English",
					"English",
				).Get(),
				text.NewText(
					"Korean",
					"한국어",
				).Get(),
				text.NewText(
					"Back",
					"돌아가기",
				).Get(),
			}
			buttonInfo, _ := button.Plan(texts, screenWidth, screenHeight, screenRatio)
			buttonActions := make([]bool, len(buttonInfo.Buttons))

			for i, btnBounds := range buttonInfo.Buttons {
				buttonActions[i] = button.DrawButtonAction(btnBounds, mousePoint, texts[i], screenRatio, currentFont)

				if buttonActions[English] {
					currentFont = rl.Font{}
					text.SetLanguageToEnglish()
				}

				if buttonActions[Korean] {
					currentFont = koreanFont
					text.SetLanguageToKorean()
				}

				if buttonActions[BackFromLanguage] {
					gameState = InSettings
				}

				if rl.IsKeyPressed(rl.KeyEscape) {
					gameState = InMainMenu
				}
			}
			break
		case InGame:
			game.Logic()
			if game.ScenarioEndCondition() {
				gameState = ScenarioOver
			}
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
		case ScenarioOver:
			rl.DrawText("success", 100, 100, 200, rl.Blue)
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
