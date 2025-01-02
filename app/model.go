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
)

var docstyle = lipgloss.NewStyle().Margin(1, 2)

type model struct {
	books         []data.Book
	state         uint
	store         *data.Store
	list          list.Model
	titleInput    textinput.Model
	authorInput   textinput.Model
	genreInput    textinput.Model
	descInput     textinput.Model
	pagesInput    textinput.Model
	chaptersInput textinput.Model
	currBook      data.Book
}

func NewModel(store *data.Store) model {
	titleInput := textinput.New()
	titleInput.Placeholder = "Enter book title"
	titleInput.Focus()

	authorInput := textinput.New()
	authorInput.Placeholder = "Enter author name"

	genreInput := textinput.New()
	genreInput.Placeholder = "Enter genre"

	descInput := textinput.New()
	descInput.Placeholder = "Enter description"

	pagesInput := textinput.New()
	pagesInput.Placeholder = "Enter total pages"

	chaptersInput := textinput.New()
	chaptersInput.Placeholder = "Enter total chapters"

	books, err := store.GetBooks()
	if err != nil {
		log.Fatalf("unable to get books: %v", err)
	}
	return model{
		store:         store,
		state:         listView,
		books:         books,
		list:          list.New(data.BookToItems(books), list.NewDefaultDelegate(), 20, 14),
		titleInput:    titleInput,
		authorInput:   authorInput,
		genreInput:    genreInput,
		descInput:     descInput,
		pagesInput:    pagesInput,
		chaptersInput: chaptersInput,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
