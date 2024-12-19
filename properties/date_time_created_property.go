package properties

import (
	"time"

	"github.com/vareversat/gics/parameters"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type DateTimeCreatedProperty interface {
	DateTimePropertyType
}

// NewDateTimeCreatedProperty create a new registries.DateTimeCreatedProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Vjournal (Optional)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.1
func NewDateTimeCreatedProperty(
	value time.Time,
	params ...parameters.Parameter,
) DateTimeCreatedProperty {
	return &dateTimePropertyType{
		PropName:   registries.DateTimeCreatedProp,
		Value:      types.NewDateTimeValue(value),
		Parameters: params,
	}
}

// NewDateTimeCreatedPropertyFromString create a new registries.DateTimeCreatedProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Vjournal (Optional)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.1
func NewDateTimeCreatedPropertyFromString(
	value string,
	params ...parameters.Parameter) DateTimeCreatedProperty {
	return &dateTimePropertyType{
		PropName:   registries.DateTimeCreatedProp,
		Value:      types.NewDateTimeValueFromString(value),
		Parameters: params,
	}
}
