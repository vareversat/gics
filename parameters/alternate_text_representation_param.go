package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.1

// Parameter used in these properties :
// - COMMENT
// - DESCRIPTION
// - LOCATION
// - RESOURCES
// - SUMMARY
// - CONTACT

import (
	"net/url"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type AlternateTextRepresentationParam interface {
	UriTypeParameter
}

func NewAlternateTextRepresentationParam(value *url.URL) AlternateTextRepresentationParam {
	return &uriParameter{
		ParamName: registries.AlternateTextRepresentation,
		Value:     types.NewUriValue(value),
	}
}
