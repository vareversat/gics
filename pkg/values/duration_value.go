package values

import (
	"github.com/vareversat/gics/pkg/registries"
)

type DurationValue interface{}

type durationValue struct {
	Value     string
	IANAToken registries.Type
}

func NewDurationValue(value string) DurationValue {
	return &durationValue{
		IANAToken: registries.DURATION,
		Value:     value,
	}
}
