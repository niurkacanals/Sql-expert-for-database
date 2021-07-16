package ischema

// Code generated by xo. DO NOT EDIT.

// DomainConstraint represents a row from 'information_schema.domain_constraints'.
type DomainConstraint struct {
	ConstraintCatalog SQLIdentifier `json:"constraint_catalog"` // constraint_catalog
	ConstraintSchema  SQLIdentifier `json:"constraint_schema"`  // constraint_schema
	ConstraintName    SQLIdentifier `json:"constraint_name"`    // constraint_name
	DomainCatalog     SQLIdentifier `json:"domain_catalog"`     // domain_catalog
	DomainSchema      SQLIdentifier `json:"domain_schema"`      // domain_schema
	DomainName        SQLIdentifier `json:"domain_name"`        // domain_name
	IsDeferrable      YesOrNo       `json:"is_deferrable"`      // is_deferrable
	InitiallyDeferred YesOrNo       `json:"initially_deferred"` // initially_deferred
}
