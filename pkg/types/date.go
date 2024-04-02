package types

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.4

// Example : 19970714

import (
	"time"

	"github.com/vareversat/gics/pkg/registries"
)

type DateValue struct {
	V
	Value     time.Time
	IANAToken registries.Type
}

func NewDateValue(value time.Time) DateValue {
	return DateValue{
		V:     NewValue(registries.DATE),
		Value: value,
	}
}

func NewDateValues(values []time.Time) []DateValue {
	dateTimeValues := make([]DateValue, 0)

	for i := 0; i < len(values); i++ {
		dateTimeValues = append(dateTimeValues, DateValue{
			V:     NewValue(registries.DATE),
			Value: values[i],
		})
	}
	return dateTimeValues
}

func (bV *DateValue) GetValue() string {
	return bV.Value.Format("20060102")
}
