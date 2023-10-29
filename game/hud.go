package game

import (
	"github.com/calvinlarimore/factory/ui"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	nonePanelIndex = iota
	buildingsPanelIndex
)

var buildings = []string{
	"Belt",
	"Miner",
}

var hudPanelStyle = ui.NewPanelStyle()

func toggleActive(c *Client, i int) bool {
	if c.activeHudPanel == nonePanelIndex {
		c.activeHudPanel = i
		return true
	}

	if c.activeHudPanel == i {
		c.activeHudPanel = nonePanelIndex
		c.hudCursor = 0
		return true
	}

	return false
}

func updateHud(c *Client, msg tea.KeyMsg) bool {
	switch msg.String() {
	case "b":
		return toggleActive(c, buildingsPanelIndex)
	}

	if c.activeHudPanel != nonePanelIndex {
		switch msg.String() {
		case "up":
			if c.hudCursor > 0 {
				c.hudCursor--
			}
			return true
		case "down":
			// TODO: fix hardcoding
			if c.hudCursor < 1 {
				c.hudCursor++
			}
			return true
		}
	}

	return false
}

func renderHud(c Client) string {
	b := ""
	for i, e := range buildings {
		// TODO: styling
		if c.activeHudPanel == buildingsPanelIndex && i == c.hudCursor {
			b += "> " + e
		} else {
			b += e
		}

		if i < len(buildings)-1 {
			b += "\n"
		}
	}

	s := hudPanelStyle.Render("Buildings", "b", b)

	return s
}
