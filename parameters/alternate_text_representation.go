package parameters

import (
	"net/url"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type AlternateTextRepresentationParam interface {
	UriTypeParameter
}

// NewAlternateTextRepresentationParam create a new registries.AlternateTextRepresentation property
// This parameter can be used in these properties :
// - registries.Comment
// - registries.Description
// - registries.Location
// - registries.Resources
// - registries.Summary
// - registries.Contact
// [See RFC-5545 ref]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.1
func NewAlternateTextRepresentationParam(value *url.URL) AlternateTextRepresentationParam {
	return &uriParameter{
		ParamName: registries.AlternateTextRepresentation,
		Value:     types.NewUriValue(value),
	}
}
