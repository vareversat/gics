package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.1

import (
	"time"

	"github.com/vareversat/gics/parameters"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type DateTimeCompletedProperty interface {
	DateTimePropertyType
}

func NewDateTimeCompletedProperty(
	value time.Time,
) DateTimeCompletedProperty {
	return &dateTimePropertyType{
		PropName: registries.CompletedProp,
		Value:    types.NewDateTimeValue(value, types.WithUtcTime),
	}
}

func NewDateTimeCompletedPropertyFromString(
	value string,
	params ...parameters.Parameter) DateTimeCompletedProperty {
	return &dateTimePropertyType{
		PropName:   registries.CompletedProp,
		Value:      types.NewDateTimeValueFromString(value, types.WithUtcTime),
		Parameters: params,
	}
}
