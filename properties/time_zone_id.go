package properties

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type TimeZoneIdProperty interface {
	TextPropertyType
}

// NewTimeZoneIdProperty create a new registries.TimeZoneIdProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vtimezone (Mandatory)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.3.1
func NewTimeZoneIdProperty(value string) TimeZoneIdProperty {
	return &textPropertyType{
		PropName: registries.TimeZoneIdProp,
		Value:    types.NewTextValue(value),
	}
}
