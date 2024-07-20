package types

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.12

// 3 forms :
// - 230000
// - 070000Z
// - TZID=America/New_York:083000

import (
	"time"

	"github.com/vareversat/gics/registries"
)

type TimeValue struct {
	V
	Value *time.Time
}

func NewTimeValue(value *time.Time) *TimeValue {
	return &TimeValue{
		V:     NewValue(registries.Time),
		Value: value,
	}
}

func (tV *TimeValue) GetValue() *time.Time {
	return tV.Value
}
