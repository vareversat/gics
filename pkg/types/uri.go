package types

import (
	"fmt"
	"net/url"

	"github.com/vareversat/gics/pkg/registries"
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
	return fmt.Sprintf("\"%s\"", uV.Value.String())
}
