package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.3.4

import (
	"github.com/vareversat/gics/pkg/parameters"
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type TimeZoneOffsetToProperty interface {
	UtcOffsetPropertyType
}

func NewTimeZoneOffsetToProperty(
	value string,
	params ...parameters.Parameter,
) TimeZoneOffsetToProperty {
	return &utcOffsetPropertyType{
		PropName:   registries.TZOFFSETTO,
		Value:      types.NewUtcOffsetValue(value),
		Parameters: params,
	}
}
