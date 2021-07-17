package sqlserver

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// BooksTag represents a row from 'django.books_tags'.
type BooksTag struct {
	ID     int   `json:"id"`      // id
	BookID int64 `json:"book_id"` // book_id
	TagID  int64 `json:"tag_id"`  // tag_id
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the BooksTag exists in the database.
func (bt *BooksTag) Exists() bool {
	return bt._exists
}

// Deleted returns true when the BooksTag has been marked for deletion from
// the database.
func (bt *BooksTag) Deleted() bool {
	return bt._deleted
}

// Insert inserts the BooksTag to the database.
func (bt *BooksTag) Insert(ctx context.Context, db DB) error {
	switch {
	case bt._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case bt._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO django.books_tags (` +
		`book_id, tag_id` +
		`) VALUES (` +
		`@p1, @p2` +
		`); SELECT ID = CONVERT(BIGINT, SCOPE_IDENTITY())`
	// run
	logf(sqlstr, bt.BookID, bt.TagID)
	rows, err := db.QueryContext(ctx, sqlstr, bt.BookID, bt.TagID)
	if err != nil {
		return logerror(err)
	}
	defer rows.Close()
	// retrieve id
	var id int64
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return logerror(err)
		}
	}
	if err := rows.Err(); err != nil {
		return logerror(err)
	} // set primary key
	bt.ID = int(id)
	// set exists
	bt._exists = true
	return nil
}

// Update updates a BooksTag in the database.
func (bt *BooksTag) Update(ctx context.Context, db DB) error {
	switch {
	case !bt._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case bt._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE django.books_tags SET ` +
		`book_id = @p1, tag_id = @p2 ` +
		`WHERE id = @p3`
	// run
	logf(sqlstr, bt.BookID, bt.TagID, bt.ID)
	if _, err := db.ExecContext(ctx, sqlstr, bt.BookID, bt.TagID, bt.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the BooksTag to the database.
func (bt *BooksTag) Save(ctx context.Context, db DB) error {
	if bt.Exists() {
		return bt.Update(ctx, db)
	}
	return bt.Insert(ctx, db)
}

// Upsert performs an upsert for BooksTag.
func (bt *BooksTag) Upsert(ctx context.Context, db DB) error {
	switch {
	case bt._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `MERGE django.books_tags AS t ` +
		`USING (` +
		`SELECT @p1 id, @p2 book_id, @p3 tag_id ` +
		`) AS s ` +
		`ON s.id = t.id ` +
		`WHEN MATCHED THEN ` +
		`UPDATE SET ` +
		`t.book_id = s.book_id, t.tag_id = s.tag_id ` +
		`WHEN NOT MATCHED THEN ` +
		`INSERT (` +
		`book_id, tag_id` +
		`) VALUES (` +
		`s.book_id, s.tag_id` +
		`);`
	// run
	logf(sqlstr, bt.ID, bt.BookID, bt.TagID)
	if _, err := db.ExecContext(ctx, sqlstr, bt.ID, bt.BookID, bt.TagID); err != nil {
		return err
	}
	// set exists
	bt._exists = true
	return nil
}

// Delete deletes the BooksTag from the database.
func (bt *BooksTag) Delete(ctx context.Context, db DB) error {
	switch {
	case !bt._exists: // doesn't exist
		return nil
	case bt._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM django.books_tags ` +
		`WHERE id = @p1`
	// run
	logf(sqlstr, bt.ID)
	if _, err := db.ExecContext(ctx, sqlstr, bt.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	bt._deleted = true
	return nil
}

// BooksTagByID retrieves a row from 'django.books_tags' as a BooksTag.
//
// Generated from index 'PK__books_ta__3213E83FEF21E54E'.
func BooksTagByID(ctx context.Context, db DB, id int) (*BooksTag, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, book_id, tag_id ` +
		`FROM django.books_tags ` +
		`WHERE id = @p1`
	// run
	logf(sqlstr, id)
	bt := BooksTag{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&bt.ID, &bt.BookID, &bt.TagID); err != nil {
		return nil, logerror(err)
	}
	return &bt, nil
}

// BooksTagsByBookID retrieves a row from 'django.books_tags' as a BooksTag.
//
// Generated from index 'books_tags_book_id_73d7d8e8'.
func BooksTagsByBookID(ctx context.Context, db DB, bookID int64) ([]*BooksTag, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, book_id, tag_id ` +
		`FROM django.books_tags ` +
		`WHERE book_id = @p1`
	// run
	logf(sqlstr, bookID)
	rows, err := db.QueryContext(ctx, sqlstr, bookID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*BooksTag
	for rows.Next() {
		bt := BooksTag{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&bt.ID, &bt.BookID, &bt.TagID); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &bt)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// BooksTagByBookIDTagID retrieves a row from 'django.books_tags' as a BooksTag.
//
// Generated from index 'books_tags_book_id_tag_id_29db9e39_uniq'.
func BooksTagByBookIDTagID(ctx context.Context, db DB, bookID, tagID int64) (*BooksTag, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, book_id, tag_id ` +
		`FROM django.books_tags ` +
		`WHERE book_id = @p1 AND tag_id = @p2`
	// run
	logf(sqlstr, bookID, tagID)
	bt := BooksTag{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, bookID, tagID).Scan(&bt.ID, &bt.BookID, &bt.TagID); err != nil {
		return nil, logerror(err)
	}
	return &bt, nil
}

// BooksTagsByTagID retrieves a row from 'django.books_tags' as a BooksTag.
//
// Generated from index 'books_tags_tag_id_8d70b40a'.
func BooksTagsByTagID(ctx context.Context, db DB, tagID int64) ([]*BooksTag, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, book_id, tag_id ` +
		`FROM django.books_tags ` +
		`WHERE tag_id = @p1`
	// run
	logf(sqlstr, tagID)
	rows, err := db.QueryContext(ctx, sqlstr, tagID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*BooksTag
	for rows.Next() {
		bt := BooksTag{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&bt.ID, &bt.BookID, &bt.TagID); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &bt)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// Book returns the Book associated with the BooksTag's (BookID).
//
// Generated from foreign key 'books_tags_book_id_73d7d8e8_fk_books_book_id'.
func (bt *BooksTag) Book(ctx context.Context, db DB) (*Book, error) {
	return BookByBookID(ctx, db, bt.BookID)
}

// Tag returns the Tag associated with the BooksTag's (TagID).
//
// Generated from foreign key 'books_tags_tag_id_8d70b40a_fk_tags_tag_id'.
func (bt *BooksTag) Tag(ctx context.Context, db DB) (*Tag, error) {
	return TagByTagID(ctx, db, bt.TagID)
}
