package postgres

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// AUniqueIndexComposite represents a row from 'public.a_unique_index_composite'.
type AUniqueIndexComposite struct {
	AKey1 sql.NullInt64 `json:"a_key1"` // a_key1
	AKey2 sql.NullInt64 `json:"a_key2"` // a_key2

}

// AUniqueIndexCompositeByAKey1AKey2 retrieves a row from 'public.a_unique_index_composite' as a AUniqueIndexComposite.
//
// Generated from index 'a_unique_index_composite_a_key1_a_key2_key'.
func AUniqueIndexCompositeByAKey1AKey2(ctx context.Context, db DB, aKey1, aKey2 sql.NullInt64) (*AUniqueIndexComposite, error) {
	// query
	const sqlstr = `SELECT ` +
		`a_key1, a_key2 ` +
		`FROM public.a_unique_index_composite ` +
		`WHERE a_key1 = $1 AND a_key2 = $2`
	// run
	logf(sqlstr, aKey1, aKey2)
	auic := AUniqueIndexComposite{}
	if err := db.QueryRowContext(ctx, sqlstr, aKey1, aKey2).Scan(&auic.AKey1, &auic.AKey2); err != nil {
		return nil, logerror(err)
	}
	return &auic, nil
}
