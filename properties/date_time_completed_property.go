package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.1

import (
	"fmt"
	"time"

	"github.com/vareversat/gics/parameters"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type DateTimeCompletedProperty interface {
	DateTimePropertyType
}

// NewDateTimeCompletedProperty create a new registries.CompletedProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vtodo (Optional)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.2
func NewDateTimeCompletedProperty(
	value time.Time,
) DateTimeCompletedProperty {
	return &dateTimePropertyType{
		PropName: registries.CompletedProp,
		Value:    types.NewDateTimeValue(value),
	}
}

// NewDateTimeCompletedPropertyFromString create a new registries.CompletedProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vtodo (Optional)
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.2
func NewDateTimeCompletedPropertyFromString(
	value string,
	params ...parameters.Parameter) (DateTimeCompletedProperty, error) {
	date, err := types.NewDateTimeValueFromString(value)
	if err != nil {
		return nil, fmt.Errorf(
			"cannot parse %s to a DATE-TIME value: %s",
			date,
			err.Error(),
		)
	}
	return &dateTimePropertyType{
		PropName:   registries.CompletedProp,
		Value:      date,
		Parameters: params,
	}, nil
}
