package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	"ft_Tool/modules/menu"
	"ft_Tool/modules/version"

	"github.com/charmbracelet/lipgloss"
	"github.com/thatisuday/clapper"
)

func main() {
	reg := clapper.NewRegistry()
	blank, _ := reg.Register("")
	blank.AddFlag("version", "v", true, "")
	reg.Register("version")
	command, err := reg.Parse(os.Args[1:])
	if err != nil {
		if strings.Contains(err.Error(), "unknown command") {
			fmt.Println(lipgloss.NewStyle().Foreground(lipgloss.Color("#ff8b94")).Bold(true).Padding(1, 1).Render("Unknown command: " + os.Args[1]))
			return
		}
	}
	switch command.Name {
	case "":
		menu.ShowMenu(command)

	case "version":
		version.ShowVersion()

	}
}
