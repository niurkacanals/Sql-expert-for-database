package sqlite3

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// Customer represents a row from 'customers'.
type Customer struct {
	CustomerID   string         `json:"customer_id"`   // customer_id
	CompanyName  string         `json:"company_name"`  // company_name
	ContactName  sql.NullString `json:"contact_name"`  // contact_name
	ContactTitle sql.NullString `json:"contact_title"` // contact_title
	Address      sql.NullString `json:"address"`       // address
	City         sql.NullString `json:"city"`          // city
	Region       sql.NullString `json:"region"`        // region
	PostalCode   sql.NullString `json:"postal_code"`   // postal_code
	Country      sql.NullString `json:"country"`       // country
	Phone        sql.NullString `json:"phone"`         // phone
	Fax          sql.NullString `json:"fax"`           // fax
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the Customer exists in the database.
func (c *Customer) Exists() bool {
	return c._exists
}

// Deleted returns true when the Customer has been marked for deletion from
// the database.
func (c *Customer) Deleted() bool {
	return c._deleted
}

// Insert inserts the Customer to the database.
func (c *Customer) Insert(ctx context.Context, db DB) error {
	switch {
	case c._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case c._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO customers (` +
		`customer_id, company_name, contact_name, contact_title, address, city, region, postal_code, country, phone, fax` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11` +
		`)`
	// run
	logf(sqlstr, c.CustomerID, c.CompanyName, c.ContactName, c.ContactTitle, c.Address, c.City, c.Region, c.PostalCode, c.Country, c.Phone, c.Fax)
	if _, err := db.ExecContext(ctx, sqlstr, c.CustomerID, c.CompanyName, c.ContactName, c.ContactTitle, c.Address, c.City, c.Region, c.PostalCode, c.Country, c.Phone, c.Fax); err != nil {
		return logerror(err)
	}
	// set exists
	c._exists = true
	return nil
}

// Update updates a Customer in the database.
func (c *Customer) Update(ctx context.Context, db DB) error {
	switch {
	case !c._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case c._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE customers SET ` +
		`company_name = $1, contact_name = $2, contact_title = $3, address = $4, city = $5, region = $6, postal_code = $7, country = $8, phone = $9, fax = $10 ` +
		`WHERE customer_id = $11`
	// run
	logf(sqlstr, c.CompanyName, c.ContactName, c.ContactTitle, c.Address, c.City, c.Region, c.PostalCode, c.Country, c.Phone, c.Fax, c.CustomerID)
	if _, err := db.ExecContext(ctx, sqlstr, c.CompanyName, c.ContactName, c.ContactTitle, c.Address, c.City, c.Region, c.PostalCode, c.Country, c.Phone, c.Fax, c.CustomerID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Customer to the database.
func (c *Customer) Save(ctx context.Context, db DB) error {
	if c.Exists() {
		return c.Update(ctx, db)
	}
	return c.Insert(ctx, db)
}

// Delete deletes the Customer from the database.
func (c *Customer) Delete(ctx context.Context, db DB) error {
	switch {
	case !c._exists: // doesn't exist
		return nil
	case c._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM customers ` +
		`WHERE customer_id = $1`
	// run
	logf(sqlstr, c.CustomerID)
	if _, err := db.ExecContext(ctx, sqlstr, c.CustomerID); err != nil {
		return logerror(err)
	}
	// set deleted
	c._deleted = true
	return nil
}

// CustomerByCustomerID retrieves a row from 'customers' as a Customer.
//
// Generated from index 'sqlite_autoindex_customers_1'.
func CustomerByCustomerID(ctx context.Context, db DB, customerID string) (*Customer, error) {
	// query
	const sqlstr = `SELECT ` +
		`customer_id, company_name, contact_name, contact_title, address, city, region, postal_code, country, phone, fax ` +
		`FROM customers ` +
		`WHERE customer_id = $1`
	// run
	logf(sqlstr, customerID)
	c := Customer{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, customerID).Scan(&c.CustomerID, &c.CompanyName, &c.ContactName, &c.ContactTitle, &c.Address, &c.City, &c.Region, &c.PostalCode, &c.Country, &c.Phone, &c.Fax); err != nil {
		return nil, logerror(err)
	}
	return &c, nil
}
