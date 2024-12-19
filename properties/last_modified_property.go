package properties

import (
	"time"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type LastModifiedProperty interface {
	DateTimePropertyType
}

// NewLastModifiedProperty create a new registries.LastModifiedProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Vjournal (Optional)
// - registries.Vtimezone (Optional)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.3
func NewLastModifiedProperty(value time.Time) LastModifiedProperty {
	return &dateTimePropertyType{
		PropName: registries.LastModifiedProp,
		Value:    types.NewDateTimeValue(value),
	}
}

// NewLastModifiedPropertyFromString create a new registries.LastModifiedProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Vjournal (Optional)
// - registries.Vtimezone (Optional)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.3
func NewLastModifiedPropertyFromString(value string) LastModifiedProperty {
	return &dateTimePropertyType{
		PropName: registries.LastModifiedProp,
		Value:    types.NewDateTimeValueFromString(value),
	}
}
