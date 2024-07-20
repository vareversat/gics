package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.3.4

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type TimeZoneOffsetToProperty interface {
	UtcOffsetPropertyType
}

func NewTimeZoneOffsetToProperty(
	value string,
	params ...parameters.Parameter,
) TimeZoneOffsetToProperty {
	return &utcOffsetPropertyType{
		PropName:   registry.TZOFFSETTO,
		Value:      types.NewUtcOffsetValue(value),
		Parameters: params,
	}
}
