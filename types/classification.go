package types

import "github.com/vareversat/gics/registry"

type ClassificationType string

const (
	PUBLIC       ClassificationType = "PUBLIC"
	PRIVATE      ClassificationType = "PRIVATE"
	CONFIDENTIAL ClassificationType = "CONFIDENTIAL"
)

type ClassificationValue struct {
	V
	Value ClassificationType
}

func NewClassificationValue(value ClassificationType) ClassificationValue {
	return ClassificationValue{
		V:     NewValue(registry.Text),
		Value: value,
	}
}
