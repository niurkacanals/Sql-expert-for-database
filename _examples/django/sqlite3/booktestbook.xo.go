package sqlite3

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// BooktestBook represents a row from 'booktest_book'.
type BooktestBook struct {
	BookID    int    `json:"book_id"`   // book_id
	Isbn      string `json:"isbn"`      // isbn
	BookType  int    `json:"book_type"` // book_type
	Title     string `json:"title"`     // title
	Year      int    `json:"year"`      // year
	Available Time   `json:"available"` // available
	AuthorID  int64  `json:"author_id"` // author_id
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the BooktestBook exists in the database.
func (bb *BooktestBook) Exists() bool {
	return bb._exists
}

// Deleted returns true when the BooktestBook has been marked for deletion from
// the database.
func (bb *BooktestBook) Deleted() bool {
	return bb._deleted
}

// Insert inserts the BooktestBook to the database.
func (bb *BooktestBook) Insert(ctx context.Context, db DB) error {
	switch {
	case bb._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case bb._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO booktest_book (` +
		`book_id, isbn, book_type, title, year, available, author_id` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`)`
	// run
	logf(sqlstr, bb.Isbn, bb.BookType, bb.Title, bb.Year, bb.Available, bb.AuthorID)
	res, err := db.ExecContext(ctx, sqlstr, bb.BookID, bb.Isbn, bb.BookType, bb.Title, bb.Year, bb.Available, bb.AuthorID)
	if err != nil {
		return err
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	} // set primary key
	bb.BookID = int(id)
	// set exists
	bb._exists = true
	return nil
}

// Update updates a BooktestBook in the database.
func (bb *BooktestBook) Update(ctx context.Context, db DB) error {
	switch {
	case !bb._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case bb._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE booktest_book SET ` +
		`isbn = $1, book_type = $2, title = $3, year = $4, available = $5, author_id = $6 ` +
		`WHERE book_id = $7`
	// run
	logf(sqlstr, bb.Isbn, bb.BookType, bb.Title, bb.Year, bb.Available, bb.AuthorID, bb.BookID)
	if _, err := db.ExecContext(ctx, sqlstr, bb.Isbn, bb.BookType, bb.Title, bb.Year, bb.Available, bb.AuthorID, bb.BookID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the BooktestBook to the database.
func (bb *BooktestBook) Save(ctx context.Context, db DB) error {
	if bb.Exists() {
		return bb.Update(ctx, db)
	}
	return bb.Insert(ctx, db)
}

// Upsert performs an upsert for BooktestBook.
func (bb *BooktestBook) Upsert(ctx context.Context, db DB) error {
	switch {
	case bb._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO booktest_book (` +
		`book_id, isbn, book_type, title, year, available, author_id` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`)` +
		` ON CONFLICT (book_id) DO ` +
		`UPDATE SET ` +
		`isbn = EXCLUDED.isbn, book_type = EXCLUDED.book_type, title = EXCLUDED.title, year = EXCLUDED.year, available = EXCLUDED.available, author_id = EXCLUDED.author_id `
	// run
	logf(sqlstr, bb.BookID, bb.Isbn, bb.BookType, bb.Title, bb.Year, bb.Available, bb.AuthorID)
	if _, err := db.ExecContext(ctx, sqlstr, bb.BookID, bb.Isbn, bb.BookType, bb.Title, bb.Year, bb.Available, bb.AuthorID); err != nil {
		return err
	}
	// set exists
	bb._exists = true
	return nil
}

// Delete deletes the BooktestBook from the database.
func (bb *BooktestBook) Delete(ctx context.Context, db DB) error {
	switch {
	case !bb._exists: // doesn't exist
		return nil
	case bb._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM booktest_book ` +
		`WHERE book_id = $1`
	// run
	logf(sqlstr, bb.BookID)
	if _, err := db.ExecContext(ctx, sqlstr, bb.BookID); err != nil {
		return logerror(err)
	}
	// set deleted
	bb._deleted = true
	return nil
}

// BooktestBookByAuthorID retrieves a row from 'booktest_book' as a BooktestBook.
//
// Generated from index 'booktest_book_author_id_d11194af'.
func BooktestBookByAuthorID(ctx context.Context, db DB, authorID int64) ([]*BooktestBook, error) {
	// query
	const sqlstr = `SELECT ` +
		`book_id, isbn, book_type, title, year, available, author_id ` +
		`FROM booktest_book ` +
		`WHERE author_id = $1`
	// run
	logf(sqlstr, authorID)
	rows, err := db.QueryContext(ctx, sqlstr, authorID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*BooktestBook
	for rows.Next() {
		bb := BooktestBook{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&bb.BookID, &bb.Isbn, &bb.BookType, &bb.Title, &bb.Year, &bb.Available, &bb.AuthorID); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &bb)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// BooktestBookByBookID retrieves a row from 'booktest_book' as a BooktestBook.
//
// Generated from index 'booktest_book_book_id_pkey'.
func BooktestBookByBookID(ctx context.Context, db DB, bookID int) (*BooktestBook, error) {
	// query
	const sqlstr = `SELECT ` +
		`book_id, isbn, book_type, title, year, available, author_id ` +
		`FROM booktest_book ` +
		`WHERE book_id = $1`
	// run
	logf(sqlstr, bookID)
	bb := BooktestBook{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, bookID).Scan(&bb.BookID, &bb.Isbn, &bb.BookType, &bb.Title, &bb.Year, &bb.Available, &bb.AuthorID); err != nil {
		return nil, logerror(err)
	}
	return &bb, nil
}

// BooktestAuthor returns the BooktestAuthor associated with the BooktestBook's (AuthorID).
//
// Generated from foreign key 'booktest_book_author_id_fkey'.
func (bb *BooktestBook) BooktestAuthor(ctx context.Context, db DB) (*BooktestAuthor, error) {
	return BooktestAuthorByAuthorID(ctx, db, int(bb.AuthorID))
}
