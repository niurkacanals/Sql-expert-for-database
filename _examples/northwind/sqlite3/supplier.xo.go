package sqlite3

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// Supplier represents a row from 'suppliers'.
type Supplier struct {
	SupplierID   int            `json:"supplier_id"`   // supplier_id
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
	Homepage     sql.NullString `json:"homepage"`      // homepage
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the Supplier exists in the database.
func (s *Supplier) Exists() bool {
	return s._exists
}

// Deleted returns true when the Supplier has been marked for deletion from
// the database.
func (s *Supplier) Deleted() bool {
	return s._deleted
}

// Insert inserts the Supplier to the database.
func (s *Supplier) Insert(ctx context.Context, db DB) error {
	switch {
	case s._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case s._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (basic)
	const sqlstr = `INSERT INTO suppliers (` +
		`supplier_id, company_name, contact_name, contact_title, address, city, region, postal_code, country, phone, fax, homepage` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`
	// run
	logf(sqlstr, s.SupplierID, s.CompanyName, s.ContactName, s.ContactTitle, s.Address, s.City, s.Region, s.PostalCode, s.Country, s.Phone, s.Fax, s.Homepage)
	if err := db.QueryRowContext(ctx, sqlstr, s.SupplierID, s.CompanyName, s.ContactName, s.ContactTitle, s.Address, s.City, s.Region, s.PostalCode, s.Country, s.Phone, s.Fax, s.Homepage).Scan(&s.SupplierID); err != nil {
		return logerror(err)
	}
	// set exists
	s._exists = true
	return nil
}

// Update updates a Supplier in the database.
func (s *Supplier) Update(ctx context.Context, db DB) error {
	switch {
	case !s._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case s._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE suppliers SET ` +
		`company_name = ?, contact_name = ?, contact_title = ?, address = ?, city = ?, region = ?, postal_code = ?, country = ?, phone = ?, fax = ?, homepage = ?` +
		` WHERE supplier_id = ?`
	// run
	logf(sqlstr, s.CompanyName, s.ContactName, s.ContactTitle, s.Address, s.City, s.Region, s.PostalCode, s.Country, s.Phone, s.Fax, s.Homepage, s.SupplierID)
	if _, err := db.ExecContext(ctx, sqlstr, s.CompanyName, s.ContactName, s.ContactTitle, s.Address, s.City, s.Region, s.PostalCode, s.Country, s.Phone, s.Fax, s.Homepage, s.SupplierID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Supplier to the database.
func (s *Supplier) Save(ctx context.Context, db DB) error {
	if s.Exists() {
		return s.Update(ctx, db)
	}
	return s.Insert(ctx, db)
}

// Delete deletes the Supplier from the database.
func (s *Supplier) Delete(ctx context.Context, db DB) error {
	switch {
	case !s._exists: // doesn't exist
		return nil
	case s._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM suppliers WHERE supplier_id = ?`
	// run
	logf(sqlstr, s.SupplierID)
	if _, err := db.ExecContext(ctx, sqlstr, s.SupplierID); err != nil {
		return logerror(err)
	}
	// set deleted
	s._deleted = true
	return nil
}

// SupplierBySupplierID retrieves a row from 'suppliers' as a Supplier.
//
// Generated from index 'sqlite_autoindex_suppliers_1'.
func SupplierBySupplierID(ctx context.Context, db DB, supplierID int) (*Supplier, error) {
	// query
	const sqlstr = `SELECT ` +
		`supplier_id, company_name, contact_name, contact_title, address, city, region, postal_code, country, phone, fax, homepage ` +
		`FROM suppliers ` +
		`WHERE supplier_id = ?`
	// run
	logf(sqlstr, supplierID)
	s := Supplier{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, supplierID).Scan(&s.SupplierID, &s.CompanyName, &s.ContactName, &s.ContactTitle, &s.Address, &s.City, &s.Region, &s.PostalCode, &s.Country, &s.Phone, &s.Fax, &s.Homepage); err != nil {
		return nil, logerror(err)
	}
	return &s, nil
}
