package parameters

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type ValueParam interface {
	TextTypeParameter
}

// NewValueParam create a new registries.ValueParam property
// This parameter can be used in these properties :
// - registries.DateTimeEndProp
// - registries.DateTimeStartProp
// - registries.DateTimeDueProp
// - registries.RecurrenceIdProp
// - registries.RecurrenceIdProp
// - registries.ExceptionDateTimesProp
// - registries.TriggerProp
// [See RFC-5545 ref]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.19
func NewValueParam(value registries.ValueTypeRegistry) ValueParam {
	return &textParameter{
		ParamName: registries.ValueParam,
		Value:     types.NewTextValue(string(value)),
	}
}
