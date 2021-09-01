package sqlserver

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"time"
)

// DjangoMigration represents a row from 'django.django_migrations'.
type DjangoMigration struct {
	ID      int64     `json:"id"`      // id
	App     string    `json:"app"`     // app
	Name    string    `json:"name"`    // name
	Applied time.Time `json:"applied"` // applied
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
		`@p1, @p2, @p3` +
		`); SELECT ID = CONVERT(BIGINT, SCOPE_IDENTITY())`
	// run
	logf(sqlstr, dm.App, dm.Name, dm.Applied)
	rows, err := db.QueryContext(ctx, sqlstr, dm.App, dm.Name, dm.Applied)
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
		`app = @p1, name = @p2, applied = @p3 ` +
		`WHERE id = @p4`
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
	const sqlstr = `MERGE django.django_migrations AS t ` +
		`USING (` +
		`SELECT @p1 id, @p2 app, @p3 name, @p4 applied ` +
		`) AS s ` +
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
		return logerror(err)
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
		`WHERE id = @p1`
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
// Generated from index 'django_migrations_id_pkey'.
func DjangoMigrationByID(ctx context.Context, db DB, id int64) (*DjangoMigration, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, app, name, applied ` +
		`FROM django.django_migrations ` +
		`WHERE id = @p1`
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
