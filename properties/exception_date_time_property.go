package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.5.1

import (
	"time"

	"github.com/vareversat/gics/utils"

	"github.com/vareversat/gics/parameters"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type ExceptionDateTimeProperty interface {
	DateTimePropertyType
	DatePropertyType
	PeriodPropertyType
}

func NewExceptionDateTimeProperty(
	values []time.Time,
	format types.DTFormat,
	params ...parameters.Parameter) ExceptionDateTimeProperty {
	valueType := string(registries.DateTime)
	for i := 0; i < len(params); i++ {
		if params[i].GetParamName() == registries.ValueParam {
			valueType = params[i].GetParamValue()
		}
	}
	switch valueType {
	case string(registries.DateTime):
		return &dateTimePropertyType{
			PropName:   registries.ExceptionDateTimesProp,
			Values:     types.NewDateTimeValues(values, format),
			Parameters: params,
		}
	case string(registries.Date):
		return &datePropertyType{
			PropName:   registries.ExceptionDateTimesProp,
			Values:     types.NewDateValues(values),
			Parameters: params,
		}
	default:
		return nil
	}
}

func NewExceptionDateTimePropertyFromString(
	values string,
	format types.DTFormat,
	params ...parameters.Parameter) ExceptionDateTimeProperty {
	// Get the VALUE param
	valueType := string(registries.DateTime)
	for i := 0; i < len(params); i++ {
		if params[i].GetParamName() == registries.ValueParam {
			valueType = params[i].GetParamValue()
		}
	}
	switch valueType {
	case string(registries.DateTime):
		return &dateTimePropertyType{
			PropName: registries.ExceptionDateTimesProp,
			Values: types.NewDateTimeValuesFromString(
				utils.StringToStringArray(values),
				format,
			),
			Parameters: params,
		}
	case string(registries.Date):
		return &datePropertyType{
			PropName:   registries.ExceptionDateTimesProp,
			Values:     types.NewDateValuesFromString(utils.StringToStringArray(values)),
			Parameters: params,
		}
	default:
		return nil
	}
}
