package types

import "github.com/vareversat/gics/registry"

type BlockDelimiterValue struct {
	V
	Value registry.ComponentRegistry
}

func NewBlockDelimiterValue(value registry.ComponentRegistry) BlockDelimiterValue {
	return BlockDelimiterValue{
		V:     NewValue(registry.Text),
		Value: value,
	}
}
