package parameters

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type LanguageParam interface {
	TextTypeParameter
}

// NewLanguageParam create a new registries.LanguageParam property
// This parameter can be used in these properties :
// - registries.CategoriesProp
// - registries.CommentProp
// - registries.DescriptionProp
// - registries.LocationProp
// - registries.ResourcesProp
// - registries.SummaryProp
// - registries.TimeZoneNameProp
// - registries.AttendeeProp
// - registries.ContactProp
// - registries.OrganizerProp
// - registries.RequestStatusProp
// [See RFC-5545 ref]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.10
func NewLanguageParam(value string) LanguageParam {
	return &textParameter{
		ParamName: registries.LanguageParam,
		Value:     types.NewTextValue(value),
	}
}
