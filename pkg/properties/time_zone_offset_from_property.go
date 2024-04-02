package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.3.3

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type TimeZoneOffsetFromProperty interface {
	UtcOffsetPropertyType
}

func NewTimeZoneOffsetFromProperty(value string) TimeZoneOffsetFromProperty {
	return &utcOffsetPropertyType{
		PropName: registries.TZOFFSETFROM,
		Value:    types.NewUtcOffsetValue(value),
	}
}
