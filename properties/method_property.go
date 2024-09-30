package properties

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type MethodProperty interface {
	TextPropertyType
}

// NewMethodProperty create a new registries.LocationProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vcalendar (Optional)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.7.2
func NewMethodProperty(value string, params ...parameters.Parameter) MethodProperty {
	return &textPropertyType{
		PropName:   registries.MethodProp,
		Value:      types.NewTextValue(value),
		Parameters: params,
	}
}
