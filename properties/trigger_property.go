package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.3

import (
	"time"

	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
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
	valueType := string(registries.DATETIME)
	for i := 0; i < len(params); i++ {
		if params[i].GetParamName() == registries.VALUE {
			valueType = params[i].GetParamValue()
		}
	}
	switch valueType {
	case string(registries.DATETIME):
		return &dateTimePropertyType{
			PropName:   registries.TRIGGER,
			Value:      types.NewDateTimeValue(dateTimeValue, format),
			Parameters: params,
		}
	case string(registries.DURATION):
		return &durationPropertyType{
			PropName:   registries.TRIGGER,
			Value:      types.NewDurationValue(durationValue),
			Parameters: params,
		}
	default:
		return nil
	}
}
