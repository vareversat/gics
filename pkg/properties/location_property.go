package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.7

import (
	"github.com/vareversat/gics/pkg/parameters"
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type LocationProperty interface {
	TextPropertyType
}

func NewLocationProperty(value string, params ...parameters.Parameter) LocationProperty {
	return &textPropertyType{
		PropName:   registries.LOCATION,
		Value:      types.NewTextValue(value),
		Parameters: params,
	}
}
