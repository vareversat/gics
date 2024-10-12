package properties

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/types"
)

type NonStandardProperty interface {
	NonStandardPropertyType
}

// NewNonStandardProperty create a new non-standard property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vcalendar (Optional)
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Vjournal (Optional)
// - registries.Vfreebusy (Optional)
// - registries.Vtimezone (Optional)
// - registries.Valarm (Optional)
// - registries.Standard (Optional)
// - registries.Daylight (Optional)
// Optional parameters :
// - any
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.8.2
func NewNonStandardProperty(
	PropName string,
	value string,
	params ...parameters.Parameter,
) NonStandardProperty {
	return &nonStandardPropertyType{
		PropName:   PropName,
		Value:      types.NewTextValue(value),
		Parameters: params,
	}
}
