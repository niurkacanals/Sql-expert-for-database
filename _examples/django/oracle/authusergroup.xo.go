package oracle

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// AuthUserGroup represents a row from 'django.auth_user_groups'.
type AuthUserGroup struct {
	ID      int64 `json:"id"`       // id
	UserID  int64 `json:"user_id"`  // user_id
	GroupID int64 `json:"group_id"` // group_id
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the AuthUserGroup exists in the database.
func (aug *AuthUserGroup) Exists() bool {
	return aug._exists
}

// Deleted returns true when the AuthUserGroup has been marked for deletion from
// the database.
func (aug *AuthUserGroup) Deleted() bool {
	return aug._deleted
}

// Insert inserts the AuthUserGroup to the database.
func (aug *AuthUserGroup) Insert(ctx context.Context, db DB) error {
	switch {
	case aug._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case aug._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO django.auth_user_groups (` +
		`user_id, group_id` +
		`) VALUES (` +
		`:1, :2` +
		`) RETURNING id /*LASTINSERTID*/ INTO :pk`
	// run
	logf(sqlstr, aug.UserID, aug.GroupID)
	var id int64
	if _, err := db.ExecContext(ctx, sqlstr, aug.UserID, aug.GroupID, sql.Named("pk", sql.Out{Dest: &id})); err != nil {
		return err
	} // set primary key
	aug.ID = int64(id)
	// set exists
	aug._exists = true
	return nil
}

// Update updates a AuthUserGroup in the database.
func (aug *AuthUserGroup) Update(ctx context.Context, db DB) error {
	switch {
	case !aug._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case aug._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE django.auth_user_groups SET ` +
		`user_id = :1, group_id = :2 ` +
		`WHERE id = :3`
	// run
	logf(sqlstr, aug.UserID, aug.GroupID, aug.ID)
	if _, err := db.ExecContext(ctx, sqlstr, aug.UserID, aug.GroupID, aug.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the AuthUserGroup to the database.
func (aug *AuthUserGroup) Save(ctx context.Context, db DB) error {
	if aug.Exists() {
		return aug.Update(ctx, db)
	}
	return aug.Insert(ctx, db)
}

// Upsert performs an upsert for AuthUserGroup.
func (aug *AuthUserGroup) Upsert(ctx context.Context, db DB) error {
	switch {
	case aug._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `MERGE django.auth_user_groupst ` +
		`USING (` +
		`SELECT :1 id, :2 user_id, :3 group_id ` +
		`FROM DUAL ) s ` +
		`ON s.id = t.id ` +
		`WHEN MATCHED THEN ` +
		`UPDATE SET ` +
		`t.user_id = s.user_id, t.group_id = s.group_id ` +
		`WHEN NOT MATCHED THEN ` +
		`INSERT (` +
		`user_id, group_id` +
		`) VALUES (` +
		`s.user_id, s.group_id` +
		`);`
	// run
	logf(sqlstr, aug.ID, aug.UserID, aug.GroupID)
	if _, err := db.ExecContext(ctx, sqlstr, aug.ID, aug.UserID, aug.GroupID); err != nil {
		return err
	}
	// set exists
	aug._exists = true
	return nil
}

// Delete deletes the AuthUserGroup from the database.
func (aug *AuthUserGroup) Delete(ctx context.Context, db DB) error {
	switch {
	case !aug._exists: // doesn't exist
		return nil
	case aug._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM django.auth_user_groups ` +
		`WHERE id = :1`
	// run
	logf(sqlstr, aug.ID)
	if _, err := db.ExecContext(ctx, sqlstr, aug.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	aug._deleted = true
	return nil
}

// AuthUserGroupsByGroupID retrieves a row from 'django.auth_user_groups' as a AuthUserGroup.
//
// Generated from index 'auth_user__group_id_97559544'.
func AuthUserGroupsByGroupID(ctx context.Context, db DB, groupID int64) ([]*AuthUserGroup, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, user_id, group_id ` +
		`FROM django.auth_user_groups ` +
		`WHERE group_id = :1`
	// run
	logf(sqlstr, groupID)
	rows, err := db.QueryContext(ctx, sqlstr, groupID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*AuthUserGroup
	for rows.Next() {
		aug := AuthUserGroup{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&aug.ID, &aug.UserID, &aug.GroupID); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &aug)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// AuthUserGroupsByUserID retrieves a row from 'django.auth_user_groups' as a AuthUserGroup.
//
// Generated from index 'auth_user__user_id_6a12ed8b'.
func AuthUserGroupsByUserID(ctx context.Context, db DB, userID int64) ([]*AuthUserGroup, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, user_id, group_id ` +
		`FROM django.auth_user_groups ` +
		`WHERE user_id = :1`
	// run
	logf(sqlstr, userID)
	rows, err := db.QueryContext(ctx, sqlstr, userID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*AuthUserGroup
	for rows.Next() {
		aug := AuthUserGroup{
			_exists: true,
		}
		// scan
		if err := rows.Scan(&aug.ID, &aug.UserID, &aug.GroupID); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &aug)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// AuthUserGroupByUserIDGroupID retrieves a row from 'django.auth_user_groups' as a AuthUserGroup.
//
// Generated from index 'auth_user_user_id_g_94350c0c_u'.
func AuthUserGroupByUserIDGroupID(ctx context.Context, db DB, userID, groupID int64) (*AuthUserGroup, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, user_id, group_id ` +
		`FROM django.auth_user_groups ` +
		`WHERE user_id = :1 AND group_id = :2`
	// run
	logf(sqlstr, userID, groupID)
	aug := AuthUserGroup{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, userID, groupID).Scan(&aug.ID, &aug.UserID, &aug.GroupID); err != nil {
		return nil, logerror(err)
	}
	return &aug, nil
}

// AuthUserGroupByID retrieves a row from 'django.auth_user_groups' as a AuthUserGroup.
//
// Generated from index 'sys_c0013446'.
func AuthUserGroupByID(ctx context.Context, db DB, id int64) (*AuthUserGroup, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, user_id, group_id ` +
		`FROM django.auth_user_groups ` +
		`WHERE id = :1`
	// run
	logf(sqlstr, id)
	aug := AuthUserGroup{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&aug.ID, &aug.UserID, &aug.GroupID); err != nil {
		return nil, logerror(err)
	}
	return &aug, nil
}

// AuthGroup returns the AuthGroup associated with the AuthUserGroup's (GroupID).
//
// Generated from foreign key 'auth_user_group_id_97559544_f'.
func (aug *AuthUserGroup) AuthGroup(ctx context.Context, db DB) (*AuthGroup, error) {
	return AuthGroupByID(ctx, db, aug.GroupID)
}

// AuthUser returns the AuthUser associated with the AuthUserGroup's (UserID).
//
// Generated from foreign key 'auth_user_user_id_6a12ed8b_f'.
func (aug *AuthUserGroup) AuthUser(ctx context.Context, db DB) (*AuthUser, error) {
	return AuthUserByID(ctx, db, aug.UserID)
}
