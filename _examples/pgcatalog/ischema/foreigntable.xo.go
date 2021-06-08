package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"github.com/xo/xo/_examples/pgcatalog/pgtypes"
)

// ForeignTable represents a row from 'information_schema.foreign_tables'.
type ForeignTable struct {
	ForeignTableCatalog  pgtypes.SQLIdentifier `json:"foreign_table_catalog"`  // foreign_table_catalog
	ForeignTableSchema   pgtypes.SQLIdentifier `json:"foreign_table_schema"`   // foreign_table_schema
	ForeignTableName     pgtypes.SQLIdentifier `json:"foreign_table_name"`     // foreign_table_name
	ForeignServerCatalog pgtypes.SQLIdentifier `json:"foreign_server_catalog"` // foreign_server_catalog
	ForeignServerName    pgtypes.SQLIdentifier `json:"foreign_server_name"`    // foreign_server_name

}
