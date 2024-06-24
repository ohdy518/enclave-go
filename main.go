package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Define a type for error messages
type errMsg error

// Model struct to hold the state of the application
type model struct {
	spinner  spinner.Model // Spinner model from Bubbles
	path     string        // Path to display
	quitting bool          // Flag to check if the application is quitting
	err      error         // Error message
}

// Key binding for quitting the application
var quitKeys = key.NewBinding(
	key.WithKeys("esc", "ctrl+c"),
	key.WithHelp("", ""),
)

// Initialize the model with a new spinner and path
func initialModel(path string) model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return model{spinner: s, path: path}
}

// Init method to initialize the spinner tick
func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

// Update method to handle messages and update the model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg: // Handle key messages
		if key.Matches(msg, quitKeys) {
			m.quitting = true
			return m, tea.Quit
		}
		return m, nil

	case errMsg: // Handle error messages
		m.err = msg
		return m, nil

	default: // Handle spinner update
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

// View method to render the UI
func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	str := fmt.Sprintf("\n\n   %s Path: %s %s\n\n", m.spinner.View(), m.path, quitKeys.Help().Desc)
	if m.quitting {
		return str + "\n"
	}
	return str
}

// Main function to run the Bubble Tea program
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./this.exe <path>")
		os.Exit(1)
	}
	path := os.Args[1]
	p := tea.NewProgram(initialModel(path))
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
