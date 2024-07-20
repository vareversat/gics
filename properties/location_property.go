package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.7

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type LocationProperty interface {
	TextPropertyType
}

func NewLocationProperty(value string, params ...parameters.Parameter) LocationProperty {
	return &textPropertyType{
		PropName:   registry.LOCATION,
		Value:      types.NewTextValue(value),
		Parameters: params,
	}
}
