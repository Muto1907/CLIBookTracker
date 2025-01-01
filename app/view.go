package app

func (m model) View() string {
	s := docstyle.Render(m.list.View())
	return s
}
