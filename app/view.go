package app

func (m model) View() string {
	switch m.state {
	case listView:
		return m.list.View()
	case addView:
		return docstyle.Render(
			"Add a new Book\n\n" +
				"Title: " + m.titleInput.View() + "\n" +
				"Author: " + m.authorInput.View() + "\n" +
				"Genre: " + m.genreInput.View() + "\n" +
				"Description: " + m.descInput.View() + "\n" +
				"Chapters: " + m.chaptersInput.View() + "\n" +
				"Pages: " + m.pagesInput.View() + "\n\n" +
				"Press [ctrl + s] tp Save, [tab] to Switch Fields, [ESC] to return",
		)
	case progressView:
		return docstyle.Render("Progress of " + m.currBook.Name)
	}
	return "Invalid State"
}
