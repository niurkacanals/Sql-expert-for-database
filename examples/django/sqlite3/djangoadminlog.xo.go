// Package sqlite3 contains the types for schema ''.
package sqlite3

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"
)

// DjangoAdminLog represents a row from 'django_admin_log'.
type DjangoAdminLog struct {
	ID            int            `json:"id"`              // id
	ObjectID      sql.NullString `json:"object_id"`       // object_id
	ObjectRepr    string         `json:"object_repr"`     // object_repr
	ActionFlag    string         `json:"action_flag"`     // action_flag
	ChangeMessage string         `json:"change_message"`  // change_message
	ContentTypeID sql.NullInt64  `json:"content_type_id"` // content_type_id
	UserID        int            `json:"user_id"`         // user_id
	ActionTime    time.Time      `json:"action_time"`     // action_time

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the DjangoAdminLog exists in the database.
func (dal *DjangoAdminLog) Exists() bool {
	return dal._exists
}

// Deleted provides information if the DjangoAdminLog has been deleted from the database.
func (dal *DjangoAdminLog) Deleted() bool {
	return dal._deleted
}

// Insert inserts the DjangoAdminLog to the database.
func (dal *DjangoAdminLog) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if dal._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO django_admin_log (` +
		`object_id, object_repr, action_flag, change_message, content_type_id, user_id, action_time` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, dal.ObjectID, dal.ObjectRepr, dal.ActionFlag, dal.ChangeMessage, dal.ContentTypeID, dal.UserID, dal.ActionTime)
	res, err := db.Exec(sqlstr, dal.ObjectID, dal.ObjectRepr, dal.ActionFlag, dal.ChangeMessage, dal.ContentTypeID, dal.UserID, dal.ActionTime)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	dal.ID = int(id)
	dal._exists = true

	return nil
}

// Update updates the DjangoAdminLog in the database.
func (dal *DjangoAdminLog) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !dal._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if dal._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE django_admin_log SET ` +
		`object_id = ?, object_repr = ?, action_flag = ?, change_message = ?, content_type_id = ?, user_id = ?, action_time = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, dal.ObjectID, dal.ObjectRepr, dal.ActionFlag, dal.ChangeMessage, dal.ContentTypeID, dal.UserID, dal.ActionTime, dal.ID)
	_, err = db.Exec(sqlstr, dal.ObjectID, dal.ObjectRepr, dal.ActionFlag, dal.ChangeMessage, dal.ContentTypeID, dal.UserID, dal.ActionTime, dal.ID)
	return err
}

// Save saves the DjangoAdminLog to the database.
func (dal *DjangoAdminLog) Save(db XODB) error {
	if dal.Exists() {
		return dal.Update(db)
	}

	return dal.Insert(db)
}

// Delete deletes the DjangoAdminLog from the database.
func (dal *DjangoAdminLog) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !dal._exists {
		return nil
	}

	// if deleted, bail
	if dal._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM django_admin_log WHERE id = ?`

	// run query
	XOLog(sqlstr, dal.ID)
	_, err = db.Exec(sqlstr, dal.ID)
	if err != nil {
		return err
	}

	// set deleted
	dal._deleted = true

	return nil
}

// DjangoContentType returns the DjangoContentType associated with the DjangoAdminLog's ContentTypeID (content_type_id).
//
// Generated from foreign key 'django_admin_log_content_type_id_fkey'.
func (dal *DjangoAdminLog) DjangoContentType(db XODB) (*DjangoContentType, error) {
	return DjangoContentTypeByID(db, int(dal.ContentTypeID.Int64))
}

// AuthUser returns the AuthUser associated with the DjangoAdminLog's UserID (user_id).
//
// Generated from foreign key 'django_admin_log_user_id_fkey'.
func (dal *DjangoAdminLog) AuthUser(db XODB) (*AuthUser, error) {
	return AuthUserByID(db, dal.UserID)
}

// DjangoAdminLogsByContentTypeID retrieves a row from 'django_admin_log' as a DjangoAdminLog.
//
// Generated from index 'django_admin_log_417f1b1c'.
func DjangoAdminLogsByContentTypeID(db XODB, contentTypeID sql.NullInt64) ([]*DjangoAdminLog, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, object_id, object_repr, action_flag, change_message, content_type_id, user_id, action_time ` +
		`FROM django_admin_log ` +
		`WHERE content_type_id = ?`

	// run query
	XOLog(sqlstr, contentTypeID)
	q, err := db.Query(sqlstr, contentTypeID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*DjangoAdminLog{}
	for q.Next() {
		dal := DjangoAdminLog{
			_exists: true,
		}

		// scan
		err = q.Scan(&dal.ID, &dal.ObjectID, &dal.ObjectRepr, &dal.ActionFlag, &dal.ChangeMessage, &dal.ContentTypeID, &dal.UserID, &dal.ActionTime)
		if err != nil {
			return nil, err
		}

		res = append(res, &dal)
	}

	return res, nil
}

// DjangoAdminLogsByUserID retrieves a row from 'django_admin_log' as a DjangoAdminLog.
//
// Generated from index 'django_admin_log_e8701ad4'.
func DjangoAdminLogsByUserID(db XODB, userID int) ([]*DjangoAdminLog, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, object_id, object_repr, action_flag, change_message, content_type_id, user_id, action_time ` +
		`FROM django_admin_log ` +
		`WHERE user_id = ?`

	// run query
	XOLog(sqlstr, userID)
	q, err := db.Query(sqlstr, userID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*DjangoAdminLog{}
	for q.Next() {
		dal := DjangoAdminLog{
			_exists: true,
		}

		// scan
		err = q.Scan(&dal.ID, &dal.ObjectID, &dal.ObjectRepr, &dal.ActionFlag, &dal.ChangeMessage, &dal.ContentTypeID, &dal.UserID, &dal.ActionTime)
		if err != nil {
			return nil, err
		}

		res = append(res, &dal)
	}

	return res, nil
}

// DjangoAdminLogByID retrieves a row from 'django_admin_log' as a DjangoAdminLog.
//
// Generated from index 'django_admin_log_id_pkey'.
func DjangoAdminLogByID(db XODB, id int) (*DjangoAdminLog, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, object_id, object_repr, action_flag, change_message, content_type_id, user_id, action_time ` +
		`FROM django_admin_log ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	dal := DjangoAdminLog{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&dal.ID, &dal.ObjectID, &dal.ObjectRepr, &dal.ActionFlag, &dal.ChangeMessage, &dal.ContentTypeID, &dal.UserID, &dal.ActionTime)
	if err != nil {
		return nil, err
	}

	return &dal, nil
}
