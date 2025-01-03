package app

import "fmt"

func (m model) View() string {
	switch m.state {
	case listView:
		return m.list.View()
	case addView:
		return docstyle.Render(
			"Add a new Book\n\n" +
				"Title: " + m.inputs[titleInput].View() + "\n" +
				"Author: " + m.inputs[authorInput].View() + "\n" +
				"Description: " + m.inputs[descInput].View() + "\n" +
				"Genre: " + m.inputs[genreInput].View() + "\n" +
				"Pages: " + m.inputs[pagesInput].View() + "\n" +
				"Chapters: " + m.inputs[chaptersInput].View() + "\n\n" +
				"Press [ctrl + s] to Save, [tab] to Switch Fields, [ESC] to return",
		)
	case progressView:
		return docstyle.Render(fmt.Sprintf("%s by %s\nGenre: %s\n%s\nChapters: %d\nTotalPages: %d", m.currBook.Name, m.currBook.Author, m.currBook.Genre, m.currBook.Descr, m.currBook.Chapters, m.currBook.Pages))
	case confirmDeleteBookView:
		return docstyle.Render("Are you sure you want to delete " + m.currBook.Name + " ? (y/n)\n\n")
	case errorView:
		return docstyle.Render(fmt.Sprintf("Error: %v\n\nPress [r] to return to Add View or [q] to return to List View", m.err))
	}

	return "Invalid State"
}
