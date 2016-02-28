// Package postgres contains the types for schema 'public'.
package postgres

// GENERATED BY XO. DO NOT EDIT.

import (
	"errors"
	"time"
)

// DjangoSession represents a row from public.django_session.
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
	const sqlstr = `INSERT INTO public.django_session (` +
		`session_data, expire_date` +
		`) VALUES (` +
		`$1, $2` +
		`) RETURNING session_key`

	// run query
	XOLog(sqlstr, ds.SessionData, ds.ExpireDate)
	err = db.QueryRow(sqlstr, ds.SessionData, ds.ExpireDate).Scan(&ds.SessionKey)
	if err != nil {
		return err
	}

	// set existence
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
	const sqlstr = `UPDATE public.django_session SET (` +
		`session_data, expire_date` +
		`) = ( ` +
		`$1, $2` +
		`) WHERE session_key = $3`

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

// Upsert performs an upsert for DjangoSession.
//
// NOTE: PostgreSQL 9.5+ only
func (ds *DjangoSession) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if ds._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.django_session (` +
		`session_key, session_data, expire_date` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) ON CONFLICT (session_key) DO UPDATE SET (` +
		`session_key, session_data, expire_date` +
		`) = (` +
		`EXCLUDED.session_key, EXCLUDED.session_data, EXCLUDED.expire_date` +
		`)`

	// run query
	XOLog(sqlstr, ds.SessionKey, ds.SessionData, ds.ExpireDate)
	_, err = db.Exec(sqlstr, ds.SessionKey, ds.SessionData, ds.ExpireDate)
	if err != nil {
		return err
	}

	// set existence
	ds._exists = true

	return nil
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
	const sqlstr = `DELETE FROM public.django_session WHERE session_key = $1`

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

// DjangoSessionsByExpireDate retrieves a row from 'public.django_session' as a DjangoSession.
//
// Generated from index 'django_session_de54fa62'.
func DjangoSessionsByExpireDate(db XODB, expireDate *time.Time) ([]*DjangoSession, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`session_key, session_data, expire_date ` +
		`FROM public.django_session ` +
		`WHERE expire_date = $1`

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

// DjangoSessionBySessionKey retrieves a row from 'public.django_session' as a DjangoSession.
//
// Generated from index 'django_session_pkey'.
func DjangoSessionBySessionKey(db XODB, sessionKey string) (*DjangoSession, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`session_key, session_data, expire_date ` +
		`FROM public.django_session ` +
		`WHERE session_key = $1`

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

// DjangoSessionsBySessionKey retrieves a row from 'public.django_session' as a DjangoSession.
//
// Generated from index 'django_session_session_key_c0390e0f_like'.
func DjangoSessionsBySessionKey(db XODB, sessionKey string) ([]*DjangoSession, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`session_key, session_data, expire_date ` +
		`FROM public.django_session ` +
		`WHERE session_key = $1`

	// run query
	XOLog(sqlstr, sessionKey)
	q, err := db.Query(sqlstr, sessionKey)
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
