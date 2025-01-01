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
	books := []data.Book{
		data.Book{Name: "Grokking Algorithms", Descr: "Algorithms and Datastructures", Pages: 43},
		data.Book{Name: "Writing an Interpreter in Go", Descr: "Theoretical Computer Science", Pages: 32},
	}
	return model{
		books:    books,
		selected: make(map[int]struct{}),
		list:     list.New(data.BookToItems(books), list.NewDefaultDelegate(), 20, 14),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}