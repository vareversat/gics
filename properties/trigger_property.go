package properties

import (
	"fmt"
	"time"

	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type TriggerProperty interface {
	DateTimePropertyType
	DurationPropertyType
}

type TriggerParameters struct {
	AlarmTriggerRelationship parameters.AlarmTriggerRelationshipParam
}

// NewTriggerFromDateTimeProperty create a new [registries.TriggerProp] property. See [RFC-5545] ref for more info
// Usage :
// - registries.Valarm (Optional)
// Optional parameters :
// - registries.ValueParam
// - registries.TimeZoneIdParam
// - registries.AlarmTriggerRelationshipParam
// - registries.AlarmTriggerRelationshipParam (Mandatory when use of registries.Duration type)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.3
func NewTriggerFromDateTimeProperty(
	dateTime time.Time,
	params ...parameters.Parameter,
) (TriggerProperty, error) {

	var valueType string
	for i := 0; i < len(params); i++ {
		if params[i].GetParamName() == registries.ValueDataTypesParam {
			valueType = params[i].GetParamValue()
		}
	}

	if valueType == "" {
		return nil, fmt.Errorf("you have to explicitly set the value type via the VALUE parameter")
	}
	if valueType != string(registries.DateTime) {
		return nil, fmt.Errorf("VALUE parameter must be set to '%s'", string(registries.DateTime))
	}

	return &dateTimePropertyType{
		PropName:   registries.TriggerProp,
		Value:      types.NewDateTimeValue(dateTime),
		Parameters: params,
	}, nil
}

// NewTriggerFromDurationProperty create a new [registries.TriggerProp] property. See [RFC-5545] ref for more info
// Usage :
// - registries.Valarm (Optional)
// Optional parameters :
// - registries.ValueParam
// - registries.TimeZoneIdParam
// - registries.AlarmTriggerRelationshipParam
// - registries.AlarmTriggerRelationshipParam (Mandatory when use of registries.Duration type)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.3
func NewTriggerFromDurationProperty(
	duration time.Duration,
	params ...parameters.Parameter,
) (TriggerProperty, error) {

	var valueType string
	for i := 0; i < len(params); i++ {
		if params[i].GetParamName() == registries.ValueDataTypesParam {
			valueType = params[i].GetParamValue()
		}
	}

	if valueType != "" && valueType != string(registries.Duration) {
		return nil, fmt.Errorf(
			"VALUE parameter must be set to '%s' or removed",
			string(registries.DateTime),
		)
	}

	return &durationPropertyType{
		PropName:   registries.TriggerProp,
		Value:      types.NewDurationValue(duration),
		Parameters: params,
	}, nil
}

// NewTriggerFromDurationProperty create a new [registries.TriggerProp] property. See [RFC-5545] ref for more info
// Usage :
// - registries.Valarm (Optional)
// Optional parameters :
// - registries.ValueParam
// - registries.TimeZoneIdParam
// - registries.AlarmTriggerRelationshipParam
// - registries.AlarmTriggerRelationshipParam (Mandatory when use of registries.Duration type)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.6.3
func NewTriggerPropertyFromString(
	value string,
	params ...parameters.Parameter,
) (TriggerProperty, error) {
	var valueType string
	for i := 0; i < len(params); i++ {
		if params[i].GetParamName() == registries.ValueDataTypesParam {
			valueType = params[i].GetParamValue()
		}
	}

	switch valueType {
	case string(registries.Duration):
		dateValue, err := types.NewDurationValueFromString(value)
		if err != nil {
			return nil, fmt.Errorf(
				"cannot parse %s to a DURATION value: %s",
				dateValue,
				err.Error(),
			)
		} else {
			return &durationPropertyType{
				PropName:   registries.TriggerProp,
				Value:      dateValue,
				Parameters: params,
			}, nil
		}
	case string(registries.DateTime):
		dateTimeValue, err := types.NewDateTimeValueFromString(value)
		if err != nil {
			return nil, fmt.Errorf(
				"cannot parse %s to a DATE-TIME value: %s",
				dateTimeValue,
				err.Error(),
			)
		} else {
			return &dateTimePropertyType{
				PropName:   registries.TriggerProp,
				Value:      dateTimeValue,
				Parameters: params,
			}, nil
		}
	// If there is no VALUE parameter specified, DURATION is the default type
	default:
		dateValue, err := types.NewDurationValueFromString(value)
		if err != nil {
			return nil, fmt.Errorf(
				"cannot parse %s to a DURATION value: %s",
				dateValue,
				err.Error(),
			)
		} else {
			return &durationPropertyType{
				PropName:   registries.TriggerProp,
				Value:      dateValue,
				Parameters: params,
			}, nil
		}
	}
}
