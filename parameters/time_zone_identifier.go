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
// - registries.DateTimeEnd
// - registries.DateTimeStart
// - registries.DateTimeDue
// - registries.RecurrenceId
// - registries.RecurrenceId
// - registries.ExceptionDateTimes
// - registries.RecurrenceDateTimes
// [See RFC-5545 ref]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.19
func NewTimeZoneIdentifierParam(value string) TimeZoneIdentifierParam {
	return &textParameter{
		ParamName: registries.TimeZoneIdParam,
		Value:     types.NewTextValue(value),
	}
}
