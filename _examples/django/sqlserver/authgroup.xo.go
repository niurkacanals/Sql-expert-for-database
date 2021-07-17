package sqlserver

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// AuthGroup represents a row from 'django.auth_group'.
type AuthGroup struct {
	ID   int    `json:"id"`   // id
	Name string `json:"name"` // name
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the AuthGroup exists in the database.
func (ag *AuthGroup) Exists() bool {
	return ag._exists
}

// Deleted returns true when the AuthGroup has been marked for deletion from
// the database.
func (ag *AuthGroup) Deleted() bool {
	return ag._deleted
}

// Insert inserts the AuthGroup to the database.
func (ag *AuthGroup) Insert(ctx context.Context, db DB) error {
	switch {
	case ag._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case ag._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO django.auth_group (` +
		`name` +
		`) VALUES (` +
		`@p1` +
		`); SELECT ID = CONVERT(BIGINT, SCOPE_IDENTITY())`
	// run
	logf(sqlstr, ag.Name)
	rows, err := db.QueryContext(ctx, sqlstr, ag.Name)
	if err != nil {
		return logerror(err)
	}
	defer rows.Close()
	// retrieve id
	var id int64
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return logerror(err)
		}
	}
	if err := rows.Err(); err != nil {
		return logerror(err)
	} // set primary key
	ag.ID = int(id)
	// set exists
	ag._exists = true
	return nil
}

// Update updates a AuthGroup in the database.
func (ag *AuthGroup) Update(ctx context.Context, db DB) error {
	switch {
	case !ag._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case ag._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE django.auth_group SET ` +
		`name = @p1 ` +
		`WHERE id = @p2`
	// run
	logf(sqlstr, ag.Name, ag.ID)
	if _, err := db.ExecContext(ctx, sqlstr, ag.Name, ag.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the AuthGroup to the database.
func (ag *AuthGroup) Save(ctx context.Context, db DB) error {
	if ag.Exists() {
		return ag.Update(ctx, db)
	}
	return ag.Insert(ctx, db)
}

// Upsert performs an upsert for AuthGroup.
func (ag *AuthGroup) Upsert(ctx context.Context, db DB) error {
	switch {
	case ag._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `MERGE django.auth_group AS t ` +
		`USING (` +
		`SELECT @p1 id, @p2 name ` +
		`) AS s ` +
		`ON s.id = t.id ` +
		`WHEN MATCHED THEN ` +
		`UPDATE SET ` +
		`t.name = s.name ` +
		`WHEN NOT MATCHED THEN ` +
		`INSERT (` +
		`name` +
		`) VALUES (` +
		`s.name` +
		`);`
	// run
	logf(sqlstr, ag.ID, ag.Name)
	if _, err := db.ExecContext(ctx, sqlstr, ag.ID, ag.Name); err != nil {
		return err
	}
	// set exists
	ag._exists = true
	return nil
}

// Delete deletes the AuthGroup from the database.
func (ag *AuthGroup) Delete(ctx context.Context, db DB) error {
	switch {
	case !ag._exists: // doesn't exist
		return nil
	case ag._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM django.auth_group ` +
		`WHERE id = @p1`
	// run
	logf(sqlstr, ag.ID)
	if _, err := db.ExecContext(ctx, sqlstr, ag.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	ag._deleted = true
	return nil
}

// AuthGroupByID retrieves a row from 'django.auth_group' as a AuthGroup.
//
// Generated from index 'PK__auth_gro__3213E83F883FA4F3'.
func AuthGroupByID(ctx context.Context, db DB, id int) (*AuthGroup, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name ` +
		`FROM django.auth_group ` +
		`WHERE id = @p1`
	// run
	logf(sqlstr, id)
	ag := AuthGroup{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&ag.ID, &ag.Name); err != nil {
		return nil, logerror(err)
	}
	return &ag, nil
}

// AuthGroupByName retrieves a row from 'django.auth_group' as a AuthGroup.
//
// Generated from index 'UQ__auth_gro__72E12F1B88861B91'.
func AuthGroupByName(ctx context.Context, db DB, name string) (*AuthGroup, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name ` +
		`FROM django.auth_group ` +
		`WHERE name = @p1`
	// run
	logf(sqlstr, name)
	ag := AuthGroup{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, name).Scan(&ag.ID, &ag.Name); err != nil {
		return nil, logerror(err)
	}
	return &ag, nil
}
