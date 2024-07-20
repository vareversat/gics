package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.10

// Parameter used in these properties :
// - CATEGORIES
// - COMMENT
// - DESCRIPTION
// - LOCATION
// - RESOURCES
// - SUMMARY
// - TZNAME
// - ATTENDEE
// - CONTACT
// - ORGANIZER
// - REQUEST-STATUS

import (
	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type LanguageParam interface {
	TextTypeParameter
}

func NewLanguageParam(value string) LanguageParam {
	return &textParameter{
		ParamName: registry.LANGUAGE,
		Value:     types.NewTextValue(value),
	}
}
