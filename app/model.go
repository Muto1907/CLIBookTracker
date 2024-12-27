package app

import (
	"github.com/Muto1907/CLIBookTracker/data"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	books    []data.Book
	cursor   int
	selected map[int]struct{}
}

func NewModel() model {
	return model{
		books: []data.Book{
			{Id: 1, Title: "Grokking Algorithms", Chapters: 10,
				Pages: 200, Genre: "CS Textbook", Author: "Aditya Bhargava"},
		},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
