package properties

import (
	"fmt"
	"time"

	"github.com/vareversat/gics/parameters"
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
func NewLastModifiedProperty(value time.Time, params ...parameters.Parameter) LastModifiedProperty {
	return &dateTimePropertyType{
		PropName:   registries.LastModifiedProp,
		Value:      types.NewDateTimeValue(value),
		Parameters: params,
	}
}

// NewLastModifiedPropertyFromString create a new registries.LastModifiedProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vevent (Optional)
// - registries.Vtodo (Optional)
// - registries.Vjournal (Optional)
// - registries.Vtimezone (Optional)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.3
func NewLastModifiedPropertyFromString(
	value string,
	params ...parameters.Parameter,
) (LastModifiedProperty, error) {
	date, err := types.NewDateTimeValueFromString(value)
	if err != nil {
		return nil, fmt.Errorf(
			"cannot parse %s to a DATE-TIME value: %s",
			date,
			err.Error(),
		)
	}
	return &dateTimePropertyType{
		PropName:   registries.LastModifiedProp,
		Value:      date,
		Parameters: params,
	}, nil
}
