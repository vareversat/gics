package types

import (
	"fmt"
	"strings"
	"time"

	"github.com/vareversat/gics/registries"
)

type PeriodOfTimeFormat string

const (
	ExplicitPeriodOfTime PeriodOfTimeFormat = "ExplicitPeriodOfTime"
	ImplicitPeriodOfTime PeriodOfTimeFormat = "ImplicitPeriodOfTime"
)

type PeriodOfTimeType interface {
	ValueType
	GetValue() (from, to time.Time)
	GetFormat() PeriodOfTimeFormat
}

type periodOfTimeType struct {
	typeName      registries.ValueTypeRegistry
	typeFromValue time.Time
	typeToValue   time.Duration
	typeFormat    PeriodOfTimeFormat
}

// GetStringValue return the string representation according to the correct format
func (d *periodOfTimeType) GetStringValue() string {
	var endToString string
	startToString := d.typeFromValue.Format("20060102T150405Z")

	switch d.typeFormat {
	case ExplicitPeriodOfTime:
		endToString = d.typeFromValue.Add(d.typeToValue).Format("20060102T150405Z")
	case ImplicitPeriodOfTime:
		endToString = parseDurationToString(d.typeToValue)
	default:
		endToString = d.typeFromValue.Add(d.typeToValue).Format("20060102T150405Z")
	}
	return fmt.Sprintf("%s/%s", startToString, endToString)
}

func (d *periodOfTimeType) GetTypeName() string {
	return string(d.typeName)
}

// GetValue get the [time.Duration] typed value
func (f *periodOfTimeType) GetValue() (from, to time.Time) {
	return f.typeFromValue, f.typeFromValue.Add(f.typeToValue)
}

func (d *periodOfTimeType) GetFormat() PeriodOfTimeFormat {
	return d.typeFormat
}

// parseStringToPeriodOfTime take a string value and return the from [time.Time] value, the to [time.Duration] value to and the corresponding [PeriodOfTimeFormat]
func parseStringToPeriodOfTime(
	value string,
) (from time.Time, to time.Duration, format PeriodOfTimeFormat) {
	splitString := strings.Split(value, "/")

	from, _ = time.Parse("20060102T150405Z", splitString[0])

	if strings.HasPrefix(splitString[1], "P") {
		// We are dealing with a duration
		to = parseStringToDuration(splitString[1])
		format = ImplicitPeriodOfTime
	} else {
		// We are dealing with a time
		toTime, _ := time.Parse("20060102T150405Z", splitString[1])
		to = toTime.Sub(from)
		format = ExplicitPeriodOfTime
	}
	return from, to, format
}

// NewPeriodOfTimeValue create a new [registries.PeriodOfTime] type value. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.9
func NewPeriodOfTimeValue(
	startValue time.Time,
	endValue time.Duration,
	format PeriodOfTimeFormat,
) PeriodOfTimeType {
	return &periodOfTimeType{
		typeName:      registries.PeriodOfTime,
		typeFromValue: startValue,
		typeToValue:   endValue,
		typeFormat:    format,
	}
}

// NewPeriodOfTimeValue create a new [registries.PeriodOfTime] type value. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.9
func NewPeriodOfTimeValueFromString(value string) PeriodOfTimeType {
	from, to, format := parseStringToPeriodOfTime(value)
	return &periodOfTimeType{
		typeName:      registries.PeriodOfTime,
		typeFromValue: from,
		typeToValue:   to,
		typeFormat:    format,
	}
}
