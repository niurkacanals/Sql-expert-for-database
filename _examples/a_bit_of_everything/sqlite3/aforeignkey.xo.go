package sqlite3

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// AForeignKey represents a row from 'a_foreign_key'.
type AForeignKey struct {
	AKey sql.NullInt64 `json:"a_key"` // a_key
}

// APrimary returns the APrimary associated with the AForeignKey's AKey (a_key).
//
// Generated from foreign key 'a_foreign_key_a_key_fkey'.
func (afk *AForeignKey) APrimary(ctx context.Context, db DB) (*APrimary, error) {
	return APrimaryByAKey(ctx, db, int(afk.AKey.Int64))
}
