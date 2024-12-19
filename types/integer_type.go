package types

import (
	"strconv"

	"github.com/vareversat/gics/registries"
)

type IntegerType interface {
	ValueType
	GetValue() int32
}

type integerType struct {
	typeName  registries.ValueTypeRegistry
	typeValue int32
}

func (i *integerType) GetStringValue() string {
	return strconv.Itoa(int(i.typeValue))
}

// GetValue get the int32 typed value
func (i *integerType) GetValue() int32 {
	return i.typeValue
}

func (i *integerType) GetTypeName() string {
	return string(i.typeName)
}

// NewIntegerValue create a new [registries.Integer] type value. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.8
func NewIntegerValue(value int32) IntegerType {
	return &integerType{
		typeName:  registries.Integer,
		typeValue: value,
	}
}

// NewIntegerValues create a new [registries.Integer] type values. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.8
func NewIntegerValues(values []int32) []IntegerType {
	var integerValues []IntegerType

	for i := 0; i < len(values); i++ {
		integerValues = append(integerValues, NewIntegerValue(
			values[i],
		))
	}
	return integerValues
}
