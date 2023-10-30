package render

import "github.com/charmbracelet/lipgloss"

type PlayerInfo struct {
	x, y  int
	color lipgloss.TerminalColor
}

func NewPlayerInfo(x, y int, color lipgloss.TerminalColor) PlayerInfo {
	return PlayerInfo{
		x: x, y: y,
		color: color,
	}
}
