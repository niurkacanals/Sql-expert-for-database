package sqlserver

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// Category represents a row from 'northwind.categories'.
type Category struct {
	CategoryID   int16          `json:"category_id"`   // category_id
	CategoryName string         `json:"category_name"` // category_name
	Description  sql.NullString `json:"description"`   // description
	Picture      []byte         `json:"picture"`       // picture
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the Category exists in the database.
func (c *Category) Exists() bool {
	return c._exists
}

// Deleted returns true when the Category has been marked for deletion from
// the database.
func (c *Category) Deleted() bool {
	return c._deleted
}

// Insert inserts the Category to the database.
func (c *Category) Insert(ctx context.Context, db DB) error {
	switch {
	case c._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case c._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO northwind.categories (` +
		`category_id, category_name, description, picture` +
		`) VALUES (` +
		`@p1, @p2, @p3, @p4` +
		`)`
	// run
	logf(sqlstr, c.CategoryID, c.CategoryName, c.Description, c.Picture)
	if _, err := db.ExecContext(ctx, sqlstr, c.CategoryID, c.CategoryName, c.Description, c.Picture); err != nil {
		return logerror(err)
	}
	// set exists
	c._exists = true
	return nil
}

// Update updates a Category in the database.
func (c *Category) Update(ctx context.Context, db DB) error {
	switch {
	case !c._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case c._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE northwind.categories SET ` +
		`category_name = @p1, description = @p2, picture = @p3` +
		` WHERE category_id = @p4`
	// run
	logf(sqlstr, c.CategoryName, c.Description, c.Picture, c.CategoryID)
	if _, err := db.ExecContext(ctx, sqlstr, c.CategoryName, c.Description, c.Picture, c.CategoryID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Category to the database.
func (c *Category) Save(ctx context.Context, db DB) error {
	if c.Exists() {
		return c.Update(ctx, db)
	}
	return c.Insert(ctx, db)
}

// Delete deletes the Category from the database.
func (c *Category) Delete(ctx context.Context, db DB) error {
	switch {
	case !c._exists: // doesn't exist
		return nil
	case c._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM northwind.categories WHERE category_id = @p1`
	// run
	logf(sqlstr, c.CategoryID)
	if _, err := db.ExecContext(ctx, sqlstr, c.CategoryID); err != nil {
		return logerror(err)
	}
	// set deleted
	c._deleted = true
	return nil
}

// CategoryByCategoryID retrieves a row from 'northwind.categories' as a Category.
//
// Generated from index 'categories_pkey'.
func CategoryByCategoryID(ctx context.Context, db DB, categoryID int16) (*Category, error) {
	// query
	const sqlstr = `SELECT ` +
		`category_id, category_name, description, picture ` +
		`FROM northwind.categories ` +
		`WHERE category_id = @p1`
	// run
	logf(sqlstr, categoryID)
	c := Category{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, categoryID).Scan(&c.CategoryID, &c.CategoryName, &c.Description, &c.Picture); err != nil {
		return nil, logerror(err)
	}
	return &c, nil
}
