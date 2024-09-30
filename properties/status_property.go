package properties

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type StatusProperty interface {
	StatusPropertyType
}

// NewStatusProperty create a new registries.StatusProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Vjournal (Optional)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.11
func NewStatusProperty(status types.StatusType, params ...parameters.Parameter) StatusProperty {
	return &statusPropertyType{
		PropName:   registries.StatusProp,
		Value:      types.NewStatusValue(status),
		Parameters: params,
	}
}

// NewStatusPropertyFromString create a new registries.StatusProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Vjournal (Optional)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.1.11
func NewStatusPropertyFromString(status string, params ...parameters.Parameter) StatusProperty {
	return &statusPropertyType{
		PropName:   registries.StatusProp,
		Value:      types.NewStatusValue(types.StatusType(status)),
		Parameters: params,
	}
}
