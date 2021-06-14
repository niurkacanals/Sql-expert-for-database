package oracle

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// AForeignKeyComposite represents a row from 'a_bit_of_everything.a_foreign_key_composite'.
type AForeignKeyComposite struct {
	AKey1 sql.NullInt64 `json:"a_key1"` // a_key1
	AKey2 sql.NullInt64 `json:"a_key2"` // a_key2
}

// APrimaryComposite returns the APrimaryComposite associated with the AForeignKeyComposite's AKey2 (a_key2).
//
// Generated from foreign key 'a_foreign_key_composite_fkey'.
func (afkc *AForeignKeyComposite) APrimaryComposite(ctx context.Context, db DB) (*APrimaryComposite, error) {
	return APrimaryCompositeByAKey2(ctx, db, int(afkc.AKey2.Int64))
}
