package game

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	IsGameInit = false
	player     Player
	enemy      Enemy
)

func Init() {
	IsGameInit = true
	player = Player{
		50,
		50,
		70,
		70,
	}
	enemy = Enemy{
		70,
		70,
		70,
		70,
	}
}

func Logic(
	font rl.Font,
	screenWidth int32,
	screenHeight int32,
) bool {
	rl.ClearBackground(rl.SkyBlue)

	text := fmt.Sprintf("x: %d y: %d", player.x, player.y)
	rl.DrawTextEx(font, text, rl.Vector2{X: 150, Y: 150}, 32, 10, rl.White)

	playerDrawX := screenWidth/2 - (int32)(player.radius/2)
	playerDrawY := screenHeight/2 - (int32)(player.radius/2)
	rl.DrawCircle(
		playerDrawX,
		playerDrawY,
		player.radius,
		rl.White,
	)
	enemyDrawX := playerDrawX + (enemy.x - player.x)
	enemyDrawY := playerDrawY + (enemy.y - player.y)
	rl.DrawCircle(
		enemyDrawX,
		enemyDrawY,
		enemy.radius,
		rl.Red,
	)

	if rl.IsKeyDown(rl.KeyA) {
		player.x -= PlayerMovementSpeed
	}
	if rl.IsKeyDown(rl.KeyD) {
		player.x += PlayerMovementSpeed
	}
	if rl.IsKeyDown(rl.KeyW) {
		player.y -= PlayerMovementSpeed
	}
	if rl.IsKeyDown(rl.KeyS) {
		player.y += PlayerMovementSpeed
	}
	mousePoint := rl.GetMousePosition()
	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		rl.DrawCircle(
			int32(mousePoint.X),
			int32(mousePoint.Y),
			player.radius,
			rl.Purple,
		)
	}
	return true
}
