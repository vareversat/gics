package types

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.4

// Example : 19970714

import (
	"fmt"
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

func NewDateValueFromString(value string) DateValue {
	return DateValue{
		V:     NewValue(registries.DATE),
		Value: parseStringToDateTime(value),
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

func NewDateValuesFromString(values []string) []DateValue {
	dateTimeValues := make([]DateValue, 0)

	for i := 0; i < len(values); i++ {
		dateTimeValues = append(dateTimeValues, DateValue{
			V:     NewValue(registries.DATE),
			Value: parseStringToDate(values[i]),
		})
	}
	return dateTimeValues
}

func parseStringToDate(value string) time.Time {
	date, err := time.Parse("20060102", value)
	if err != nil {
		fmt.Println(err)
	}
	return date
}

func (bV *DateValue) GetValue() string {
	return bV.Value.Format("20060102")
}
