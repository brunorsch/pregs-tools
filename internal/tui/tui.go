package tui

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct{}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	return "Hello, World!\nPressione 'q' para sair."
}

func RunTUI() {
	p := tea.NewProgram(model{})
	if err := p.Start(); err != nil {
		os.Exit(1)
	}
} 