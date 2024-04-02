package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.11

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type StatusProperty interface {
	StatusPropertyType
}

func NewStatusProperty(status types.StatusType) StatusProperty {
	return &statusPropertyType{
		PropName: registries.STATUS,
		Value:    types.NewStatusValue(status),
	}
}
