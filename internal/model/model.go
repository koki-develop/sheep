package model

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

var (
	_ tea.Model = (*Model)(nil)
)

type keymap struct {
	Abort key.Binding
}

type Model struct {
	Abort  bool
	keymap keymap
}

func New() *Model {
	return &Model{
		Abort: false,
		keymap: keymap{
			Abort: key.NewBinding(key.WithKeys(tea.KeyCtrlC.String())),
		},
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) View() string {
	return "sheep"
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.Abort):
			m.Abort = true
			return m, tea.Quit
		}
	}

	return m, nil
}
