package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.9

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type PriorityProperty interface {
	IntegerPropertyType
}

func NewPriorityProperty(value int32) PriorityProperty {
	return &integerPropertyType{
		PropName: registries.PRIORITY,
		Value:    types.NewIntegerValue(value),
	}
}
