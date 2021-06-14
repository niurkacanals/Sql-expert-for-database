package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"github.com/xo/xo/_examples/pgcatalog/pgtypes"
)

// Schematum represents a row from 'information_schema.schemata'.
type Schematum struct {
	CatalogName                pgtypes.SQLIdentifier `json:"catalog_name"`                  // catalog_name
	SchemaName                 pgtypes.SQLIdentifier `json:"schema_name"`                   // schema_name
	SchemaOwner                pgtypes.SQLIdentifier `json:"schema_owner"`                  // schema_owner
	DefaultCharacterSetCatalog pgtypes.SQLIdentifier `json:"default_character_set_catalog"` // default_character_set_catalog
	DefaultCharacterSetSchema  pgtypes.SQLIdentifier `json:"default_character_set_schema"`  // default_character_set_schema
	DefaultCharacterSetName    pgtypes.SQLIdentifier `json:"default_character_set_name"`    // default_character_set_name
	SQLPath                    pgtypes.CharacterData `json:"sql_path"`                      // sql_path
}
