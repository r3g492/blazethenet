package game

import (
	"blazethenet/text"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Logic(font rl.Font) bool {
	rl.ClearBackground(rl.SkyBlue)

	node := NewNode(
		*text.NewText(
			"node1",
			"노드1",
		),
	)
	linkedNode := NewNode(
		*text.NewText(
			"node2",
			"노드2",
		),
	)
	node.linkNode(linkedNode)
	rl.DrawTextEx(font, node.String(), rl.Vector2{X: 150, Y: 150}, 32, 10, rl.White)
	return true
}
