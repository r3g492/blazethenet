package game1

import rl "github.com/gen2brain/raylib-go/raylib"

type LinkMap struct {
	linkedMaps []LinkMap
}

func CreateLinkMap() LinkMap {
	return LinkMap{}
}

func (m *LinkMap) Render() {
}

func (m *LinkMap) Control(
	mousePoint rl.Vector2,
) {
}
