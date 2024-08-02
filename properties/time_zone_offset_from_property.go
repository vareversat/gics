package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.3.3

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type TimeZoneOffsetFromProperty interface {
	UtcOffsetPropertyType
}

func NewTimeZoneOffsetFromProperty(
	value string,
	params ...parameters.Parameter,
) TimeZoneOffsetFromProperty {
	return &utcOffsetPropertyType{
		PropName:   registries.TimeZoneOffsetFromProp,
		Value:      types.NewUtcOffsetValue(value),
		Parameters: params,
	}
}
