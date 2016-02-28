// Package mysql contains the types for schema 'django'.
package mysql

// GENERATED BY XO. DO NOT EDIT.

import "errors"

// AuthPermission represents a row from 'django.auth_permission'.
type AuthPermission struct {
	ID            int    // id
	Name          string // name
	ContentTypeID int    // content_type_id
	Codename      string // codename

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the AuthPermission exists in the database.
func (ap *AuthPermission) Exists() bool {
	return ap._exists
}

// Deleted provides information if the AuthPermission has been deleted from the database.
func (ap *AuthPermission) Deleted() bool {
	return ap._deleted
}

// Insert inserts the AuthPermission to the database.
func (ap *AuthPermission) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if ap._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO django.auth_permission (` +
		`name, content_type_id, codename` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, ap.Name, ap.ContentTypeID, ap.Codename)
	res, err := db.Exec(sqlstr, ap.Name, ap.ContentTypeID, ap.Codename)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	ap.ID = int(id)
	ap._exists = true

	return nil
}

// Update updates the AuthPermission in the database.
func (ap *AuthPermission) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ap._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if ap._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE django.auth_permission SET ` +
		`name = ?, content_type_id = ?, codename = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, ap.Name, ap.ContentTypeID, ap.Codename, ap.ID)
	_, err = db.Exec(sqlstr, ap.Name, ap.ContentTypeID, ap.Codename, ap.ID)
	return err
}

// Save saves the AuthPermission to the database.
func (ap *AuthPermission) Save(db XODB) error {
	if ap.Exists() {
		return ap.Update(db)
	}

	return ap.Insert(db)
}

// Delete deletes the AuthPermission from the database.
func (ap *AuthPermission) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !ap._exists {
		return nil
	}

	// if deleted, bail
	if ap._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM django.auth_permission WHERE id = ?`

	// run query
	XOLog(sqlstr, ap.ID)
	_, err = db.Exec(sqlstr, ap.ID)
	if err != nil {
		return err
	}

	// set deleted
	ap._deleted = true

	return nil
}

// DjangoContentType returns the DjangoContentType associated with the AuthPermission's ContentTypeID (content_type_id).
//
// Generated from foreign key 'auth_permissi_content_type_id_2f476e4b_fk_django_content_type_id'.
func (ap *AuthPermission) DjangoContentType(db XODB) (*DjangoContentType, error) {
	return DjangoContentTypeByID(db, ap.ContentTypeID)
}

// AuthPermissionByContentTypeIDCodename retrieves a row from 'django.auth_permission' as a AuthPermission.
//
// Generated from index 'auth_permission_content_type_id_01ab375a_uniq'.
func AuthPermissionByContentTypeIDCodename(db XODB, contentTypeID int, codename string) (*AuthPermission, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, content_type_id, codename ` +
		`FROM django.auth_permission ` +
		`WHERE content_type_id = ? AND codename = ?`

	// run query
	XOLog(sqlstr, contentTypeID, codename)
	ap := AuthPermission{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, contentTypeID, codename).Scan(&ap.ID, &ap.Name, &ap.ContentTypeID, &ap.Codename)
	if err != nil {
		return nil, err
	}

	return &ap, nil
}

// AuthPermissionByID retrieves a row from 'django.auth_permission' as a AuthPermission.
//
// Generated from index 'auth_permission_id_pkey'.
func AuthPermissionByID(db XODB, id int) (*AuthPermission, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, content_type_id, codename ` +
		`FROM django.auth_permission ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	ap := AuthPermission{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&ap.ID, &ap.Name, &ap.ContentTypeID, &ap.Codename)
	if err != nil {
		return nil, err
	}

	return &ap, nil
}
