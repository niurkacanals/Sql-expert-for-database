package oracle

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// APrimaryComposite represents a row from 'a_bit_of_everything.a_primary_composite'.
type APrimaryComposite struct {
	AKey1 int `json:"a_key1"` // a_key1
	AKey2 int `json:"a_key2"` // a_key2
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the APrimaryComposite exists in the database.
func (apc *APrimaryComposite) Exists() bool {
	return apc._exists
}

// Deleted returns true when the APrimaryComposite has been marked for deletion from
// the database.
func (apc *APrimaryComposite) Deleted() bool {
	return apc._deleted
}

// Insert inserts the APrimaryComposite to the database.
func (apc *APrimaryComposite) Insert(ctx context.Context, db DB) error {
	switch {
	case apc._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case apc._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO a_bit_of_everything.a_primary_composite (` +
		`a_key1` +
		`) VALUES (` +
		`:1` +
		`) RETURNING a_key2 /*LASTINSERTID*/ INTO :pk`
	// run
	logf(sqlstr, apc.AKey1, nil)
	var id int64
	if _, err := db.ExecContext(ctx, sqlstr, apc.AKey1, sql.Named("pk", sql.Out{Dest: &id})); err != nil {
		return err
	} // set primary key
	apc.AKey2 = int(id)
	// set exists
	apc._exists = true
	return nil
}

// ------ NOTE: Update statements omitted due to lack of fields other than primary key ------

// Delete deletes the APrimaryComposite from the database.
func (apc *APrimaryComposite) Delete(ctx context.Context, db DB) error {
	switch {
	case !apc._exists: // doesn't exist
		return nil
	case apc._deleted: // deleted
		return nil
	}
	// delete with composite primary key
	const sqlstr = `DELETE FROM a_bit_of_everything.a_primary_composite WHERE a_key1 = :1 AND a_key2 = :2`
	// run
	logf(sqlstr, apc.AKey1, apc.AKey2)
	if _, err := db.ExecContext(ctx, sqlstr, apc.AKey1, apc.AKey2); err != nil {
		return logerror(err)
	}
	// set deleted
	apc._deleted = true
	return nil
}

// APrimaryCompositeByAKey1AKey2 retrieves a row from 'a_bit_of_everything.a_primary_composite' as a APrimaryComposite.
//
// Generated from index 'a_primary_composite_pkey'.
func APrimaryCompositeByAKey1AKey2(ctx context.Context, db DB, aKey1, aKey2 int) (*APrimaryComposite, error) {
	// query
	const sqlstr = `SELECT ` +
		`a_key1, a_key2 ` +
		`FROM a_bit_of_everything.a_primary_composite ` +
		`WHERE a_key1 = :1 AND a_key2 = :2`
	// run
	logf(sqlstr, aKey1, aKey2)
	apc := APrimaryComposite{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, aKey1, aKey2).Scan(&apc.AKey1, &apc.AKey2); err != nil {
		return nil, logerror(err)
	}
	return &apc, nil
}
