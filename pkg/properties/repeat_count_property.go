package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.2

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type RepeatCountProperty interface {
	IntegerPropertyType
}

func NewRepeatCountProperty(value int32) RepeatCountProperty {
	return &integerPropertyType{
		PropName: registries.REPEAT,
		Value:    types.NewIntegerValue(value),
	}
}
