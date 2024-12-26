package types

import (
	"fmt"
	"time"

	"github.com/vareversat/gics/registries"
)

type DateType interface {
	ValueType
	GetValue() time.Time
}

type dateType struct {
	typeName  registries.ValueTypeRegistry
	typeValue time.Time
}

func (d *dateType) GetStringValue() string {
	return d.typeValue.Format("20060102")
}

func (d *dateType) GetTypeName() string {
	return string(d.typeName)
}

// GetValue get the [time.Time] typed value
func (d *dateType) GetValue() time.Time {
	return d.typeValue
}

// NewDateValue create a new [registries.Date] type value. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.4
func NewDateValue(value time.Time) DateType {
	return &dateType{
		typeName:  registries.Date,
		typeValue: value,
	}
}

// NewDateValueFromString create a new [registries.Date] type value from string. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.4
func NewDateValueFromString(value string) (DateType, error) {
	date, err := time.Parse("20060102", value)

	if err != nil {
		return nil, fmt.Errorf("%s is not a valid DATE value: %s", value, err.Error())
	}

	return &dateType{
		typeName:  registries.Date,
		typeValue: date,
	}, nil
}

// NewDateValues create new [registries.Date] type values. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.4
func NewDateValues(values []time.Time) []DateType {
	var dateTimeValues []DateType

	for i := 0; i < len(values); i++ {
		dateTimeValues = append(dateTimeValues, NewDateValue(values[i]))
	}
	return dateTimeValues
}

// NewDateValuesFromString create new [registries.Date] type values from string. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.4s
func NewDateValuesFromString(values []string) ([]DateType, error) {
	var dateValues []DateType

	for i := 0; i < len(values); i++ {
		date, err := NewDateTimeValueFromString(values[i])
		if err != nil {
			return nil, fmt.Errorf(
				"unable to format the %d-th element of the list: %s",
				i, err.Error(),
			)
		} else {
			dateValues = append(dateValues, date)
		}
	}
	return dateValues, nil
}
