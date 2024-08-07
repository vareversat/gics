package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.13

// Parameter used in these properties :
// - RECURRENCE-ID

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type RangeParam interface {
	TextTypeParameter
}

func NewRangeParam() RangeParam {
	return &textParameter{
		ParamName: registries.Range,
		Value:     types.NewTextValue("THISANDFUTURE"),
	}
}
