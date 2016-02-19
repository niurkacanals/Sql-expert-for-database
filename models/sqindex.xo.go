// Package models contains the types for schema ''.
package models

// GENERATED BY XO. DO NOT EDIT.

// SqIndex represents a row from .
type SqIndex struct {
	Seq       int    // seq
	IndexName string // index_name
	IsUnique  bool   // is_unique
	Origin    string // origin
	IsPartial bool   // is_partial
}

// SqIndicesByTable runs a custom query, returning results as SqIndex.
func SqIndicesByTable(db XODB, table string) ([]*SqIndex, error) {
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
	res := []*SqIndex{}
	for q.Next() {
		si := SqIndex{}

		// scan
		err = q.Scan(&si.Seq, &si.IndexName, &si.IsUnique, &si.Origin, &si.IsPartial)
		if err != nil {
			return nil, err
		}

		res = append(res, &si)
	}

	return res, nil
}
