package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// Index represents an index.
type Index struct {
	IndexName string `json:"index_name"` // index_name
	IsUnique  bool   `json:"is_unique"`  // is_unique
	IsPrimary bool   `json:"is_primary"` // is_primary
}

// PostgresTableIndexes runs a custom query, returning results as Index.
func PostgresTableIndexes(ctx context.Context, db DB, schema, table string) ([]*Index, error) {
	// query
	const sqlstr = `SELECT ` +
		`DISTINCT ic.relname, ` + // ::varchar AS index_name
		`i.indisunique, ` + // ::boolean AS is_unique
		`i.indisprimary ` + // ::boolean AS is_primary
		`FROM pg_index i ` +
		`JOIN ONLY pg_class c ON c.oid = i.indrelid ` +
		`JOIN ONLY pg_namespace n ON n.oid = c.relnamespace ` +
		`JOIN ONLY pg_class ic ON ic.oid = i.indexrelid ` +
		`WHERE i.indkey <> '0' AND n.nspname = $1 AND c.relname = $2`
	// run
	logf(sqlstr, schema, table)
	rows, err := db.QueryContext(ctx, sqlstr, schema, table)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*Index
	for rows.Next() {
		var i Index
		// scan
		if err := rows.Scan(&i.IndexName, &i.IsUnique, &i.IsPrimary); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// MysqlTableIndexes runs a custom query, returning results as Index.
func MysqlTableIndexes(ctx context.Context, db DB, schema, table string) ([]*Index, error) {
	// query
	const sqlstr = `SELECT ` +
		`DISTINCT index_name, ` +
		`NOT non_unique AS is_unique ` +
		`FROM information_schema.statistics ` +
		`WHERE index_name <> 'PRIMARY' AND index_schema = ? AND table_name = ?`
	// run
	logf(sqlstr, schema, table)
	rows, err := db.QueryContext(ctx, sqlstr, schema, table)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*Index
	for rows.Next() {
		var i Index
		// scan
		if err := rows.Scan(&i.IndexName, &i.IsUnique); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// Sqlite3TableIndexes runs a custom query, returning results as Index.
func Sqlite3TableIndexes(ctx context.Context, db DB, table string) ([]*Index, error) {
	// query
	const sqlstr = `SELECT ` +
		`name AS index_name, ` +
		`"unique" AS is_unique, ` +
		`CAST(origin = 'pk' AS boolean) AS is_primary ` +
		`FROM pragma_index_list(?)`
	// run
	logf(sqlstr, table)
	rows, err := db.QueryContext(ctx, sqlstr, table)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*Index
	for rows.Next() {
		var i Index
		// scan
		if err := rows.Scan(&i.IndexName, &i.IsUnique, &i.IsPrimary); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// SqlserverTableIndexes runs a custom query, returning results as Index.
func SqlserverTableIndexes(ctx context.Context, db DB, schema, table string) ([]*Index, error) {
	// query
	const sqlstr = `SELECT ` +
		`i.name AS index_name, ` +
		`i.is_primary_key AS is_primary, ` +
		`i.is_unique ` +
		`FROM sys.indexes i ` +
		`INNER JOIN sysobjects o ON i.object_id = o.id ` +
		`WHERE i.name IS NOT NULL AND o.type = 'U' AND SCHEMA_NAME(o.uid) = @p1 AND o.name = @p2`
	// run
	logf(sqlstr, schema, table)
	rows, err := db.QueryContext(ctx, sqlstr, schema, table)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*Index
	for rows.Next() {
		var i Index
		// scan
		if err := rows.Scan(&i.IndexName, &i.IsPrimary, &i.IsUnique); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// OracleTableIndexes runs a custom query, returning results as Index.
func OracleTableIndexes(ctx context.Context, db DB, schema, table string) ([]*Index, error) {
	// query
	const sqlstr = `SELECT ` +
		`LOWER(index_name) AS index_name, ` +
		`CASE WHEN uniqueness = 'UNIQUE' THEN '1' ELSE '0' END AS is_unique ` +
		`FROM all_indexes ` +
		`WHERE owner = UPPER(:1) AND table_name = UPPER(:2)`
	// run
	logf(sqlstr, schema, table)
	rows, err := db.QueryContext(ctx, sqlstr, schema, table)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*Index
	for rows.Next() {
		var i Index
		// scan
		if err := rows.Scan(&i.IndexName, &i.IsUnique); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}
