package types

import (
	"fmt"
	"strings"
	"time"

	"github.com/vareversat/gics/registries"
)

type PeriodOfTimeFormat string

const (
	ImplicitPeriodOfTime PeriodOfTimeFormat = "ImplicitPeriodOfTime"
	ExplicitPeriodOfTime PeriodOfTimeFormat = "ExplicitPeriodOfTime"
)

type PeriodOfTimeType interface {
	ValueType
	GetValue() (from, to time.Time)
	GetFormat() PeriodOfTimeFormat
}

type implicitPeriodOfTimeType struct {
	typeName      registries.ValueTypeRegistry
	typeFromValue time.Time
	typeToValue   time.Duration
	typeFormat    PeriodOfTimeFormat
}

type explicitPeriodOfTimeType struct {
	typeName      registries.ValueTypeRegistry
	typeFromValue time.Time
	typeToValue   time.Time
	typeFormat    PeriodOfTimeFormat
}

// GetStringValue return the string representation according to the correct format
func (d *implicitPeriodOfTimeType) GetStringValue() string {
	fromToString := d.typeFromValue.Format("20060102T150405Z")
	toToString := parseDurationToString(d.typeToValue)

	return fmt.Sprintf("%s/%s", fromToString, toToString)
}

// GetStringValue return the string representation according to the correct format
func (d *explicitPeriodOfTimeType) GetStringValue() string {
	fromToString := d.typeFromValue.Format("20060102T150405Z")
	toToString := d.typeToValue.Format("20060102T150405Z")

	return fmt.Sprintf("%s/%s", fromToString, toToString)
}

func (d *implicitPeriodOfTimeType) GetTypeName() string {
	return string(d.typeName)
}

func (d *explicitPeriodOfTimeType) GetTypeName() string {
	return string(d.typeName)
}

// GetValue get the [time.Duration] typed value
func (f *implicitPeriodOfTimeType) GetValue() (from, to time.Time) {
	return f.typeFromValue, f.typeFromValue.Add(f.typeToValue)
}

// GetValue get the [time.Duration] typed value
func (f *explicitPeriodOfTimeType) GetValue() (from, to time.Time) {
	return f.typeFromValue, f.typeFromValue.Add(f.typeToValue.Sub(f.typeFromValue))
}

func (d *implicitPeriodOfTimeType) GetFormat() PeriodOfTimeFormat {
	return d.typeFormat
}

func (d *explicitPeriodOfTimeType) GetFormat() PeriodOfTimeFormat {
	return d.typeFormat
}

// NewImplicitPeriodOfTimeValue create a new [registries.PeriodOfTime] type value. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.9
func NewImplicitPeriodOfTimeValue(
	fromValue time.Time,
	toValue time.Duration,
) PeriodOfTimeType {
	return &implicitPeriodOfTimeType{
		typeName:      registries.PeriodOfTime,
		typeFromValue: fromValue,
		typeToValue:   toValue,
		typeFormat:    ImplicitPeriodOfTime,
	}
}

// NewExplicitPeriodOfTimeValue create a new [registries.PeriodOfTime] type value used for [registries.FreeBusyTimeProp]. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.9
func NewExplicitPeriodOfTimeValue(
	fromValue time.Time,
	toValue time.Time,
) PeriodOfTimeType {
	return &explicitPeriodOfTimeType{
		typeName:      registries.PeriodOfTime,
		typeFromValue: fromValue,
		typeToValue:   toValue,
		typeFormat:    ExplicitPeriodOfTime,
	}
}

// NewPeriodOfTimeValueFromString create a new [registries.PeriodOfTime] type value. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.9
func NewPeriodOfTimeValueFromString(value string) (PeriodOfTimeType, error) {
	var err error
	var from time.Time

	splitString := strings.Split(value, "/")
	from, err = time.Parse("20060102T150405Z", splitString[0])

	if err != nil {
		return nil, fmt.Errorf("%s is not a valid period of time", value)
	}

	if strings.HasPrefix(splitString[1], "P") {
		// We are dealing with an implicit period of time
		var to *time.Duration
		to, err = parseStringToDuration(splitString[1])

		if err != nil {
			return nil, fmt.Errorf("%s is not a valid implicit period of time", value)
		}

		return &implicitPeriodOfTimeType{
			typeName:      registries.PeriodOfTime,
			typeFromValue: from,
			typeToValue:   *to,
			typeFormat:    ImplicitPeriodOfTime,
		}, nil

	} else {
		// We are dealing with an explicit period of time
		var to time.Time
		to, err = time.Parse("20060102T150405Z", splitString[1])

		if err != nil {
			return nil, fmt.Errorf("%s is not a valid explicit period of time", value)
		}

		return &explicitPeriodOfTimeType{
			typeName:      registries.PeriodOfTime,
			typeFromValue: from,
			typeToValue:   to,
			typeFormat:    ImplicitPeriodOfTime,
		}, nil
	}
}
