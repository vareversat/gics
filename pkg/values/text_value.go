package values

import (
	"github.com/vareversat/gics/pkg/registries"
)

type TextValue interface{}

type textValue struct {
	Value     string
	IANAToken registries.Type
}

func NewTextValue(value string) TextValue {
	return &textValue{
		IANAToken: registries.TEXT,
		Value:     value,
	}
}
