package oracle

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// AuthUserUserPermission represents a row from 'django.auth_user_user_permissions'.
type AuthUserUserPermission struct {
	ID           int64 `json:"id"`            // id
	UserID       int64 `json:"user_id"`       // user_id
	PermissionID int64 `json:"permission_id"` // permission_id
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the AuthUserUserPermission exists in the database.
func (auup *AuthUserUserPermission) Exists() bool {
	return auup._exists
}

// Deleted returns true when the AuthUserUserPermission has been marked for deletion from
// the database.
func (auup *AuthUserUserPermission) Deleted() bool {
	return auup._deleted
}

// Insert inserts the AuthUserUserPermission to the database.
func (auup *AuthUserUserPermission) Insert(ctx context.Context, db DB) error {
	switch {
	case auup._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case auup._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO django.auth_user_user_permissions (` +
		`user_id, permission_id` +
		`) VALUES (` +
		`:1, :2` +
		`) RETURNING id /*LASTINSERTID*/ INTO :pk`
	// run
	logf(sqlstr, auup.UserID, auup.PermissionID)
	var id int64
	if _, err := db.ExecContext(ctx, sqlstr, auup.UserID, auup.PermissionID, sql.Named("pk", sql.Out{Dest: &id})); err != nil {
		return err
	} // set primary key
	auup.ID = int64(id)
	// set exists
	auup._exists = true
	return nil
}

// Update updates a AuthUserUserPermission in the database.
func (auup *AuthUserUserPermission) Update(ctx context.Context, db DB) error {
	switch {
	case !auup._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case auup._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE django.auth_user_user_permissions SET ` +
		`user_id = :1, permission_id = :2 ` +
		`WHERE id = :3`
	// run
	logf(sqlstr, auup.UserID, auup.PermissionID, auup.ID)
	if _, err := db.ExecContext(ctx, sqlstr, auup.UserID, auup.PermissionID, auup.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the AuthUserUserPermission to the database.
func (auup *AuthUserUserPermission) Save(ctx context.Context, db DB) error {
	if auup.Exists() {
		return auup.Update(ctx, db)
	}
	return auup.Insert(ctx, db)
}

// Upsert performs an upsert for AuthUserUserPermission.
func (auup *AuthUserUserPermission) Upsert(ctx context.Context, db DB) error {
	switch {
	case auup._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `MERGE django.auth_user_user_permissionst ` +
		`USING (` +
		`SELECT :1 id, :2 user_id, :3 permission_id ` +
		`FROM DUAL ) s ` +
		`ON s.id = t.id ` +
		`WHEN MATCHED THEN ` +
		`UPDATE SET ` +
		`t.user_id = s.user_id, t.permission_id = s.permission_id ` +
		`WHEN NOT MATCHED THEN ` +
		`INSERT (` +
		`user_id, permission_id` +
		`) VALUES (` +
		`s.user_id, s.permission_id` +
		`);`
	// run
	logf(sqlstr, auup.ID, auup.UserID, auup.PermissionID)
	if _, err := db.ExecContext(ctx, sqlstr, auup.ID, auup.UserID, auup.PermissionID); err != nil {
		return err
	}
	// set exists
	auup._exists = true
	return nil
}

// Delete deletes the AuthUserUserPermission from the database.
func (auup *AuthUserUserPermission) Delete(ctx context.Context, db DB) error {
	switch {
	case !auup._exists: // doesn't exist
		return nil
	case auup._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM django.auth_user_user_permissions ` +
		`WHERE id = :1`
	// run
	logf(sqlstr, auup.ID)
	if _, err := db.ExecContext(ctx, sqlstr, auup.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	auup._deleted = true
	return nil
}

// AuthUserUserPermissionsByPermissionID retrieves a row from 'django.auth_user_user_permissions' as a AuthUserUserPermission.
//
// Generated from index 'auth_user__permission_1fbb5f2c'.
func AuthUserUserPermissionsByPermissionID(ctx context.Context, db DB, permissionID int64) ([]*AuthUserUserPermission, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, user_id, permission_id ` +
		`FROM django.auth_user_user_permissions ` +
		`WHERE permission_id = :1`
	// run
	logf(sqlstr, permissionID)
	rows, err := db.QueryContext(ctx, sqlstr, permissionID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*AuthUserUserPermission
	for rows.Next() {
		auup := AuthUserUserPermission{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&auup.ID, &auup.UserID, &auup.PermissionID); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &auup)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// AuthUserUserPermissionsByUserID retrieves a row from 'django.auth_user_user_permissions' as a AuthUserUserPermission.
//
// Generated from index 'auth_user__user_id_a95ead1b'.
func AuthUserUserPermissionsByUserID(ctx context.Context, db DB, userID int64) ([]*AuthUserUserPermission, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, user_id, permission_id ` +
		`FROM django.auth_user_user_permissions ` +
		`WHERE user_id = :1`
	// run
	logf(sqlstr, userID)
	rows, err := db.QueryContext(ctx, sqlstr, userID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*AuthUserUserPermission
	for rows.Next() {
		auup := AuthUserUserPermission{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&auup.ID, &auup.UserID, &auup.PermissionID); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &auup)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// AuthUserUserPermissionByUserIDPermissionID retrieves a row from 'django.auth_user_user_permissions' as a AuthUserUserPermission.
//
// Generated from index 'auth_user_user_id_p_14a6b632_u'.
func AuthUserUserPermissionByUserIDPermissionID(ctx context.Context, db DB, userID, permissionID int64) (*AuthUserUserPermission, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, user_id, permission_id ` +
		`FROM django.auth_user_user_permissions ` +
		`WHERE user_id = :1 AND permission_id = :2`
	// run
	logf(sqlstr, userID, permissionID)
	auup := AuthUserUserPermission{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, userID, permissionID).Scan(&auup.ID, &auup.UserID, &auup.PermissionID); err != nil {
		return nil, logerror(err)
	}
	return &auup, nil
}

// AuthUserUserPermissionByID retrieves a row from 'django.auth_user_user_permissions' as a AuthUserUserPermission.
//
// Generated from index 'sys_c0013711'.
func AuthUserUserPermissionByID(ctx context.Context, db DB, id int64) (*AuthUserUserPermission, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, user_id, permission_id ` +
		`FROM django.auth_user_user_permissions ` +
		`WHERE id = :1`
	// run
	logf(sqlstr, id)
	auup := AuthUserUserPermission{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&auup.ID, &auup.UserID, &auup.PermissionID); err != nil {
		return nil, logerror(err)
	}
	return &auup, nil
}

// AuthPermission returns the AuthPermission associated with the AuthUserUserPermission's (PermissionID).
//
// Generated from foreign key 'auth_user_permissio_1fbb5f2c_f'.
func (auup *AuthUserUserPermission) AuthPermission(ctx context.Context, db DB) (*AuthPermission, error) {
	return AuthPermissionByID(ctx, db, auup.PermissionID)
}

// AuthUser returns the AuthUser associated with the AuthUserUserPermission's (UserID).
//
// Generated from foreign key 'auth_user_user_id_a95ead1b_f'.
func (auup *AuthUserUserPermission) AuthUser(ctx context.Context, db DB) (*AuthUser, error) {
	return AuthUserByID(ctx, db, auup.UserID)
}
