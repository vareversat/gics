package properties

import (
	"fmt"
	"time"

	"github.com/vareversat/gics/parameters"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type DateTimeStartProperty interface {
	DateTimePropertyType
	DatePropertyType
}

// NewDateTimeStartProperty create a new registries.DateTimeStartProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional, mandatory if registries.RecurrenceRuleProp is set or registries.MethodProp is not set)
// - registries.Vtodo (Optional, mandatory if registries.RecurrenceRuleProp)
// - registries.Vfreebusy (Optional, mandatory if registries.RecurrenceRuleProp)
// - registries.Standard (Optional)
// - registries.Daylight (Optional)
// Optional parameters :
// - registries.ValueParam
// - registries.TimeZoneIdParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.4
func NewDateTimeStartProperty(
	value time.Time,
	params ...parameters.Parameter) DateTimeStartProperty {
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
			PropName:   registries.DateTimeStartProp,
			Value:      types.NewDateTimeValue(value),
			Parameters: params,
		}
	case string(registries.Date):
		return &datePropertyType{
			PropName:   registries.DateTimeStartProp,
			Value:      types.NewDateValue(value),
			Parameters: params,
		}
	default:
		return nil
	}
}

// NewDateTimeStartPropertyFromString create a new registries.DateTimeStartProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional, mandatory if registries.RecurrenceRuleProp is set or registries.MethodProp is not set)
// - registries.Vtodo (Optional, mandatory if registries.RecurrenceRuleProp)
// - registries.Vfreebusy (Optional, mandatory if registries.RecurrenceRuleProp)
// - registries.Standard (Optional)
// - registries.Daylight (Optional)
// Optional parameters :
// - registries.ValueParam
// - registries.TimeZoneIdParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.4
func NewDateTimeStartPropertyFromString(
	value string,
	params ...parameters.Parameter) (DateTimeStartProperty, error) {
	// Get the VALUE param
	valueType := string(registries.DateTime)
	for i := 0; i < len(params); i++ {
		if params[i].GetParamName() == registries.ValueDataTypesParam {
			valueType = params[i].GetParamValue()
		}
	}
	switch valueType {
	case string(registries.Date):
		dateValue, err := types.NewDateValueFromString(value)
		if err != nil {
			return nil, fmt.Errorf("cannot parse %s to a DATE value: %s", dateValue, err.Error())
		} else {
			return &datePropertyType{
				PropName:   registries.DateTimeStartProp,
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
				PropName:   registries.DateTimeStartProp,
				Value:      dateTimeValue,
				Parameters: params,
			}, nil
		}
	// If there is no VALUE parameter specified, DATE-TIME is the default type
	default:
		dateTimeValue, err := types.NewDateTimeValueFromString(value)
		if err != nil {
			return nil, fmt.Errorf(
				"cannot parse %s to a DATE-TIME value: %s",
				dateTimeValue,
				err.Error(),
			)
		} else {
			return &dateTimePropertyType{
				PropName:   registries.DateTimeStartProp,
				Value:      dateTimeValue,
				Parameters: params,
			}, nil
		}

	}
}
