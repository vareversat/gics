package properties

import (
	"fmt"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type DurationProperty interface {
	DurationPropertyType
}

// NewDurationPropertyFromString create a new registries.DurationProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Valarm (Optional)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.5
func NewDurationPropertyFromString(value string) (DurationProperty, error) {

	duration, err := types.NewDurationValueFromString(value)

	if err != nil {
		return nil, fmt.Errorf("%s is not a valid duration", value)
	}

	return &durationPropertyType{
		PropName: registries.DurationProp,
		Value:    duration,
	}, nil
}
