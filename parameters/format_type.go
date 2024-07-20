package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.8

// Parameter used in these properties :
// - ATTACH

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type FormatTypeParam interface {
	TextTypeParameter
}

func NewFormatTypeParam(value string) FormatTypeParam {
	return &textParameter{
		ParamName: registries.FormatType,
		Value:     types.NewTextValue(value),
	}
}
