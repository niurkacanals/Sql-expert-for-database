package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
)

// UdtPrivilege represents a row from 'information_schema.udt_privileges'.
type UdtPrivilege struct {
	Grantor       sql.NullString `json:"grantor"`        // grantor
	Grantee       sql.NullString `json:"grantee"`        // grantee
	UdtCatalog    sql.NullString `json:"udt_catalog"`    // udt_catalog
	UdtSchema     sql.NullString `json:"udt_schema"`     // udt_schema
	UdtName       sql.NullString `json:"udt_name"`       // udt_name
	PrivilegeType sql.NullString `json:"privilege_type"` // privilege_type
	IsGrantable   sql.NullString `json:"is_grantable"`   // is_grantable
}
