package types

import "github.com/vareversat/gics/registries"

type ActionValue struct {
	V
	Value registries.ActionRegistry
}

func NewActionValue(value registries.ActionRegistry) ActionValue {
	return ActionValue{
		V:     NewValue(registries.Text),
		Value: value,
	}
}
