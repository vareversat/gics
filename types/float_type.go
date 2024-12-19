package types

import (
	"fmt"
	"strconv"

	"github.com/vareversat/gics/registries"
)

type FloatType interface {
	ValueType
	GetValue() float64
}

type floatType struct {
	typeName  registries.ValueTypeRegistry
	typeValue float64
}

func (f *floatType) GetStringValue() string {
	return fmt.Sprintf("%f", f.typeValue)
}

func (f *floatType) GetTypeName() string {
	return string(f.typeName)
}

// GetValue get the float64 typed value
func (f *floatType) GetValue() float64 {
	return f.typeValue
}

// NewFloatValue create a new [registries.Float] type value. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.7
func NewFloatValue(value float64) FloatType {
	return &floatType{
		typeName:  registries.Float,
		typeValue: value,
	}
}

// NewFloatValueFromString create a new [registries.Float] type value from a string. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.7
func NewFloatValueFromString(value string) FloatType {
	float, _ := strconv.ParseFloat(value, 64)
	return &floatType{
		typeName:  registries.Float,
		typeValue: float,
	}
}
