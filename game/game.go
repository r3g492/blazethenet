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

	playerDrawX := screenWidth/2 - player.boxLen/2
	playerDrawY := screenHeight/2 - player.boxLen/2
	rl.DrawRectangle(
		playerDrawX,
		playerDrawY,
		player.boxLen,
		player.boxLen,
		rl.White,
	)

	enemyDrawX := playerDrawX + (enemy.x - player.x)
	enemyDrawY := playerDrawY + (enemy.y - player.y)
	rl.DrawRectangle(
		enemyDrawX,
		enemyDrawY,
		enemy.boxLen,
		enemy.boxLen,
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

	return true
}
