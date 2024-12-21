package parameters

import (
	"fmt"
	"time"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type TimeZoneIdentifierParam interface {
	TextTypeParameter
}

// NewTimeZoneIdentifierParam create a new registries.TimeZoneIdParam property. See [RFC-5545] ref for more info
// This parameter can be used in these properties :
// - registries.DateTimeEndProp
// - registries.DateTimeStartProp
// - registries.DateTimeDueProp
// - registries.RecurrenceIdProp
// - registries.RecurrenceIdProp
// - registries.ExceptionDateTimesProp
// - registries.RecurrenceDateTimesProp
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.2.19
func NewTimeZoneIdentifierParam(value string) (TimeZoneIdentifierParam, error) {
	_, err := time.LoadLocation(value)
	if err != nil {
		return nil, fmt.Errorf("%s is not a valid TZDATA", value)
	}
	return &textParameter{
		ParamName: registries.TimeZoneIdParam,
		Value:     types.NewTextValue(value),
	}, nil
}
