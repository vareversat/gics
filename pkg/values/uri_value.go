package values

import (
	"net/url"

	"github.com/vareversat/gics/pkg/registries"
)

type UriValue interface{}

type uriValue struct {
	Value     url.URL
	IANAToken registries.Type
}

func NewUriValue(value url.URL) UriValue {
	return &uriValue{
		IANAToken: registries.URI,
		Value:     value,
	}
}
