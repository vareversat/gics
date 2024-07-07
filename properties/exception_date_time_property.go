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
	valueType := string(registries.DATETIME)
	for i := 0; i < len(params); i++ {
		if params[i].GetParamName() == registries.VALUE {
			valueType = params[i].GetParamValue()
		}
	}
	switch valueType {
	case string(registries.DATETIME):
		return &dateTimePropertyType{
			PropName:   registries.EXDATE,
			Values:     types.NewDateTimeValues(values, format),
			Parameters: params,
		}
	case string(registries.DATE):
		return &datePropertyType{
			PropName:   registries.EXDATE,
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
	valueType := string(registries.DATETIME)
	for i := 0; i < len(params); i++ {
		if params[i].GetParamName() == registries.VALUE {
			valueType = params[i].GetParamValue()
		}
	}
	switch valueType {
	case string(registries.DATETIME):
		return &dateTimePropertyType{
			PropName: registries.EXDATE,
			Values: types.NewDateTimeValuesFromString(
				utils.StringToStringArray(values),
				format,
			),
			Parameters: params,
		}
	case string(registries.DATE):
		return &datePropertyType{
			PropName:   registries.EXDATE,
			Values:     types.NewDateValuesFromString(utils.StringToStringArray(values)),
			Parameters: params,
		}
	default:
		return nil
	}
}
