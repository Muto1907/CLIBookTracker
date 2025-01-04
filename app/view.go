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
		return docstyle.Render(fmt.Sprintf("%s by %s\nGenre: %s\n%s\nChapters: %d\nTotalPages: %d\n\nDid you read any pages today?\n%s\n\n%s", m.currBook.Name, m.currBook.Author,
			m.currBook.Genre, m.currBook.Descr, m.currBook.Chapters, m.currBook.Pages, m.progressInput.View(), m.noteInput.View()))
	case confirmDeleteBookView:
		return docstyle.Render("Are you sure you want to delete " + m.currBook.Name + " ? (y/n)\n\n")
	case confirmDeleteNoteView:
		return docstyle.Render("Are you sure you want to delete your progress from " + m.currProgress.Date + " ? (y/n)\n\n")
	case errorView:
		return docstyle.Render(fmt.Sprintf("Error: %v\n\nPress [r] to return to Add View or [q] to return to List View", m.err))
	case noteslistView:
		return docstyle.Render(m.noteslist.View())
	case notesEditView:
		return docstyle.Render("Edit note from " + m.currProgress.Date + "\n\n" +
			m.progressInput.View() + "\n\n" + m.noteInput.View() + "\n\n" +
			"Press [ctrl + s] to Save, [tab] to Switch Fields, [ESC] to return")

	}

	return "Invalid State"
}
