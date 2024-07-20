package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.3

// Parameter used in these properties :
// - ATTENDEE

import (
	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type CalendarUserTypeParam interface {
	TextTypeParameter
}

func NewCalendarUserTypeParam(value registry.CalendarUserTypeRegistry) CalendarUserTypeParam {
	return &textParameter{
		ParamName: registry.CalendarUserType,
		Value:     types.NewTextValue(string(value)),
	}
}
