package sqlserver

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// AuthGroupPermission represents a row from 'django.auth_group_permissions'.
type AuthGroupPermission struct {
	ID           int64 `json:"id"`            // id
	GroupID      int   `json:"group_id"`      // group_id
	PermissionID int   `json:"permission_id"` // permission_id
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the AuthGroupPermission exists in the database.
func (agp *AuthGroupPermission) Exists() bool {
	return agp._exists
}

// Deleted returns true when the AuthGroupPermission has been marked for deletion from
// the database.
func (agp *AuthGroupPermission) Deleted() bool {
	return agp._deleted
}

// Insert inserts the AuthGroupPermission to the database.
func (agp *AuthGroupPermission) Insert(ctx context.Context, db DB) error {
	switch {
	case agp._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case agp._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO django.auth_group_permissions (` +
		`group_id, permission_id` +
		`) VALUES (` +
		`@p1, @p2` +
		`); SELECT ID = CONVERT(BIGINT, SCOPE_IDENTITY())`
	// run
	logf(sqlstr, agp.GroupID, agp.PermissionID)
	rows, err := db.QueryContext(ctx, sqlstr, agp.GroupID, agp.PermissionID)
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
	agp.ID = int64(id)
	// set exists
	agp._exists = true
	return nil
}

// Update updates a AuthGroupPermission in the database.
func (agp *AuthGroupPermission) Update(ctx context.Context, db DB) error {
	switch {
	case !agp._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case agp._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE django.auth_group_permissions SET ` +
		`group_id = @p1, permission_id = @p2 ` +
		`WHERE id = @p3`
	// run
	logf(sqlstr, agp.GroupID, agp.PermissionID, agp.ID)
	if _, err := db.ExecContext(ctx, sqlstr, agp.GroupID, agp.PermissionID, agp.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the AuthGroupPermission to the database.
func (agp *AuthGroupPermission) Save(ctx context.Context, db DB) error {
	if agp.Exists() {
		return agp.Update(ctx, db)
	}
	return agp.Insert(ctx, db)
}

// Upsert performs an upsert for AuthGroupPermission.
func (agp *AuthGroupPermission) Upsert(ctx context.Context, db DB) error {
	switch {
	case agp._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `MERGE django.auth_group_permissions AS t ` +
		`USING (` +
		`SELECT @p1 id, @p2 group_id, @p3 permission_id ` +
		`) AS s ` +
		`ON s.id = t.id ` +
		`WHEN MATCHED THEN ` +
		`UPDATE SET ` +
		`t.group_id = s.group_id, t.permission_id = s.permission_id ` +
		`WHEN NOT MATCHED THEN ` +
		`INSERT (` +
		`group_id, permission_id` +
		`) VALUES (` +
		`s.group_id, s.permission_id` +
		`);`
	// run
	logf(sqlstr, agp.ID, agp.GroupID, agp.PermissionID)
	if _, err := db.ExecContext(ctx, sqlstr, agp.ID, agp.GroupID, agp.PermissionID); err != nil {
		return err
	}
	// set exists
	agp._exists = true
	return nil
}

// Delete deletes the AuthGroupPermission from the database.
func (agp *AuthGroupPermission) Delete(ctx context.Context, db DB) error {
	switch {
	case !agp._exists: // doesn't exist
		return nil
	case agp._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM django.auth_group_permissions ` +
		`WHERE id = @p1`
	// run
	logf(sqlstr, agp.ID)
	if _, err := db.ExecContext(ctx, sqlstr, agp.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	agp._deleted = true
	return nil
}

// AuthGroupPermissionByID retrieves a row from 'django.auth_group_permissions' as a AuthGroupPermission.
//
// Generated from index 'PK__auth_gro__3213E83F98B3D98E'.
func AuthGroupPermissionByID(ctx context.Context, db DB, id int64) (*AuthGroupPermission, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, group_id, permission_id ` +
		`FROM django.auth_group_permissions ` +
		`WHERE id = @p1`
	// run
	logf(sqlstr, id)
	agp := AuthGroupPermission{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&agp.ID, &agp.GroupID, &agp.PermissionID); err != nil {
		return nil, logerror(err)
	}
	return &agp, nil
}

// AuthGroupPermissionsByGroupID retrieves a row from 'django.auth_group_permissions' as a AuthGroupPermission.
//
// Generated from index 'auth_group_permissions_group_id_b120cbf9'.
func AuthGroupPermissionsByGroupID(ctx context.Context, db DB, groupID int) ([]*AuthGroupPermission, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, group_id, permission_id ` +
		`FROM django.auth_group_permissions ` +
		`WHERE group_id = @p1`
	// run
	logf(sqlstr, groupID)
	rows, err := db.QueryContext(ctx, sqlstr, groupID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*AuthGroupPermission
	for rows.Next() {
		agp := AuthGroupPermission{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&agp.ID, &agp.GroupID, &agp.PermissionID); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &agp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// AuthGroupPermissionByGroupIDPermissionID retrieves a row from 'django.auth_group_permissions' as a AuthGroupPermission.
//
// Generated from index 'auth_group_permissions_group_id_permission_id_0cd325b0_uniq'.
func AuthGroupPermissionByGroupIDPermissionID(ctx context.Context, db DB, groupID, permissionID int) (*AuthGroupPermission, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, group_id, permission_id ` +
		`FROM django.auth_group_permissions ` +
		`WHERE group_id = @p1 AND permission_id = @p2`
	// run
	logf(sqlstr, groupID, permissionID)
	agp := AuthGroupPermission{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, groupID, permissionID).Scan(&agp.ID, &agp.GroupID, &agp.PermissionID); err != nil {
		return nil, logerror(err)
	}
	return &agp, nil
}

// AuthGroupPermissionsByPermissionID retrieves a row from 'django.auth_group_permissions' as a AuthGroupPermission.
//
// Generated from index 'auth_group_permissions_permission_id_84c5c92e'.
func AuthGroupPermissionsByPermissionID(ctx context.Context, db DB, permissionID int) ([]*AuthGroupPermission, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, group_id, permission_id ` +
		`FROM django.auth_group_permissions ` +
		`WHERE permission_id = @p1`
	// run
	logf(sqlstr, permissionID)
	rows, err := db.QueryContext(ctx, sqlstr, permissionID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*AuthGroupPermission
	for rows.Next() {
		agp := AuthGroupPermission{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&agp.ID, &agp.GroupID, &agp.PermissionID); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &agp)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// AuthGroup returns the AuthGroup associated with the AuthGroupPermission's (GroupID).
//
// Generated from foreign key 'auth_group_permissions_group_id_b120cbf9_fk_auth_group_id'.
func (agp *AuthGroupPermission) AuthGroup(ctx context.Context, db DB) (*AuthGroup, error) {
	return AuthGroupByID(ctx, db, agp.GroupID)
}

// AuthPermission returns the AuthPermission associated with the AuthGroupPermission's (PermissionID).
//
// Generated from foreign key 'auth_group_permissions_permission_id_84c5c92e_fk_auth_permission_id'.
func (agp *AuthGroupPermission) AuthPermission(ctx context.Context, db DB) (*AuthPermission, error) {
	return AuthPermissionByID(ctx, db, agp.PermissionID)
}
