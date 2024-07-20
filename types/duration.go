package types

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.6

// Example : P15DT5H0M20S (A duration of 15 days, 5 hours, and 20 seconds)

import (
	"github.com/vareversat/gics/registry"
)

type DurationValue struct {
	V
	Value string
}

func NewDurationValue(value string) DurationValue {
	return DurationValue{
		V:     NewValue(registry.DURATION),
		Value: value,
	}
}

func (bV *DurationValue) GetValue() string {
	return bV.Value
}
