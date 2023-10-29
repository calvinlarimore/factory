package game

import (
	"fmt"

	"github.com/calvinlarimore/factory/render"
	"github.com/calvinlarimore/factory/ui"
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
	c := Client{
		ch: tickChannel,

		term:   pty.Term,
		width:  pty.Window.Width,
		height: pty.Window.Height,
	}
	return c, []tea.ProgramOption{tea.WithAltScreen()}
}

var panelStyle = ui.NewPanelStyle()

type Client struct {
	ch chan struct{}

	term   string
	width  int
	height int

	x, y int

	activeHudPanel int
	hudCursor      int
}

func (c Client) Init() tea.Cmd {
	return waitForTick(c.ch)
}

func (c Client) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		c.height = msg.Height
		c.width = msg.Width
	case tea.KeyMsg:
		if updateHud(&c, msg) {
			break
		}

		switch msg.String() {
		case "up":
			c.y--
		case "down":
			c.y++
		case "left":
			c.x--
		case "right":
			c.x++

		case "q", "ctrl+c":
			return c, tea.Quit
		}

	case tickMsg:
		return c, waitForTick(c.ch)
	}

	return c, nil
}

func (c Client) View() string {
	s := "Your term is %s\n"
	s += "Your window size is x: %d y: %d\n\n"
	s += "Press 'q' to quit"
	left := panelStyle.Render("Terminal Info", "", fmt.Sprintf(s, c.term, c.width, c.height)) + "\n" +
		renderHud(c)

	world := render.RenderWorld(c.width, c.height, c.x, c.y, machines, belts)

	return render.Composite(world, left)
}
