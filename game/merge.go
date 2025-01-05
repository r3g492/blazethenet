package game

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"strconv"
)

type MergeMap struct {
	mergeWidth  int32
	mergeHeight int32
	mergeList   []int32
}

func CreateMerge(
	mergeWidth int32,
	mergeHeight int32,
) MergeMap {
	return MergeMap{
		mergeWidth:  mergeWidth,
		mergeHeight: mergeHeight,
		mergeList:   make([]int32, mergeWidth*mergeHeight),
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
	for i := range m.mergeList {
		rl.DrawTextEx(font,
			strconv.Itoa(i),
			rl.Vector2{X: float32(int32(i)%m.mergeWidth*50.0) + 300, Y: float32(screenHeight/2 + (int32(i)/m.mergeWidth)*100)},
			fontSize,
			0,
			rl.White,
		)
	}
}

func (m *MergeMap) Control() {
}
