package properties

import (
	"fmt"
	"time"

	"github.com/vareversat/gics/parameters"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type RecurrenceIdProperty interface {
	DateTimePropertyType
	DatePropertyType
}

// NewRecurrenceIdProperty create a new registries.RecurrenceIdProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vcalendar (Mandatory if recurring calendar components)
// Optional parameters :
// - registries.ValueParam
// - registries.TimeZoneIdParam
// - registries.RecurrenceIdRangeParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.4
func NewRecurrenceIdProperty(
	value time.Time,
	params ...parameters.Parameter) RecurrenceIdProperty {
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
			PropName:   registries.RecurrenceIdProp,
			Value:      types.NewDateTimeValue(value),
			Parameters: params,
		}
	case string(registries.Date):
		return &datePropertyType{
			PropName:   registries.RecurrenceIdProp,
			Value:      types.NewDateValue(value),
			Parameters: params,
		}
	default:
		return nil
	}
}

// NewRecurrenceIdPropertyFromString create a new registries.RecurrenceIdProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vcalendar (Mandatory if recurring calendar components)
// Optional parameters :
// - registries.ValueParam
// - registries.TimeZoneIdParam
// - registries.RecurrenceIdRangeParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.4
func NewRecurrenceIdPropertyFromString(
	value string,
	params ...parameters.Parameter) (RecurrenceIdProperty, error) {
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
				PropName:   registries.RecurrenceDateTimesProp,
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
				PropName:   registries.RecurrenceDateTimesProp,
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
				PropName:   registries.RecurrenceDateTimesProp,
				Value:      dateTimeValue,
				Parameters: params,
			}, nil
		}

	}
}
