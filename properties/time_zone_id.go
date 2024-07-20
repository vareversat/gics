package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.3.1

import (
	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type TimeZoneIdProperty interface {
	TextPropertyType
}

func NewTimeZoneIdProperty(value string) TimeZoneIdProperty {
	return &textPropertyType{
		PropName: registry.PROP_TZID,
		Value:    types.NewTextValue(value),
	}
}
