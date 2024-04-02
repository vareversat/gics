package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.13

// Parameter used in these properties :
// - RECURRENCE-ID

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type RangeParam interface {
	TextParameter
}

func NewRangeParam() RangeParam {
	return &textParameter{
		ParamName: registries.RANGE,
		Value:     types.NewTextValue("THISANDFUTURE"),
	}
}
