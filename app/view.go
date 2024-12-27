package app

import "fmt"

func (m model) View() string {
	var s string
	for i, book := range m.books {
		s += fmt.Sprintf("%d, %s by %s (%s) - %d chapters, %d pages", i+1, book.Title, book.Author, book.Genre, book.Chapters, book.Pages)
	}
	return s + "\nPress 'q' to quit."
}
