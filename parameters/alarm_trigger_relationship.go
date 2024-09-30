package parameters

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type AlarmTriggerRelationshipParam interface {
	TextTypeParameter
}

// NewAlarmTriggerRelationshipParam create a new registries.AlarmTriggerRelationshipParam property. See [RFC-5545] ref for more info
// This parameter can be used in this property :
// - registries.TriggerProp
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.14
func NewAlarmTriggerRelationshipParam(
	value registries.AlarmTriggerRelationshipTypeRegistry,
) AlarmTriggerRelationshipParam {
	return &textParameter{
		ParamName: registries.AlarmTriggerRelationshipParam,
		Value:     types.NewTextValue(string(value)),
	}
}
