package properties

import (
	"time"

	"github.com/vareversat/gics/parameters"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type DateTimeDueProperty interface {
	DateTimePropertyType
	DatePropertyType
}

// NewDateTimeDueProperty create a new registries.DateTimeDueProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vtodo (Optional)
// Optional parameters :
// - registries.ValueParam
// - registries.TimeZoneIdParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.3
func NewDateTimeDueProperty(
	value time.Time,
	format types.DTFormat,
	params ...parameters.Parameter) DateTimeDueProperty {
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
			PropName:   registries.DateTimeDueProp,
			Value:      types.NewDateTimeValue(value, format),
			Parameters: params,
		}
	case string(registries.Date):
		return &datePropertyType{
			PropName:   registries.DateTimeDueProp,
			Value:      types.NewDateValue(value),
			Parameters: params,
		}
	default:
		return nil
	}
}

// NewDateTimeDuePropertyFromString create a new registries.DateTimeDueProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vtodo (Optional)
// Optional parameters :
// - registries.ValueParam
// - registries.TimeZoneIdParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.3
func NewDateTimeDuePropertyFromString(
	value string,
	format types.DTFormat,
	params ...parameters.Parameter) DateTimeDueProperty {
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
			PropName:   registries.DateTimeDueProp,
			Value:      types.NewDateTimeValueFromString(value, format),
			Parameters: params,
		}
	case string(registries.Date):
		return &datePropertyType{
			PropName:   registries.DateTimeDueProp,
			Value:      types.NewDateValueFromString(value),
			Parameters: params,
		}
	default:
		return nil
	}
}
