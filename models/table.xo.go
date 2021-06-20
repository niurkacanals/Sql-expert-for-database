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
		`(CASE c.relkind ` +
		`WHEN 'r' THEN 'table' ` +
		`WHEN 'v' THEN 'view' ` +
		`END), ` + // ::varchar AS type
		`c.relname, ` + // ::varchar AS table_name
		`false ` + // ::boolean AS manual_pk
		`FROM pg_class c ` +
		`JOIN ONLY pg_namespace n ON n.oid = c.relnamespace ` +
		`WHERE n.nspname = $1 ` +
		`AND (CASE c.relkind ` +
		`WHEN 'r' THEN 'table' ` +
		`WHEN 'v' THEN 'view' ` +
		`END) = LOWER($2)`
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
		var row Table
		// scan
		if err := rows.Scan(&row.Type, &row.TableName, &row.ManualPk); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &row)
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
		`(CASE table_type ` +
		`WHEN 'BASE TABLE' THEN 'table' ` +
		`WHEN 'VIEW' THEN 'view' ` +
		`END) AS type, ` +
		`table_name ` +
		`FROM information_schema.tables ` +
		`WHERE table_schema = ? ` +
		`AND (CASE table_type ` +
		`WHEN 'BASE TABLE' THEN 'table' ` +
		`WHEN 'VIEW' THEN 'view' ` +
		`END) = LOWER(?)`
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
		var row Table
		// scan
		if err := rows.Scan(&row.Type, &row.TableName); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &row)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// Sqlite3Tables runs a custom query, returning results as Table.
func Sqlite3Tables(ctx context.Context, db DB, schema, kind string) ([]*Table, error) {
	// query
	sqlstr := `/* ` + schema + ` */ ` +
		`SELECT ` +
		`type, ` +
		`tbl_name AS table_name ` +
		`FROM sqlite_master ` +
		`WHERE tbl_name NOT LIKE 'sqlite_%' ` +
		`AND LOWER(type) = LOWER($1)`
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
		var row Table
		// scan
		if err := rows.Scan(&row.Type, &row.TableName); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &row)
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
		`(CASE xtype ` +
		`WHEN 'U' THEN 'table' ` +
		`WHEN 'V' THEN 'view' ` +
		`END) AS type, ` +
		`name AS table_name ` +
		`FROM sysobjects ` +
		`WHERE SCHEMA_NAME(uid) = @p1 ` +
		`AND (CASE xtype ` +
		`WHEN 'U' THEN 'table' ` +
		`WHEN 'V' THEN 'view' ` +
		`END) = LOWER(@p2)`
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
		var row Table
		// scan
		if err := rows.Scan(&row.Type, &row.TableName); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &row)
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
		`LOWER(object_type) AS type, ` +
		`LOWER(object_name) AS table_name ` +
		`FROM all_objects ` +
		`WHERE object_name NOT LIKE '%$%' ` +
		`AND object_name NOT LIKE 'LOGMNR%_%' ` +
		`AND object_name NOT LIKE 'REDO_%' ` +
		`AND object_name NOT LIKE 'SCHEDULER_%_TBL' ` +
		`AND object_name NOT LIKE 'SQLPLUS_%' ` +
		`AND owner = UPPER(:1) ` +
		`AND object_type = UPPER(:2)`
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
		var row Table
		// scan
		if err := rows.Scan(&row.Type, &row.TableName); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &row)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}
