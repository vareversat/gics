package values

import (
	"github.com/vareversat/gics/pkg/registries"
)

type FloatValue interface{}

type floatValue struct {
	Value     float32
	IANAToken registries.Type
}

func NewFloatValue(value float32) FloatValue {
	return &floatValue{
		IANAToken: registries.FLOAT,
		Value:     value,
	}
}
