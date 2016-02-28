// Package mysql contains the types for schema 'django'.
package mysql

// GENERATED BY XO. DO NOT EDIT.

import (
	"errors"
	"time"
)

// DjangoSession represents a row from 'django.django_session'.
type DjangoSession struct {
	SessionKey  string     // session_key
	SessionData string     // session_data
	ExpireDate  *time.Time // expire_date

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the DjangoSession exists in the database.
func (ds *DjangoSession) Exists() bool {
	return ds._exists
}

// Deleted provides information if the DjangoSession has been deleted from the database.
func (ds *DjangoSession) Deleted() bool {
	return ds._deleted
}

// Insert inserts the DjangoSession to the database.
func (ds *DjangoSession) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if ds._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO django.django_session (` +
		`session_data, expire_date` +
		`) VALUES (` +
		`?, ?` +
		`)`

	// run query
	XOLog(sqlstr, ds.SessionData, ds.ExpireDate)
	res, err := db.Exec(sqlstr, ds.SessionData, ds.ExpireDate)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	ds.SessionKey = string(id)
	ds._exists = true

	return nil
}

// Update updates the DjangoSession in the database.
func (ds *DjangoSession) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ds._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if ds._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE django.django_session SET ` +
		`session_data = ?, expire_date = ?` +
		` WHERE session_key = ?`

	// run query
	XOLog(sqlstr, ds.SessionData, ds.ExpireDate, ds.SessionKey)
	_, err = db.Exec(sqlstr, ds.SessionData, ds.ExpireDate, ds.SessionKey)
	return err
}

// Save saves the DjangoSession to the database.
func (ds *DjangoSession) Save(db XODB) error {
	if ds.Exists() {
		return ds.Update(db)
	}

	return ds.Insert(db)
}

// Delete deletes the DjangoSession from the database.
func (ds *DjangoSession) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ds._exists {
		return nil
	}

	// if deleted, bail
	if ds._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM django.django_session WHERE session_key = ?`

	// run query
	XOLog(sqlstr, ds.SessionKey)
	_, err = db.Exec(sqlstr, ds.SessionKey)
	if err != nil {
		return err
	}

	// set deleted
	ds._deleted = true

	return nil
}

// DjangoSessionsByExpireDate retrieves a row from 'django.django_session' as a DjangoSession.
//
// Generated from index 'django_session_de54fa62'.
func DjangoSessionsByExpireDate(db XODB, expireDate *time.Time) ([]*DjangoSession, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`session_key, session_data, expire_date ` +
		`FROM django.django_session ` +
		`WHERE expire_date = ?`

	// run query
	XOLog(sqlstr, expireDate)
	q, err := db.Query(sqlstr, expireDate)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*DjangoSession{}
	for q.Next() {
		ds := DjangoSession{
			_exists: true,
		}

		// scan
		err = q.Scan(&ds.SessionKey, &ds.SessionData, &ds.ExpireDate)
		if err != nil {
			return nil, err
		}

		res = append(res, &ds)
	}

	return res, nil
}

// DjangoSessionBySessionKey retrieves a row from 'django.django_session' as a DjangoSession.
//
// Generated from index 'django_session_session_key_pkey'.
func DjangoSessionBySessionKey(db XODB, sessionKey string) (*DjangoSession, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`session_key, session_data, expire_date ` +
		`FROM django.django_session ` +
		`WHERE session_key = ?`

	// run query
	XOLog(sqlstr, sessionKey)
	ds := DjangoSession{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, sessionKey).Scan(&ds.SessionKey, &ds.SessionData, &ds.ExpireDate)
	if err != nil {
		return nil, err
	}

	return &ds, nil
}
