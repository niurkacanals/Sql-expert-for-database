package mysql

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// Shipper represents a row from 'northwind.shippers'.
type Shipper struct {
	ShipperID   int16          `json:"shipper_id"`   // shipper_id
	CompanyName string         `json:"company_name"` // company_name
	Phone       sql.NullString `json:"phone"`        // phone
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the Shipper exists in the database.
func (s *Shipper) Exists() bool {
	return s._exists
}

// Deleted returns true when the Shipper has been marked for deletion from
// the database.
func (s *Shipper) Deleted() bool {
	return s._deleted
}

// Insert inserts the Shipper to the database.
func (s *Shipper) Insert(ctx context.Context, db DB) error {
	switch {
	case s._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case s._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (basic)
	const sqlstr = `INSERT INTO northwind.shippers (` +
		`shipper_id, company_name, phone` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`
	// run
	logf(sqlstr, s.ShipperID, s.CompanyName, s.Phone)
	if err := db.QueryRowContext(ctx, sqlstr, s.ShipperID, s.CompanyName, s.Phone).Scan(&s.ShipperID); err != nil {
		return logerror(err)
	}
	// set exists
	s._exists = true
	return nil
}

// Update updates a Shipper in the database.
func (s *Shipper) Update(ctx context.Context, db DB) error {
	switch {
	case !s._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case s._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE northwind.shippers SET ` +
		`company_name = ?, phone = ?` +
		` WHERE shipper_id = ?`
	// run
	logf(sqlstr, s.CompanyName, s.Phone, s.ShipperID)
	if _, err := db.ExecContext(ctx, sqlstr, s.CompanyName, s.Phone, s.ShipperID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Shipper to the database.
func (s *Shipper) Save(ctx context.Context, db DB) error {
	if s.Exists() {
		return s.Update(ctx, db)
	}
	return s.Insert(ctx, db)
}

// Delete deletes the Shipper from the database.
func (s *Shipper) Delete(ctx context.Context, db DB) error {
	switch {
	case !s._exists: // doesn't exist
		return nil
	case s._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM northwind.shippers WHERE shipper_id = ?`
	// run
	logf(sqlstr, s.ShipperID)
	if _, err := db.ExecContext(ctx, sqlstr, s.ShipperID); err != nil {
		return logerror(err)
	}
	// set deleted
	s._deleted = true
	return nil
}

// ShipperByShipperID retrieves a row from 'northwind.shippers' as a Shipper.
//
// Generated from index 'shippers_shipper_id_pkey'.
func ShipperByShipperID(ctx context.Context, db DB, shipperID int16) (*Shipper, error) {
	// query
	const sqlstr = `SELECT ` +
		`shipper_id, company_name, phone ` +
		`FROM northwind.shippers ` +
		`WHERE shipper_id = ?`
	// run
	logf(sqlstr, shipperID)
	s := Shipper{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, shipperID).Scan(&s.ShipperID, &s.CompanyName, &s.Phone); err != nil {
		return nil, logerror(err)
	}
	return &s, nil
}
