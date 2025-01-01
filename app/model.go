package app

import (
	"github.com/Muto1907/CLIBookTracker/data"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docstyle = lipgloss.NewStyle().Margin(1, 2)

type model struct {
	books    []data.Book
	cursor   int
	selected map[int]struct{}
	list     list.Model
}

func NewModel() model {
	items := []list.Item{
		data.Book{Name: "Grokking Algorithms", Descr: "Algorithms and Datastructures", Pages: 43},
		data.Book{Name: "Writing an Interpreter in Go", Descr: "Theoretical Computer Science", Pages: 32},
	}
	return model{
		books: []data.Book{
			{Id: 1, Name: "Grokking Algorithms", Chapters: 10,
				Pages: 200, Genre: "CS Textbook", Author: "Aditya Bhargava"},
		},
		selected: make(map[int]struct{}),
		list:     list.New(items, list.NewDefaultDelegate(), 20, 14),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
