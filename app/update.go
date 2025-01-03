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
				m.state = confirmDeleteBookView
				return m, nil
			}
		case confirmDeleteBookView:
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
					return m, sendErrorMsg(err)
				}
				chapters, err := strconv.Atoi(m.inputs[chaptersInput].Value())
				if err != nil {
					return m, sendErrorMsg(err)
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
			switch key {
			case "ctrl+c":
				m.state = listView
				return m, nil
			case "tab":
				if m.progressInput.Focused() {
					m.progressInput.Blur()
					m.noteInput.Focus()
				} else if m.noteInput.Focused() {
					m.noteInput.Blur()
					m.progressInput.Focus()
				}
			case "ctrl+s":
				m.state = noteCreateView
			}
		case errorView:
			if key == "q" {
				m.state = listView
			} else if key == "r" {
				m.state = addView
			}
			m.err = nil
			return m, nil

		}

	case tea.WindowSizeMsg:
		h, v := docstyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
		return m, nil

	case ErrMsg:
		m.err = msg
		m.state = errorView
		return m, nil

	case BooksMsg:
		m.books = msg.Books
		m.list.SetItems(data.BookToItems(m.books))
		return m, nil
	}

	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch m.state {
	case listView:
		m.list, cmd = m.list.Update(msg)
	case addView:
		for i := range m.inputs {
			m.inputs[i], cmd = m.inputs[i].Update(msg)
			cmds = append(cmds, cmd)
		}
		return m, tea.Batch(cmds...)
	case progressView:
		m.progressInput, cmd = m.progressInput.Update(msg)
		cmds = append(cmds, cmd)
		m.noteInput, cmd = m.noteInput.Update(msg)
		cmds = append(cmds, cmd)
		return m, tea.Batch(cmds...)
	}

	return m, cmd
}
