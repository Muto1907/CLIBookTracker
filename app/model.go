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
type item struct {
	title, desc string
	pages       int
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }
func (i item) Pages() int          { return i.pages }

func NewModel() model {
	items := []list.Item{
		item{title: "Grokking Algorithms", desc: "Algorithms and Datastructures", pages: 43},
		item{title: "Writing an Interpreter in Go", desc: "Theoretical Computer Science", pages: 32},
	}
	return model{
		books: []data.Book{
			{Id: 1, Title: "Grokking Algorithms", Chapters: 10,
				Pages: 200, Genre: "CS Textbook", Author: "Aditya Bhargava"},
		},
		selected: make(map[int]struct{}),
		list:     list.New(items, list.NewDefaultDelegate(), 20, 14),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
