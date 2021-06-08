package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"context"

	"github.com/xo/xo/_examples/pgcatalog/pgtypes"
)

// PgTruetypid calls the stored procedure 'information_schema._pg_truetypid(pg_attribute, pg_type) oid' on db.
func PgTruetypid(ctx context.Context, db DB, v0 pgtypes.PgAttribute, v1 pgtypes.PgType) (pgtypes.Oid, error) {
	// call information_schema._pg_truetypid
	const sqlstr = `SELECT information_schema._pg_truetypid($1, $2)`
	// run
	var ret pgtypes.Oid
	logf(sqlstr, v0, v1)
	if err := db.QueryRowContext(ctx, sqlstr, v0, v1).Scan(&ret); err != nil {
		return pgtypes.Oid{}, logerror(err)
	}
	return ret, nil
}
