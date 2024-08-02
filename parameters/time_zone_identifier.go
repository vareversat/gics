package parameters

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type TimeZoneIdentifierParam interface {
	TextTypeParameter
}

// NewTimeZoneIdentifierParam create a new registries.TimeZoneIdParam property
// This parameter can be used in these properties :
// - registries.DateTimeEndProp
// - registries.DateTimeStartProp
// - registries.DateTimeDueProp
// - registries.RecurrenceIdProp
// - registries.RecurrenceIdProp
// - registries.ExceptionDateTimesProp
// - registries.RecurrenceDateTimesProp
// [See RFC-5545 ref]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.19
func NewTimeZoneIdentifierParam(value string) TimeZoneIdentifierParam {
	return &textParameter{
		ParamName: registries.TimeZoneIdParam,
		Value:     types.NewTextValue(value),
	}
}