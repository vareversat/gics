package types

import (
	"github.com/vareversat/gics/registry"
)

type UtcOffsetValue struct {
	V
	Value string
}

func NewUtcOffsetValue(value string) UtcOffsetValue {
	return UtcOffsetValue{
		V:     NewValue(registry.UTCOFFSET),
		Value: value,
	}
}

func (uOV *UtcOffsetValue) GetValue() string {
	return uOV.Value
}
