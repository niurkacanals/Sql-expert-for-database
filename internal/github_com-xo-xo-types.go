// Code generated by 'yaegi extract github.com/xo/xo/types'. DO NOT EDIT.

package internal

import (
	"github.com/xo/xo/types"
	"reflect"
)

func init() {
	Symbols["github.com/xo/xo/types/types"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"DbKey":          reflect.ValueOf(types.DbKey),
		"DriverDbSchema": reflect.ValueOf(types.DriverDbSchema),
		"DriverKey":      reflect.ValueOf(types.DriverKey),
		"NewValue":       reflect.ValueOf(types.NewValue),
		"Out":            reflect.ValueOf(types.Out),
		"OutKey":         reflect.ValueOf(types.OutKey),
		"ParseType":      reflect.ValueOf(types.ParseType),
		"SchemaKey":      reflect.ValueOf(types.SchemaKey),
		"Single":         reflect.ValueOf(types.Single),
		"SingleKey":      reflect.ValueOf(types.SingleKey),

		// type definitions
		"ContextKey":   reflect.ValueOf((*types.ContextKey)(nil)),
		"Enum":         reflect.ValueOf((*types.Enum)(nil)),
		"Field":        reflect.ValueOf((*types.Field)(nil)),
		"Flag":         reflect.ValueOf((*types.Flag)(nil)),
		"FlagSet":      reflect.ValueOf((*types.FlagSet)(nil)),
		"ForeignKey":   reflect.ValueOf((*types.ForeignKey)(nil)),
		"Index":        reflect.ValueOf((*types.Index)(nil)),
		"Proc":         reflect.ValueOf((*types.Proc)(nil)),
		"Query":        reflect.ValueOf((*types.Query)(nil)),
		"Schema":       reflect.ValueOf((*types.Schema)(nil)),
		"Set":          reflect.ValueOf((*types.Set)(nil)),
		"Table":        reflect.ValueOf((*types.Table)(nil)),
		"Template":     reflect.ValueOf((*types.Template)(nil)),
		"TemplateType": reflect.ValueOf((*types.TemplateType)(nil)),
		"Type":         reflect.ValueOf((*types.Type)(nil)),
		"Value":        reflect.ValueOf((*types.Value)(nil)),
	}
}
