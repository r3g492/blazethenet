package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Player struct {
	x         int32
	y         int32
	radius    float32
	health    int32
	direction rl.Vector2
}
