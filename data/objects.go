package data

import (
	"github.com/charmbracelet/bubbles/list"
)

type Book struct {
	Id       uint
	Name     string
	Descr    string
	Chapters int
	Pages    int
	Genre    string
	Author   string
}

func (b Book) Title() string {
	return b.Name
}
func (b Book) Description() string {
	return b.Descr
}
func (b Book) FilterValue() string {
	return b.Name
}

func BookToItems(books []Book) []list.Item {
	var items []list.Item
	for _, book := range books {
		items = append(items, book)
	}
	return items
}
