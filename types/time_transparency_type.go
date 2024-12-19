package types

import "github.com/vareversat/gics/registries"

type TimeTransparencyType interface {
	ValueType
	GetValue() registries.TimeTransparencyRegistry
}

type timeTransparencyType struct {
	typeName  registries.ValueTypeRegistry
	typeValue registries.TimeTransparencyRegistry
}

func (c *timeTransparencyType) GetStringValue() string {
	return string(c.typeValue)
}

func (c *timeTransparencyType) GetTypeName() string {
	return string(c.typeName)
}

// GetValue get the [registries.TimeTransparencyRegistry] typed value
func (c *timeTransparencyType) GetValue() registries.TimeTransparencyRegistry {
	return c.typeValue
}

// NewTimeTransparencyValue create a new [registries.TimeTransparencyRegistry] type value. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.7
func NewTimeTransparencyValue(value registries.TimeTransparencyRegistry) TimeTransparencyType {
	return &timeTransparencyType{
		typeName:  registries.Text,
		typeValue: value,
	}
}

// NewTimeTransparencyValueFromString create a new [registries.TimeTransparencyRegistry] type value from a string. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.7
func NewTimeTransparencyValueFromString(value string) TimeTransparencyType {
	return &timeTransparencyType{
		typeName:  registries.Text,
		typeValue: registries.TimeTransparencyRegistry(value),
	}
}
