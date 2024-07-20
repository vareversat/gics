package parameters

import (
	"github.com/vareversat/gics/registry"
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

func NewValueParam(value registry.ValueTypeRegistry) ValueParam {
	return &textParameter{
		ParamName: registry.Value,
		Value:     types.NewTextValue(string(value)),
	}
}
