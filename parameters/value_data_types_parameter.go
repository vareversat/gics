package parameters

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type ValueParam interface {
	TextTypeParameter
}

// NewValueDataTypesParam create a new registries.ValueParam property. See [RFC-5545] ref for more info
// This parameter can be used in these properties :
// - registries.DateTimeEndProp
// - registries.DateTimeStartProp
// - registries.DateTimeDueProp
// - registries.RecurrenceIdProp
// - registries.RecurrenceIdProp
// - registries.ExceptionDateTimesProp
// - registries.TriggerProp
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.19
func NewValueDataTypesParam(value registries.ValueTypeRegistry) ValueParam {
	return &textParameter{
		ParamName: registries.ValueDataTypesParam,
		Value:     types.NewTextValue(string(value)),
	}
}

func NewValueDataTypesParamFromString(value string) ValueParam {
	return &textParameter{
		ParamName: registries.ValueDataTypesParam,
		Value:     types.NewTextValue(value),
	}
}
