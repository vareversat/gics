package types

import "github.com/vareversat/gics/registry"

type ActionType string

const (
	AUDIO   ActionType = "AUDIO"
	DISPLAY ActionType = "DISPLAY"
	EMAIL   ActionType = "EMAIL"
)

type ActionValue struct {
	V
	Value ActionType
}

func NewActionValue(value ActionType) ActionValue {
	return ActionValue{
		V:     NewValue(registry.Text),
		Value: value,
	}
}
