package values

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.8

import (
	"github.com/vareversat/gics/pkg/registries"
)

type IntegerValue interface{}

type integerValue struct {
	Value     int32
	IANAToken registries.Type
}

func NewIntegerValue(value int32) IntegerValue {
	return &integerValue{
		IANAToken: registries.INTEGER,
		Value:     value,
	}
}
