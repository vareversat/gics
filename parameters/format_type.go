package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.8

// Parameter used in these properties :
// - ATTACH

import (
	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type FormatTypeParam interface {
	TextTypeParameter
}

func NewFormatTypeParam(value string) FormatTypeParam {
	return &textParameter{
		ParamName: registry.FmtType,
		Value:     types.NewTextValue(value),
	}
}
