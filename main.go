package main

import (
	"log"
	"os"

	"github.com/Muto1907/CLIBookTracker/app"
	"github.com/Muto1907/CLIBookTracker/data"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	store := &data.Store{}
	if err := store.Init(); err != nil {
		log.Fatalf("Unable to init store: %v", err)
	}

	p := tea.NewProgram(app.NewModel(store), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatalf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
