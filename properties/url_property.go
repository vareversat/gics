package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.3

import (
	"fmt"
	"net/url"

	"github.com/vareversat/gics/parameters"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type UrlProperty interface {
	UriPropertyType
}

func NewUrlProperty(uri *url.URL, params ...parameters.Parameter) UrlProperty {
	return &uriPropertyType{
		PropName:   registries.URL,
		Value:      types.NewUriValue(uri),
		Parameters: params,
	}
}

func NewUrlPropertyFromString(uri string, params ...parameters.Parameter) (UrlProperty, error) {
	urlValue, err := url.Parse(uri)
	if err != nil {
		return nil, fmt.Errorf("%s is not a valid URL format", uri)
	}

	return &uriPropertyType{
		PropName:   registries.URL,
		Value:      types.NewUriValue(urlValue),
		Parameters: params,
	}, nil
}
