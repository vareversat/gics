package parameters

import (
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
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
	TextParameter
}

func NewValueParam(value registries.Type) ValueParam {
	return &textParameter{
		ParamName: registries.VALUE,
		Value:     types.NewTextValue(string(value)),
	}
}
