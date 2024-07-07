package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.19

// Parameter used in these properties :
// - DTEND
// - DSTART
// - DUE
// - RECURRENCE-ID
// - EXDATE
// - RDATE

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type TimeZoneIdentifierParam interface {
	TextParameter
}

func NewTimeZoneIdentifierParam(value string) TimeZoneIdentifierParam {
	return &textParameter{
		ParamName: registries.PARAM_TZID,
		Value:     types.NewTextValue(value),
	}
}
