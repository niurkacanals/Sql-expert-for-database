package sqlserver

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// DjangoContentType represents a row from 'django.django_content_type'.
type DjangoContentType struct {
	ID       int    `json:"id"`        // id
	AppLabel string `json:"app_label"` // app_label
	Model    string `json:"model"`     // model
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the DjangoContentType exists in the database.
func (dct *DjangoContentType) Exists() bool {
	return dct._exists
}

// Deleted returns true when the DjangoContentType has been marked for deletion from
// the database.
func (dct *DjangoContentType) Deleted() bool {
	return dct._deleted
}

// Insert inserts the DjangoContentType to the database.
func (dct *DjangoContentType) Insert(ctx context.Context, db DB) error {
	switch {
	case dct._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case dct._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO django.django_content_type (` +
		`app_label, model` +
		`) VALUES (` +
		`@p1, @p2` +
		`); SELECT ID = CONVERT(BIGINT, SCOPE_IDENTITY())`
	// run
	logf(sqlstr, dct.AppLabel, dct.Model)
	rows, err := db.QueryContext(ctx, sqlstr, dct.AppLabel, dct.Model)
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
	dct.ID = int(id)
	// set exists
	dct._exists = true
	return nil
}

// Update updates a DjangoContentType in the database.
func (dct *DjangoContentType) Update(ctx context.Context, db DB) error {
	switch {
	case !dct._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case dct._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with primary key
	const sqlstr = `UPDATE django.django_content_type SET ` +
		`app_label = @p1, model = @p2 ` +
		`WHERE id = @p3`
	// run
	logf(sqlstr, dct.AppLabel, dct.Model, dct.ID)
	if _, err := db.ExecContext(ctx, sqlstr, dct.AppLabel, dct.Model, dct.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the DjangoContentType to the database.
func (dct *DjangoContentType) Save(ctx context.Context, db DB) error {
	if dct.Exists() {
		return dct.Update(ctx, db)
	}
	return dct.Insert(ctx, db)
}

// Upsert performs an upsert for DjangoContentType.
func (dct *DjangoContentType) Upsert(ctx context.Context, db DB) error {
	switch {
	case dct._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `MERGE django.django_content_type AS t ` +
		`USING (` +
		`SELECT @p1 id, @p2 app_label, @p3 model ` +
		`) AS s ` +
		`ON s.id = t.id ` +
		`WHEN MATCHED THEN ` +
		`UPDATE SET ` +
		`t.app_label = s.app_label, t.model = s.model ` +
		`WHEN NOT MATCHED THEN ` +
		`INSERT (` +
		`app_label, model` +
		`) VALUES (` +
		`s.app_label, s.model` +
		`);`
	// run
	logf(sqlstr, dct.ID, dct.AppLabel, dct.Model)
	if _, err := db.ExecContext(ctx, sqlstr, dct.ID, dct.AppLabel, dct.Model); err != nil {
		return err
	}
	// set exists
	dct._exists = true
	return nil
}

// Delete deletes the DjangoContentType from the database.
func (dct *DjangoContentType) Delete(ctx context.Context, db DB) error {
	switch {
	case !dct._exists: // doesn't exist
		return nil
	case dct._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM django.django_content_type ` +
		`WHERE id = @p1`
	// run
	logf(sqlstr, dct.ID)
	if _, err := db.ExecContext(ctx, sqlstr, dct.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	dct._deleted = true
	return nil
}

// DjangoContentTypeByID retrieves a row from 'django.django_content_type' as a DjangoContentType.
//
// Generated from index 'PK__django_c__3213E83F8BDC15B8'.
func DjangoContentTypeByID(ctx context.Context, db DB, id int) (*DjangoContentType, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, app_label, model ` +
		`FROM django.django_content_type ` +
		`WHERE id = @p1`
	// run
	logf(sqlstr, id)
	dct := DjangoContentType{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&dct.ID, &dct.AppLabel, &dct.Model); err != nil {
		return nil, logerror(err)
	}
	return &dct, nil
}

// DjangoContentTypeByAppLabelModel retrieves a row from 'django.django_content_type' as a DjangoContentType.
//
// Generated from index 'django_content_type_app_label_model_76bd3d3b_uniq'.
func DjangoContentTypeByAppLabelModel(ctx context.Context, db DB, appLabel, model string) (*DjangoContentType, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, app_label, model ` +
		`FROM django.django_content_type ` +
		`WHERE app_label = @p1 AND model = @p2`
	// run
	logf(sqlstr, appLabel, model)
	dct := DjangoContentType{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, appLabel, model).Scan(&dct.ID, &dct.AppLabel, &dct.Model); err != nil {
		return nil, logerror(err)
	}
	return &dct, nil
}
