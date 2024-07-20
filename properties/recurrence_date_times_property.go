package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.2

import (
	"time"

	"github.com/vareversat/gics/parameters"

	"github.com/vareversat/gics/registry"
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
	valueType := string(registry.DATETIME)
	for i := 0; i < len(params); i++ {
		if params[i].GetParamName() == registry.Value {
			valueType = params[i].GetParamValue()
		}
	}
	switch valueType {
	case string(registry.PERIOD):
		return &periodPropertyType{
			PropName:   registry.RDATE,
			Values:     types.NewPeriodValues(startValues, endValues),
			Parameters: params,
		}
	case string(registry.DATETIME):
		return &dateTimePropertyType{
			PropName:   registry.RDATE,
			Values:     types.NewDateTimeValues(startValues, format),
			Parameters: params,
		}
	case string(registry.DATE):
		return &datePropertyType{
			PropName:   registry.RDATE,
			Values:     types.NewDateValues(startValues),
			Parameters: params,
		}
	default:
		return nil
	}
}
