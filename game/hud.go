package game

import (
	"slices"
	"strings"

	"github.com/calvinlarimore/factory/ui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
		var style lipgloss.Style

		if c.activeHudPanel == buildingsPanelIndex && i == c.hudCursor {
			style = lipgloss.NewStyle().
				Background(lipgloss.ANSIColor(7)).
				Foreground(lipgloss.ANSIColor(0))
		} else {
			style = lipgloss.NewStyle().
				Background(lipgloss.ANSIColor(0)).
				Foreground(lipgloss.ANSIColor(7))
		}

		b += style.
			Width(hudPanelStyle.InnerWidth()).
			Render(e)

		if i < len(buildings)-1 {
			b += "\n"
		}
	}

	names := make([]string, 0, len(players))
	for name := range players {
		names = append(names, name)
	}

	slices.Sort(names)
	for i, name := range names {
		client := players[name]
		style := lipgloss.NewStyle().
			Foreground(client.color).
			Bold(name == c.name)

		names[i] = style.Inline(true).Render(name)
	}

	s := hudPanelStyle.Render("Players", "", strings.Join(names, "\n")) + "\n"
	s += hudPanelStyle.Render("Buildings", "b", b)

	return s
}
