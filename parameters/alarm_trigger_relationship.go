package parameters

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type AlarmTriggerRelationshipParam interface {
	TextTypeParameter
}

// NewAlarmTriggerRelationshipParam create a new registries.Related property
// This parameter can be used in this property :
// - registries.Trigger
// [See RFC-5545 ref]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.14
func NewAlarmTriggerRelationshipParam(
	value registries.AlarmTriggerRelationshipType,
) AlarmTriggerRelationshipParam {
	return &textParameter{
		ParamName: registries.Related,
		Value:     types.NewTextValue(string(value)),
	}
}
