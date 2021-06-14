package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"github.com/xo/xo/_examples/pgcatalog/pgtypes"
)

// ApplicableRole represents a row from 'information_schema.applicable_roles'.
type ApplicableRole struct {
	Grantee     pgtypes.SQLIdentifier `json:"grantee"`      // grantee
	RoleName    pgtypes.SQLIdentifier `json:"role_name"`    // role_name
	IsGrantable pgtypes.YesOrNo       `json:"is_grantable"` // is_grantable
}
