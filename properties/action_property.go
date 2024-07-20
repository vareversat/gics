package properties

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type ActionProperty interface {
	ActionPropertyType
}

// NewActionProperty create a new ACTION property
// This property MUST be seen in VALARM component
// [See RFC-5545 ref]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.1
func NewActionProperty(
	action registry.ActionRegistry,
	params ...parameters.Parameter,
) ActionProperty {
	return &actionPropertyType{
		PropName:   registry.ACTION,
		Value:      types.NewActionValue(action),
		Parameters: params,
	}
}

func NewActionPropertyFromString(
	actionString string,
	params ...parameters.Parameter,
) BlockDelimiterProperty {
	return &actionPropertyType{
		PropName:   registry.ACTION,
		Value:      types.NewActionValue(registry.ActionRegistry(actionString)),
		Parameters: params,
	}
}

func (aP *actionPropertyType) GetActionValue() registry.ActionRegistry {
	return aP.Value.Value
}
