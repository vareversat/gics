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
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type LanguageParam interface {
	TextParameter
}

func NewLanguageParam(value string) LanguageParam {
	return &textParameter{
		ParamName: registries.LANGUAGE,
		Value:     types.NewTextValue(value),
	}
}
