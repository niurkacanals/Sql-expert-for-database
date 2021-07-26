package oracle

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"time"
)

// DjangoMigration represents a row from 'django.django_migrations'.
type DjangoMigration struct {
	ID      int64          `json:"id"`      // id
	App     sql.NullString `json:"app"`     // app
	Name    sql.NullString `json:"name"`    // name
	Applied time.Time      `json:"applied"` // applied
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the DjangoMigration exists in the database.
func (dm *DjangoMigration) Exists() bool {
	return dm._exists
}

// Deleted returns true when the DjangoMigration has been marked for deletion from
// the database.
func (dm *DjangoMigration) Deleted() bool {
	return dm._deleted
}

// Insert inserts the DjangoMigration to the database.
func (dm *DjangoMigration) Insert(ctx context.Context, db DB) error {
	switch {
	case dm._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case dm._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO django.django_migrations (` +
		`app, name, applied` +
		`) VALUES (` +
		`:1, :2, :3` +
		`) RETURNING id /*LASTINSERTID*/ INTO :pk`
	// run
	logf(sqlstr, dm.App, dm.Name, dm.Applied)
	var id int64
	if _, err := db.ExecContext(ctx, sqlstr, dm.App, dm.Name, dm.Applied, sql.Named("pk", sql.Out{Dest: &id})); err != nil {
		return err
	} // set primary key
	dm.ID = int64(id)
	// set exists
	dm._exists = true
	return nil
}

// Update updates a DjangoMigration in the database.
func (dm *DjangoMigration) Update(ctx context.Context, db DB) error {
	switch {
	case !dm._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case dm._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE django.django_migrations SET ` +
		`app = :1, name = :2, applied = :3 ` +
		`WHERE id = :4`
	// run
	logf(sqlstr, dm.App, dm.Name, dm.Applied, dm.ID)
	if _, err := db.ExecContext(ctx, sqlstr, dm.App, dm.Name, dm.Applied, dm.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the DjangoMigration to the database.
func (dm *DjangoMigration) Save(ctx context.Context, db DB) error {
	if dm.Exists() {
		return dm.Update(ctx, db)
	}
	return dm.Insert(ctx, db)
}

// Upsert performs an upsert for DjangoMigration.
func (dm *DjangoMigration) Upsert(ctx context.Context, db DB) error {
	switch {
	case dm._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `MERGE django.django_migrationst ` +
		`USING (` +
		`SELECT :1 id, :2 app, :3 name, :4 applied ` +
		`FROM DUAL ) s ` +
		`ON s.id = t.id ` +
		`WHEN MATCHED THEN ` +
		`UPDATE SET ` +
		`t.app = s.app, t.name = s.name, t.applied = s.applied ` +
		`WHEN NOT MATCHED THEN ` +
		`INSERT (` +
		`app, name, applied` +
		`) VALUES (` +
		`s.app, s.name, s.applied` +
		`);`
	// run
	logf(sqlstr, dm.ID, dm.App, dm.Name, dm.Applied)
	if _, err := db.ExecContext(ctx, sqlstr, dm.ID, dm.App, dm.Name, dm.Applied); err != nil {
		return err
	}
	// set exists
	dm._exists = true
	return nil
}

// Delete deletes the DjangoMigration from the database.
func (dm *DjangoMigration) Delete(ctx context.Context, db DB) error {
	switch {
	case !dm._exists: // doesn't exist
		return nil
	case dm._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM django.django_migrations ` +
		`WHERE id = :1`
	// run
	logf(sqlstr, dm.ID)
	if _, err := db.ExecContext(ctx, sqlstr, dm.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	dm._deleted = true
	return nil
}

// DjangoMigrationByID retrieves a row from 'django.django_migrations' as a DjangoMigration.
//
// Generated from index 'sys_c0014032'.
func DjangoMigrationByID(ctx context.Context, db DB, id int64) (*DjangoMigration, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, app, name, applied ` +
		`FROM django.django_migrations ` +
		`WHERE id = :1`
	// run
	logf(sqlstr, id)
	dm := DjangoMigration{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&dm.ID, &dm.App, &dm.Name, &dm.Applied); err != nil {
		return nil, logerror(err)
	}
	return &dm, nil
}
