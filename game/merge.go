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
) MergeMap {
	mergeRectangles := make([]rl.Rectangle, mergeWidth*mergeHeight)
	for i := range mergeWidth * mergeHeight {
		mergeRectangles[i] = rl.Rectangle{
			0,
			0,
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
		rl.DrawTextEx(font,
			strconv.Itoa(i),
			rl.Vector2{X: float32(int32(i)%m.mergeWidth*50.0) + 300, Y: float32(screenHeight/2 + (int32(i)/m.mergeWidth)*100)},
			fontSize,
			0,
			rl.White,
		)
	}
}

func (m *MergeMap) Control(
	mousePoint rl.Vector2,
) {
}
