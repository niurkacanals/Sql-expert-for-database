package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
)

// ForeignDataWrapperOption represents a row from 'information_schema.foreign_data_wrapper_options'.
type ForeignDataWrapperOption struct {
	ForeignDataWrapperCatalog sql.NullString `json:"foreign_data_wrapper_catalog"` // foreign_data_wrapper_catalog
	ForeignDataWrapperName    sql.NullString `json:"foreign_data_wrapper_name"`    // foreign_data_wrapper_name
	OptionName                sql.NullString `json:"option_name"`                  // option_name
	OptionValue               sql.NullString `json:"option_value"`                 // option_value
}
