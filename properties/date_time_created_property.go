package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.1

import (
	"time"

	"github.com/vareversat/gics/parameters"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type DateTimeCreatedProperty interface {
	DateTimePropertyType
}

func NewDateTimeCreatedProperty(
	timeValue time.Time,
	params ...parameters.Parameter,
) DateTimeCreatedProperty {
	return &dateTimePropertyType{
		PropName:   registries.CREATED,
		Value:      types.NewDateTimeValue(timeValue, types.WithUtcTime),
		Parameters: params,
	}
}

func NewDateTimeCreatedPropertyFromString(
	value string,
	params ...parameters.Parameter) DateTimeCreatedProperty {
	return &dateTimePropertyType{
		PropName:   registries.CREATED,
		Value:      types.NewDateTimeValueFromString(value, types.WithUtcTime),
		Parameters: params,
	}
}
