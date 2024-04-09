package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.3.2

import (
	"github.com/vareversat/gics/pkg/parameters"
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type TimeZoneNameProperty interface {
	TextPropertyType
}

func NewTimeZoneNameProperty(value string, params ...parameters.Parameter) TimeZoneNameProperty {
	return &textPropertyType{
		PropName:   registries.TZNAME,
		Value:      types.NewTextValue(value),
		Parameters: params,
	}
}
