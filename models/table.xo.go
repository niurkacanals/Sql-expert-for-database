package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// Table is a table.
type Table struct {
	Type      string `json:"type"`       // type
	TableName string `json:"table_name"` // table_name
	ManualPk  bool   `json:"manual_pk"`  // manual_pk
	ViewDef   string `json:"view_def"`   // view_def
}

// PostgresTables runs a custom query, returning results as Table.
func PostgresTables(ctx context.Context, db DB, schema, typ string) ([]*Table, error) {
	// query
	const sqlstr = `SELECT ` +
		`(CASE c.relkind ` +
		`WHEN 'r' THEN 'table' ` +
		`WHEN 'v' THEN 'view' ` +
		`END), ` + // ::varchar AS type
		`c.relname, ` + // ::varchar AS table_name
		`false, ` + // ::boolean AS manual_pk
		`CASE c.relkind ` +
		`WHEN 'r' THEN '' ` +
		`WHEN 'v' THEN v.definition ` +
		`END AS view_def ` +
		`FROM pg_class c ` +
		`JOIN ONLY pg_namespace n ON n.oid = c.relnamespace ` +
		`LEFT JOIN pg_views v ON n.nspname = v.schemaname ` +
		`AND v.viewname = c.relname ` +
		`WHERE n.nspname = $1 ` +
		`AND (CASE c.relkind ` +
		`WHEN 'r' THEN 'table' ` +
		`WHEN 'v' THEN 'view' ` +
		`END) = LOWER($2)`
	// run
	logf(sqlstr, schema, typ)
	rows, err := db.QueryContext(ctx, sqlstr, schema, typ)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*Table
	for rows.Next() {
		var t Table
		// scan
		if err := rows.Scan(&t.Type, &t.TableName, &t.ManualPk, &t.ViewDef); err != nil {
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
func MysqlTables(ctx context.Context, db DB, schema, typ string) ([]*Table, error) {
	// query
	const sqlstr = `SELECT ` +
		`(CASE t.table_type ` +
		`WHEN 'BASE TABLE' THEN 'table' ` +
		`WHEN 'VIEW' THEN 'view' ` +
		`END) AS type, ` +
		`t.table_name, ` +
		`CASE t.table_type ` +
		`WHEN 'BASE TABLE' THEN '' ` +
		`WHEN 'VIEW' then v.view_definition ` +
		`END AS view_def ` +
		`FROM information_schema.tables t ` +
		`LEFT JOIN information_schema.views v ON t.table_schema = v.table_schema ` +
		`AND t.table_name = v.table_name ` +
		`WHERE t.table_schema = ? ` +
		`AND (CASE t.table_type ` +
		`WHEN 'BASE TABLE' THEN 'table' ` +
		`WHEN 'VIEW' THEN 'view' ` +
		`END) = LOWER(?)`
	// run
	logf(sqlstr, schema, typ)
	rows, err := db.QueryContext(ctx, sqlstr, schema, typ)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*Table
	for rows.Next() {
		var t Table
		// scan
		if err := rows.Scan(&t.Type, &t.TableName, &t.ViewDef); err != nil {
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
func Sqlite3Tables(ctx context.Context, db DB, schema, typ string) ([]*Table, error) {
	// query
	sqlstr := `/* ` + schema + ` */ ` +
		`SELECT ` +
		`type, ` +
		`tbl_name AS table_name, ` +
		`CASE LOWER(type) ` +
		`WHEN 'table' THEN '' ` +
		`WHEN 'view' THEN sql ` +
		`END AS view_def ` +
		`FROM sqlite_master ` +
		`WHERE tbl_name NOT LIKE 'sqlite_%' ` +
		`AND LOWER(type) = LOWER($1)`
	// run
	logf(sqlstr, typ)
	rows, err := db.QueryContext(ctx, sqlstr, typ)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*Table
	for rows.Next() {
		var t Table
		// scan
		if err := rows.Scan(&t.Type, &t.TableName, &t.ViewDef); err != nil {
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
func SqlserverTables(ctx context.Context, db DB, schema, typ string) ([]*Table, error) {
	// query
	const sqlstr = `SELECT ` +
		`(CASE xtype ` +
		`WHEN 'U' THEN 'table' ` +
		`WHEN 'V' THEN 'view' ` +
		`END) AS type, ` +
		`name AS table_name, ` +
		`CASE xtype ` +
		`WHEN 'U' THEN '' ` +
		`WHEN 'V' THEN OBJECT_DEFINITION(id) ` +
		`END AS view_def ` +
		`FROM sysobjects ` +
		`WHERE SCHEMA_NAME(uid) = @p1 ` +
		`AND (CASE xtype ` +
		`WHEN 'U' THEN 'table' ` +
		`WHEN 'V' THEN 'view' ` +
		`END) = LOWER(@p2)`
	// run
	logf(sqlstr, schema, typ)
	rows, err := db.QueryContext(ctx, sqlstr, schema, typ)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*Table
	for rows.Next() {
		var t Table
		// scan
		if err := rows.Scan(&t.Type, &t.TableName, &t.ViewDef); err != nil {
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
func OracleTables(ctx context.Context, db DB, schema, typ string) ([]*Table, error) {
	// query
	const sqlstr = `SELECT ` +
		`LOWER(o.object_type) AS type, ` +
		`LOWER(o.object_name) AS table_name, ` +
		`CASE o.object_type ` +
		`WHEN 'TABLE' THEN ' ' ` +
		`WHEN 'VIEW' THEN v.text_vc ` +
		`END AS view_def ` +
		`FROM all_objects o ` +
		`LEFT JOIN all_views v ON o.owner = v.owner ` +
		`AND o.object_name = v.view_name ` +
		`WHERE o.object_name NOT LIKE '%$%' ` +
		`AND o.object_name NOT LIKE 'LOGMNR%_%' ` +
		`AND o.object_name NOT LIKE 'REDO_%' ` +
		`AND o.object_name NOT LIKE 'SCHEDULER_%_TBL' ` +
		`AND o.object_name NOT LIKE 'SQLPLUS_%' ` +
		`AND o.owner = UPPER(:1) ` +
		`AND o.object_type = UPPER(:2)`
	// run
	logf(sqlstr, schema, typ)
	rows, err := db.QueryContext(ctx, sqlstr, schema, typ)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*Table
	for rows.Next() {
		var t Table
		// scan
		if err := rows.Scan(&t.Type, &t.TableName, &t.ViewDef); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &t)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}
