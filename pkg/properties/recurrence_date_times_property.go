package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.2

import (
	"time"

	"github.com/vareversat/gics/pkg/parameters"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type RecurrenceDateTimesProperty interface {
	DateTimePropertyType
	DatePropertyType
	PeriodPropertyType
}

func NewRecurrenceDateTimesProperty(
	startValues []time.Time,
	endValues []time.Time,
	format types.DTFormat,
	params ...parameters.Parameter,
) RecurrenceDateTimesProperty {
	valueType := string(registries.DATETIME)
	for i := 0; i < len(params); i++ {
		if params[i].GetParamName() == registries.VALUE {
			valueType = params[i].GetParamValue()
		}
	}
	switch valueType {
	case string(registries.PERIOD):
		return &periodPropertyType{
			PropName:   registries.RDATE,
			Values:     types.NewPeriodValues(startValues, endValues),
			Parameters: params,
		}
	case string(registries.DATETIME):
		return &dateTimePropertyType{
			PropName:   registries.RDATE,
			Values:     types.NewDateTimeValues(startValues, format),
			Parameters: params,
		}
	case string(registries.DATE):
		return &datePropertyType{
			PropName:   registries.RDATE,
			Values:     types.NewDateValues(startValues),
			Parameters: params,
		}
	default:
		return nil
	}
}
