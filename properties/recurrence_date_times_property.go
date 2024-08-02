package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.2

import (
	"time"

	"github.com/vareversat/gics/parameters"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
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
	valueType := string(registries.DateTime)
	for i := 0; i < len(params); i++ {
		if params[i].GetParamName() == registries.ValueParam {
			valueType = params[i].GetParamValue()
		}
	}
	switch valueType {
	case string(registries.Period):
		return &periodPropertyType{
			PropName:   registries.RecurrenceDateTimes,
			Values:     types.NewPeriodValues(startValues, endValues),
			Parameters: params,
		}
	case string(registries.DateTime):
		return &dateTimePropertyType{
			PropName:   registries.RecurrenceDateTimes,
			Values:     types.NewDateTimeValues(startValues, format),
			Parameters: params,
		}
	case string(registries.Date):
		return &datePropertyType{
			PropName:   registries.RecurrenceDateTimes,
			Values:     types.NewDateValues(startValues),
			Parameters: params,
		}
	default:
		return nil
	}
}
