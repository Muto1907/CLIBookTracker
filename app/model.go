package app

import (
	"github.com/Muto1907/CLIBookTracker/data"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	listView uint = iota
	addView
	progressView
)

var docstyle = lipgloss.NewStyle().Margin(1, 2)

type model struct {
	books    []data.Book
	state    uint
	selected map[int]struct{}
	list     list.Model
}

func NewModel() model {
	books := []data.Book{
		{Name: "Grokking Algorithms", Descr: "Algorithms and Datastructures", Pages: 43},
		{Name: "Writing an Interpreter in Go", Descr: "Theoretical Computer Science", Pages: 32},
	}
	return model{
		state:    listView,
		books:    books,
		selected: make(map[int]struct{}),
		list:     list.New(data.BookToItems(books), list.NewDefaultDelegate(), 20, 14),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
