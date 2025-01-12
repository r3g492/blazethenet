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
	mergeContents   []string
	isDragging      bool
	dragContent     string
}

func CreateMerge(
	mergeWidth int32,
	mergeHeight int32,
	screenWidth int32,
	screenHeight int32,
) MergeMap {
	mergeRectangles := make([]rl.Rectangle, mergeWidth*mergeHeight)
	mergeContents := make([]string, mergeWidth*mergeHeight)
	var iconLen int32 = 50
	for i := range mergeWidth * mergeHeight {
		xPos := float32(i%mergeWidth*iconLen) + float32(screenWidth/2-mergeWidth*iconLen/2)
		yPos := float32(screenHeight/2 + (i/mergeWidth)*iconLen)
		mergeRectangles[i] = rl.Rectangle{
			X:      xPos,
			Y:      yPos,
			Width:  float32(iconLen),
			Height: float32(iconLen),
		}
		mergeContents[i] = strconv.Itoa(int(i))
	}

	return MergeMap{
		mergeWidth:      mergeWidth,
		mergeHeight:     mergeHeight,
		mergeRectangles: mergeRectangles,
		mergeContents:   mergeContents,
		isDragging:      false,
		dragContent:     "",
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
	mousePoint rl.Vector2,
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
			rl.Vector2{
				X: m.mergeRectangles[i].X,
				Y: m.mergeRectangles[i].Y,
			},
			fontSize,
			0.1,
			rl.Red,
		)
	}

	if m.isDragging {
		rl.DrawText(
			m.dragContent,
			int32(mousePoint.X-50),
			int32(mousePoint.Y-50),
			100,
			rl.Blue,
		)
	}
}

func (m *MergeMap) Control(
	mousePoint rl.Vector2,
) {
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		for i, rect := range m.mergeRectangles {
			if rl.CheckCollisionPointRec(mousePoint, rect) {
				m.isDragging = true
				m.dragContent = m.mergeContents[i]
				break
			}
		}
	}
	if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
		m.isDragging = false
		m.dragContent = ""
	}
}
