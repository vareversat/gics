package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.3

import (
	"time"

	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type TriggerProperty interface {
	DateTimePropertyType
	DatePropertyType
}

type TriggerParameters struct {
	AlarmTriggerRelationship parameters.AlarmTriggerRelationshipParam
}

func NewTriggerProperty(
	dateTimeValue time.Time,
	durationValue string,
	format types.DTFormat,
	params ...parameters.Parameter,
) TriggerProperty {
	// Get the VALUE param
	valueType := string(registry.DATETIME)
	for i := 0; i < len(params); i++ {
		if params[i].GetParamName() == registry.VALUE {
			valueType = params[i].GetParamValue()
		}
	}
	switch valueType {
	case string(registry.DATETIME):
		return &dateTimePropertyType{
			PropName:   registry.TRIGGER,
			Value:      types.NewDateTimeValue(dateTimeValue, format),
			Parameters: params,
		}
	case string(registry.DURATION):
		return &durationPropertyType{
			PropName:   registry.TRIGGER,
			Value:      types.NewDurationValue(durationValue),
			Parameters: params,
		}
	default:
		return nil
	}
}
