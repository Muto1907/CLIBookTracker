package app

import (
	"log"

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
	store    *data.Store
	list     list.Model
	currBook data.Book
}

func NewModel(store *data.Store) model {
	/*books := []data.Book{
		{Name: "Grokking Algorithms", Descr: "Algorithms and Datastructures", Pages: 43},
		{Name: "Writing an Interpreter in Go", Descr: "Theoretical Computer Science", Pages: 32},
	}*/
	books, err := store.GetBooks()
	if err != nil {
		log.Fatalf("unable to get books: %v", err)
	}
	return model{
		store: store,
		state: listView,
		books: books,
		list:  list.New(data.BookToItems(books), list.NewDefaultDelegate(), 20, 14),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
