package game

type LinkMap struct {
	linkedMaps []LinkMap
}

func CreateLinkMap() LinkMap {
	return LinkMap{}
}

func (m *LinkMap) Render() {
}

func (m *LinkMap) Control() {
}
