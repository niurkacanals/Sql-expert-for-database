package mysql

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// Tag represents a row from 'django.tags'.
type Tag struct {
	TagID int64  `json:"tag_id"` // tag_id
	Tag   string `json:"tag"`    // tag
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the Tag exists in the database.
func (t *Tag) Exists() bool {
	return t._exists
}

// Deleted returns true when the Tag has been marked for deletion from
// the database.
func (t *Tag) Deleted() bool {
	return t._deleted
}

// Insert inserts the Tag to the database.
func (t *Tag) Insert(ctx context.Context, db DB) error {
	switch {
	case t._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case t._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO django.tags (` +
		`tag` +
		`) VALUES (` +
		`?` +
		`)`
	// run
	logf(sqlstr, t.Tag)
	res, err := db.ExecContext(ctx, sqlstr, t.Tag)
	if err != nil {
		return err
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	} // set primary key
	t.TagID = int64(id)
	// set exists
	t._exists = true
	return nil
}

// Update updates a Tag in the database.
func (t *Tag) Update(ctx context.Context, db DB) error {
	switch {
	case !t._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case t._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE django.tags SET ` +
		`tag = ? ` +
		`WHERE tag_id = ?`
	// run
	logf(sqlstr, t.Tag, t.TagID)
	if _, err := db.ExecContext(ctx, sqlstr, t.Tag, t.TagID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Tag to the database.
func (t *Tag) Save(ctx context.Context, db DB) error {
	if t.Exists() {
		return t.Update(ctx, db)
	}
	return t.Insert(ctx, db)
}

// Upsert performs an upsert for Tag.
func (t *Tag) Upsert(ctx context.Context, db DB) error {
	switch {
	case t._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO django.tags (` +
		`tag_id, tag` +
		`) VALUES (` +
		`?, ?` +
		`)` +
		` ON DUPLICATE KEY UPDATE ` +
		`tag = VALUES(tag)`
	// run
	logf(sqlstr, t.TagID, t.Tag)
	if _, err := db.ExecContext(ctx, sqlstr, t.TagID, t.Tag); err != nil {
		return err
	}
	// set exists
	t._exists = true
	return nil
}

// Delete deletes the Tag from the database.
func (t *Tag) Delete(ctx context.Context, db DB) error {
	switch {
	case !t._exists: // doesn't exist
		return nil
	case t._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM django.tags ` +
		`WHERE tag_id = ?`
	// run
	logf(sqlstr, t.TagID)
	if _, err := db.ExecContext(ctx, sqlstr, t.TagID); err != nil {
		return logerror(err)
	}
	// set deleted
	t._deleted = true
	return nil
}

// TagByTagID retrieves a row from 'django.tags' as a Tag.
//
// Generated from index 'tags_tag_id_pkey'.
func TagByTagID(ctx context.Context, db DB, tagID int64) (*Tag, error) {
	// query
	const sqlstr = `SELECT ` +
		`tag_id, tag ` +
		`FROM django.tags ` +
		`WHERE tag_id = ?`
	// run
	logf(sqlstr, tagID)
	t := Tag{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, tagID).Scan(&t.TagID, &t.Tag); err != nil {
		return nil, logerror(err)
	}
	return &t, nil
}
