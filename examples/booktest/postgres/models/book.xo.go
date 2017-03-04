// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

import (
	"errors"
	"time"
)

// Book represents a row from 'public.books'.
type Book struct {
	BookID    int         `json:"book_id"`   // book_id
	AuthorID  int         `json:"author_id"` // author_id
	Isbn      string      `json:"isbn"`      // isbn
	Booktype  BookType    `json:"booktype"`  // booktype
	Title     string      `json:"title"`     // title
	Year      int         `json:"year"`      // year
	Available time.Time   `json:"available"` // available
	Tags      StringSlice `json:"tags"`      // tags

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

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.books (` +
		`author_id, isbn, booktype, title, year, available, tags` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`) RETURNING book_id`

	// run query
	XOLog(sqlstr, b.AuthorID, b.Isbn, b.Booktype, b.Title, b.Year, b.Available, b.Tags)
	err = db.QueryRow(sqlstr, b.AuthorID, b.Isbn, b.Booktype, b.Title, b.Year, b.Available, b.Tags).Scan(&b.BookID)
	if err != nil {
		return err
	}

	// set existence
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
	const sqlstr = `UPDATE public.books SET (` +
		`author_id, isbn, booktype, title, year, available, tags` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`) WHERE book_id = $8`

	// run query
	XOLog(sqlstr, b.AuthorID, b.Isbn, b.Booktype, b.Title, b.Year, b.Available, b.Tags, b.BookID)
	_, err = db.Exec(sqlstr, b.AuthorID, b.Isbn, b.Booktype, b.Title, b.Year, b.Available, b.Tags, b.BookID)
	return err
}

// Save saves the Book to the database.
func (b *Book) Save(db XODB) error {
	if b.Exists() {
		return b.Update(db)
	}

	return b.Insert(db)
}

// Upsert performs an upsert for Book.
//
// NOTE: PostgreSQL 9.5+ only
func (b *Book) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if b._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.books (` +
		`book_id, author_id, isbn, booktype, title, year, available, tags` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8` +
		`) ON CONFLICT (book_id) DO UPDATE SET (` +
		`book_id, author_id, isbn, booktype, title, year, available, tags` +
		`) = (` +
		`EXCLUDED.book_id, EXCLUDED.author_id, EXCLUDED.isbn, EXCLUDED.booktype, EXCLUDED.title, EXCLUDED.year, EXCLUDED.available, EXCLUDED.tags` +
		`)`

	// run query
	XOLog(sqlstr, b.BookID, b.AuthorID, b.Isbn, b.Booktype, b.Title, b.Year, b.Available, b.Tags)
	_, err = db.Exec(sqlstr, b.BookID, b.AuthorID, b.Isbn, b.Booktype, b.Title, b.Year, b.Available, b.Tags)
	if err != nil {
		return err
	}

	// set existence
	b._exists = true

	return nil
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
	const sqlstr = `DELETE FROM public.books WHERE book_id = $1`

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
// Generated from foreign key 'books_author_id_fkey'.
func (b *Book) Author(db XODB) (*Author, error) {
	return AuthorByAuthorID(db, b.AuthorID)
}

// BookByIsbn retrieves a row from 'public.books' as a Book.
//
// Generated from index 'books_isbn_key'.
func BookByIsbn(db XODB, isbn string) (*Book, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`book_id, author_id, isbn, booktype, title, year, available, tags ` +
		`FROM public.books ` +
		`WHERE isbn = $1`

	// run query
	XOLog(sqlstr, isbn)
	b := Book{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, isbn).Scan(&b.BookID, &b.AuthorID, &b.Isbn, &b.Booktype, &b.Title, &b.Year, &b.Available, &b.Tags)
	if err != nil {
		return nil, err
	}

	return &b, nil
}

// BookByBookID retrieves a row from 'public.books' as a Book.
//
// Generated from index 'books_pkey'.
func BookByBookID(db XODB, bookID int) (*Book, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`book_id, author_id, isbn, booktype, title, year, available, tags ` +
		`FROM public.books ` +
		`WHERE book_id = $1`

	// run query
	XOLog(sqlstr, bookID)
	b := Book{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, bookID).Scan(&b.BookID, &b.AuthorID, &b.Isbn, &b.Booktype, &b.Title, &b.Year, &b.Available, &b.Tags)
	if err != nil {
		return nil, err
	}

	return &b, nil
}

// BooksByTitle retrieves a row from 'public.books' as a Book.
//
// Generated from index 'books_title_idx'.
func BooksByTitle(db XODB, title string, year int) ([]*Book, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`book_id, author_id, isbn, booktype, title, year, available, tags ` +
		`FROM public.books ` +
		`WHERE title = $1 AND year = $2`

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
		err = q.Scan(&b.BookID, &b.AuthorID, &b.Isbn, &b.Booktype, &b.Title, &b.Year, &b.Available, &b.Tags)
		if err != nil {
			return nil, err
		}

		res = append(res, &b)
	}

	return res, nil
}

// BooksByTitleLower retrieves a row from 'public.books' as a Book.
//
// Generated from index 'books_title_lower_idx'.
func BooksByTitleLower(db XODB, title string) ([]*Book, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`book_id, author_id, isbn, booktype, title, year, available, tags ` +
		`FROM public.books ` +
		`WHERE title = $1`

	// run query
	XOLog(sqlstr, title)
	q, err := db.Query(sqlstr, title)
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
		err = q.Scan(&b.BookID, &b.AuthorID, &b.Isbn, &b.Booktype, &b.Title, &b.Year, &b.Available, &b.Tags)
		if err != nil {
			return nil, err
		}

		res = append(res, &b)
	}

	return res, nil
}
