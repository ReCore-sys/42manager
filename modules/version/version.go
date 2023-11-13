package version

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func ShowVersion() {

	versionStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ff99e6")).
		Align(lipgloss.Center).
		Bold(true).
		Padding(1, 1)
	fmt.Printf("%s\n", versionStyle.Render("Version: ✨ v0.0.1 ✨"))
}
