// Code generated by 'yaegi extract os/exec'. DO NOT EDIT.

//go:build go1.17
// +build go1.17

package internal

import (
	"os/exec"
	"reflect"
)

func init() {
	Symbols["os/exec/exec"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Command":        reflect.ValueOf(exec.Command),
		"CommandContext": reflect.ValueOf(exec.CommandContext),
		"ErrNotFound":    reflect.ValueOf(&exec.ErrNotFound).Elem(),
		"LookPath":       reflect.ValueOf(exec.LookPath),

		// type definitions
		"Cmd":       reflect.ValueOf((*exec.Cmd)(nil)),
		"Error":     reflect.ValueOf((*exec.Error)(nil)),
		"ExitError": reflect.ValueOf((*exec.ExitError)(nil)),
	}
}
