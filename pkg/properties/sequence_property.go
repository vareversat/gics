package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.4

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type SequenceProperty interface {
	IntegerPropertyType
}

func NewSequenceProperty(value int32) SequenceProperty {
	return &integerPropertyType{
		PropName: registries.SEQUENCE,
		Value:    types.NewIntegerValue(value),
	}
}
