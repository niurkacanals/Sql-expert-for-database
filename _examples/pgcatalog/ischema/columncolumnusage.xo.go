package ischema

// Code generated by xo. DO NOT EDIT.

// ColumnColumnUsage represents a row from 'information_schema.column_column_usage'.
type ColumnColumnUsage struct {
	TableCatalog    SQLIdentifier `json:"table_catalog"`    // table_catalog
	TableSchema     SQLIdentifier `json:"table_schema"`     // table_schema
	TableName       SQLIdentifier `json:"table_name"`       // table_name
	ColumnName      SQLIdentifier `json:"column_name"`      // column_name
	DependentColumn SQLIdentifier `json:"dependent_column"` // dependent_column
}
