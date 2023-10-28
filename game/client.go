package game

import (
	"fmt"

	"github.com/calvinlarimore/factory/render"
	"github.com/calvinlarimore/factory/ui"
	"github.com/calvinlarimore/factory/ui/hud"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
)

type tickMsg struct{}

func waitForTick(ch chan struct{}) tea.Cmd {
	return func() tea.Msg {
		return tickMsg(<-ch)
	}
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, active := s.Pty()
	if !active {
		wish.Fatalln(s, "no active terminal, skipping")
		return nil, nil
	}
	m := model{
		ch:     tickChannel,
		term:   pty.Term,
		width:  pty.Window.Width,
		height: pty.Window.Height,
	}
	return m, []tea.ProgramOption{tea.WithAltScreen()}
}

var panel = ui.NewPanel()

type model struct {
	ch     chan struct{}
	term   string
	width  int
	height int

	x, y int
}

func (m model) Init() tea.Cmd {
	return waitForTick(m.ch)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case tea.KeyMsg:
		if hud.Update(msg) {
			break
		}

		fmt.Println(msg.String())

		switch msg.String() {
		case "up":
			m.y--
		case "down":
			m.y++
		case "left":
			m.x--
		case "right":
			m.x++

		case "q", "ctrl+c":
			return m, tea.Quit
		}
	case tickMsg:
		return m, waitForTick(m.ch)
	}

	return m, nil
}

func (m model) View() string {
	s := "Your term is %s\n"
	s += "Your window size is x: %d y: %d\n\n"
	s += "Press 'q' to quit"
	left := panel.Render("Terminal Info", "", fmt.Sprintf(s, m.term, m.width, m.height)) + "\n" +
		hud.View()

	world := render.RenderWorld(m.width, m.height, m.x, m.y, machines, belts)

	return render.Composite(world, left)
}
