package properties

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type ActionProperty interface {
	ActionPropertyType
}

// NewActionProperty create a new [registries.ActionProp] property. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.1
func NewActionProperty(
	action registries.ActionRegistry,
	parameters ...parameters.Parameter,
) ActionProperty {
	return &actionPropertyType{
		PropName:   registries.ActionProp,
		PropValue:  types.NewActionValue(action),
		Parameters: parameters,
	}
}

// NewActionPropertyFromString create a new [registries.ActionProp] property from a string value. See [RFC-5545] ref for more info
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.1
func NewActionPropertyFromString(
	action string,
	parameters ...parameters.Parameter,
) ActionProperty {
	return &actionPropertyType{
		PropName:   registries.ActionProp,
		PropValue:  types.NewActionValue(registries.ActionRegistry(action)),
		Parameters: parameters,
	}
}
