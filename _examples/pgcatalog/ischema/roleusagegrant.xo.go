package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
)

// RoleUsageGrant represents a row from 'information_schema.role_usage_grants'.
type RoleUsageGrant struct {
	Grantor       sql.NullString `json:"grantor"`        // grantor
	Grantee       sql.NullString `json:"grantee"`        // grantee
	ObjectCatalog sql.NullString `json:"object_catalog"` // object_catalog
	ObjectSchema  sql.NullString `json:"object_schema"`  // object_schema
	ObjectName    sql.NullString `json:"object_name"`    // object_name
	ObjectType    sql.NullString `json:"object_type"`    // object_type
	PrivilegeType sql.NullString `json:"privilege_type"` // privilege_type
	IsGrantable   sql.NullString `json:"is_grantable"`   // is_grantable
}
