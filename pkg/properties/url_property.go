package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.3

import (
	"net/url"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/values"
)

type UrlProperty interface {
}

type urlProperty struct {
	IANAToken registries.Properties
	Value     values.CalendarUserAddressValue
}

func NewUrlProperty(uri url.URL) UrlProperty {
	return &urlProperty{
		IANAToken: registries.URL,
		Value:     values.NewUriValue(uri),
	}
}
