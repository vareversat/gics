package parameters

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type RecurrenceIdRangeParam interface {
	TextTypeParameter
}

// NewRecurrenceIdRangeParam create a new registries.RecurrenceIdRangeParam property. See [RFC-5545] ref for more info
// This parameter can be used in this property :
// - registries.RecurrenceIdProp
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.13
func NewRecurrenceIdRangeParam() RecurrenceIdRangeParam {
	return &textParameter{
		ParamName: registries.RecurrenceIdRangeParam,
		Value:     types.NewTextValue(string(registries.ThisAndFuture)),
	}
}
