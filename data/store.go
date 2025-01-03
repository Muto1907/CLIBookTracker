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
		chapters integer not null,
		pages integer not null,
		genre text not null,
		author text not null,
		completed boolean not null
	);`

	if _, err = s.conn.Exec(createTableStmt); err != nil {
		return err
	}

	createProgressTableStmt := `CREATE TABLE IF NOT EXISTS progress (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		book_id INTEGER NOT NULL,
		start_page INTEGER NOT NULL,
		end_page INTEGER NOT NULL,
		note TEXT,
		FOREIGN KEY (book_id) REFERENCES books (id)
	);`

	if _, err = s.conn.Exec(createProgressTableStmt); err != nil {
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
		if err := rows.Scan(&book.Id, &book.Name, &book.Descr, &book.Chapters, &book.Pages,
			&book.Genre, &book.Author, &book.Completed); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (s *Store) SaveBook(book Book) error {
	if book.Id == 0 {
		insertQuery := `INSERT INTO books (name, descr, chapters, pages, genre, author, completed)
	VALUES (?, ?, ?, ?, ?, ?, ?);`
		_, err := s.conn.Exec(insertQuery, book.Name, book.Descr, book.Chapters, book.Pages, book.Genre, book.Author, book.Completed)
		return err
	}
	upsertQuery := `INSERT INTO books (id, name, descr, chapters, pages, genre, author, completed)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	ON CONFLICT(id) DO UPDATE
	SET 
		name = excluded.name, 
		descr = excluded.descr, 
		chapters = excluded.chapters,
		pages = excluded.pages, 
		genre = excluded.genre, 
		author = excluded.author, 
		completed = excluded.completed;
	`

	if _, err := s.conn.Exec(upsertQuery, book.Id, book.Name, book.Descr, book.Chapters, book.Pages, book.Genre, book.Author, book.Completed); err != nil {
		return err
	}
	return nil
}

func (s *Store) DeleteBook(book Book) error {
	deleteQuery := `DELETE FROM books WHERE id = ?`

	if _, err := s.conn.Exec(deleteQuery, book.Id); err != nil {
		return err
	}
	return nil
}

func (s *Store) GetProgress() ([]Progress, error) {
	rows, err := s.conn.Query("SELECT * FROM progress")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	progress := []Progress{}
	for rows.Next() {
		var prog Progress
		if err := rows.Scan(&prog.Id, &prog.Book_id, &prog.Start_Page, &prog.End_Page, &prog.Note); err != nil {
			return nil, err
		}
		progress = append(progress, prog)
	}
	return progress, nil
}

func (s *Store) GetLatestProgress(book Book) (Progress, error) {
	row := s.conn.QueryRow("SELECT to_page FROM progress WHERE book_id = ? ORDER BY id DESC LIMIT 1;", book.Id)
	var prog Progress
	if err := row.Scan(&prog.Id, &prog.Book_id, &prog.Start_Page, &prog.End_Page, &prog.Note); err != nil {
		return prog, err
	}
	return prog, nil
}

func (s *Store) SaveProgress(prog Progress) error {
	if prog.Id == 0 {
		insertQuery := `INSERT INTO progress (book_id, start_page, end_page, Note)
	VALUES (?, ?, ?, ?);`
		_, err := s.conn.Exec(insertQuery, prog.Book_id, prog.Start_Page, prog.End_Page, prog.Note)
		return err
	}
	upsertQuery := `INSERT INTO progress (id, book_id, start_page, end_page, Note)
	VALUES (?, ?, ?, ?, ?)
	ON CONFLICT(id) DO UPDATE
	SET 
		from = excluded.start_page,
		to = excluded.end_page, 
		note = excluded.note;
	`

	if _, err := s.conn.Exec(upsertQuery, prog.Id, prog.Book_id, prog.Start_Page, prog.End_Page, prog.Note); err != nil {
		return err
	}
	return nil
}

func (s *Store) DeleteProgress(prog Progress) error {
	deleteQuery := `DELETE FROM progress WHERE id = ?`

	if _, err := s.conn.Exec(deleteQuery, prog.Id); err != nil {
		return err
	}
	return nil
}
