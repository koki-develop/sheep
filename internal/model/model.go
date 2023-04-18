package model

import (
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

var (
	_ tea.Model = (*Model)(nil)
)

type keymap struct {
	Abort key.Binding
}

type Config struct {
	Duration time.Duration
}

type Model struct {
	aborted     bool
	duration    time.Duration
	windowWidth int
	keymap      keymap
	frame       int
	sheep       sheep
	awake       bool
}

func New(cfg *Config) *Model {
	return &Model{
		aborted:     false,
		duration:    cfg.Duration,
		windowWidth: 0,
		keymap: keymap{
			Abort: key.NewBinding(key.WithKeys(tea.KeyCtrlC.String())),
		},
		frame: 0,
		sheep: simpleSheep,
		awake: false,
	}
}

func (m *Model) Aborted() bool {
	return m.aborted
}

func (m *Model) Init() tea.Cmd {
	return tea.Batch(
		m.sleep(),
		m.tick(),
	)
}

func (m *Model) View() string {
	if m.awake {
		return m.sheep.Awake
	} else {
		return m.sheep.Asleep[m.frame]
	}
}

type sleptMsg struct{}
type tickMsg struct{}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowWidth = msg.Width
	case sleptMsg:
		m.awake = true
		return m, tea.Quit
	case tickMsg:
		m.frame = (m.frame + 1) % len(m.sheep.Asleep)
		return m, m.tick()
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.Abort):
			m.aborted = true
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m *Model) sleep() tea.Cmd {
	return tea.Tick(m.duration, func(t time.Time) tea.Msg {
		return sleptMsg{}
	})
}

func (m *Model) tick() tea.Cmd {
	return tea.Tick(500*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
}
