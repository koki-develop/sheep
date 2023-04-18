package model

import tea "github.com/charmbracelet/bubbletea"

var (
	_ tea.Model = (*Model)(nil)
)

type Model struct{}

func New() *Model {
	return &Model{}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) View() string {
	return "sheep"
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, tea.Quit
}
