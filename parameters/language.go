package parameters

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type LanguageParam interface {
	TextTypeParameter
}

// NewLanguageParam create a new registries.Language property
// This parameter can be used in these properties :
// - registries.Categories
// - registries.Comment
// - registries.Description
// - registries.Location
// - registries.Resources
// - registries.Summary
// - registries.TimeZoneName
// - registries.Attendee
// - registries.Contact
// - registries.Organizer
// - registries.RequestStatus
// [See RFC-5545 ref]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.10
func NewLanguageParam(value string) LanguageParam {
	return &textParameter{
		ParamName: registries.Language,
		Value:     types.NewTextValue(value),
	}
}
