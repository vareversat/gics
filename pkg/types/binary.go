package types

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.1

import (
	"github.com/vareversat/gics/pkg/registries"
)

type BinaryValue struct {
	V
	Base64Value string
}

func NewBinaryValue(value string) BinaryValue {
	return BinaryValue{
		V:           NewValue(registries.BINARY),
		Base64Value: value,
	}
}

func (tV *BinaryValue) GetRFCValue() string {
	return tV.Base64Value
}
