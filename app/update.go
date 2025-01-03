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
				m.currBook = m.list.SelectedItem().(data.Book)
				return m, nil
			case "ctrl+d":
				m.currBook = m.list.SelectedItem().(data.Book)
				m.state = confirmDeleteView
				return m, nil
			}
		case confirmDeleteView:
			switch key {
			case "n":
				m.state = listView
				return m, nil
			case "y":
				cmd := deleteBookCmd(m.currBook, m.store)
				m.state = listView
				return m, cmd
			}
		case addView:
			switch key {
			case "esc":
				m.state = listView
				return m, nil
			case "tab":
				for i := range m.inputs {
					if m.inputs[i].Focused() {
						m.inputs[i].Blur()
						m.inputs[(i+1)%len(m.inputs)].Focus()
						break
					}
				}
				return m, nil
			case "ctrl+s":
				pages, err := strconv.Atoi(m.inputs[pagesInput].Value())
				if err != nil {
					m.err = ErrMsg{err}
					return m, nil
				}
				chapters, err := strconv.Atoi(m.inputs[chaptersInput].Value())
				if err != nil {
					m.err = ErrMsg{err}
					return m, nil
				}

				newBook := data.Book{
					Name:      m.inputs[titleInput].Value(),
					Author:    m.inputs[authorInput].Value(),
					Descr:     m.inputs[descInput].Value(),
					Genre:     m.inputs[genreInput].Value(),
					Chapters:  chapters,
					Pages:     pages,
					Completed: false,
				}
				cmd := saveBookCmd(newBook, m.store)
				for i, input := range m.inputs {
					input.SetValue("")
					m.inputs[i] = input
				}

				m.state = listView
				return m, cmd
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

	case ErrMsg:
		m.err = msg
		return m, tea.Quit

	case BooksMsg:
		m.books = msg.Books
		m.list.SetItems(data.BookToItems(m.books))
		return m, nil
	}

	var cmd tea.Cmd
	switch m.state {
	case listView:
		m.list, cmd = m.list.Update(msg)
	case addView:
		var cmds []tea.Cmd
		for i := range m.inputs {
			m.inputs[i], cmd = m.inputs[i].Update(msg)
			cmds = append(cmds, cmd)
		}
		return m, tea.Batch(cmds...)
	}
	return m, cmd
}
