package ischema

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
)

// CheckConstraint represents a row from 'information_schema.check_constraints'.
type CheckConstraint struct {
	ConstraintCatalog sql.NullString `json:"constraint_catalog"` // constraint_catalog
	ConstraintSchema  sql.NullString `json:"constraint_schema"`  // constraint_schema
	ConstraintName    sql.NullString `json:"constraint_name"`    // constraint_name
	CheckClause       sql.NullString `json:"check_clause"`       // check_clause
}
