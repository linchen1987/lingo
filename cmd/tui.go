package cmd

import (
	"fmt"
	"strings"

	"lin/internal/tools"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tool struct {
	name        string
	description string
	actions     []string
}

var toolsList = []tool{
	{name: "datetime", description: "Convert timestamp <-> datetime", actions: []string{}},
	{name: "base58", description: "Base58 encode/decode", actions: []string{"encode", "decode"}},
	{name: "base64", description: "Base64 encode/decode", actions: []string{"encode", "decode"}},
}

type model struct {
	cursor   int
	selected int
	input    string
	output   string
	mode     string
	err      error
}

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#7C3AED")).
			MarginBottom(1)

	toolStyle = lipgloss.NewStyle().
			Padding(0, 2)

	selectedStyle = lipgloss.NewStyle().
			Padding(0, 2).
			Foreground(lipgloss.Color("#7C3AED")).
			Bold(true)

	inputStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7C3AED")).
			Padding(0, 1).
			Width(60)

	outputStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#10B981")).
			Padding(0, 1).
			Width(60)

	errorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#EF4444"))
)

func initialModel() model {
	return model{
		cursor:   0,
		selected: -1,
		mode:     "select_tool",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			if m.mode == "input" {
				m.mode = "select_tool"
				m.input = ""
				m.output = ""
				m.selected = -1
				return m, nil
			}
			return m, tea.Quit

		case "up", "k":
			if m.mode == "select_tool" && m.cursor > 0 {
				m.cursor--
			} else if m.mode == "select_action" && m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.mode == "select_tool" && m.cursor < len(toolsList)-1 {
				m.cursor++
			} else if m.mode == "select_action" && m.cursor < len(toolsList[m.selected].actions)-1 {
				m.cursor++
			}

		case "enter":
			if m.mode == "select_tool" {
				m.selected = m.cursor
				if len(toolsList[m.selected].actions) == 0 {
					m.mode = "input"
				} else {
					m.mode = "select_action"
					m.cursor = 0
				}
			} else if m.mode == "select_action" {
				m.mode = "input"
			} else if m.mode == "input" {
				m.processInput()
			}

		case "backspace":
			if m.mode == "input" && len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}

		case "esc":
			if m.mode == "select_action" {
				m.mode = "select_tool"
				m.cursor = m.selected
				m.selected = -1
			} else if m.mode == "input" {
				if len(toolsList[m.selected].actions) > 0 {
					m.mode = "select_action"
					m.input = ""
					m.output = ""
					m.cursor = 0
				} else {
					m.mode = "select_tool"
					m.cursor = m.selected
					m.selected = -1
					m.input = ""
					m.output = ""
				}
			}

		default:
			if m.mode == "input" && len(msg.String()) == 1 {
				m.input += msg.String()
			}
		}
	}

	return m, nil
}

func (m *model) processInput() {
	if m.selected < 0 {
		return
	}

	t := toolsList[m.selected]

	switch t.name {
	case "datetime":
		result, err := tools.ParseDatetime(m.input)
		if err != nil {
			m.err = err
			m.output = ""
		} else {
			m.err = nil
			m.output = result
		}

	case "base58":
		action := t.actions[m.cursor]
		if action == "encode" {
			m.output = tools.Base58Encode(m.input)
		} else {
			result, err := tools.Base58Decode(m.input)
			if err != nil {
				m.err = err
				m.output = ""
			} else {
				m.err = nil
				m.output = result
			}
		}

	case "base64":
		action := t.actions[m.cursor]
		if action == "encode" {
			m.output = tools.Base64Encode(m.input)
		} else {
			result, err := tools.Base64Decode(m.input)
			if err != nil {
				m.err = err
				m.output = ""
			} else {
				m.err = nil
				m.output = result
			}
		}
	}
}

func (m model) View() string {
	var b strings.Builder

	b.WriteString(titleStyle.Render("🛠  Lin CLI Tools"))
	b.WriteString("\n\n")

	if m.mode == "select_tool" {
		b.WriteString("Select a tool:\n\n")
		for i, t := range toolsList {
			if m.cursor == i {
				b.WriteString(selectedStyle.Render("▶ " + t.name + " - " + t.description))
			} else {
				b.WriteString(toolStyle.Render("  " + t.name + " - " + t.description))
			}
			b.WriteString("\n")
		}
		b.WriteString("\n  ↑/k: up  ↓/j: down  enter: select  q: quit\n")
	} else if m.mode == "select_action" {
		t := toolsList[m.selected]
		b.WriteString(fmt.Sprintf("Select action for %s:\n\n", t.name))
		for i, action := range t.actions {
			if m.cursor == i {
				b.WriteString(selectedStyle.Render("▶ " + action))
			} else {
				b.WriteString(toolStyle.Render("  " + action))
			}
			b.WriteString("\n")
		}
		b.WriteString("\n  ↑/k: up  ↓/j: down  enter: select  esc: back  q: quit\n")
	} else if m.mode == "input" {
		t := toolsList[m.selected]
		var action string
		if len(t.actions) > 0 {
			action = t.actions[m.cursor]
			b.WriteString(fmt.Sprintf("%s %s\n\n", t.name, action))
		} else {
			b.WriteString(fmt.Sprintf("%s\n\n", t.name))
		}

		b.WriteString("Input:\n")
		b.WriteString(inputStyle.Render(m.input + "▏"))
		b.WriteString("\n\n")

		if m.output != "" {
			b.WriteString("Output:\n")
			b.WriteString(outputStyle.Render(m.output))
			b.WriteString("\n")
		}

		if m.err != nil {
			b.WriteString("\n")
			b.WriteString(errorStyle.Render(fmt.Sprintf("Error: %v", m.err)))
			b.WriteString("\n")
		}

		b.WriteString("\n  type input, enter: process  esc: back  q: quit\n")
	}

	return b.String()
}

func runTUI() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running TUI: %v\n", err)
	}
}
