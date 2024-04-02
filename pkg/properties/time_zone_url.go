package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.3.5

import (
	"net/url"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type TimeZoneUrlProperty interface {
	UriPropertyType
}

func NewTimeZoneUrlProperty(value *url.URL) TimeZoneUrlProperty {
	return &uriPropertyType{
		PropName: registries.TZURL,
		Value:    types.NewUriValue(value),
	}
}
