package render

import (
	"strings"

	"github.com/calvinlarimore/factory/belt"
	"github.com/calvinlarimore/factory/machine"
	"github.com/charmbracelet/lipgloss"
)

func RenderWorld(width, height, cx, cy int, machines []machine.Machine, belts []*belt.Belt, players []PlayerInfo) string {
	line := strings.Repeat(".", width)
	world := strings.Split((strings.Repeat(line+"\n", height-1) + line), "\n")

	for _, b := range belts {
		x, y := b.Pos()
		x -= cx
		y -= cy

		if x >= 0 && x < width && y >= 0 && y < height {
			char := "?"

			if b.Item() == 0 {
				char = []string{"▲", "►", "▼", "◄"}[b.Rotation()]
			} else {
				char = getItemSprite(b.Item())
			}

			world[y] = splice(world[y], x, char)
		}
	}

	for _, m := range machines {
		x, y := m.Pos()
		x -= cx
		y -= cy

		for i, line := range strings.Split(m.Sprite(), "\n") {
			if y+i >= 0 && y+i < height {
				for j, char := range line {
					if x+j >= 0 && x+j < width {
						world[y+i] = splice(world[y+i], x+j, string(char))
					}
				}
			}
		}
	}

	for _, p := range players {
		p.x -= cx
		p.y -= cy

		if p.x >= 0 && p.x < width && p.y >= 0 && p.y < height {
			style := lipgloss.NewStyle().Foreground(p.color)
			world[p.y] = splice(world[p.y], p.x, style.Inline(true).Render("☻"))
		}
	}

	// Cursor
	{
		x := width / 2
		y := height / 2

		style := lipgloss.NewStyle().Foreground(lipgloss.ANSIColor(15)).Bold(true)
		world[y] = splice(world[y], x, style.Inline(true).Render("○"))
	}

	return strings.Join(world, "\n")
}

func Composite(world string, hudLeft string) string {
	final := strings.Split(world, "\n")

	for i, s := range strings.Split(hudLeft, "\n") {
		width := lipgloss.Width(s)
		final[i] = s + after(final[i], width)
	}

	return strings.Join(final, "\n")
}

func visualIndex(str string, width int) int {
	for i := lipgloss.Width(str); i >= 0; i-- {
		if lipgloss.Width(str[:i]) == width {
			return i
		}
	}

	return len(str)
}

func before(str string, i int) string {
	return str[:visualIndex(str, i)]
}

func after(str string, i int) string {
	return str[visualIndex(str, i):]
}

func splice(base string, i int, str string) string {
	return before(base, i) + str + after(base, i+lipgloss.Width(str))
}
