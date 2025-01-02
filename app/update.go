package app

import (
	"strconv"

	"github.com/Muto1907/CLIBookTracker/data"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()
		switch m.state {
		case listView:
			switch key {
			case "q":
				return m, tea.Quit
			case "a":
				m.state = addView
				return m, nil
			case "enter":
				m.state = progressView
				return m, nil
			}
		case addView:
			switch key {
			case "esc":
				m.state = listView
				return m, nil
			case "tab":
				if m.titleInput.Focused() {
					m.titleInput.Blur()
					m.authorInput.Focus()
				} else if m.authorInput.Focused() {
					m.authorInput.Blur()
					m.genreInput.Focus()
				} else if m.genreInput.Focused() {
					m.genreInput.Blur()
					m.descInput.Focus()
				} else if m.descInput.Focused() {
					m.descInput.Blur()
					m.chaptersInput.Focus()
				} else if m.chaptersInput.Focused() {
					m.chaptersInput.Blur()
					m.pagesInput.Focus()
				} else if m.pagesInput.Focused() {
					m.pagesInput.Blur()
					m.titleInput.Focus()
				}
				return m, nil
			case "ctrl+s":
				pages, _ := strconv.Atoi(m.pagesInput.Value())
				chapters, _ := strconv.Atoi(m.chaptersInput.Value())

				newBook := data.Book{
					Name:      m.titleInput.Value(),
					Author:    m.authorInput.Value(),
					Descr:     m.descInput.Value(),
					Genre:     m.genreInput.Value(),
					Chapters:  chapters,
					Pages:     pages,
					Completed: false,
				}
				if err := m.store.SaveBook(newBook); err != nil {
					panic(err)
				}
				m.titleInput.SetValue("")
				m.authorInput.SetValue("")
				m.genreInput.SetValue("")
				m.descInput.SetValue("")
				m.pagesInput.SetValue("")
				m.chaptersInput.SetValue("")
				m.state = listView
				m.books, _ = m.store.GetBooks()
				m.list.SetItems(data.BookToItems(m.books))
				return m, nil
			}
		case progressView:
			if key == "q" {
				m.state = listView
				return m, nil
			}
		}

	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil
	}

	var cmd tea.Cmd
	switch m.state {
	case listView:
		m.list, cmd = m.list.Update(msg)
	case addView:
		var cmds []tea.Cmd
		m.titleInput, cmd = m.titleInput.Update(msg)
		cmds = append(cmds, cmd)
		m.authorInput, cmd = m.authorInput.Update(msg)
		cmds = append(cmds, cmd)
		m.genreInput, cmd = m.genreInput.Update(msg)
		cmds = append(cmds, cmd)
		m.descInput, cmd = m.descInput.Update(msg)
		cmds = append(cmds, cmd)
		m.pagesInput, cmd = m.pagesInput.Update(msg)
		cmds = append(cmds, cmd)
		m.chaptersInput, cmd = m.chaptersInput.Update(msg)
		cmds = append(cmds, cmd)
		return m, tea.Batch(cmds...)
	}
	return m, cmd
}
