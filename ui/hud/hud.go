package hud

import (
	"github.com/calvinlarimore/factory/ui"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	nonePanelIndex = iota
	buildingsPanelIndex
)

var panel = ui.NewPanel()
var (
	active = -1
	cursor = 0
)

func toggleActive(i int) bool {
	if active == -1 {
		active = i
		return true
	}

	if active == i {
		active = -1
		cursor = 0
		return true
	}

	return false
}

func Update(msg tea.KeyMsg) bool {
	switch msg.String() {
	case "b":
		return toggleActive(buildingsPanelIndex)
	}

	if active != -1 {

		switch msg.String() {
		case "up":
			if cursor > 0 {
				cursor--
			}
			return true
		case "down":
			// TODO: fix hardcoding
			if cursor < 1 {
				cursor++
			}
			return true
		}
	}

	return false
}

func View() string {
	s := panel.Render("Buildings", "b", renderBuildings())

	return s
}
