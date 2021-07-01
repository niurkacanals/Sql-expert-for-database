package oracle

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// CustomerCustomerDemo represents a row from 'northwind.customer_customer_demo'.
type CustomerCustomerDemo struct {
	CustomerID     string `json:"customer_id"`      // customer_id
	CustomerTypeID string `json:"customer_type_id"` // customer_type_id
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the CustomerCustomerDemo exists in the database.
func (ccd *CustomerCustomerDemo) Exists() bool {
	return ccd._exists
}

// Deleted returns true when the CustomerCustomerDemo has been marked for deletion from
// the database.
func (ccd *CustomerCustomerDemo) Deleted() bool {
	return ccd._deleted
}

// Insert inserts the CustomerCustomerDemo to the database.
func (ccd *CustomerCustomerDemo) Insert(ctx context.Context, db DB) error {
	switch {
	case ccd._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case ccd._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO northwind.customer_customer_demo (` +
		`customer_id, customer_type_id` +
		`) VALUES (` +
		`:1, :2` +
		`)`
	// run
	logf(sqlstr, ccd.CustomerID, ccd.CustomerTypeID)
	if _, err := db.ExecContext(ctx, sqlstr, ccd.CustomerID, ccd.CustomerTypeID); err != nil {
		return logerror(err)
	}
	// set exists
	ccd._exists = true
	return nil
}

// ------ NOTE: Update statements omitted due to lack of fields other than primary key ------

// Delete deletes the CustomerCustomerDemo from the database.
func (ccd *CustomerCustomerDemo) Delete(ctx context.Context, db DB) error {
	switch {
	case !ccd._exists: // doesn't exist
		return nil
	case ccd._deleted: // deleted
		return nil
	}
	// delete with composite primary key
	const sqlstr = `DELETE FROM northwind.customer_customer_demo ` +
		`WHERE customer_id = :1 AND customer_type_id = :2`
	// run
	logf(sqlstr, ccd.CustomerID, ccd.CustomerTypeID)
	if _, err := db.ExecContext(ctx, sqlstr, ccd.CustomerID, ccd.CustomerTypeID); err != nil {
		return logerror(err)
	}
	// set deleted
	ccd._deleted = true
	return nil
}

// CustomerCustomerDemoByCustomerIDCustomerTypeID retrieves a row from 'northwind.customer_customer_demo' as a CustomerCustomerDemo.
//
// Generated from index 'customer_customer_demo_pkey'.
func CustomerCustomerDemoByCustomerIDCustomerTypeID(ctx context.Context, db DB, customerID, customerTypeID string) (*CustomerCustomerDemo, error) {
	// query
	const sqlstr = `SELECT ` +
		`customer_id, customer_type_id ` +
		`FROM northwind.customer_customer_demo ` +
		`WHERE customer_id = :1 AND customer_type_id = :2`
	// run
	logf(sqlstr, customerID, customerTypeID)
	ccd := CustomerCustomerDemo{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, customerID, customerTypeID).Scan(&ccd.CustomerID, &ccd.CustomerTypeID); err != nil {
		return nil, logerror(err)
	}
	return &ccd, nil
}

// Customer returns the Customer associated with the CustomerCustomerDemo's CustomerID (customer_id).
//
// Generated from foreign key 'customer_customer_demo_customer_id_fkey'.
func (ccd *CustomerCustomerDemo) Customer(ctx context.Context, db DB) (*Customer, error) {
	return CustomerByCustomerID(ctx, db, ccd.CustomerID)
}

// CustomerDemographic returns the CustomerDemographic associated with the CustomerCustomerDemo's CustomerTypeID (customer_type_id).
//
// Generated from foreign key 'customer_customer_demo_customer_type_id_fkey'.
func (ccd *CustomerCustomerDemo) CustomerDemographic(ctx context.Context, db DB) (*CustomerDemographic, error) {
	return CustomerDemographicByCustomerTypeID(ctx, db, ccd.CustomerTypeID)
}
