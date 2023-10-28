package hud

var buildings = []string{
	"Belt",
	"Miner",
}

func renderBuildings() string {
	s := ""
	for i, e := range buildings {
		// TODO: styling
		if active == buildingsPanelIndex && i == cursor {
			s += "> " + e
		} else {
			s += e
		}

		if i < len(buildings)-1 {
			s += "\n"
		}
	}

	return s
}
