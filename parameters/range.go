package parameters

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type RangeParam interface {
	TextTypeParameter
}

// NewRangeParam create a new registries.RangeParam property
// This parameter can be used in this property :
// - registries.RecurrenceId
// [See RFC-5545 ref]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.13
func NewRangeParam() RangeParam {
	return &textParameter{
		ParamName: registries.RangeParam,
		Value:     types.NewTextValue(string(registries.ThisAndFuture)),
	}
}
