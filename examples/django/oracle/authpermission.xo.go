// Package oracle contains the types for schema 'django'.
package oracle

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// AuthPermission represents a row from 'django.auth_permission'.
type AuthPermission struct {
	ID            float64        // id
	Name          sql.NullString // name
	ContentTypeID float64        // content_type_id
	Codename      sql.NullString // codename

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
		`:1, :2, :3` +
		`) RETURNING id /*lastInsertId*/ INTO :pk`

	// run query
	XOLog(sqlstr, ap.Name, ap.ContentTypeID, ap.Codename, nil)
	res, err := db.Exec(sqlstr, ap.Name, ap.ContentTypeID, ap.Codename, nil)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	ap.ID = float64(id)
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
		`name = :1, content_type_id = :2, codename = :3` +
		` WHERE id = :4`

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
	const sqlstr = `DELETE FROM django.auth_permission WHERE id = :1`

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
// Generated from foreign key 'cca0f0681577d4054ed10f5351689a'.
func (ap *AuthPermission) DjangoContentType(db XODB) (*DjangoContentType, error) {
	return DjangoContentTypeByID(db, ap.ContentTypeID)
}

// AuthPermissionsByContentTypeID retrieves a row from 'django.auth_permission' as a AuthPermission.
//
// Generated from index 'auth_permission_417f1b1c'.
func AuthPermissionsByContentTypeID(db XODB, contentTypeID float64) ([]*AuthPermission, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, content_type_id, codename ` +
		`FROM django.auth_permission ` +
		`WHERE content_type_id = :1`

	// run query
	XOLog(sqlstr, contentTypeID)
	q, err := db.Query(sqlstr, contentTypeID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AuthPermission{}
	for q.Next() {
		ap := AuthPermission{
			_exists: true,
		}

		// scan
		err = q.Scan(&ap.ID, &ap.Name, &ap.ContentTypeID, &ap.Codename)
		if err != nil {
			return nil, err
		}

		res = append(res, &ap)
	}

	return res, nil
}

// AuthPermissionByContentTypeIDCodename retrieves a row from 'django.auth_permission' as a AuthPermission.
//
// Generated from index 'content_type_id_01ab375a_uniq'.
func AuthPermissionByContentTypeIDCodename(db XODB, contentTypeID float64, codename sql.NullString) (*AuthPermission, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, content_type_id, codename ` +
		`FROM django.auth_permission ` +
		`WHERE content_type_id = :1 AND codename = :2`

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
// Generated from index 'sys_c004959'.
func AuthPermissionByID(db XODB, id float64) (*AuthPermission, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, content_type_id, codename ` +
		`FROM django.auth_permission ` +
		`WHERE id = :1`

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
