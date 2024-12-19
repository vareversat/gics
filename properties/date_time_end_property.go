package properties

import (
	"time"

	"github.com/vareversat/gics/parameters"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type DateTimeEndProperty interface {
	DateTimePropertyType
	DatePropertyType
}

// NewDateTimeEndProperty create a new registries.DateTimeEndProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional & must not appear at the same time as registries.PropertyRegistry)
// - registries.Vfreebusy (Optional)
// Optional parameters :
// - registries.ValueParam
// - registries.TimeZoneIdParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.2
func NewDateTimeEndProperty(
	value time.Time,
	params ...parameters.Parameter) DateTimeEndProperty {
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
			PropName:   registries.DateTimeEndProp,
			Value:      types.NewDateTimeValue(value),
			Parameters: params,
		}
	case string(registries.Date):
		return &datePropertyType{
			PropName:   registries.DateTimeEndProp,
			Value:      types.NewDateValue(value),
			Parameters: params,
		}
	default:
		return nil
	}
}

// NewDateTimeEndPropertyFromString create a new registries.DateTimeEndProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional & must not appear at the same time as registries.PropertyRegistry)
// - registries.Vfreebusy (Optional)
// Optional parameters :
// - registries.ValueParam
// - registries.TimeZoneIdParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.2
func NewDateTimeEndPropertyFromString(
	value string,
	params ...parameters.Parameter) DateTimeEndProperty {
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
			PropName:   registries.DateTimeEndProp,
			Value:      types.NewDateTimeValueFromString(value),
			Parameters: params,
		}
	case string(registries.Date):
		return &datePropertyType{
			PropName:   registries.DateTimeEndProp,
			Value:      types.NewDateValueFromString(value),
			Parameters: params,
		}
	default:
		return nil
	}
}
