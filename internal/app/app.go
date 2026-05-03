package app

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"

	"liun-dots/internal/ui"
)

func Run() error {
	defer clearTerminal()
	p := tea.NewProgram(ui.NewModel())
	_, err := p.Run()
	return err
}

func clearTerminal() {
	// Limpieza al salir para no dejar residuos visuales en consola.
	fmt.Print("\x1b[2J\x1b[H")
}
