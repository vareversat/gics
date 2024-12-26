package properties

import (
	"fmt"
	"time"

	"github.com/vareversat/gics/parameters"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type RecurrenceDateTimesProperty interface {
	DatePropertyType
	DateTimePropertyType
	PeriodPropertyType
}

// NewRecurrenceDateTimesFromDateProperty create a new registries.RecurrenceDateTimesProp property from a DATE ([time.Time]) value . See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Vjournal (Optional)
// - registries.Standard (Optional)
// - registries.Daylight (Optional)
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.2
func NewRecurrenceDateTimesFromDateProperty(
	date time.Time,
	params ...parameters.Parameter,
) (RecurrenceDateTimesProperty, error) {

	var valueType string
	for i := 0; i < len(params); i++ {
		if params[i].GetParamName() == registries.ValueDataTypesParam {
			valueType = params[i].GetParamValue()
		}
	}

	if valueType == "" {
		return nil, fmt.Errorf("you have to explicitly set the value type via the VALUE parameter")
	}
	if valueType != string(registries.Date) {
		return nil, fmt.Errorf("VALUE parameter must be set to '%s'", string(registries.Date))
	}

	return &datePropertyType{
		PropName:   registries.RecurrenceDateTimesProp,
		Value:      types.NewDateValue(date),
		Parameters: params,
	}, nil
}

// NewRecurrenceDateTimesFromDateTimeProperty create a new registries.RecurrenceDateTimesProp property from a DATE-TIME ([time.Time]) value. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Vjournal (Optional)
// - registries.Standard (Optional)
// - registries.Daylight (Optional)
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.2
func NewRecurrenceDateTimesFromDateTimeProperty(
	dateTime time.Time,
	params ...parameters.Parameter,
) (RecurrenceDateTimesProperty, error) {

	var valueType string
	for i := 0; i < len(params); i++ {
		if params[i].GetParamName() == registries.ValueDataTypesParam {
			valueType = params[i].GetParamValue()
		}
	}

	if valueType != "" && valueType != string(registries.DateTime) {
		return nil, fmt.Errorf(
			"VALUE parameter must be set to '%s' or removed",
			string(registries.DateTime),
		)
	}

	return &dateTimePropertyType{
		PropName:   registries.RecurrenceDateTimesProp,
		Value:      types.NewDateTimeValue(dateTime),
		Parameters: params,
	}, nil
}

// NewRecurrenceDateTimesFromImplicitPeriodProperty create a new registries.RecurrenceDateTimesProp property from an implicit PERIOD ([time.Time]/[time.Duration]) value. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Vjournal (Optional)
// - registries.Standard (Optional)
// - registries.Daylight (Optional)
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.2
func NewRecurrenceDateTimesFromImplicitPeriodProperty(
	from time.Time,
	to time.Duration,
	params ...parameters.Parameter,
) (RecurrenceDateTimesProperty, error) {
	var valueType string
	for i := 0; i < len(params); i++ {
		if params[i].GetParamName() == registries.ValueDataTypesParam {
			valueType = params[i].GetParamValue()
		}
	}

	if valueType == "" {
		return nil, fmt.Errorf("you have to explicitly set the value type via the VALUE parameter")
	}
	if valueType != string(registries.PeriodOfTime) {
		return nil, fmt.Errorf(
			"VALUE parameter must be set to '%s'",
			string(registries.PeriodOfTime),
		)
	}

	return &periodPropertyType{
		PropName:   registries.RecurrenceDateTimesProp,
		Value:      types.NewImplicitPeriodOfTimeValue(from, to),
		Parameters: params,
	}, nil
}

// NewRecurrenceDateTimesFromExplicitPeriodProperty create a new registries.RecurrenceDateTimesProp property from an implicit PERIOD ([time.Time]/[time.Time]) value. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Vjournal (Optional)
// - registries.Standard (Optional)
// - registries.Daylight (Optional)
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.2
func NewRecurrenceDateTimesFromExplicitPeriodProperty(
	from time.Time,
	to time.Time,
	params ...parameters.Parameter,
) (RecurrenceDateTimesProperty, error) {
	var valueType string
	for i := 0; i < len(params); i++ {
		if params[i].GetParamName() == registries.ValueDataTypesParam {
			valueType = params[i].GetParamValue()
		}
	}

	if valueType == "" {
		return nil, fmt.Errorf("you have to explicitly set the value type via the VALUE parameter")
	}
	if valueType != string(registries.PeriodOfTime) {
		return nil, fmt.Errorf(
			"VALUE parameter must be set to '%s'",
			string(registries.PeriodOfTime),
		)
	}

	return &periodPropertyType{
		PropName:   registries.RecurrenceDateTimesProp,
		Value:      types.NewExplicitPeriodOfTimeValue(from, to),
		Parameters: params,
	}, nil
}

// NewRecurrenceDateTimesPropertyFromString create a new registries.RecurrenceDateTimesProp property from string. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Vjournal (Optional)
// - registries.Standard (Optional)
// - registries.Daylight (Optional)
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.2
func NewRecurrenceDateTimesPropertyFromString(
	value string,
	params ...parameters.Parameter,
) (RecurrenceDateTimesProperty, error) {
	var valueType string
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
	case string(registries.PeriodOfTime):
		periodValue, err := types.NewPeriodOfTimeValueFromString(value)
		if err != nil {
			return nil, fmt.Errorf(
				"cannot parse %s to a PERIOD value: %s",
				periodValue,
				err.Error(),
			)
		} else {
			return &periodPropertyType{
				PropName:   registries.RecurrenceDateTimesProp,
				Value:      periodValue,
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
