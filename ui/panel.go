package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Panel struct {
	MainStyle  lipgloss.Style
	TitleStyle lipgloss.Style
}

func NewPanel() Panel {
	return Panel{
		MainStyle: lipgloss.NewStyle().
			Padding(0, 1).
			Border(lipgloss.NormalBorder()).
			Width(24),

		TitleStyle: lipgloss.NewStyle().
			Bold(true),
	}
}

func (p Panel) Render(title string, subtitle string, strs ...string) string {
	var (
		border lipgloss.Border = p.MainStyle.GetBorderStyle()
		style  lipgloss.Style  = lipgloss.NewStyle().
			Foreground(p.MainStyle.GetBorderTopForeground()).
			Background(p.MainStyle.GetBorderTopBackground())

		left   string = style.Render(border.TopLeft)
		middle string = style.Render(border.Top)
		right  string = style.Render(border.TopRight)
		gap    string = style.Render(" ")
	)

	main := p.MainStyle.Copy().BorderTop(false).Render(strs...)

	width := lipgloss.Width(main)
	top := left + middle + right + gap
	top += p.TitleStyle.Inline(true).Render(title)
	top += gap + left

	if subtitle != "" {
		end := p.TitleStyle.Inline(true).Render("[")
		end += p.TitleStyle.Inline(true).Render(subtitle)
		end += p.TitleStyle.Inline(true).Render("]")
		end += middle + right

		top += strings.Repeat(middle, width-lipgloss.Width(top)-lipgloss.Width(end))
		top += end
	} else {
		top += strings.Repeat(middle, width-lipgloss.Width(top)-1)
		top += right
	}

	return top + "\n" + main
}
