package sqlserver

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"time"
)

// DjangoSession represents a row from 'django.django_session'.
type DjangoSession struct {
	SessionKey  string    `json:"session_key"`  // session_key
	SessionData string    `json:"session_data"` // session_data
	ExpireDate  time.Time `json:"expire_date"`  // expire_date
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the DjangoSession exists in the database.
func (ds *DjangoSession) Exists() bool {
	return ds._exists
}

// Deleted returns true when the DjangoSession has been marked for deletion from
// the database.
func (ds *DjangoSession) Deleted() bool {
	return ds._deleted
}

// Insert inserts the DjangoSession to the database.
func (ds *DjangoSession) Insert(ctx context.Context, db DB) error {
	switch {
	case ds._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case ds._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO django.django_session (` +
		`session_key, session_data, expire_date` +
		`) VALUES (` +
		`@p1, @p2, @p3` +
		`)`
	// run
	logf(sqlstr, ds.SessionKey, ds.SessionData, ds.ExpireDate)
	if _, err := db.ExecContext(ctx, sqlstr, ds.SessionKey, ds.SessionData, ds.ExpireDate); err != nil {
		return logerror(err)
	}
	// set exists
	ds._exists = true
	return nil
}

// Update updates a DjangoSession in the database.
func (ds *DjangoSession) Update(ctx context.Context, db DB) error {
	switch {
	case !ds._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case ds._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE django.django_session SET ` +
		`session_data = @p1, expire_date = @p2 ` +
		`WHERE session_key = @p3`
	// run
	logf(sqlstr, ds.SessionData, ds.ExpireDate, ds.SessionKey)
	if _, err := db.ExecContext(ctx, sqlstr, ds.SessionData, ds.ExpireDate, ds.SessionKey); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the DjangoSession to the database.
func (ds *DjangoSession) Save(ctx context.Context, db DB) error {
	if ds.Exists() {
		return ds.Update(ctx, db)
	}
	return ds.Insert(ctx, db)
}

// Upsert performs an upsert for DjangoSession.
func (ds *DjangoSession) Upsert(ctx context.Context, db DB) error {
	switch {
	case ds._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `MERGE django.django_session AS t ` +
		`USING (` +
		`SELECT @p1 session_key, @p2 session_data, @p3 expire_date ` +
		`) AS s ` +
		`ON s.session_key = t.session_key ` +
		`WHEN MATCHED THEN ` +
		`UPDATE SET ` +
		`t.session_data = s.session_data, t.expire_date = s.expire_date ` +
		`WHEN NOT MATCHED THEN ` +
		`INSERT (` +
		`session_key, session_data, expire_date` +
		`) VALUES (` +
		`s.session_key, s.session_data, s.expire_date` +
		`);`
	// run
	logf(sqlstr, ds.SessionKey, ds.SessionData, ds.ExpireDate)
	if _, err := db.ExecContext(ctx, sqlstr, ds.SessionKey, ds.SessionData, ds.ExpireDate); err != nil {
		return err
	}
	// set exists
	ds._exists = true
	return nil
}

// Delete deletes the DjangoSession from the database.
func (ds *DjangoSession) Delete(ctx context.Context, db DB) error {
	switch {
	case !ds._exists: // doesn't exist
		return nil
	case ds._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM django.django_session ` +
		`WHERE session_key = @p1`
	// run
	logf(sqlstr, ds.SessionKey)
	if _, err := db.ExecContext(ctx, sqlstr, ds.SessionKey); err != nil {
		return logerror(err)
	}
	// set deleted
	ds._deleted = true
	return nil
}

// DjangoSessionBySessionKey retrieves a row from 'django.django_session' as a DjangoSession.
//
// Generated from index 'PK__django_s__B3BA0F1FC45FFD1B'.
func DjangoSessionBySessionKey(ctx context.Context, db DB, sessionKey string) (*DjangoSession, error) {
	// query
	const sqlstr = `SELECT ` +
		`session_key, session_data, expire_date ` +
		`FROM django.django_session ` +
		`WHERE session_key = @p1`
	// run
	logf(sqlstr, sessionKey)
	ds := DjangoSession{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, sessionKey).Scan(&ds.SessionKey, &ds.SessionData, &ds.ExpireDate); err != nil {
		return nil, logerror(err)
	}
	return &ds, nil
}

// DjangoSessionByExpireDate retrieves a row from 'django.django_session' as a DjangoSession.
//
// Generated from index 'django_session_expire_date_a5c62663'.
func DjangoSessionByExpireDate(ctx context.Context, db DB, expireDate time.Time) ([]*DjangoSession, error) {
	// query
	const sqlstr = `SELECT ` +
		`session_key, session_data, expire_date ` +
		`FROM django.django_session ` +
		`WHERE expire_date = @p1`
	// run
	logf(sqlstr, expireDate)
	rows, err := db.QueryContext(ctx, sqlstr, expireDate)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*DjangoSession
	for rows.Next() {
		ds := DjangoSession{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&ds.SessionKey, &ds.SessionData, &ds.ExpireDate); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &ds)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}
