package properties

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type TimeTransparencyProperty interface {
	TimeTransparencyPropertyType
}

// NewTimeTransparencyProperty create a new registries.TimeTransparencyProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vcalendar (Mandatory)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.7.4
func NewTimeTransparencyProperty(
	value types.TimeTransparencyType,
	params ...parameters.Parameter,
) TimeTransparencyProperty {
	return &timeTransparencyPropertyType{
		PropName:   registries.TimeTransparencyProp,
		Value:      types.NewTimeTransparencyValue(value),
		Parameters: params,
	}
}

// NewTimeTransparencyPropertyFromString create a new registries.TimeTransparencyProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vcalendar (Mandatory)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.7.4
func NewTimeTransparencyPropertyFromString(
	value string,
	params ...parameters.Parameter,
) TimeTransparencyProperty {
	return &timeTransparencyPropertyType{
		PropName:   registries.TimeTransparencyProp,
		Value:      types.NewTimeTransparencyValue(types.TimeTransparencyType(value)),
		Parameters: params,
	}
}
