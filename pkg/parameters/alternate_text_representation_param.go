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

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type AlternateTextRepresentationParam interface {
	UriParameter
}

func NewAlternateTextRepresentationParam(value *url.URL) AlternateTextRepresentationParam {
	return &uriParameter{
		ParamName: registries.ALTREP,
		Value:     types.NewUriValue(value),
	}
}
