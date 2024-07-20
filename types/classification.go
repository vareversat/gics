package types

import "github.com/vareversat/gics/registries"

type ClassificationValue struct {
	V
	Value registries.ClassificationRegistry
}

func NewClassificationValue(value registries.ClassificationRegistry) ClassificationValue {
	return ClassificationValue{
		V:     NewValue(registries.Text),
		Value: value,
	}
}
