package properties

import (
	"time"

	"github.com/vareversat/gics/utils"

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
	params ...parameters.Parameter) ExceptionDateTimeProperty {
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
			PropName: registries.ExceptionDateTimesProp,
			Values: types.NewDateTimeValueFromStrings(
				utils.StringToStringArray(value),
			),
			Parameters: params,
		}
	case string(registries.Date):
		return &datePropertyType{
			PropName:   registries.ExceptionDateTimesProp,
			Values:     types.NewDateValuesFromString(utils.StringToStringArray(value)),
			Parameters: params,
		}
	default:
		return nil
	}
}
