package game

import "fmt"

type MergeMap struct {
	mergeWidth  int32
	mergeHeight int32
}

func CreateMerge(
	mergeWidth int32,
	mergeHeight int32,
) MergeMap {
	return MergeMap{
		mergeWidth:  mergeWidth,
		mergeHeight: mergeHeight,
	}
}

func (m *MergeMap) String() string {
	return fmt.Sprintf("mergeWidth: %d, mergeHeight: %d", m.mergeWidth, m.mergeHeight)
}

func (m *MergeMap) Render() {
}

func (m *MergeMap) Control() {
}
