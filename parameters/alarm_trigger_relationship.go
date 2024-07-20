package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.14

// Parameter used in these properties :
// - TRIGGER

import (
	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type AlarmTriggerRelationshipParam interface {
	TextTypeParameter
}

func NewAlarmTriggerRelationshipParam(
	value AlarmTriggerRelationshipType,
) AlarmTriggerRelationshipParam {
	return &textParameter{
		ParamName: registry.RELATED,
		Value:     types.NewTextValue(string(value)),
	}
}

type AlarmTriggerRelationshipType string

const (
	START AlarmTriggerRelationshipType = "START"
	END   AlarmTriggerRelationshipType = "END"
)
