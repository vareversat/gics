package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.3.5

import (
	"fmt"
	"net/url"

	"github.com/vareversat/gics/pkg/parameters"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type TimeZoneUrlProperty interface {
	UriPropertyType
}

func NewTimeZoneUrlProperty(value *url.URL, params ...parameters.Parameter) TimeZoneUrlProperty {
	return &uriPropertyType{
		PropName:   registries.TZURL,
		Value:      types.NewUriValue(value),
		Parameters: params,
	}
}

func NewTimeZoneUrlPropertyFromString(
	uri string,
	params ...parameters.Parameter,
) (TimeZoneUrlProperty, error) {
	urlValue, err := url.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("%s is not a valid URL format", uri)
	}

	return &uriPropertyType{
		PropName:   registries.TZURL,
		Value:      types.NewUriValue(urlValue),
		Parameters: params,
	}, nil
}
