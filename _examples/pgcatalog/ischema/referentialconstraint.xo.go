package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
)

// ReferentialConstraint represents a row from 'information_schema.referential_constraints'.
type ReferentialConstraint struct {
	ConstraintCatalog       sql.NullString `json:"constraint_catalog"`        // constraint_catalog
	ConstraintSchema        sql.NullString `json:"constraint_schema"`         // constraint_schema
	ConstraintName          sql.NullString `json:"constraint_name"`           // constraint_name
	UniqueConstraintCatalog sql.NullString `json:"unique_constraint_catalog"` // unique_constraint_catalog
	UniqueConstraintSchema  sql.NullString `json:"unique_constraint_schema"`  // unique_constraint_schema
	UniqueConstraintName    sql.NullString `json:"unique_constraint_name"`    // unique_constraint_name
	MatchOption             sql.NullString `json:"match_option"`              // match_option
	UpdateRule              sql.NullString `json:"update_rule"`               // update_rule
	DeleteRule              sql.NullString `json:"delete_rule"`               // delete_rule
}
