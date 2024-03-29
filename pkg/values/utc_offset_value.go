package values

import (
	"github.com/vareversat/gics/pkg/registries"
)

type UtcOffsetValue interface{}

type utcOffsetValue struct {
	Value     string
	IANAToken registries.Type
}

func NewUtcOffsetValue(value string) UtcOffsetValue {
	return &utcOffsetValue{
		IANAToken: registries.UTCOFFSET,
		Value:     value,
	}
}
