package game

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"strconv"
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
	var iconLen int32 = 50
	for i := range mergeWidth * mergeHeight {
		xPos := float32(i%mergeWidth*iconLen) + float32(screenWidth/2-mergeWidth*iconLen/2)
		yPos := float32(screenHeight/2 + (i/mergeWidth)*iconLen)
		mergeRectangles[i] = rl.Rectangle{
			xPos,
			yPos,
			float32(iconLen),
			float32(iconLen),
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
	for i := range m.mergeRectangles {
		rl.DrawTextEx(
			font,
			strconv.Itoa(i),
			rl.Vector2{m.mergeRectangles[i].X, m.mergeRectangles[i].Y},
			fontSize,
			0.1,
			rl.Red,
		)
	}
}

func (m *MergeMap) Control(
	mousePoint rl.Vector2,
) {
}
