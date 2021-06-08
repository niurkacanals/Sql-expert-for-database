package oracle

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// AuthorBookResult is the result of a search.
type AuthorBookResult struct {
	AuthorID   int64  `json:"author_id"`   // author_id
	AuthorName string `json:"author_name"` // author_name
	BookID     int64  `json:"book_id"`     // book_id
	BookIsbn   string `json:"book_isbn"`   // book_isbn
	BookTitle  string `json:"book_title"`  // book_title
	BookTags   string `json:"book_tags"`   // book_tags
}

// AuthorBookResultsByTags runs a custom query, returning results as AuthorBookResult.
func AuthorBookResultsByTags(ctx context.Context, db DB, tags string) ([]*AuthorBookResult, error) {
	// query
	const sqlstr = `SELECT ` +
		`a.author_id AS author_id, ` +
		`a.name AS author_name, ` +
		`b.book_id AS book_id, ` +
		`b.isbn AS book_isbn, ` +
		`b.title AS book_title, ` +
		`b.tags AS book_tags ` +
		`FROM books b ` +
		`JOIN authors a ON a.author_id = b.author_id ` +
		`WHERE b.tags LIKE '%' || :1 || '%'`
	// run
	logf(sqlstr, tags)
	rows, err := db.QueryContext(ctx, sqlstr, tags)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*AuthorBookResult
	for rows.Next() {
		var abr AuthorBookResult
		// scan
		if err := rows.Scan(&abr.AuthorID, &abr.AuthorName, &abr.BookID, &abr.BookIsbn, &abr.BookTitle, &abr.BookTags); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &abr)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}
