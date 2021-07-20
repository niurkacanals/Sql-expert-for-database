package sqlserver

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// Author represents a row from 'django.authors'.
type Author struct {
	AuthorID int64  `json:"author_id"` // author_id
	Name     string `json:"name"`      // name
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the Author exists in the database.
func (a *Author) Exists() bool {
	return a._exists
}

// Deleted returns true when the Author has been marked for deletion from
// the database.
func (a *Author) Deleted() bool {
	return a._deleted
}

// Insert inserts the Author to the database.
func (a *Author) Insert(ctx context.Context, db DB) error {
	switch {
	case a._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case a._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO django.authors (` +
		`name` +
		`) VALUES (` +
		`@p1` +
		`); SELECT ID = CONVERT(BIGINT, SCOPE_IDENTITY())`
	// run
	logf(sqlstr, a.Name)
	rows, err := db.QueryContext(ctx, sqlstr, a.Name)
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
	a.AuthorID = int64(id)
	// set exists
	a._exists = true
	return nil
}

// Update updates a Author in the database.
func (a *Author) Update(ctx context.Context, db DB) error {
	switch {
	case !a._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case a._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE django.authors SET ` +
		`name = @p1 ` +
		`WHERE author_id = @p2`
	// run
	logf(sqlstr, a.Name, a.AuthorID)
	if _, err := db.ExecContext(ctx, sqlstr, a.Name, a.AuthorID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Author to the database.
func (a *Author) Save(ctx context.Context, db DB) error {
	if a.Exists() {
		return a.Update(ctx, db)
	}
	return a.Insert(ctx, db)
}

// Upsert performs an upsert for Author.
func (a *Author) Upsert(ctx context.Context, db DB) error {
	switch {
	case a._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `MERGE django.authors AS t ` +
		`USING (` +
		`SELECT @p1 author_id, @p2 name ` +
		`) AS s ` +
		`ON s.author_id = t.author_id ` +
		`WHEN MATCHED THEN ` +
		`UPDATE SET ` +
		`t.name = s.name ` +
		`WHEN NOT MATCHED THEN ` +
		`INSERT (` +
		`name` +
		`) VALUES (` +
		`s.name` +
		`);`
	// run
	logf(sqlstr, a.AuthorID, a.Name)
	if _, err := db.ExecContext(ctx, sqlstr, a.AuthorID, a.Name); err != nil {
		return err
	}
	// set exists
	a._exists = true
	return nil
}

// Delete deletes the Author from the database.
func (a *Author) Delete(ctx context.Context, db DB) error {
	switch {
	case !a._exists: // doesn't exist
		return nil
	case a._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM django.authors ` +
		`WHERE author_id = @p1`
	// run
	logf(sqlstr, a.AuthorID)
	if _, err := db.ExecContext(ctx, sqlstr, a.AuthorID); err != nil {
		return logerror(err)
	}
	// set deleted
	a._deleted = true
	return nil
}

// AuthorByAuthorID retrieves a row from 'django.authors' as a Author.
//
// Generated from index 'PK__authors__86516BCF8B6E070E'.
func AuthorByAuthorID(ctx context.Context, db DB, authorID int64) (*Author, error) {
	// query
	const sqlstr = `SELECT ` +
		`author_id, name ` +
		`FROM django.authors ` +
		`WHERE author_id = @p1`
	// run
	logf(sqlstr, authorID)
	a := Author{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, authorID).Scan(&a.AuthorID, &a.Name); err != nil {
		return nil, logerror(err)
	}
	return &a, nil
}
