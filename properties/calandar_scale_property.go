package properties

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

// CalendarScaleValue This is the default value of the CALSCALE property
const CalendarScaleValue = "GREGORIAN"

type CalendarScaleProperty interface {
	TextPropertyType
}

// NewCalScaleProperty create a new registries.CalendarScaleProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Valarm (Optional)
// - registries.Vcalendar (Optional)
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Vfreebusy (Optional)
// - registries.Valarm (Optional)
// - registries.Vjournal (Optional)
// - registries.Standard (Optional)
// - registries.Daylight (Optional)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.7.1
func NewCalScaleProperty(params ...parameters.Parameter) CalendarScaleProperty {
	return &textPropertyType{
		Parameters: params,
		PropName:   registries.CalendarScaleProp,
		Value: types.NewTextValue(
			CalendarScaleValue),
	}
}
