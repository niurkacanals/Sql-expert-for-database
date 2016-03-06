// Package models contains the types for schema 'booktest'.
package models

// GENERATED BY XO. DO NOT EDIT.

import (
	"errors"
	"time"
)

// Book represents a row from 'booktest.books'.
type Book struct {
	BookID    int        // book_id
	AuthorID  int        // author_id
	Isbn      string     // isbn
	BookType  BookType   // book_type
	Title     string     // title
	Year      int        // year
	Available *time.Time // available
	Tags      string     // tags

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Book exists in the database.
func (b *Book) Exists() bool {
	return b._exists
}

// Deleted provides information if the Book has been deleted from the database.
func (b *Book) Deleted() bool {
	return b._deleted
}

// Insert inserts the Book to the database.
func (b *Book) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if b._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO booktest.books (` +
		`author_id, isbn, book_type, title, year, available, tags` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, b.AuthorID, b.Isbn, b.BookType, b.Title, b.Year, b.Available, b.Tags)
	res, err := db.Exec(sqlstr, b.AuthorID, b.Isbn, b.BookType, b.Title, b.Year, b.Available, b.Tags)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	b.BookID = int(id)
	b._exists = true

	return nil
}

// Update updates the Book in the database.
func (b *Book) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !b._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if b._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE booktest.books SET ` +
		`author_id = ?, isbn = ?, book_type = ?, title = ?, year = ?, available = ?, tags = ?` +
		` WHERE book_id = ?`

	// run query
	XOLog(sqlstr, b.AuthorID, b.Isbn, b.BookType, b.Title, b.Year, b.Available, b.Tags, b.BookID)
	_, err = db.Exec(sqlstr, b.AuthorID, b.Isbn, b.BookType, b.Title, b.Year, b.Available, b.Tags, b.BookID)
	return err
}

// Save saves the Book to the database.
func (b *Book) Save(db XODB) error {
	if b.Exists() {
		return b.Update(db)
	}

	return b.Insert(db)
}

// Delete deletes the Book from the database.
func (b *Book) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !b._exists {
		return nil
	}

	// if deleted, bail
	if b._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM booktest.books WHERE book_id = ?`

	// run query
	XOLog(sqlstr, b.BookID)
	_, err = db.Exec(sqlstr, b.BookID)
	if err != nil {
		return err
	}

	// set deleted
	b._deleted = true

	return nil
}

// Author returns the Author associated with the Book's AuthorID (author_id).
//
// Generated from foreign key 'books_ibfk_1'.
func (b *Book) Author(db XODB) (*Author, error) {
	return AuthorByAuthorID(db, b.AuthorID)
}

// BooksByAuthorID retrieves a row from 'booktest.books' as a Book.
//
// Generated from index 'author_id'.
func BooksByAuthorID(db XODB, authorID int) ([]*Book, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`book_id, author_id, isbn, book_type, title, year, available, tags ` +
		`FROM booktest.books ` +
		`WHERE author_id = ?`

	// run query
	XOLog(sqlstr, authorID)
	q, err := db.Query(sqlstr, authorID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Book{}
	for q.Next() {
		b := Book{
			_exists: true,
		}

		// scan
		err = q.Scan(&b.BookID, &b.AuthorID, &b.Isbn, &b.BookType, &b.Title, &b.Year, &b.Available, &b.Tags)
		if err != nil {
			return nil, err
		}

		res = append(res, &b)
	}

	return res, nil
}

// BookByBookID retrieves a row from 'booktest.books' as a Book.
//
// Generated from index 'books_book_id_pkey'.
func BookByBookID(db XODB, bookID int) (*Book, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`book_id, author_id, isbn, book_type, title, year, available, tags ` +
		`FROM booktest.books ` +
		`WHERE book_id = ?`

	// run query
	XOLog(sqlstr, bookID)
	b := Book{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, bookID).Scan(&b.BookID, &b.AuthorID, &b.Isbn, &b.BookType, &b.Title, &b.Year, &b.Available, &b.Tags)
	if err != nil {
		return nil, err
	}

	return &b, nil
}

// BooksByTitle retrieves a row from 'booktest.books' as a Book.
//
// Generated from index 'books_title_idx'.
func BooksByTitle(db XODB, title string, year int) ([]*Book, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`book_id, author_id, isbn, book_type, title, year, available, tags ` +
		`FROM booktest.books ` +
		`WHERE title = ? AND year = ?`

	// run query
	XOLog(sqlstr, title, year)
	q, err := db.Query(sqlstr, title, year)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Book{}
	for q.Next() {
		b := Book{
			_exists: true,
		}

		// scan
		err = q.Scan(&b.BookID, &b.AuthorID, &b.Isbn, &b.BookType, &b.Title, &b.Year, &b.Available, &b.Tags)
		if err != nil {
			return nil, err
		}

		res = append(res, &b)
	}

	return res, nil
}

// BookByIsbn retrieves a row from 'booktest.books' as a Book.
//
// Generated from index 'isbn'.
func BookByIsbn(db XODB, isbn string) (*Book, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`book_id, author_id, isbn, book_type, title, year, available, tags ` +
		`FROM booktest.books ` +
		`WHERE isbn = ?`

	// run query
	XOLog(sqlstr, isbn)
	b := Book{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, isbn).Scan(&b.BookID, &b.AuthorID, &b.Isbn, &b.BookType, &b.Title, &b.Year, &b.Available, &b.Tags)
	if err != nil {
		return nil, err
	}

	return &b, nil
}
