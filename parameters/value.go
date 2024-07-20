package parameters

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

// Used for the properties to determine their types
// - DSTART
// - DTEND
// - DUE
// - RECURRENCE-ID
// - EDATE
// - RDATE
// - TRIGGER

type ValueParam interface {
	TextTypeParameter
}

func NewValueParam(value registries.ValueTypeRegistry) ValueParam {
	return &textParameter{
		ParamName: registries.Value,
		Value:     types.NewTextValue(string(value)),
	}
}
