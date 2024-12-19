package types

import (
	"github.com/vareversat/gics/registries"
)

type ClassificationType interface {
	ValueType
	GetValue() registries.ClassificationRegistry
}

type classificationType struct {
	typeName  registries.ValueTypeRegistry
	typeValue registries.ClassificationRegistry
}

func (c *classificationType) GetStringValue() string {
	return string(c.typeValue)
}

func (c *classificationType) GetTypeName() string {
	return string(c.typeName)
}

// GetValue get the registries.ClassificationRegistry typed value
func (c *classificationType) GetValue() registries.ClassificationRegistry {
	return c.typeValue
}

// NewClassificationValue create a new [registries.Text] type value. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-8.3.11
func NewClassificationValue(value registries.ClassificationRegistry) ClassificationType {
	return &classificationType{
		typeName:  registries.Text,
		typeValue: value,
	}
}
