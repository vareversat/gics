package types

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.8

import (
	"strconv"

	"github.com/vareversat/gics/registry"
)

type IntegerValue struct {
	V
	Value int32
}

func NewIntegerValue(value int32) IntegerValue {
	return IntegerValue{
		V:     NewValue(registry.INTEGER),
		Value: value,
	}
}

func NewIntegerValues(values []int32) []IntegerValue {
	integerValues := make([]IntegerValue, 0)

	for i := 0; i < len(values); i++ {
		integerValues = append(integerValues, IntegerValue{
			Value: values[i],
		})
	}
	return integerValues
}

func (iV *IntegerValue) GetStringValue() string {
	return strconv.Itoa(int(iV.Value))
}

func (iV *IntegerValue) GetValue() int32 {
	return iV.Value
}
