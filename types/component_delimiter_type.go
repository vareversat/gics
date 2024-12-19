package types

import "github.com/vareversat/gics/registries"

type ComponentDelimiterType interface {
	ValueType
	GetValue() registries.ComponentRegistry
}

type componentDelimiterType struct {
	typeName  registries.ValueTypeRegistry
	typeValue registries.ComponentRegistry
}

func (c *componentDelimiterType) GetStringValue() string {
	return string(c.typeValue)
}

func (c *componentDelimiterType) GetTypeName() string {
	return string(c.typeName)
}

// GetValue get the [registries.ComponentRegistry] typed value
func (c *componentDelimiterType) GetValue() registries.ComponentRegistry {
	return c.typeValue
}

func NewComponentDelimiterValue(value registries.ComponentRegistry) ComponentDelimiterType {
	return &componentDelimiterType{
		typeName:  registries.Text,
		typeValue: value,
	}
}
