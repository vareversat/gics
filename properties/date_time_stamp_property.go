package properties

import (
	"time"

	"github.com/vareversat/gics/parameters"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type DateTimeStampProperty interface {
	DateTimePropertyType
}

// NewDateTimeStampProperty create a new registries.DateTimeStampProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Mandatory)
// - registries.Vtodo (Mandatory)
// - registries.Vjournal (Mandatory)
// - registries.Vfreebusy (Mandatory)
// Optional parameters :
// - registries.ValueParam
// - registries.TimeZoneIdParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.2
func NewDateTimeStampProperty(
	timeValue time.Time,
	params ...parameters.Parameter,
) DateTimeStampProperty {
	return &dateTimePropertyType{
		PropName:   registries.DateTimeStampProp,
		Value:      types.NewDateTimeValue(timeValue),
		Parameters: params,
	}
}

// NewDateTimeStampPropertyFromString create a new registries.DateTimeStampProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Mandatory)
// - registries.Vtodo (Mandatory)
// - registries.Vjournal (Mandatory)
// - registries.Vfreebusy (Mandatory)
// Optional parameters :
// - registries.ValueParam
// - registries.TimeZoneIdParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.2
func NewDateTimeStampPropertyFromString(
	value string,
	params ...parameters.Parameter) DateTimeStampProperty {
	return &dateTimePropertyType{
		PropName:   registries.DateTimeStampProp,
		Value:      types.NewDateTimeValueFromString(value),
		Parameters: params,
	}
}
