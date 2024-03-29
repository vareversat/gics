package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.8

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type PercentCompleteProperty interface {
}

type percentCompleteProperty struct {
	IANAToken registries.Properties
	Value     values.IntegerValue
}

func NewPercentCompleteProperty(value int32) PercentCompleteProperty {
	return &percentCompleteProperty{
		IANAToken: registries.PERCENTCOMPLETE,
		Value:     values.NewIntegerValue(value),
	}
}
