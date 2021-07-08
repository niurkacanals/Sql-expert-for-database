package postgres

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// ASequenceMulti represents a row from 'public.a_sequence_multi'.
type ASequenceMulti struct {
	ASeq  int            `json:"a_seq"`  // a_seq
	AText sql.NullString `json:"a_text"` // a_text
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the ASequenceMulti exists in the database.
func (asm *ASequenceMulti) Exists() bool {
	return asm._exists
}

// Deleted returns true when the ASequenceMulti has been marked for deletion from
// the database.
func (asm *ASequenceMulti) Deleted() bool {
	return asm._deleted
}

// Insert inserts the ASequenceMulti to the database.
func (asm *ASequenceMulti) Insert(ctx context.Context, db DB) error {
	switch {
	case asm._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case asm._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO public.a_sequence_multi (` +
		`a_text` +
		`) VALUES (` +
		`$1` +
		`) RETURNING a_seq`
	// run
	logf(sqlstr, asm.AText)
	if err := db.QueryRowContext(ctx, sqlstr, asm.AText).Scan(&asm.ASeq); err != nil {
		return logerror(err)
	}
	// set exists
	asm._exists = true
	return nil
}

// Update updates a ASequenceMulti in the database.
func (asm *ASequenceMulti) Update(ctx context.Context, db DB) error {
	switch {
	case !asm._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case asm._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.a_sequence_multi SET ` +
		`a_text = $1 ` +
		`WHERE a_seq = $2`
	// run
	logf(sqlstr, asm.AText, asm.ASeq)
	if _, err := db.ExecContext(ctx, sqlstr, asm.AText, asm.ASeq); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the ASequenceMulti to the database.
func (asm *ASequenceMulti) Save(ctx context.Context, db DB) error {
	if asm.Exists() {
		return asm.Update(ctx, db)
	}
	return asm.Insert(ctx, db)
}

// Upsert performs an upsert for ASequenceMulti.
func (asm *ASequenceMulti) Upsert(ctx context.Context, db DB) error {
	switch {
	case asm._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.a_sequence_multi (` +
		`a_text` +
		`) VALUES (` +
		`$1` +
		`)` +
		` ON CONFLICT (a_seq) DO ` +
		`UPDATE SET ` +
		`a_text = EXCLUDED.a_text `
	// run
	logf(sqlstr, asm.ASeq, asm.AText)
	if _, err := db.ExecContext(ctx, sqlstr, asm.ASeq, asm.AText); err != nil {
		return err
	}
	// set exists
	asm._exists = true
	return nil
}

// Delete deletes the ASequenceMulti from the database.
func (asm *ASequenceMulti) Delete(ctx context.Context, db DB) error {
	switch {
	case !asm._exists: // doesn't exist
		return nil
	case asm._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.a_sequence_multi ` +
		`WHERE a_seq = $1`
	// run
	logf(sqlstr, asm.ASeq)
	if _, err := db.ExecContext(ctx, sqlstr, asm.ASeq); err != nil {
		return logerror(err)
	}
	// set deleted
	asm._deleted = true
	return nil
}

// ASequenceMultiByASeq retrieves a row from 'public.a_sequence_multi' as a ASequenceMulti.
//
// Generated from index 'a_sequence_multi_pkey'.
func ASequenceMultiByASeq(ctx context.Context, db DB, aSeq int) (*ASequenceMulti, error) {
	// query
	const sqlstr = `SELECT ` +
		`a_seq, a_text ` +
		`FROM public.a_sequence_multi ` +
		`WHERE a_seq = $1`
	// run
	logf(sqlstr, aSeq)
	asm := ASequenceMulti{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, aSeq).Scan(&asm.ASeq, &asm.AText); err != nil {
		return nil, logerror(err)
	}
	return &asm, nil
}
