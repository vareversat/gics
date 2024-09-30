package properties

import (
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
	format types.DTFormat,
	params ...parameters.Parameter) RecurrenceIdProperty {
	// Get the VALUE param
	valueType := string(registries.DateTime)
	for i := 0; i < len(params); i++ {
		if params[i].GetParamName() == registries.ValueParam {
			valueType = params[i].GetParamValue()
		}
	}
	switch valueType {
	case string(registries.DateTime):
		return &dateTimePropertyType{
			PropName:   registries.RecurrenceIdProp,
			Value:      types.NewDateTimeValue(value, format),
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
	format types.DTFormat,
	params ...parameters.Parameter) RecurrenceIdProperty {
	// Get the VALUE param
	valueType := string(registries.DateTime)
	for i := 0; i < len(params); i++ {
		if params[i].GetParamName() == registries.ValueParam {
			valueType = params[i].GetParamValue()
		}
	}
	switch valueType {
	case string(registries.DateTime):
		return &dateTimePropertyType{
			PropName:   registries.RecurrenceIdProp,
			Value:      types.NewDateTimeValueFromString(value, format),
			Parameters: params,
		}
	case string(registries.Date):
		return &datePropertyType{
			PropName:   registries.RecurrenceIdProp,
			Value:      types.NewDateValueFromString(value),
			Parameters: params,
		}
	default:
		return nil
	}
}
