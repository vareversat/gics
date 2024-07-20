package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.3.2

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type TimeZoneNameProperty interface {
	TextPropertyType
}

func NewTimeZoneNameProperty(value string, params ...parameters.Parameter) TimeZoneNameProperty {
	return &textPropertyType{
		PropName:   registry.TZNAME,
		Value:      types.NewTextValue(value),
		Parameters: params,
	}
}
