package app

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
		return docstyle.Render("Progress of " + m.currBook.Name)
	}
	return "Invalid State"
}
