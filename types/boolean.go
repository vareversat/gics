package types

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.2

import (
	"github.com/vareversat/gics/registry"
)

type BooleanValue struct {
	V
	Value bool
}

func NewBooleanValue(value bool) BooleanValue {
	return BooleanValue{
		V:     NewValue(registry.Boolean),
		Value: value,
	}
}

func (bV *BooleanValue) GetValue() string {
	if bV.Value {
		return "TRUE"
	} else {
		return "FALSE"
	}
}
