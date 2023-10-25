package client

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
)

func TeaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, active := s.Pty()
	if !active {
		wish.Fatalln(s, "no active terminal, skipping")
		return nil, nil
	}
	m := model{
		term:   pty.Term,
		width:  pty.Window.Width,
		height: pty.Window.Height,
	}
	return m, []tea.ProgramOption{tea.WithAltScreen()}
}

var style = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.ANSIColor(15)).
	Background(lipgloss.ANSIColor(6)).
	Padding(2, 4).
	Width(24)

type model struct {
	term   string
	width  int
	height int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Your term is %s\n"
	s += "Your window size is x: %d y: %d\n\n"
	s += "Press 'q' to quit"
	return style.Render(fmt.Sprintf(s, m.term, m.width, m.height))
}
