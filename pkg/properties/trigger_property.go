package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.3

import (
	"time"

	"github.com/vareversat/gics/pkg/parameters"
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type TriggerProperty interface {
	DateTimePropertyType
	DatePropertyType
}

type triggerProperty struct {
	TriggerParameters
	PropName      registries.Properties
	DateTimeValue types.DateTimeValue
	DurationValue types.DurationValue
}

type TriggerParameters struct {
	AlarmTriggerRelationship parameters.AlarmTriggerRelationshipParam
}

func NewTriggerProperty(
	dateTimeValue time.Time,
	durationValue string,
	format types.DTFormat,
	valueType registries.Type,
	alarmTriggerRelationshipParam parameters.AlarmTriggerRelationshipType,
) TriggerProperty {

	paramSlice := make(parameters.Parameters, 0)
	paramSlice = append(paramSlice, parameters.NewValueParam(valueType))
	paramSlice = append(
		paramSlice,
		parameters.NewAlarmTriggerRelationshipParam(alarmTriggerRelationshipParam),
	)
	switch valueType {
	case registries.DATETIME:
		return &dateTimePropertyType{
			PropName:   registries.TRIGGER,
			Value:      types.NewDateTimeValue(dateTimeValue, format),
			Parameters: paramSlice,
		}
	case registries.DURATION:
		return &durationPropertyType{
			PropName:   registries.TRIGGER,
			Value:      types.NewDurationValue(durationValue),
			Parameters: paramSlice,
		}
	default:
		return nil
	}
}
