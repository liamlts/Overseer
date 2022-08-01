package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	IPs []string
	err error
}

func checkLogs() tea.Msg {
	ips := MonitLogs()
	return logMsg(ips)
}

type logMsg []string

type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }

func (m model) Init() tea.Cmd {
	return checkLogs
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case logMsg:
		m.IPs = []string(msg)
		return m, tea.Quit
	case errMsg:
		m.err = msg
		return m, tea.Quit
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("\nWe had some trouble: %v", m.err)
	}

	s := "Checking log files ... "
	for i := range m.IPs {
		s += m.IPs[i] + "  "
	}
	return s
}

func main() {
	if err := tea.NewProgram(model{}).Start(); err != nil {
		fmt.Printf("There was an error starting the program: %v", err)
		os.Exit(1)
	}
}
