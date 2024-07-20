package types

import "github.com/vareversat/gics/registry"

type ActionValue struct {
	V
	Value registry.ActionRegistry
}

func NewActionValue(value registry.ActionRegistry) ActionValue {
	return ActionValue{
		V:     NewValue(registry.Text),
		Value: value,
	}
}
