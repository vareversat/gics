package parameters

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type CalendarUserTypeParam interface {
	TextTypeParameter
}

// NewCalendarUserTypeParam create a new registries.CalendarUserTypeParam property
// This parameter can be used in this property :
// - registries.AttendeeProp
// [See RFC-5545 ref]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.3
func NewCalendarUserTypeParam(value registries.CalendarUserTypeRegistry) CalendarUserTypeParam {
	return &textParameter{
		ParamName: registries.CalendarUserTypeParam,
		Value:     types.NewTextValue(string(value)),
	}
}
