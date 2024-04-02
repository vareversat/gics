package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.3

import (
	"net/url"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type UrlProperty interface {
	UriPropertyType
}

func NewUrlProperty(uri *url.URL) UrlProperty {
	return &uriPropertyType{
		PropName: registries.URL,
		Value:    types.NewUriValue(uri),
	}
}
