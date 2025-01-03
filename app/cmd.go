package app

import (
	"github.com/Muto1907/CLIBookTracker/data"
	tea "github.com/charmbracelet/bubbletea"
)

func deleteBookCmd(book data.Book, store *data.Store) tea.Cmd {
	return func() tea.Msg {
		if err := store.DeleteBook(book); err != nil {
			return ErrMsg{err}
		}
		books, err := store.GetBooks()
		if err != nil {
			return ErrMsg{err}
		}
		return BooksMsg{books}
	}
}

func saveBookCmd(book data.Book, store *data.Store) tea.Cmd {
	return func() tea.Msg {
		if err := store.SaveBook(book); err != nil {
			return ErrMsg{err}
		}
		books, err := store.GetBooks()
		if err != nil {
			return ErrMsg{err}
		}
		return BooksMsg{books}
	}
}

type BooksMsg struct {
	Books []data.Book
}

func sendErrorMsg(err error) tea.Cmd {
	return func() tea.Msg {
		return ErrMsg{err}
	}
}

func deleteProgressCmd(prog data.Progress, store *data.Store) tea.Cmd {
	return func() tea.Msg {
		if err := store.DeleteProgress(prog); err != nil {
			return ErrMsg{err}
		}
		progresses, err := store.GetProgress()
		if err != nil {
			return ErrMsg{err}
		}
		return ProgressMsg{progresses}
	}
}

func saveProgressCmd(prog data.Progress, store *data.Store) tea.Cmd {
	return func() tea.Msg {
		if err := store.SaveProgress(prog); err != nil {
			return ErrMsg{err}
		}
		progresses, err := store.GetProgress()
		if err != nil {
			return ErrMsg{err}
		}
		return ProgressMsg{progresses}
	}
}

type ProgressMsg struct {
	Progresses []data.Progress
}
