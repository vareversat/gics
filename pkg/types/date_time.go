package types

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.5

// 2 forms :
// - 19980118T230000
// - 19980119T070000Z

import (
	"time"

	"github.com/vareversat/gics/pkg/registries"
)

type DTFormat string

const (
	WithLocalTime DTFormat = "WithLocalTime"
	WithUtcTime   DTFormat = "WithUtcTime"
)

type DateTimeValue struct {
	V
	Value  time.Time
	Format DTFormat
}

func NewDateTimeValues(values []time.Time, format DTFormat) []DateTimeValue {
	dateTimeValues := make([]DateTimeValue, 0)

	for i := 0; i < len(values); i++ {
		dateTimeValues = append(dateTimeValues, DateTimeValue{
			V:      NewValue(registries.DATETIME),
			Value:  values[i],
			Format: format,
		})
	}
	return dateTimeValues
}

func NewDateTimeValue(value time.Time, format DTFormat) DateTimeValue {
	return DateTimeValue{
		V:      NewValue(registries.DATETIME),
		Value:  value,
		Format: format,
	}
}

func (bV *DateTimeValue) GetValue() string {
	switch bV.Format {
	case WithLocalTime:
		return bV.Value.Format("20060102T150405")
	case WithUtcTime:
		return bV.Value.Format("20060102T150405Z")
	default:
		return "Error: DTFormat is missing"
	}

}
