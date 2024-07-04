package properties

import (
	"github.com/vareversat/gics/pkg/parameters"
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type ActionProperty interface {
	ActionPropertyType
	// GetActionValue return the types.ActionValue associated to this ActionProperty
	GetActionValue() types.ActionType
}

// NewActionProperty create a new ACTION property
// This property MUST be seen in VALARM component
// [See RFC-5545 ref]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.1
func NewActionProperty(action types.ActionType, params ...parameters.Parameter) ActionProperty {
	return &actionPropertyType{
		PropName:   registries.ACTION,
		Value:      types.NewActionValue(action),
		Parameters: params,
	}
}

func NewActionPropertyFromString(
	actionString string,
	params ...parameters.Parameter,
) BlockDelimiterProperty {
	return &actionPropertyType{
		PropName:   registries.ACTION,
		Value:      types.NewActionValue(types.ActionType(actionString)),
		Parameters: params,
	}
}

func (aP *actionPropertyType) GetActionValue() types.ActionType {
	return aP.Value.Value
}
