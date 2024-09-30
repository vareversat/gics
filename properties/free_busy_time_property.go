package properties

import (
	"time"

	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type FreeBusyTimeProperty interface {
	PeriodPropertyType
}

// NewFreeBusyTimeProperty create a new registries.FreeBusyTimeProp property. See [RFC-5545] ref for more info
// Usage :
// - registries.Vfreebusy (Optional)
// Optional parameters :
// - registries.FreeBusyTimeTypeParam
// [RFC-5545]: https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.6
func NewFreeBusyTimeProperty(
	startTime time.Time,
	endTime time.Time,
	freeBusyTimeParam parameters.FreeBusyTimeParam,
) FreeBusyTimeProperty {
	paramSlice := make(parameters.Parameters, 0)
	paramSlice = append(paramSlice, freeBusyTimeParam)
	return &periodPropertyType{
		PropName:   registries.FreeBusyTimeProp,
		Value:      types.NewPeriodValue(startTime, endTime),
		Parameters: paramSlice,
	}
}
