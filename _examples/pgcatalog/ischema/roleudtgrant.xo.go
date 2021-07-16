package ischema

// Code generated by xo. DO NOT EDIT.

// RoleUdtGrant represents a row from 'information_schema.role_udt_grants'.
type RoleUdtGrant struct {
	Grantor       SQLIdentifier `json:"grantor"`        // grantor
	Grantee       SQLIdentifier `json:"grantee"`        // grantee
	UdtCatalog    SQLIdentifier `json:"udt_catalog"`    // udt_catalog
	UdtSchema     SQLIdentifier `json:"udt_schema"`     // udt_schema
	UdtName       SQLIdentifier `json:"udt_name"`       // udt_name
	PrivilegeType CharacterData `json:"privilege_type"` // privilege_type
	IsGrantable   YesOrNo       `json:"is_grantable"`   // is_grantable
}
