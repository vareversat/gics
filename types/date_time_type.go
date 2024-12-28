package types

import (
	"fmt"
	"time"

	"github.com/vareversat/gics/registries"
)

type DateTimeFormat string

const (
	WithLocalTime DateTimeFormat = "WithLocalTime"
	WithUtcTime   DateTimeFormat = "WithUtcTime"
)

type DateTimeType interface {
	ValueType
	GetValue() time.Time
	GetFormat() DateTimeFormat
}

type dateTimeType struct {
	typeName   registries.ValueTypeRegistry
	typeValue  time.Time
	typeFormat DateTimeFormat
}

func (d *dateTimeType) GetStringValue() string {
	switch d.typeFormat {
	case WithLocalTime:
		return d.typeValue.Format("20060102T150405")
	case WithUtcTime:
		return d.typeValue.Format("20060102T150405Z")
	default:
		return ""
	}
}

func (d *dateTimeType) GetTypeName() string {
	return string(d.typeName)
}

// GetValue get the [time.Time] typed value
func (d *dateTimeType) GetValue() time.Time {
	return d.typeValue
}

func (d *dateTimeType) GetFormat() DateTimeFormat {
	return d.typeFormat
}

// NewDurationValue create a new [registries.Duration] type value. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.6
func NewDateTimeValue(value time.Time) DateTimeType {
	zone, _ := value.Zone()
	format := WithLocalTime
	// Determine the correct format representation
	if zone == "UTC" {
		format = WithUtcTime
	}
	return &dateTimeType{
		typeName:   registries.DateTime,
		typeValue:  value,
		typeFormat: format,
	}
}

// NewDurationValue create a new [registries.Duration] type value array. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.6
func NewDateTimeValues(values []time.Time) []DateTimeType {
	var dateTimeValues []DateTimeType

	for i := 0; i < len(values); i++ {
		dateTimeValues = append(dateTimeValues, NewDateTimeValue(
			values[i],
		))
	}
	return dateTimeValues
}

// NewDurationValue create a new [registries.Duration] type value from a string one (ex: 19980119T020000 or 19980119T070000Z). See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.6
func NewDateTimeValueFromString(value string) (DateTimeType, error) {
	// Try to parse it into UTC format
	date, errUTC := time.Parse("20060102T150405Z", value)
	format := WithUtcTime
	if errUTC != nil {
		// If failed, parse to local time format
		var errLocal error
		date, errLocal = time.Parse("20060102T150405", value)
		format = WithLocalTime

		if errLocal != nil {
			return nil, fmt.Errorf(
				"%s is not a valid value DATE-TIME (either UTC or Local formatted) value: %s",
				value, errLocal.Error(),
			)
		}
	}
	return &dateTimeType{
		typeName:   registries.DateTime,
		typeValue:  date,
		typeFormat: format,
	}, nil
}

// NewDurationValue create a new [registries.Duration] type value from multiple strings (ex: 19980119T020000,20000119T020000). See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.6
func NewDateTimeValueFromStrings(values []string) ([]DateTimeType, error) {
	var dateTimeValues []DateTimeType

	for i := 0; i < len(values); i++ {
		dateTime, err := NewDateTimeValueFromString(values[i])
		if err != nil {
			return nil, fmt.Errorf(
				"unable to format the %d-th element of the list: %s",
				i, err.Error(),
			)
		} else {
			dateTimeValues = append(dateTimeValues, dateTime)
		}
	}
	return dateTimeValues, nil
}
