package sqlserver

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// Customer represents a row from 'northwind.customers'.
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
	const sqlstr = `INSERT INTO northwind.customers (` +
		`customer_id, company_name, contact_name, contact_title, address, city, region, postal_code, country, phone, fax` +
		`) VALUES (` +
		`@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9, @p10, @p11` +
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
	const sqlstr = `UPDATE northwind.customers SET ` +
		`company_name = @p1, contact_name = @p2, contact_title = @p3, address = @p4, city = @p5, region = @p6, postal_code = @p7, country = @p8, phone = @p9, fax = @p10 ` +
		`WHERE customer_id = @p11`
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

// Upsert performs an upsert for Customer.
func (c *Customer) Upsert(ctx context.Context, db DB) error {
	switch {
	case c._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `MERGE northwind.customers AS t ` +
		`USING (` +
		`SELECT @p1 customer_id, @p2 company_name, @p3 contact_name, @p4 contact_title, @p5 address, @p6 city, @p7 region, @p8 postal_code, @p9 country, @p10 phone, @p11 fax ` +
		`) AS s ` +
		`ON s.customer_id = t.customer_id ` +
		`WHEN MATCHED THEN ` +
		`UPDATE SET ` +
		`t.company_name = s.company_name, t.contact_name = s.contact_name, t.contact_title = s.contact_title, t.address = s.address, t.city = s.city, t.region = s.region, t.postal_code = s.postal_code, t.country = s.country, t.phone = s.phone, t.fax = s.fax ` +
		`WHEN NOT MATCHED THEN ` +
		`INSERT (` +
		`customer_id, company_name, contact_name, contact_title, address, city, region, postal_code, country, phone, fax` +
		`) VALUES (` +
		`s.customer_id, s.company_name, s.contact_name, s.contact_title, s.address, s.city, s.region, s.postal_code, s.country, s.phone, s.fax` +
		`);`
	// run
	logf(sqlstr, c.CustomerID, c.CompanyName, c.ContactName, c.ContactTitle, c.Address, c.City, c.Region, c.PostalCode, c.Country, c.Phone, c.Fax)
	if _, err := db.ExecContext(ctx, sqlstr, c.CustomerID, c.CompanyName, c.ContactName, c.ContactTitle, c.Address, c.City, c.Region, c.PostalCode, c.Country, c.Phone, c.Fax); err != nil {
		return err
	}
	// set exists
	c._exists = true
	return nil
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
	const sqlstr = `DELETE FROM northwind.customers ` +
		`WHERE customer_id = @p1`
	// run
	logf(sqlstr, c.CustomerID)
	if _, err := db.ExecContext(ctx, sqlstr, c.CustomerID); err != nil {
		return logerror(err)
	}
	// set deleted
	c._deleted = true
	return nil
}

// CustomerByCustomerID retrieves a row from 'northwind.customers' as a Customer.
//
// Generated from index 'customers_pkey'.
func CustomerByCustomerID(ctx context.Context, db DB, customerID string) (*Customer, error) {
	// query
	const sqlstr = `SELECT ` +
		`customer_id, company_name, contact_name, contact_title, address, city, region, postal_code, country, phone, fax ` +
		`FROM northwind.customers ` +
		`WHERE customer_id = @p1`
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
