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
	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type TimeZoneIdentifierParam interface {
	TextTypeParameter
}

func NewTimeZoneIdentifierParam(value string) TimeZoneIdentifierParam {
	return &textParameter{
		ParamName: registry.TZIDParameter,
		Value:     types.NewTextValue(value),
	}
}
