package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.8

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type PercentCompleteProperty interface {
	IntegerPropertyType
}

func NewPercentCompleteProperty(value int32) IntegerPropertyType {
	return &integerPropertyType{
		PropName: registries.PERCENTCOMPLETE,
		Value:    types.NewIntegerValue(value),
	}
}
