package helpcommand

import (
	_ "embed"
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

//go:embed commands.txt
var commands string

func ShowHelp() {

	commandsStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#869ef4")).
		Bold(true).
		Padding(1, 1)
	fmt.Printf("%s\n", commandsStyle.Render(commands))
}
