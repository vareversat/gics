package properties

import (
	"time"

	"github.com/vareversat/gics/parameters"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type RecurrenceDateTimesProperty interface {
	DateTimePropertyType
	DatePropertyType
	PeriodPropertyType
}

// NewRecurrenceDateTimesProperty create a new registries.RecurrenceDateTimesProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Vjournal (Optional)
// - registries.Standard (Optional)
// - registries.Daylight (Optional)
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.2
func NewRecurrenceDateTimesProperty(
	startValues []time.Time,
	endValues []time.Time,
	params ...parameters.Parameter,
) RecurrenceDateTimesProperty {
	// Default valueType
	var valueType = string(registries.DateTime)
	for i := 0; i < len(params); i++ {
		if params[i].GetParamName() == registries.ValueDataTypesParam {
			valueType = params[i].GetParamValue()
		}
	}
	switch valueType {
	// case string(registries.PeriodOfTime):
	// 	return &periodPropertyType{
	// 		PropName:   registries.RecurrenceDateTimesProp,
	// 		Values:     types.NewPeriodValues(startValues, endValues),
	// 		Parameters: params,
	// 	}
	case string(registries.DateTime):
		return &dateTimePropertyType{
			PropName:   registries.RecurrenceDateTimesProp,
			Values:     types.NewDateTimeValues(startValues),
			Parameters: params,
		}
	case string(registries.Date):
		return &datePropertyType{
			PropName:   registries.RecurrenceDateTimesProp,
			Values:     types.NewDateValues(startValues),
			Parameters: params,
		}
	default:
		return nil
	}
}
