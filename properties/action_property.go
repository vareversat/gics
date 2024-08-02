package properties

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type ActionProperty interface {
	ActionPropertyType
}

// NewActionProperty create a new ACTION property
// This property MUST be seen in VALARM component
// [See RFC-5545 ref]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.1
func NewActionProperty(
	action registries.ActionRegistry,
	params ...parameters.Parameter,
) ActionProperty {
	return &actionPropertyType{
		PropName:   registries.ActionProp,
		Value:      types.NewActionValue(action),
		Parameters: params,
	}
}

func NewActionPropertyFromString(
	actionString string,
	params ...parameters.Parameter,
) BlockDelimiterProperty {
	return &actionPropertyType{
		PropName:   registries.ActionProp,
		Value:      types.NewActionValue(registries.ActionRegistry(actionString)),
		Parameters: params,
	}
}

func (aP *actionPropertyType) GetActionValue() registries.ActionRegistry {
	return aP.Value.Value
}
