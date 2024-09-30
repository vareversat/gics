package properties

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type DurationProperty interface {
	DurationPropertyType
}

// NewDurationProperty create a new registries.DurationProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Valarm (Optional)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.5
func NewDurationProperty(duration string) DurationProperty {
	return &durationPropertyType{
		PropName: registries.DurationProp,
		Value:    types.NewDurationValue(duration),
	}
}
