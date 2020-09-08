package store

import "example.com/m/v2/db"

// Repository is ...
type Repository struct{}

// GetBooks is ...
func (r Repository) GetBooks() ([]Book, error) {
	rows, err := db.DBCon.Query(
		"SELECT id, isbn, title, price  FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []Book{}
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Isbn, &b.Title, &b.Price); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil

}

// GetBook is ...
func (r Repository) GetBook(id int) (Book, error) {
	var b Book
	row := db.DBCon.QueryRow("SELECT id, isbn, title, price  FROM books WHERE id=$1", id)
	if err := row.Scan(&b.ID, &b.Isbn, &b.Title, &b.Price); err != nil {
		return Book{}, err
	}
	return b, nil
}

// CreateBook is ...
func (r Repository) CreateBook(book Book) (int, error) {
	var id int
	err := db.DBCon.QueryRow(
		"INSERT INTO books(isbn,title,price) VALUES($1, $2, $3) RETURNING id", book.Isbn, book.Title, book.Price).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

//UpdateBook is ...
func (r Repository) UpdateBook(book Book) error {
	_, err :=
		db.DBCon.Exec("UPDATE books SET isbn=$1, title=$2, price=$3 WHERE id=$4",
			book.Isbn, book.Title, book.Price, book.ID)

	return err
}

//DeleteBook is ...
func (r Repository) DeleteBook(id int) error {
	_, err := db.DBCon.Exec("DELETE FROM books WHERE id=$1", id)

	return err
}
