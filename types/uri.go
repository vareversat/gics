package types

import (
	"net/url"

	"github.com/vareversat/gics/registries"
)

type UriValue struct {
	V
	Value *url.URL
}

func NewUriValue(value *url.URL) UriValue {
	return UriValue{
		V:     NewValue(registries.URI),
		Value: value,
	}
}

func (uV *UriValue) GetValue() string {
	return uV.Value.String()
}
