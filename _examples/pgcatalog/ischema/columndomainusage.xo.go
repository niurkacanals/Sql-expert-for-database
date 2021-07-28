package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
)

// ColumnDomainUsage represents a row from 'information_schema.column_domain_usage'.
type ColumnDomainUsage struct {
	DomainCatalog sql.NullString `json:"domain_catalog"` // domain_catalog
	DomainSchema  sql.NullString `json:"domain_schema"`  // domain_schema
	DomainName    sql.NullString `json:"domain_name"`    // domain_name
	TableCatalog  sql.NullString `json:"table_catalog"`  // table_catalog
	TableSchema   sql.NullString `json:"table_schema"`   // table_schema
	TableName     sql.NullString `json:"table_name"`     // table_name
	ColumnName    sql.NullString `json:"column_name"`    // column_name
}
