package types

import "github.com/vareversat/gics/registries"

type TimeTransparencyType string

const (
	OPAQUE      TimeTransparencyType = "OPAQUE"
	TRANSPARENT TimeTransparencyType = "TRANSPARENT"
)

type TimeTransparencyValue struct {
	V
	Value TimeTransparencyType
}

func NewTimeTransparencyValue(value TimeTransparencyType) TimeTransparencyValue {
	return TimeTransparencyValue{
		V:     NewValue(registries.TEXT),
		Value: value,
	}
}
