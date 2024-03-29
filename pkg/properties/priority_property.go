package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.9

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type PriorityProperty interface {
}

type priorityProperty struct {
	IANAToken registries.Properties
	Value     values.IntegerValue
}

func NewPriorityProperty(value int32) PriorityProperty {
	return &priorityProperty{
		IANAToken: registries.PRIORITY,
		Value:     values.NewIntegerValue(value),
	}
}
