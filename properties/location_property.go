package properties

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type LocationProperty interface {
	TextPropertyType
}

// NewLocationProperty create a new registries.LocationProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// Optional parameters :
// - registries.AlternateTextRepresentationParam
// - registries.LanguageParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.7
func NewLocationProperty(value string, params ...parameters.Parameter) LocationProperty {
	return &textPropertyType{
		PropName:   registries.LocationProp,
		Value:      types.NewTextValue(value),
		Parameters: params,
	}
}
