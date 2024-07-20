package types

import "github.com/vareversat/gics/registries"

type BlockDelimiterValue struct {
	V
	Value registries.ComponentRegistry
}

func NewBlockDelimiterValue(value registries.ComponentRegistry) BlockDelimiterValue {
	return BlockDelimiterValue{
		V:     NewValue(registries.Text),
		Value: value,
	}
}
