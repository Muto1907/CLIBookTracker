package data

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
)

type Book struct {
	Id        int64
	Name      string
	Descr     string
	Chapters  int
	Pages     int
	Genre     string
	Author    string
	Completed bool
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

type Progress struct {
	Id         int64
	Book_id    int64
	Start_Page int
	End_Page   int
	Note       string
	Date       string
}

func (p Progress) Title() string {
	if len(p.Note) > 30 {
		return p.Note[:30]
	}
	return p.Note
}

func (p Progress) Description() string {
	return fmt.Sprintf("From %d To %d", p.Start_Page, p.End_Page)
}

func (p Progress) FilterValue() string {
	return p.Note
}

func ProgressToItems(prog []Progress) []list.Item {
	var items []list.Item
	for _, progr := range prog {
		items = append(items, progr)
	}
	return items
}
