package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// IndexColumn represents index column info.
type IndexColumn struct {
	SeqNo      int    `json:"seq_no"`      // seq_no
	Cid        int    `json:"cid"`         // cid
	ColumnName string `json:"column_name"` // column_name
}

// PostgresIndexColumns runs a custom query, returning results as IndexColumn.
func PostgresIndexColumns(ctx context.Context, db DB, schema, index string) ([]*IndexColumn, error) {
	// query
	const sqlstr = `SELECT ` +
		`(row_number() over()), ` + // ::integer AS seq_no
		`a.attnum, ` + // ::integer AS cid
		`a.attname ` + // ::varchar AS column_name
		`FROM pg_index i ` +
		`JOIN ONLY pg_class c ON c.oid = i.indrelid ` +
		`JOIN ONLY pg_namespace n ON n.oid = c.relnamespace ` +
		`JOIN ONLY pg_class ic ON ic.oid = i.indexrelid ` +
		`LEFT JOIN pg_attribute a ON i.indrelid = a.attrelid ` +
		`AND a.attnum = ANY(i.indkey) ` +
		`AND a.attisdropped = false ` +
		`WHERE i.indkey <> '0' ` +
		`AND n.nspname = $1 ` +
		`AND ic.relname = $2`
	// run
	logf(sqlstr, schema, index)
	rows, err := db.QueryContext(ctx, sqlstr, schema, index)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*IndexColumn
	for rows.Next() {
		var ic IndexColumn
		// scan
		if err := rows.Scan(&ic.SeqNo, &ic.Cid, &ic.ColumnName); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ic)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// MysqlIndexColumns runs a custom query, returning results as IndexColumn.
func MysqlIndexColumns(ctx context.Context, db DB, schema, table, index string) ([]*IndexColumn, error) {
	// query
	const sqlstr = `SELECT ` +
		`seq_in_index AS seq_no, ` +
		`column_name ` +
		`FROM information_schema.statistics ` +
		`WHERE index_schema = ? ` +
		`AND table_name = ? ` +
		`AND index_name = ? ` +
		`ORDER BY seq_in_index`
	// run
	logf(sqlstr, schema, table, index)
	rows, err := db.QueryContext(ctx, sqlstr, schema, table, index)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*IndexColumn
	for rows.Next() {
		var ic IndexColumn
		// scan
		if err := rows.Scan(&ic.SeqNo, &ic.ColumnName); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ic)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// Sqlite3IndexColumns runs a custom query, returning results as IndexColumn.
func Sqlite3IndexColumns(ctx context.Context, db DB, schema, table, index string) ([]*IndexColumn, error) {
	// query
	sqlstr := `/* ` + schema + ` ` + table + ` */ ` +
		`SELECT ` +
		`seqno AS seq_no, ` +
		`cid, ` +
		`name AS column_name ` +
		`FROM pragma_index_info($1)`
	// run
	logf(sqlstr, index)
	rows, err := db.QueryContext(ctx, sqlstr, index)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*IndexColumn
	for rows.Next() {
		var ic IndexColumn
		// scan
		if err := rows.Scan(&ic.SeqNo, &ic.Cid, &ic.ColumnName); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ic)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// SqlserverIndexColumns runs a custom query, returning results as IndexColumn.
func SqlserverIndexColumns(ctx context.Context, db DB, schema, table, index string) ([]*IndexColumn, error) {
	// query
	const sqlstr = `SELECT ` +
		`k.keyno AS seq_no, ` +
		`k.colid AS cid, ` +
		`c.name AS column_name ` +
		`FROM sysindexes i ` +
		`INNER JOIN sysobjects o ON i.id = o.id ` +
		`INNER JOIN sysindexkeys k ON k.id = o.id ` +
		`AND k.indid = i.indid ` +
		`INNER JOIN syscolumns c ON c.id = o.id ` +
		`AND c.colid = k.colid ` +
		`WHERE o.type = 'U' ` +
		`AND SCHEMA_NAME(o.uid) = @p1 ` +
		`AND o.name = @p2 ` +
		`AND i.name = @p3 ` +
		`ORDER BY k.keyno`
	// run
	logf(sqlstr, schema, table, index)
	rows, err := db.QueryContext(ctx, sqlstr, schema, table, index)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*IndexColumn
	for rows.Next() {
		var ic IndexColumn
		// scan
		if err := rows.Scan(&ic.SeqNo, &ic.Cid, &ic.ColumnName); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ic)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// OracleIndexColumns runs a custom query, returning results as IndexColumn.
func OracleIndexColumns(ctx context.Context, db DB, schema, table, index string) ([]*IndexColumn, error) {
	// query
	const sqlstr = `SELECT ` +
		`column_position AS seq_no, ` +
		`LOWER(column_name) AS column_name ` +
		`FROM all_ind_columns ` +
		`WHERE index_owner = UPPER(:1) ` +
		`AND table_name = UPPER(:2) ` +
		`AND index_name = UPPER(:3) ` +
		`ORDER BY column_position`
	// run
	logf(sqlstr, schema, table, index)
	rows, err := db.QueryContext(ctx, sqlstr, schema, table, index)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*IndexColumn
	for rows.Next() {
		var ic IndexColumn
		// scan
		if err := rows.Scan(&ic.SeqNo, &ic.ColumnName); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ic)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}
