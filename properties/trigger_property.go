package properties

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

// NewTriggerProperty create a new registries.TriggerProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Valarm (Optional)
// Optional parameters :
// - registries.ValueParam
// - registries.TimeZoneIdParam
// - registries.AlarmTriggerRelationshipParam
// - registries.AlarmTriggerRelationshipParam (Mandatory when use of registries.Duration type)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.3
func NewTriggerProperty(
	dateTimeValue time.Time,
	durationValue string,
	params ...parameters.Parameter,
) TriggerProperty {
	// Get the VALUE param
	valueType := string(registries.DateTime)
	for i := 0; i < len(params); i++ {
		if params[i].GetParamName() == registries.ValueDataTypesParam {
			valueType = params[i].GetParamValue()
		}
	}
	switch valueType {
	case string(registries.DateTime):
		return &dateTimePropertyType{
			PropName:   registries.TriggerProp,
			Value:      types.NewDateTimeValue(dateTimeValue),
			Parameters: params,
		}
	case string(registries.Duration):
		duration, _ := types.NewDurationValueFromString(durationValue)
		return &durationPropertyType{
			PropName:   registries.TriggerProp,
			Value:      duration,
			Parameters: params,
		}
	default:
		return nil
	}
}
