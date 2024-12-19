package types

import (
	"fmt"
	"time"

	"github.com/vareversat/gics/registries"
)

type TimeFormat string

const (
	TimeWithLocalTime TimeFormat = "TimeWithLocalTime"
	TimeWithUtcTime   TimeFormat = "TimeWithUtcTime"
)

type TimeType interface {
	ValueType
	GetValue() time.Time
	GetFormat() TimeFormat
}

type timeType struct {
	typeName   registries.ValueTypeRegistry
	typeValue  time.Time
	typeFormat TimeFormat
}

// GetStringValue return the string representation according to the correct format
func (d *timeType) GetStringValue() string {
	switch d.typeFormat {
	case TimeWithLocalTime:
		return d.typeValue.Format("150405")
	case TimeWithUtcTime:
		return d.typeValue.Format("150405Z")
	default:
		return ""
	}
}

func (d *timeType) GetTypeName() string {
	return string(d.typeName)
}

// GetValue get the time.Time typed value
func (f *timeType) GetValue() time.Time {
	return f.typeValue
}

func (d *timeType) GetFormat() TimeFormat {
	return d.typeFormat
}

// parseStringToTime take a string value and return a [time.Time] and the corresponding [TimeFormat]
func parseStringToTime(value string) (time.Time, TimeFormat) {
	// Try to parse it into UTC format
	date, err := time.Parse("150405Z", value)
	format := TimeWithUtcTime
	if err != nil {
		// If failed, parse to local time format
		date, err = time.Parse("150405", value)
		format = TimeWithLocalTime
	}
	return date, format
}

// NewTimeValue create a new [registries.Time] type value. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.12
func NewTimeValue(value time.Time) TimeType {
	zone, _ := value.Zone()
	fmt.Print(zone)
	format := TimeWithLocalTime
	// Determine the correct format representation
	if zone == "UTC" {
		format = TimeWithUtcTime
	}
	return &timeType{
		typeName:   registries.Time,
		typeValue:  value,
		typeFormat: format,
	}
}

// NewTimeValueFromString create a new [registries.Time] type value from string. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.12
func NewTimeValueFromString(value string) TimeType {
	timeValue, format := parseStringToTime(value)
	return &timeType{
		typeName:   registries.Time,
		typeValue:  timeValue,
		typeFormat: format,
	}
}

// NewTimeValues create a new [registries.Time] type value array. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.12
func NewTimeValues(values []time.Time) []TimeType {
	var timeValues []TimeType

	for i := 0; i < len(values); i++ {
		timeValues = append(timeValues, NewTimeValue(
			values[i],
		))
	}
	return timeValues
}

// NewTimeValuesFromString create new [registries.Time] type values from string. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.12
func NewTimeValuesFromString(values []string) []TimeType {
	var timeValues []TimeType

	for i := 0; i < len(values); i++ {
		timeValues = append(timeValues, NewTimeValueFromString(values[i]))
	}
	return timeValues
}
