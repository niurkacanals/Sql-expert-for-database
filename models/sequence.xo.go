package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// Sequence represents a table that references a sequence.
type Sequence struct {
	TableName string `json:"table_name"` // table_name
}

// PostgresTableSequences runs a custom query, returning results as Sequence.
func PostgresTableSequences(ctx context.Context, db DB, schema string) ([]*Sequence, error) {
	// query
	const sqlstr = `SELECT ` +
		`t.relname ` + // ::varchar AS table_name
		`FROM pg_class s ` +
		`JOIN pg_depend d ON d.objid = s.oid ` +
		`JOIN pg_class t ON d.objid = s.oid AND d.refobjid = t.oid ` +
		`JOIN pg_namespace n ON n.oid = s.relnamespace ` +
		`WHERE n.nspname = $1 AND s.relkind = 'S'`
	// run
	logf(sqlstr, schema)
	rows, err := db.QueryContext(ctx, sqlstr, schema)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*Sequence
	for rows.Next() {
		var s Sequence
		// scan
		if err := rows.Scan(&s.TableName); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &s)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// MysqlTableSequences runs a custom query, returning results as Sequence.
func MysqlTableSequences(ctx context.Context, db DB, schema string) ([]*Sequence, error) {
	// query
	const sqlstr = `SELECT ` +
		`table_name ` +
		`FROM information_schema.tables ` +
		`WHERE auto_increment IS NOT NULL AND table_schema = ?`
	// run
	logf(sqlstr, schema)
	rows, err := db.QueryContext(ctx, sqlstr, schema)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*Sequence
	for rows.Next() {
		var s Sequence
		// scan
		if err := rows.Scan(&s.TableName); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &s)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// Sqlite3TableSequences runs a custom query, returning results as Sequence.
func Sqlite3TableSequences(ctx context.Context, db DB, schema string) ([]*Sequence, error) {
	// query
	sqlstr := `/* ` + schema + ` */ ` +
		`SELECT ` +
		`name AS table_name ` +
		`FROM sqlite_master ` +
		`WHERE type='table' AND LOWER(sql) LIKE '%autoincrement%' ` +
		`ORDER BY name`
	// run
	logf(sqlstr)
	rows, err := db.QueryContext(ctx, sqlstr)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*Sequence
	for rows.Next() {
		var s Sequence
		// scan
		if err := rows.Scan(&s.TableName); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &s)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// SqlserverTableSequences runs a custom query, returning results as Sequence.
func SqlserverTableSequences(ctx context.Context, db DB, schema string) ([]*Sequence, error) {
	// query
	const sqlstr = `SELECT ` +
		`o.name AS table_name ` +
		`FROM sys.objects o ` +
		`INNER JOIN sys.columns c ON o.object_id = c.object_id ` +
		`WHERE c.is_identity = 1 AND SCHEMA_NAME(o.schema_id) = @p1 AND o.type = 'U'`
	// run
	logf(sqlstr, schema)
	rows, err := db.QueryContext(ctx, sqlstr, schema)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*Sequence
	for rows.Next() {
		var s Sequence
		// scan
		if err := rows.Scan(&s.TableName); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &s)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// OracleTableSequences runs a custom query, returning results as Sequence.
func OracleTableSequences(ctx context.Context, db DB, schema string) ([]*Sequence, error) {
	// query
	const sqlstr = `SELECT ` +
		`LOWER(c.table_name) AS table_name ` +
		`FROM all_tab_columns c ` +
		`WHERE c.identity_column='YES' AND c.owner = UPPER(:1)`
	// run
	logf(sqlstr, schema)
	rows, err := db.QueryContext(ctx, sqlstr, schema)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*Sequence
	for rows.Next() {
		var s Sequence
		// scan
		if err := rows.Scan(&s.TableName); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &s)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}
