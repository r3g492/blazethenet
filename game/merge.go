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
	dragIdx         int
}

const (
	Fire1 = "fire 1"
	Fire2 = "fire 2"
	Fire3 = "fire 3"
	Fire4 = "fire 4"
)

var FireMap = map[int]string{
	1: Fire1,
	2: Fire2,
	3: Fire3,
	4: Fire4,
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
		dragIdx:         -1,
	}
}

func (m *MergeMap) AddFire(
	rectangleIdx int,
) {
	var result = deriveMerged(FireMap[1])
	if result != "" {
		m.mergeContents[rectangleIdx] = result
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
		var color rl.Color
		if m.dragIdx == i {
			color = rl.White
		} else if m.mergeContents[i] == FireMap[1] {
			color = rl.Red
		} else if m.mergeContents[i] == FireMap[2] {
			color = rl.Blue
		} else if m.mergeContents[i] == FireMap[3] {
			color = rl.Purple
		} else if m.mergeContents[i] == FireMap[4] {
			color = rl.Yellow
		} else {
			color = rl.White
		}
		rl.DrawRectangle(
			int32(m.mergeRectangles[i].X),
			int32(m.mergeRectangles[i].Y),
			int32(m.mergeRectangles[i].Width),
			int32(m.mergeRectangles[i].Height),
			color,
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
				m.dragIdx = i
				return
			}
		}
	}
	if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
		for i, rect := range m.mergeRectangles {
			if rl.CheckCollisionPointRec(mousePoint, rect) {
				if !m.isDragging {
					return
				}
				if m.dragContent == m.mergeContents[i] {
					m.isDragging = false
					m.dragContent = ""
					m.mergeContents[m.dragIdx] = ""
					m.dragIdx = -1
					var updated = deriveMerged(m.mergeContents[i])
					if updated != "" {
						m.mergeContents[i] = updated
					}
				}
				if m.dragContent != "" && notInDictionary(m.mergeContents[i]) {
					m.mergeContents[i] = m.dragContent
					m.mergeContents[m.dragIdx] = ""
				}
				m.isDragging = false
				m.dragContent = ""
				m.dragIdx = -1
				return
			}
		}
		m.isDragging = false
		m.dragContent = ""
		m.dragIdx = -1
	}
}

func notInDictionary(
	content string,
) bool {
	if content == FireMap[1] || content == FireMap[2] || content == FireMap[3] || content == FireMap[4] {
		return false
	}
	return true
}

func deriveMerged(
	content string,
) string {
	if content == FireMap[1] {
		return FireMap[2]
	}
	if content == FireMap[2] {
		return FireMap[3]
	}
	if content == FireMap[3] {
		return FireMap[4]
	}
	return ""
}
