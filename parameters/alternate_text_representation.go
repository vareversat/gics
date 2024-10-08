package parameters

import (
	"net/url"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type AlternateTextRepresentationParam interface {
	UriTypeParameter
}

// NewAlternateTextRepresentationParam create a new registries.AlternateTextRepresentationParam property. See [RFC-5545] ref for more info
// This parameter can be used in these properties :
// - registries.CommentProp
// - registries.DescriptionProp
// - registries.LocationProp
// - registries.ResourcesProp
// - registries.SummaryProp
// - registries.ContactProp
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.1
func NewAlternateTextRepresentationParam(value *url.URL) AlternateTextRepresentationParam {
	return &uriParameter{
		ParamName: registries.AlternateTextRepresentationParam,
		Value:     types.NewUriValue(value),
	}
}
