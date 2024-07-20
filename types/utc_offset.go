package types

import (
	"github.com/vareversat/gics/registries"
)

type UtcOffsetValue struct {
	V
	Value string
}

func NewUtcOffsetValue(value string) UtcOffsetValue {
	return UtcOffsetValue{
		V:     NewValue(registries.UTCOffset),
		Value: value,
	}
}

func (uOV *UtcOffsetValue) GetValue() string {
	return uOV.Value
}
