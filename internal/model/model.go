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
	aborted  bool
	duration time.Duration
	keymap   keymap
}

func New(cfg *Config) *Model {
	return &Model{
		aborted:  false,
		duration: cfg.Duration,
		keymap: keymap{
			Abort: key.NewBinding(key.WithKeys(tea.KeyCtrlC.String())),
		},
	}
}

func (m *Model) Aborted() bool {
	return m.aborted
}

func (m *Model) Init() tea.Cmd {
	return tea.Batch(
		m.sleep(),
	)
}

func (m *Model) View() string {
	return "sheep"
}

type sleptMsg struct{}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case sleptMsg:
		return m, tea.Quit
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
