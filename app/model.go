package app

import (
	"log"

	"github.com/Muto1907/CLIBookTracker/data"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	listView uint = iota
	addView
	progressView
	confirmDeleteView
)

const (
	titleInput uint = iota
	authorInput
	descInput
	genreInput
	pagesInput
	chaptersInput
)

var docstyle = lipgloss.NewStyle().Margin(1, 2)

type model struct {
	books    []data.Book
	state    uint
	store    *data.Store
	list     list.Model
	inputs   []textinput.Model
	currBook data.Book
	err      error
}

type ErrMsg struct{ err error }

func (e ErrMsg) Error() string { return e.err.Error() }

func NewModel(store *data.Store) model {

	books, err := store.GetBooks()
	if err != nil {
		log.Fatalf("unable to get books: %v", err)
	}
	return model{
		store: store,
		state: listView,
		books: books,
		list:  list.New(data.BookToItems(books), list.NewDefaultDelegate(), 20, 14),
		inputs: []textinput.Model{
			newTextInput("Enter book title", true), newTextInput("Enter author name", false),
			newTextInput("Enter description", false), newTextInput("Enter genre", false),
			newTextInput("Enter total pages", false), newTextInput("Enter total chapters", false),
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func newTextInput(placeholder string, focus bool) textinput.Model {
	txt := textinput.New()
	txt.Placeholder = placeholder
	if focus {
		txt.Focus()
	}
	return txt
}
