package properties

import (
	"fmt"
	"time"

	"github.com/vareversat/gics/parameters"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type ExceptionDateTimeProperty interface {
	DateTimePropertyType
	DatePropertyType
	PeriodPropertyType
}

// NewExceptionDateTimeProperty create a new registries.ExceptionDateTimesProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Valarm (Optional)
// - registries.Vjournal (Optional)
// - registries.Standard (Optional)
// - registries.Daylight (Optional)
// Optional parameters :
// - registries.ValueParam
// - registries.TimeZoneIdParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.1
func NewExceptionDateTimeProperty(
	values []time.Time,
	params ...parameters.Parameter) ExceptionDateTimeProperty {
	valueType := string(registries.DateTime)
	for i := 0; i < len(params); i++ {
		if params[i].GetParamName() == registries.ValueDataTypesParam {
			valueType = params[i].GetParamValue()
		}
	}
	switch valueType {
	case string(registries.DateTime):
		return &dateTimePropertyType{
			PropName:   registries.ExceptionDateTimesProp,
			Values:     types.NewDateTimeValues(values),
			Parameters: params,
		}
	case string(registries.Date):
		return &datePropertyType{
			PropName:   registries.ExceptionDateTimesProp,
			Values:     types.NewDateValues(values),
			Parameters: params,
		}
	default:
		return nil
	}
}

// NewExceptionDateTimePropertyFromString create a new registries.ExceptionDateTimesProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Valarm (Optional)
// - registries.Vjournal (Optional)
// - registries.Standard (Optional)
// - registries.Daylight (Optional)
// Optional parameters :
// - registries.ValueParam
// - registries.TimeZoneIdParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.1
func NewExceptionDateTimePropertyFromString(
	value string,
	params ...parameters.Parameter) (ExceptionDateTimeProperty, error) {
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
				PropName:   registries.ExceptionDateTimesProp,
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
				PropName:   registries.ExceptionDateTimesProp,
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
				PropName:   registries.ExceptionDateTimesProp,
				Value:      dateTimeValue,
				Parameters: params,
			}, nil
		}

	}
}
