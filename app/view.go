package app

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	listStyle    = lipgloss.NewStyle().Padding(1, 2).Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("62"))
	inputStyle   = lipgloss.NewStyle().Padding(1, 2).AlignHorizontal(lipgloss.Center).Bold(true).Border(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("63"))
	errorStyle   = lipgloss.NewStyle().Padding(1, 2).AlignHorizontal(lipgloss.Center).AlignVertical(lipgloss.Center).Blink(true).Background(lipgloss.Color("9")).Border(lipgloss.ThickBorder()).BorderForeground(lipgloss.Color("59"))
	confirmStyle = lipgloss.NewStyle().Padding(1, 2).AlignHorizontal(lipgloss.Center).AlignVertical(lipgloss.Center).Blink(true).Background(lipgloss.Color("63")).Border(lipgloss.ThickBorder()).BorderForeground(lipgloss.Color("59"))
)

func (m model) View() string {
	switch m.state {
	case listView:
		return listStyle.Render(m.list.View())
	case addView:
		return inputStyle.Render(lipgloss.JoinVertical(lipgloss.Left,
			"Add a new Book\n",
			"Title: "+m.inputs[titleInput].View(),
			"Author: "+m.inputs[authorInput].View(),
			"Description: "+m.inputs[descInput].View(),
			"Genre: "+m.inputs[genreInput].View(),
			"Pages: "+m.inputs[pagesInput].View(),
			"Chapters: "+m.inputs[chaptersInput].View(),
			"\n\nPress [ctrl + s] to Save, [tab] to Switch Fields, [ESC] to return",
		))
	case progressView:
		return inputStyle.Render(fmt.Sprintf("%s by %s\nGenre: %s\n%s\nChapters: %d\nTotalPages: %d\n\nDid you read any pages today?\n%s\n\n%s", m.currBook.Name, m.currBook.Author,
			m.currBook.Genre, m.currBook.Descr, m.currBook.Chapters, m.currBook.Pages, m.progressInput.View(), m.noteInput.View()))
	case confirmDeleteBookView:
		return confirmStyle.Render("Are you sure you want to delete " + m.currBook.Name + " ? \n(y/n)")
	case confirmDeleteNoteView:
		return confirmStyle.Render("Are you sure you want to delete your progress from " + m.currProgress.Date + " ?\n(y/n)")
	case errorView:
		return errorStyle.Render(fmt.Sprintf("Error: %v\n\nPress [r] to return to Add View or [q] to return to List View", m.err))
	case noteslistView:
		return listStyle.Render(m.noteslist.View())
	case notesEditView:
		return inputStyle.Render("Edit note from " + m.currProgress.Date + "\n\n" +
			m.progressInput.View() + "\n\n" + m.noteInput.View() + "\n\n" +
			"Press [ctrl + s] to Save, [tab] to Switch Fields, [ESC] to return")

	}

	return "Invalid State"
}
