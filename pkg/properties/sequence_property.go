package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.4

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type SequenceProperty interface {
}

type sequenceProperty struct {
	IANAToken registries.Properties
	Value     values.IntegerValue
}

func NewSequenceProperty(value int32) SequenceProperty {
	return &sequenceProperty{
		IANAToken: registries.SEQUENCE,
		Value:     values.NewIntegerValue(value),
	}
}
