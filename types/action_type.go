package types

import (
	"github.com/vareversat/gics/registries"
)

type ActionType interface {
	ValueType
	GetValue() registries.ActionRegistry
}

type actionType struct {
	typeName  registries.ValueTypeRegistry
	typeValue registries.ActionRegistry
}

// GetValue get the [registries.ActionRegistry] typed value
func (aT *actionType) GetValue() registries.ActionRegistry {
	return aT.typeValue
}

func (aT *actionType) GetStringValue() string {
	return string(aT.typeValue)
}

func (aT *actionType) GetTypeName() string {
	return string(aT.typeName)
}

// NewActionValue create a new [registries.Text] type value. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.11
func NewActionValue(value registries.ActionRegistry) ActionType {
	return &actionType{
		typeName:  registries.Text,
		typeValue: registries.ActionRegistry(value),
	}
}
