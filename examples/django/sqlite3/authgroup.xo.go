// Package sqlite3 contains the types for schema ''.
package sqlite3

// GENERATED BY XO. DO NOT EDIT.

import "errors"

// AuthGroup represents a row from 'auth_group'.
type AuthGroup struct {
	ID   int    `json:"id"`   // id
	Name string `json:"name"` // name

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the AuthGroup exists in the database.
func (ag *AuthGroup) Exists() bool {
	return ag._exists
}

// Deleted provides information if the AuthGroup has been deleted from the database.
func (ag *AuthGroup) Deleted() bool {
	return ag._deleted
}

// Insert inserts the AuthGroup to the database.
func (ag *AuthGroup) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if ag._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO auth_group (` +
		`name` +
		`) VALUES (` +
		`?` +
		`)`

	// run query
	XOLog(sqlstr, ag.Name)
	res, err := db.Exec(sqlstr, ag.Name)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	ag.ID = int(id)
	ag._exists = true

	return nil
}

// Update updates the AuthGroup in the database.
func (ag *AuthGroup) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ag._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if ag._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE auth_group SET ` +
		`name = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, ag.Name, ag.ID)
	_, err = db.Exec(sqlstr, ag.Name, ag.ID)
	return err
}

// Save saves the AuthGroup to the database.
func (ag *AuthGroup) Save(db XODB) error {
	if ag.Exists() {
		return ag.Update(db)
	}

	return ag.Insert(db)
}

// Delete deletes the AuthGroup from the database.
func (ag *AuthGroup) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ag._exists {
		return nil
	}

	// if deleted, bail
	if ag._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM auth_group WHERE id = ?`

	// run query
	XOLog(sqlstr, ag.ID)
	_, err = db.Exec(sqlstr, ag.ID)
	if err != nil {
		return err
	}

	// set deleted
	ag._deleted = true

	return nil
}

// AuthGroupByID retrieves a row from 'auth_group' as a AuthGroup.
//
// Generated from index 'auth_group_id_pkey'.
func AuthGroupByID(db XODB, id int) (*AuthGroup, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name ` +
		`FROM auth_group ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	ag := AuthGroup{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&ag.ID, &ag.Name)
	if err != nil {
		return nil, err
	}

	return &ag, nil
}

// AuthGroupByName retrieves a row from 'auth_group' as a AuthGroup.
//
// Generated from index 'sqlite_autoindex_auth_group_1'.
func AuthGroupByName(db XODB, name string) (*AuthGroup, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name ` +
		`FROM auth_group ` +
		`WHERE name = ?`

	// run query
	XOLog(sqlstr, name)
	ag := AuthGroup{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, name).Scan(&ag.ID, &ag.Name)
	if err != nil {
		return nil, err
	}

	return &ag, nil
}
