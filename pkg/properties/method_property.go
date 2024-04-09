package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.7.2

import (
	"github.com/vareversat/gics/pkg/parameters"
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type MethodProperty interface {
	TextPropertyType
}

func NewMethodProperty(value string, params ...parameters.Parameter) MethodProperty {
	return &textPropertyType{
		PropName:   registries.METHOD,
		Value:      types.NewTextValue(value),
		Parameters: params,
	}
}
