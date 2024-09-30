package properties

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type TimeZoneNameProperty interface {
	TextPropertyType
}

// NewTimeZoneNameProperty create a new registries.TimeZoneNameProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Standard (Optional)
// - registries.Daylight (Optional)
// Optional parameters :
// - registries.LanguageParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.3.2
func NewTimeZoneNameProperty(value string, params ...parameters.Parameter) TimeZoneNameProperty {
	return &textPropertyType{
		PropName:   registries.TimeZoneNameProp,
		Value:      types.NewTextValue(value),
		Parameters: params,
	}
}
