package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.8

// Parameter used in these properties :
// - ATTACH

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type FormatTypeParam interface {
	TextParameter
}

func NewFormatTypeParam(value string) FormatTypeParam {
	return &textParameter{
		ParamName: registries.FMTTYPE,
		Value:     types.NewTextValue(value),
	}
}
