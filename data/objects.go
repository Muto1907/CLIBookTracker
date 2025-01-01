package data

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
