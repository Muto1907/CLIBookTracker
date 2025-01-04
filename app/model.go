package app

import (
	"log"

	"github.com/Muto1907/CLIBookTracker/data"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	listView uint = iota
	addView
	progressView
	confirmDeleteBookView
	confirmDeleteNoteView
	errorView
	noteCreateView
	noteslistView
	notesEditView
)

const (
	titleInput uint = iota
	authorInput
	descInput
	genreInput
	pagesInput
	chaptersInput
)

type model struct {
	books         []data.Book
	progresses    []data.Progress
	state         uint
	store         *data.Store
	list          list.Model
	noteslist     list.Model
	inputs        []textinput.Model
	progressInput textinput.Model
	currBook      data.Book
	currProgress  data.Progress
	err           error
	progressBar   progress.Model
	noteInput     textarea.Model
}

type ErrMsg struct{ err error }

func (e ErrMsg) Error() string { return e.err.Error() }

func NewModel(store *data.Store) model {

	books, err := store.GetBooks()
	if err != nil {
		log.Fatalf("unable to get books: %v", err)
	}

	progressBar := progress.New(
		progress.WithScaledGradient("#DF6020", "#331005"),
		progress.WithWidth(40),
	)

	notesArea := textarea.New()
	notesArea.Placeholder = "What did you read today?"
	notesArea.SetWidth(40)
	notesArea.SetHeight(10)
	model := model{
		store:     store,
		state:     listView,
		books:     books,
		list:      list.New(data.BookToItems(books), list.NewDefaultDelegate(), 0, 0),
		noteslist: list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0),
		inputs: []textinput.Model{
			newTextInput("Enter book title", true), newTextInput("Enter author name", false),
			newTextInput("Enter description", false), newTextInput("Enter genre", false),
			newTextInput("Enter total pages", false), newTextInput("Enter total chapters", false),
		},
		progressInput: newTextInput("Until which page did you read today?", true),
		progressBar:   progressBar,
		noteInput:     notesArea,
	}
	model.list.Title = "Books"
	return model
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
