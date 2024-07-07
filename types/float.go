package types

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.7

import (
	"github.com/vareversat/gics/registries"
)

type FloatValue struct {
	V
	Value float32
}

func NewFloatValue(value float32) FloatValue {
	return FloatValue{
		V:     NewValue(registries.FLOAT),
		Value: value,
	}
}

func (fV *FloatValue) GetValue() float32 {
	return fV.Value
}
