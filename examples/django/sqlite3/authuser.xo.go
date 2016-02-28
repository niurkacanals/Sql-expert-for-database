// Package sqlite3 contains the types for schema ''.
package sqlite3

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// AuthUser represents a row from 'auth_user'.
type AuthUser struct {
	ID          int            // id
	Password    string         // password
	LastLogin   sql.NullString // last_login
	IsSuperuser bool           // is_superuser
	FirstName   string         // first_name
	LastName    string         // last_name
	Email       string         // email
	IsStaff     bool           // is_staff
	IsActive    bool           // is_active
	DateJoined  string         // date_joined
	Username    string         // username

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the AuthUser exists in the database.
func (au *AuthUser) Exists() bool {
	return au._exists
}

// Deleted provides information if the AuthUser has been deleted from the database.
func (au *AuthUser) Deleted() bool {
	return au._deleted
}

// Insert inserts the AuthUser to the database.
func (au *AuthUser) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if au._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO auth_user (` +
		`password, last_login, is_superuser, first_name, last_name, email, is_staff, is_active, date_joined, username` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, au.Password, au.LastLogin, au.IsSuperuser, au.FirstName, au.LastName, au.Email, au.IsStaff, au.IsActive, au.DateJoined, au.Username)
	res, err := db.Exec(sqlstr, au.Password, au.LastLogin, au.IsSuperuser, au.FirstName, au.LastName, au.Email, au.IsStaff, au.IsActive, au.DateJoined, au.Username)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	au.ID = int(id)
	au._exists = true

	return nil
}

// Update updates the AuthUser in the database.
func (au *AuthUser) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !au._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if au._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE auth_user SET ` +
		`password = ?, last_login = ?, is_superuser = ?, first_name = ?, last_name = ?, email = ?, is_staff = ?, is_active = ?, date_joined = ?, username = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, au.Password, au.LastLogin, au.IsSuperuser, au.FirstName, au.LastName, au.Email, au.IsStaff, au.IsActive, au.DateJoined, au.Username, au.ID)
	_, err = db.Exec(sqlstr, au.Password, au.LastLogin, au.IsSuperuser, au.FirstName, au.LastName, au.Email, au.IsStaff, au.IsActive, au.DateJoined, au.Username, au.ID)
	return err
}

// Save saves the AuthUser to the database.
func (au *AuthUser) Save(db XODB) error {
	if au.Exists() {
		return au.Update(db)
	}

	return au.Insert(db)
}

// Delete deletes the AuthUser from the database.
func (au *AuthUser) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !au._exists {
		return nil
	}

	// if deleted, bail
	if au._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM auth_user WHERE id = ?`

	// run query
	XOLog(sqlstr, au.ID)
	_, err = db.Exec(sqlstr, au.ID)
	if err != nil {
		return err
	}

	// set deleted
	au._deleted = true

	return nil
}

// AuthUserByID retrieves a row from 'auth_user' as a AuthUser.
//
// Generated from index 'auth_user_id_pkey'.
func AuthUserByID(db XODB, id int) (*AuthUser, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, password, last_login, is_superuser, first_name, last_name, email, is_staff, is_active, date_joined, username ` +
		`FROM auth_user ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	au := AuthUser{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&au.ID, &au.Password, &au.LastLogin, &au.IsSuperuser, &au.FirstName, &au.LastName, &au.Email, &au.IsStaff, &au.IsActive, &au.DateJoined, &au.Username)
	if err != nil {
		return nil, err
	}

	return &au, nil
}

// AuthUserByUsername retrieves a row from 'auth_user' as a AuthUser.
//
// Generated from index 'sqlite_autoindex_auth_user_1'.
func AuthUserByUsername(db XODB, username string) (*AuthUser, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, password, last_login, is_superuser, first_name, last_name, email, is_staff, is_active, date_joined, username ` +
		`FROM auth_user ` +
		`WHERE username = ?`

	// run query
	XOLog(sqlstr, username)
	au := AuthUser{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, username).Scan(&au.ID, &au.Password, &au.LastLogin, &au.IsSuperuser, &au.FirstName, &au.LastName, &au.Email, &au.IsStaff, &au.IsActive, &au.DateJoined, &au.Username)
	if err != nil {
		return nil, err
	}

	return &au, nil
}
