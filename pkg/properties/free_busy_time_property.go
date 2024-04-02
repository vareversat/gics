package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.6

import (
	"time"

	"github.com/vareversat/gics/pkg/parameters"
	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type FreeBusyTimeProperty interface {
	PeriodPropertyType
}

func NewFreeBusyTimeProperty(
	startTime time.Time,
	endTime time.Time,
	freeBusyTimeParam parameters.FreeBusyTimeParam,
) FreeBusyTimeProperty {
	paramSlice := make(parameters.Parameters, 0)
	paramSlice = append(paramSlice, freeBusyTimeParam)
	return &periodPropertyType{
		PropName:   registries.FREEBUSY,
		Value:      types.NewPeriodValue(startTime, endTime),
		Parameters: paramSlice,
	}
}
