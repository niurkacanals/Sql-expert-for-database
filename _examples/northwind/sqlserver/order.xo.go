package sqlserver

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// Order represents a row from 'northwind.orders'.
type Order struct {
	OrderID        int16           `json:"order_id"`         // order_id
	CustomerID     sql.NullString  `json:"customer_id"`      // customer_id
	EmployeeID     sql.NullInt64   `json:"employee_id"`      // employee_id
	OrderDate      sql.NullTime    `json:"order_date"`       // order_date
	RequiredDate   sql.NullTime    `json:"required_date"`    // required_date
	ShippedDate    sql.NullTime    `json:"shipped_date"`     // shipped_date
	Freight        sql.NullFloat64 `json:"freight"`          // freight
	ShipName       sql.NullString  `json:"ship_name"`        // ship_name
	ShipAddress    sql.NullString  `json:"ship_address"`     // ship_address
	ShipCity       sql.NullString  `json:"ship_city"`        // ship_city
	ShipRegion     sql.NullString  `json:"ship_region"`      // ship_region
	ShipPostalCode sql.NullString  `json:"ship_postal_code"` // ship_postal_code
	ShipCountry    sql.NullString  `json:"ship_country"`     // ship_country
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the Order exists in the database.
func (o *Order) Exists() bool {
	return o._exists
}

// Deleted returns true when the Order has been marked for deletion from
// the database.
func (o *Order) Deleted() bool {
	return o._deleted
}

// Insert inserts the Order to the database.
func (o *Order) Insert(ctx context.Context, db DB) error {
	switch {
	case o._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case o._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO northwind.orders (` +
		`order_id, customer_id, employee_id, order_date, required_date, shipped_date, freight, ship_name, ship_address, ship_city, ship_region, ship_postal_code, ship_country` +
		`) VALUES (` +
		`@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9, @p10, @p11, @p12, @p13` +
		`)`
	// run
	logf(sqlstr, o.OrderID, o.CustomerID, o.EmployeeID, o.OrderDate, o.RequiredDate, o.ShippedDate, o.Freight, o.ShipName, o.ShipAddress, o.ShipCity, o.ShipRegion, o.ShipPostalCode, o.ShipCountry)
	if _, err := db.ExecContext(ctx, sqlstr, o.OrderID, o.CustomerID, o.EmployeeID, o.OrderDate, o.RequiredDate, o.ShippedDate, o.Freight, o.ShipName, o.ShipAddress, o.ShipCity, o.ShipRegion, o.ShipPostalCode, o.ShipCountry); err != nil {
		return logerror(err)
	}
	// set exists
	o._exists = true
	return nil
}

// Update updates a Order in the database.
func (o *Order) Update(ctx context.Context, db DB) error {
	switch {
	case !o._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case o._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE northwind.orders SET ` +
		`customer_id = @p1, employee_id = @p2, order_date = @p3, required_date = @p4, shipped_date = @p5, freight = @p6, ship_name = @p7, ship_address = @p8, ship_city = @p9, ship_region = @p10, ship_postal_code = @p11, ship_country = @p12 ` +
		`WHERE order_id = @p13`
	// run
	logf(sqlstr, o.CustomerID, o.EmployeeID, o.OrderDate, o.RequiredDate, o.ShippedDate, o.Freight, o.ShipName, o.ShipAddress, o.ShipCity, o.ShipRegion, o.ShipPostalCode, o.ShipCountry, o.OrderID)
	if _, err := db.ExecContext(ctx, sqlstr, o.CustomerID, o.EmployeeID, o.OrderDate, o.RequiredDate, o.ShippedDate, o.Freight, o.ShipName, o.ShipAddress, o.ShipCity, o.ShipRegion, o.ShipPostalCode, o.ShipCountry, o.OrderID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the Order to the database.
func (o *Order) Save(ctx context.Context, db DB) error {
	if o.Exists() {
		return o.Update(ctx, db)
	}
	return o.Insert(ctx, db)
}

// Delete deletes the Order from the database.
func (o *Order) Delete(ctx context.Context, db DB) error {
	switch {
	case !o._exists: // doesn't exist
		return nil
	case o._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM northwind.orders ` +
		`WHERE order_id = @p1`
	// run
	logf(sqlstr, o.OrderID)
	if _, err := db.ExecContext(ctx, sqlstr, o.OrderID); err != nil {
		return logerror(err)
	}
	// set deleted
	o._deleted = true
	return nil
}

// OrderByOrderID retrieves a row from 'northwind.orders' as a Order.
//
// Generated from index 'orders_pkey'.
func OrderByOrderID(ctx context.Context, db DB, orderID int16) (*Order, error) {
	// query
	const sqlstr = `SELECT ` +
		`order_id, customer_id, employee_id, order_date, required_date, shipped_date, freight, ship_name, ship_address, ship_city, ship_region, ship_postal_code, ship_country ` +
		`FROM northwind.orders ` +
		`WHERE ` +
		`order_id = @p1`
	// run
	logf(sqlstr, orderID)
	o := Order{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, orderID).Scan(&o.OrderID, &o.CustomerID, &o.EmployeeID, &o.OrderDate, &o.RequiredDate, &o.ShippedDate, &o.Freight, &o.ShipName, &o.ShipAddress, &o.ShipCity, &o.ShipRegion, &o.ShipPostalCode, &o.ShipCountry); err != nil {
		return nil, logerror(err)
	}
	return &o, nil
}

// Customer returns the Customer associated with the Order's CustomerID (customer_id).
//
// Generated from foreign key 'orders_customer_id_fkey'.
func (o *Order) Customer(ctx context.Context, db DB) (*Customer, error) {
	return CustomerByCustomerID(ctx, db, o.CustomerID.String)
}

// Employee returns the Employee associated with the Order's EmployeeID (employee_id).
//
// Generated from foreign key 'orders_employee_id_fkey'.
func (o *Order) Employee(ctx context.Context, db DB) (*Employee, error) {
	return EmployeeByEmployeeID(ctx, db, int16(o.EmployeeID.Int64))
}
