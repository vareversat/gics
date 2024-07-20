package types

import (
	"net/url"

	"github.com/vareversat/gics/registry"
)

type UriValue struct {
	V
	Value *url.URL
}

func NewUriValue(value *url.URL) UriValue {
	return UriValue{
		V:     NewValue(registry.Uri),
		Value: value,
	}
}

func (uV *UriValue) GetValue() string {
	return uV.Value.String()
}
