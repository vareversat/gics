package types

import (
	"fmt"
	"net/url"

	"github.com/vareversat/gics/registries"
)

type UriType interface {
	ValueType
	GetValue() *url.URL
}

type uriType struct {
	typeName  registries.ValueTypeRegistry
	typeValue *url.URL
}

// GetStringValue implements UriType.
func (u *uriType) GetStringValue() string {
	return u.typeValue.String()
}

// GetTypeName implements UriType.
func (u *uriType) GetTypeName() string {
	return string(u.typeName)
}

// GetValue implements UriType.
func (u *uriType) GetValue() *url.URL {
	return u.typeValue
}

// parseStringToDate take a string value and return an int (offset representation in seconds)
func parseStringToUrl(value string) *url.URL {
	url, err := url.Parse(value)
	if err != nil {
		fmt.Println(err)
	}
	return url
}

// NewUriValue create a new [registries.Uri] type value. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.14
func NewUriValue(value *url.URL) UriType {
	return &uriType{
		typeName:  registries.Uri,
		typeValue: value,
	}
}

// NewUriValue create a new [registries.Uri] type value. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.14
func NewUriValueFromString(value string) UriType {
	return &uriType{
		typeName:  registries.Uri,
		typeValue: parseStringToUrl(value),
	}
}
