package types

import "github.com/vareversat/gics/registry"

type ClassificationValue struct {
	V
	Value registry.ClassificationRegistry
}

func NewClassificationValue(value registry.ClassificationRegistry) ClassificationValue {
	return ClassificationValue{
		V:     NewValue(registry.Text),
		Value: value,
	}
}
