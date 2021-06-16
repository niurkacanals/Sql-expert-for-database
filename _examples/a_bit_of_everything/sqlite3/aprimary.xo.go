package sqlite3

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// APrimary represents a row from 'a_primary'.
type APrimary struct {
	AKey int `json:"a_key"` // a_key
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the APrimary exists in the database.
func (ap *APrimary) Exists() bool {
	return ap._exists
}

// Deleted returns true when the APrimary has been marked for deletion from
// the database.
func (ap *APrimary) Deleted() bool {
	return ap._deleted
}

// Insert inserts the APrimary to the database.
func (ap *APrimary) Insert(ctx context.Context, db DB) error {
	switch {
	case ap._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case ap._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (basic)
	const sqlstr = `INSERT INTO a_primary (` +
		`a_key` +
		`) VALUES (` +
		`?` +
		`)`
	// run
	logf(sqlstr, ap.AKey)
	if err := db.QueryRowContext(ctx, sqlstr, ap.AKey).Scan(&ap.AKey); err != nil {
		return logerror(err)
	}
	// set exists
	ap._exists = true
	return nil
}

// ------ NOTE: Update statements omitted due to lack of fields other than primary key ------

// Delete deletes the APrimary from the database.
func (ap *APrimary) Delete(ctx context.Context, db DB) error {
	switch {
	case !ap._exists: // doesn't exist
		return nil
	case ap._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM a_primary WHERE a_key = ?`
	// run
	logf(sqlstr, ap.AKey)
	if _, err := db.ExecContext(ctx, sqlstr, ap.AKey); err != nil {
		return logerror(err)
	}
	// set deleted
	ap._deleted = true
	return nil
}

// APrimaryByAKey retrieves a row from 'a_primary' as a APrimary.
//
// Generated from index 'a_primary_a_key_pkey'.
func APrimaryByAKey(ctx context.Context, db DB, aKey int) (*APrimary, error) {
	// query
	const sqlstr = `SELECT ` +
		`a_key ` +
		`FROM a_primary ` +
		`WHERE a_key = ?`
	// run
	logf(sqlstr, aKey)
	ap := APrimary{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, aKey).Scan(&ap.AKey); err != nil {
		return nil, logerror(err)
	}
	return &ap, nil
}