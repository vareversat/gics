package types

import "github.com/vareversat/gics/registries"

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
		V:     NewValue(registries.TEXT),
		Value: value,
	}
}
