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

// parseStringToDateTime take a string value and return a [time.Time] and the corresponding [DateTimeFormat]
func parseStringToDateTime(value string) (time.Time, DateTimeFormat) {
	// Try to parse it into UTC format
	date, err := time.Parse("20060102T150405Z", value)
	format := WithUtcTime
	if err != nil {
		// If failed, parse to local time format
		date, err = time.Parse("20060102T150405", value)
		format = WithLocalTime
	}
	return date, format
}

// NewDurationValue create a new [registries.Duration] type value. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.6
func NewDateTimeValue(value time.Time) DateTimeType {
	zone, _ := value.Zone()
	fmt.Print(zone)
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
func NewDateTimeValueFromString(value string) DateTimeType {
	dateTime, format := parseStringToDateTime(value)
	return &dateTimeType{
		typeName:   registries.DateTime,
		typeValue:  dateTime,
		typeFormat: format,
	}
}

// NewDurationValue create a new [registries.Duration] type value from multiple strings (ex: 19980119T020000,20000119T020000). See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.6
func NewDateTimeValueFromStrings(values []string) []DateTimeType {
	var dateTimeValues []DateTimeType

	for i := 0; i < len(values); i++ {
		dateTimeValues = append(dateTimeValues, NewDateTimeValueFromString(
			values[i],
		))
	}
	return dateTimeValues
}
