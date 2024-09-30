package properties

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type TimeZoneOffsetFromProperty interface {
	UtcOffsetPropertyType
}

// NewTimeZoneOffsetFromProperty create a new registries.TimeZoneOffsetFromProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Standard (Mandatory)
// - registries.Daylight (Mandatory)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.3.3
func NewTimeZoneOffsetFromProperty(
	value string,
	params ...parameters.Parameter,
) TimeZoneOffsetFromProperty {
	return &utcOffsetPropertyType{
		PropName:   registries.TimeZoneOffsetFromProp,
		Value:      types.NewUtcOffsetValue(value),
		Parameters: params,
	}
}
