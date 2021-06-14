package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"github.com/xo/xo/_examples/pgcatalog/pgtypes"
)

// AdministrableRoleAuthorization represents a row from 'information_schema.administrable_role_authorizations'.
type AdministrableRoleAuthorization struct {
	Grantee     pgtypes.SQLIdentifier `json:"grantee"`      // grantee
	RoleName    pgtypes.SQLIdentifier `json:"role_name"`    // role_name
	IsGrantable pgtypes.YesOrNo       `json:"is_grantable"` // is_grantable
}
