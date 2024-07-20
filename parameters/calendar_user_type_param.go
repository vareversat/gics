package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.3

// Parameter used in these properties :
// - ATTENDEE

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type CalendarUserTypeParam interface {
	TextTypeParameter
}

func NewCalendarUserTypeParam(value registries.CalendarUserTypeRegistry) CalendarUserTypeParam {
	return &textParameter{
		ParamName: registries.CalendarUserType,
		Value:     types.NewTextValue(string(value)),
	}
}
