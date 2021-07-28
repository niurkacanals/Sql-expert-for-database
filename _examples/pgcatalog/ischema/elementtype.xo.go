package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
)

// ElementType represents a row from 'information_schema.element_types'.
type ElementType struct {
	ObjectCatalog            sql.NullString `json:"object_catalog"`             // object_catalog
	ObjectSchema             sql.NullString `json:"object_schema"`              // object_schema
	ObjectName               sql.NullString `json:"object_name"`                // object_name
	ObjectType               sql.NullString `json:"object_type"`                // object_type
	CollectionTypeIdentifier sql.NullString `json:"collection_type_identifier"` // collection_type_identifier
	DataType                 sql.NullString `json:"data_type"`                  // data_type
	CharacterMaximumLength   sql.NullInt64  `json:"character_maximum_length"`   // character_maximum_length
	CharacterOctetLength     sql.NullInt64  `json:"character_octet_length"`     // character_octet_length
	CharacterSetCatalog      sql.NullString `json:"character_set_catalog"`      // character_set_catalog
	CharacterSetSchema       sql.NullString `json:"character_set_schema"`       // character_set_schema
	CharacterSetName         sql.NullString `json:"character_set_name"`         // character_set_name
	CollationCatalog         sql.NullString `json:"collation_catalog"`          // collation_catalog
	CollationSchema          sql.NullString `json:"collation_schema"`           // collation_schema
	CollationName            sql.NullString `json:"collation_name"`             // collation_name
	NumericPrecision         sql.NullInt64  `json:"numeric_precision"`          // numeric_precision
	NumericPrecisionRadix    sql.NullInt64  `json:"numeric_precision_radix"`    // numeric_precision_radix
	NumericScale             sql.NullInt64  `json:"numeric_scale"`              // numeric_scale
	DatetimePrecision        sql.NullInt64  `json:"datetime_precision"`         // datetime_precision
	IntervalType             sql.NullString `json:"interval_type"`              // interval_type
	IntervalPrecision        sql.NullInt64  `json:"interval_precision"`         // interval_precision
	DomainDefault            sql.NullString `json:"domain_default"`             // domain_default
	UdtCatalog               sql.NullString `json:"udt_catalog"`                // udt_catalog
	UdtSchema                sql.NullString `json:"udt_schema"`                 // udt_schema
	UdtName                  sql.NullString `json:"udt_name"`                   // udt_name
	ScopeCatalog             sql.NullString `json:"scope_catalog"`              // scope_catalog
	ScopeSchema              sql.NullString `json:"scope_schema"`               // scope_schema
	ScopeName                sql.NullString `json:"scope_name"`                 // scope_name
	MaximumCardinality       sql.NullInt64  `json:"maximum_cardinality"`        // maximum_cardinality
	DtdIdentifier            sql.NullString `json:"dtd_identifier"`             // dtd_identifier
}
