package game

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MergeMap struct {
	mergeWidth      int32
	mergeHeight     int32
	mergeRectangles []rl.Rectangle
}

func CreateMerge(
	mergeWidth int32,
	mergeHeight int32,
	screenWidth int32,
	screenHeight int32,
) MergeMap {
	mergeRectangles := make([]rl.Rectangle, mergeWidth*mergeHeight)
	for i := range mergeWidth * mergeHeight {
		xPos := float32(i%mergeWidth*50.0) + 300
		yPos := float32(screenHeight/2 + (i/mergeWidth)*100)
		mergeRectangles[i] = rl.Rectangle{
			xPos,
			yPos,
			10.0,
			10.0,
		}
	}

	return MergeMap{
		mergeWidth:      mergeWidth,
		mergeHeight:     mergeHeight,
		mergeRectangles: mergeRectangles,
	}
}

func (m *MergeMap) String() string {
	return fmt.Sprintf("mergeWidth: %d, mergeHeight: %d", m.mergeWidth, m.mergeHeight)
}

func (m *MergeMap) Render(
	font rl.Font,
	fontSize float32,
	screenWidth int32,
	screenHeight int32,
) {
	for i := range m.mergeRectangles {
		rl.DrawRectangle(
			int32(m.mergeRectangles[i].X),
			int32(m.mergeRectangles[i].Y),
			int32(m.mergeRectangles[i].Width),
			int32(m.mergeRectangles[i].Height),
			rl.White,
		)
	}
}

func (m *MergeMap) Control(
	mousePoint rl.Vector2,
) {
}
