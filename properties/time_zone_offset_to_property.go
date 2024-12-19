package properties

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type TimeZoneOffsetToProperty interface {
	UtcOffsetPropertyType
}

// NewTimeZoneOffsetToProperty create a new registries.TimeZoneOffsetToProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Standard (Mandatory)
// - registries.Daylight (Mandatory)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.3.4
func NewTimeZoneOffsetToProperty(
	value string,
	params ...parameters.Parameter,
) TimeZoneOffsetToProperty {
	return &utcOffsetPropertyType{
		PropName:   registries.TimeZoneOffsetToProp,
		Value:      types.NewUtcOffsetValueFromString(value),
		Parameters: params,
	}
}
