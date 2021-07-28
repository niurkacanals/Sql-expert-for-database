package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
)

// AdministrableRoleAuthorization represents a row from 'information_schema.administrable_role_authorizations'.
type AdministrableRoleAuthorization struct {
	Grantee     sql.NullString `json:"grantee"`      // grantee
	RoleName    sql.NullString `json:"role_name"`    // role_name
	IsGrantable sql.NullString `json:"is_grantable"` // is_grantable
}
