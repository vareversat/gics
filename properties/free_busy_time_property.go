package properties

import (
	"fmt"
	"time"

	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type FreeBusyTimeProperty interface {
	PeriodPropertyType
}

// NewImplicitFreeBusyTimeProperty create a new [registries.FreeBusyTimeProp] property with a [time.Time] as "from" and [time.Duration] as "to". See [RFC-5545] ref for more info
// Usage :
// - [registries.Vfreebusy] (Optional)
// Optional parameters :
// - [registries.FreeBusyTimeTypeParam]
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.6
func NewImplicitFreeBusyTimeProperty(
	from time.Time,
	to time.Duration,
	format types.PeriodOfTimeFormat,
	params ...parameters.Parameter,
) FreeBusyTimeProperty {
	return &periodPropertyType{
		PropName:   registries.FreeBusyTimeProp,
		Value:      types.NewImplicitPeriodOfTimeValue(from, to),
		Parameters: params,
	}
}

// NewFreeBusyTimeProperty create a [registries.FreeBusyTimeProp] property with a [time.Time] as "from" and [time.Time] as "to". See [RFC-5545] ref for more info
// Usage :
// - [registries.Vfreebusy] (Optional)
// Optional parameters :
// - [registries.FreeBusyTimeTypeParam]
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.6
func NewExplicitFreeBusyTimeProperty(
	from time.Time,
	to time.Time,
	params ...parameters.Parameter,
) FreeBusyTimeProperty {
	return &periodPropertyType{
		PropName:   registries.FreeBusyTimeProp,
		Value:      types.NewExplicitPeriodOfTimeValue(from, to),
		Parameters: params,
	}
}

// NewFreeBusyTimePropertyFromString create a [registries.FreeBusyTimeProp] property from a string. See [RFC-5545] ref for more info
// Usage :
// - [registries.Vfreebusy] (Optional)
// Optional parameters :
// - [registries.FreeBusyTimeTypeParam]
//
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.6
func NewFreeBusyTimePropertyFromString(
	value string,
	params ...parameters.Parameter,
) (FreeBusyTimeProperty, error) {
	periodOfTime, err := types.NewPeriodOfTimeValueFromString(value)

	if err != nil {
		return nil, fmt.Errorf("%s is not a valid free busy time property", value)
	}

	return &periodPropertyType{
		PropName:   registries.FreeBusyTimeProp,
		Value:      periodOfTime,
		Parameters: params,
	}, nil
}
