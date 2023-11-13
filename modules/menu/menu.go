package menu

import (
	_ "embed"
	"ft_Tool/modules/helpcommand"
	"ft_Tool/modules/version"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/thatisuday/clapper"
	"golang.org/x/term"
)

//go:embed textbig.txt
var textbig string

//go:embed textsmall.txt
var textsmall string

func ShowMenu(cmd *clapper.CommandConfig) {
	if val, ok := cmd.Flags["version"]; ok && val.Value == "true" {
		version.ShowVersion()
	} else {

		width, _, err := term.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			panic(err)
		}
		style := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ff99e6")).
			BorderForeground(lipgloss.Color("#b5a6f9")).
			Border(lipgloss.RoundedBorder()).
			Padding(1, 2).
			Align(lipgloss.Center).
			Bold(true)
		if width >= 105 {
			println(style.Render(textbig))
		} else {
			println(style.Render(textsmall))
		}
		helpcommand.ShowHelp()
	}
}
