package data

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Store struct {
	conn *sql.DB
}

func (s *Store) Init() error {
	var err error
	s.conn, err = sql.Open("sqlite3", "./books.db")
	if err != nil {
		return err
	}

	createTableStmt := `CREATE TABLE IF NOT EXISTS books (
		id integer not null primary key AUTOINCREMENT,
		name text not null,
		descr text not null,
		pages integer not null,
		genre text not null,
		author text not null,
		completed boolean not null
	);`

	if _, err = s.conn.Exec(createTableStmt); err != nil {
		return err
	}
	return nil
}

func (s *Store) GetBooks() ([]Book, error) {
	rows, err := s.conn.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	books := []Book{}
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.Id, &book.Name, &book.Descr, &book.Pages,
			&book.Genre, &book.Author, &book.Completed); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (s *Store) SaveBook(book Book) error {
	upsertQuery := `INSERT INTO books (name, descr, pages, genre, author, completed)
	VALUES (?, ?, ?, ?, ?, ?)
	ON CONFLICT(id) DO UPDATE
	SET 
		name = excluded.name, 
		descr = excluded.descr, 
		pages = excluded.pages, 
		genre = excluded.genre, 
		author = excluded.author, 
		completed = excluded.completed;
	`

	if _, err := s.conn.Exec(upsertQuery, book.Id, book.Name, book.Descr, book.Pages, book.Genre, book.Author, book.Completed); err != nil {
		return err
	}
	return nil
}
