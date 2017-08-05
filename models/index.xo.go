// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

// Index represents an index.
type Index struct {
	IndexName string // index_name
	IsUnique  bool   // is_unique
	IsPrimary bool   // is_primary
	SeqNo     int    // seq_no
	Origin    string // origin
	IsPartial bool   // is_partial
}

// PgTableIndexes runs a custom query, returning results as Index.
func PgTableIndexes(db XODB, schema string, table string) ([]*Index, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`DISTINCT ic.relname, ` + // ::varchar AS index_name
		`i.indisunique, ` + // ::boolean AS is_unique
		`i.indisprimary, ` + // ::boolean AS is_primary
		`0, ` + // ::integer AS seq_no
		`'', ` + // ::varchar AS origin
		`false ` + // ::boolean AS is_partial
		`FROM pg_index i ` +
		`JOIN ONLY pg_class c ON c.oid = i.indrelid ` +
		`JOIN ONLY pg_namespace n ON n.oid = c.relnamespace ` +
		`JOIN ONLY pg_class ic ON ic.oid = i.indexrelid ` +
		`WHERE i.indkey <> '0' AND n.nspname = $1 AND c.relname = $2`

	// run query
	XOLog(sqlstr, schema, table)
	q, err := db.Query(sqlstr, schema, table)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Index{}
	for q.Next() {
		i := Index{}

		// scan
		err = q.Scan(&i.IndexName, &i.IsUnique, &i.IsPrimary, &i.SeqNo, &i.Origin, &i.IsPartial)
		if err != nil {
			return nil, err
		}

		res = append(res, &i)
	}

	return res, nil
}

// MyTableIndexes runs a custom query, returning results as Index.
func MyTableIndexes(db XODB, schema string, table string) ([]*Index, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`DISTINCT index_name, ` +
		`NOT non_unique AS is_unique ` +
		`FROM information_schema.statistics ` +
		`WHERE index_name <> 'PRIMARY' AND index_schema = ? AND table_name = ?`

	// run query
	XOLog(sqlstr, schema, table)
	q, err := db.Query(sqlstr, schema, table)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Index{}
	for q.Next() {
		i := Index{}

		// scan
		err = q.Scan(&i.IndexName, &i.IsUnique)
		if err != nil {
			return nil, err
		}

		res = append(res, &i)
	}

	return res, nil
}

// SqTableIndexes runs a custom query, returning results as Index.
func SqTableIndexes(db XODB, table string) ([]*Index, error) {
	var err error

	// sql query
	var sqlstr = `PRAGMA index_list(` + table + `)`

	// run query
	XOLog(sqlstr)
	q, err := db.Query(sqlstr, table)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Index{}
	for q.Next() {
		i := Index{}

		// scan
		err = q.Scan(&i.SeqNo, &i.IndexName, &i.IsUnique, &i.Origin, &i.IsPartial)
		if err != nil {
			return nil, err
		}

		res = append(res, &i)
	}

	return res, nil
}

// MsTableIndexes runs a custom query, returning results as Index.
func MsTableIndexes(db XODB, schema string, table string) ([]*Index, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`i.name AS index_name, ` +
		`i.is_primary_key AS is_primary, ` +
		`i.is_unique ` +
		`FROM sys.indexes i ` +
		`INNER JOIN sysobjects o ON i.object_id = o.id ` +
		`WHERE i.name IS NOT NULL AND o.type = 'U' AND SCHEMA_NAME(o.uid) = $1 AND o.name = $2`

	// run query
	XOLog(sqlstr, schema, table)
	q, err := db.Query(sqlstr, schema, table)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Index{}
	for q.Next() {
		i := Index{}

		// scan
		err = q.Scan(&i.IndexName, &i.IsPrimary, &i.IsUnique)
		if err != nil {
			return nil, err
		}

		res = append(res, &i)
	}

	return res, nil
}

// OrTableIndexes runs a custom query, returning results as Index.
func OrTableIndexes(db XODB, schema string, table string) ([]*Index, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`LOWER(index_name) AS index_name, ` +
		`CASE WHEN uniqueness = 'UNIQUE' THEN '1' ELSE '0' END AS is_unique ` +
		`FROM all_indexes ` +
		`WHERE owner = UPPER(:1) AND table_name = UPPER(:2)`

	// run query
	XOLog(sqlstr, schema, table)
	q, err := db.Query(sqlstr, schema, table)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Index{}
	for q.Next() {
		i := Index{}

		// scan
		err = q.Scan(&i.IndexName, &i.IsUnique)
		if err != nil {
			return nil, err
		}

		res = append(res, &i)
	}

	return res, nil
}
