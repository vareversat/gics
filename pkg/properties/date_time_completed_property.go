package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.1

import (
	"time"

	"github.com/vareversat/gics/pkg/parameters"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type DateTimeCompletedProperty interface {
	DateTimePropertyType
}

func NewDateTimeCompletedProperty(
	value time.Time,
) DateTimeCompletedProperty {
	return &dateTimePropertyType{
		PropName: registries.COMPLETED_PROP,
		Value:    types.NewDateTimeValue(value, types.WithUtcTime),
	}
}

func NewDateTimeCompletedPropertyFromString(
	value string,
	params ...parameters.Parameter) DateTimeCompletedProperty {
	return &dateTimePropertyType{
		PropName:   registries.COMPLETED_PROP,
		Value:      types.NewDateTimeValueFromString(value, types.WithUtcTime),
		Parameters: params,
	}
}
