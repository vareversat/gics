package parameters

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.14

// Parameter used in these properties :
// - TRIGGER

import (
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type AlarmTriggerRelationshipParam interface {
	TextParameter
}

func NewAlarmTriggerRelationshipParam(
	value AlarmTriggerRelationshipType,
) AlarmTriggerRelationshipParam {
	return &textParameter{
		ParamName: registries.RELATED,
		Value:     types.NewTextValue(string(value)),
	}
}

type AlarmTriggerRelationshipType string

const (
	START AlarmTriggerRelationshipType = "START"
	END   AlarmTriggerRelationshipType = "END"
)
