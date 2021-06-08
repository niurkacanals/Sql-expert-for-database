package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// Table represents table info.
type Table struct {
	Type      string `json:"type"`       // type
	TableName string `json:"table_name"` // table_name
	ManualPk  bool   `json:"manual_pk"`  // manual_pk
}

// PostgresTables runs a custom query, returning results as Table.
func PostgresTables(ctx context.Context, db DB, schema, kind string) ([]*Table, error) {
	// query
	const sqlstr = `SELECT ` +
		`c.relkind, ` + // ::varchar AS type
		`c.relname, ` + // ::varchar AS table_name
		`false ` + // ::boolean AS manual_pk
		`FROM pg_class c ` +
		`JOIN ONLY pg_namespace n ON n.oid = c.relnamespace ` +
		`WHERE n.nspname = $1 AND c.relkind = $2`
	// run
	logf(sqlstr, schema, kind)
	rows, err := db.QueryContext(ctx, sqlstr, schema, kind)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*Table
	for rows.Next() {
		var t Table
		// scan
		if err := rows.Scan(&t.Type, &t.TableName, &t.ManualPk); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &t)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// MysqlTables runs a custom query, returning results as Table.
func MysqlTables(ctx context.Context, db DB, schema, kind string) ([]*Table, error) {
	// query
	const sqlstr = `SELECT ` +
		`table_name ` +
		`FROM information_schema.tables ` +
		`WHERE table_schema = ? AND table_type = ?`
	// run
	logf(sqlstr, schema, kind)
	rows, err := db.QueryContext(ctx, sqlstr, schema, kind)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*Table
	for rows.Next() {
		var t Table
		// scan
		if err := rows.Scan(&t.TableName); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &t)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// Sqlite3Tables runs a custom query, returning results as Table.
func Sqlite3Tables(ctx context.Context, db DB, kind string) ([]*Table, error) {
	// query
	const sqlstr = `SELECT ` +
		`tbl_name AS table_name ` +
		`FROM sqlite_master ` +
		`WHERE tbl_name NOT LIKE 'sqlite_%' AND type = ?`
	// run
	logf(sqlstr, kind)
	rows, err := db.QueryContext(ctx, sqlstr, kind)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*Table
	for rows.Next() {
		var t Table
		// scan
		if err := rows.Scan(&t.TableName); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &t)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// SqlserverTables runs a custom query, returning results as Table.
func SqlserverTables(ctx context.Context, db DB, schema, kind string) ([]*Table, error) {
	// query
	const sqlstr = `SELECT ` +
		`xtype AS type, ` +
		`name AS table_name ` +
		`FROM sysobjects ` +
		`WHERE SCHEMA_NAME(uid) = @p1 AND xtype = @p2`
	// run
	logf(sqlstr, schema, kind)
	rows, err := db.QueryContext(ctx, sqlstr, schema, kind)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*Table
	for rows.Next() {
		var t Table
		// scan
		if err := rows.Scan(&t.Type, &t.TableName); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &t)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// OracleTables runs a custom query, returning results as Table.
func OracleTables(ctx context.Context, db DB, schema, kind string) ([]*Table, error) {
	// query
	const sqlstr = `SELECT ` +
		`LOWER(object_name) AS table_name ` +
		`FROM all_objects ` +
		`WHERE ` +
		`owner = UPPER(:1) ` +
		`AND object_type = UPPER(:2) ` +
		`AND object_name NOT LIKE '%$%' ` +
		`AND object_name NOT LIKE 'LOGMNR%_%' ` +
		`AND object_name NOT LIKE 'REDO_%' ` +
		`AND object_name NOT LIKE 'SCHEDULER_%_TBL' ` +
		`AND object_name NOT LIKE 'SQLPLUS_%'`
	// run
	logf(sqlstr, schema, kind)
	rows, err := db.QueryContext(ctx, sqlstr, schema, kind)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*Table
	for rows.Next() {
		var t Table
		// scan
		if err := rows.Scan(&t.TableName); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &t)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}
