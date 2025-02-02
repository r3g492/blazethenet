package game1

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"strconv"
	"strings"
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
	Fire1    = "fire 1"
	Fire2    = "fire 2"
	Fire3    = "fire 3"
	Fire4    = "fire 4"
	FireGen1 = "fire generator 1"
	FireGen2 = "fire generator 2"
	FireGen3 = "fire generator 3"
	FireGen4 = "fire generator 4"
)

var FireMap = map[int]string{
	1: Fire1,
	2: Fire2,
	3: Fire3,
	4: Fire4,
}

var FireGeneratorMap = map[int]string{
	1: FireGen1,
	2: FireGen2,
	3: FireGen3,
	4: FireGen4,
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
		mergeContents[i] = ""
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
	var result = deriveMerged(
		FireMap[1],
		m.mergeContents[rectangleIdx],
	)
	if result != "" {
		m.mergeContents[rectangleIdx] = result
	}
}

func (m *MergeMap) AddFireGen(
	rectangleIdx int,
) {
	m.mergeContents[rectangleIdx] = FireGeneratorMap[1]
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
		} else if m.mergeContents[i] == FireGeneratorMap[1] {
			color = rl.Black
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
	if rl.IsMouseButtonPressed(rl.MouseRightButton) {
		m.isDragging = false
		m.dragContent = ""
		m.dragIdx = -1
	}
	if m.isDragging && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		for i, rect := range m.mergeRectangles {
			if rl.CheckCollisionPointRec(mousePoint, rect) {
				if !m.isDragging || m.dragIdx == i {
					break
				}
				if m.dragContent == m.mergeContents[i] {
					var updated = deriveMerged(
						m.dragContent,
						m.mergeContents[i],
					)
					if updated != "" {
						m.mergeContents[i] = updated
						m.isDragging = false
						m.dragContent = ""
						m.mergeContents[m.dragIdx] = ""
						m.dragIdx = -1
						break
					}
				}
				if m.dragContent != "" && notInDictionary(m.mergeContents[i]) && i != m.dragIdx {
					m.mergeContents[i] = m.dragContent
					m.mergeContents[m.dragIdx] = ""
					m.isDragging = false
					m.dragContent = ""
					m.dragIdx = -1
					break
				}
				break
			}
		}
		m.isDragging = false
		m.dragContent = ""
		m.dragIdx = -1
	}
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		for i, rect := range m.mergeRectangles {
			if rl.CheckCollisionPointRec(mousePoint, rect) && m.isDragging == false {
				m.isDragging = true
				m.dragContent = m.mergeContents[i]
				m.dragIdx = i
				break
			}
		}
	}
	if rl.IsMouseButtonReleased(rl.MouseLeftButton) {
		for i, rect := range m.mergeRectangles {
			if rl.CheckCollisionPointRec(mousePoint, rect) {
				if !m.isDragging || m.dragIdx == i {
					break
				}
				if m.dragContent == m.mergeContents[i] {
					var updated = deriveMerged(
						m.dragContent,
						m.mergeContents[i],
					)
					if updated != "" {
						m.mergeContents[i] = updated
						m.isDragging = false
						m.dragContent = ""
						m.mergeContents[m.dragIdx] = ""
						m.dragIdx = -1
						break
					}
				}
				if m.dragContent != "" && notInDictionary(m.mergeContents[i]) && i != m.dragIdx {
					m.mergeContents[i] = m.dragContent
					m.mergeContents[m.dragIdx] = ""
					m.isDragging = false
					m.dragContent = ""
					m.dragIdx = -1
					break
				}
				break
			}
		}
		m.isDragging = false
		m.dragContent = ""
		m.dragIdx = -1
	}
}

func (m *MergeMap) ProcessTurn() {
	for contentIdx := range m.mergeContents {
		for fireGenIdx := range FireGeneratorMap {
			if m.mergeContents[contentIdx] == FireGeneratorMap[fireGenIdx] {
				m.AddFire(contentIdx + 1)
			}
		}
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
	targetContent string,
) string {
	if targetContent == "" {
		return content
	}
	for fireKey := range FireMap {
		if strings.Contains(content, FireMap[fireKey]) && content == targetContent {
			fmt.Println("Merge happened")
			return FireMap[fireKey+1]
		}
	}
	return ""
}
