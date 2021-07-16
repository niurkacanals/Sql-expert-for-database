package ischema

// Code generated by xo. DO NOT EDIT.

// UserMapping represents a row from 'information_schema.user_mappings'.
type UserMapping struct {
	AuthorizationIdentifier SQLIdentifier `json:"authorization_identifier"` // authorization_identifier
	ForeignServerCatalog    SQLIdentifier `json:"foreign_server_catalog"`   // foreign_server_catalog
	ForeignServerName       SQLIdentifier `json:"foreign_server_name"`      // foreign_server_name
}
