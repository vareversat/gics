package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.11

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type StatusProperty interface {
	StatusPropertyType
}

func NewStatusProperty(status types.StatusType, params ...parameters.Parameter) StatusProperty {
	return &statusPropertyType{
		PropName:   registries.STATUS,
		Value:      types.NewStatusValue(status),
		Parameters: params,
	}
}

func NewStatusPropertyFromString(status string, params ...parameters.Parameter) StatusProperty {
	return &statusPropertyType{
		PropName:   registries.STATUS,
		Value:      types.NewStatusValue(types.StatusType(status)),
		Parameters: params,
	}
}
