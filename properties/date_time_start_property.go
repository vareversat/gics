package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.2

import (
	"time"

	"github.com/vareversat/gics/parameters"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type DateTimeStartProperty interface {
	DateTimePropertyType
	DatePropertyType
}

func NewDateTimeStartProperty(
	value time.Time,
	format types.DTFormat,
	params ...parameters.Parameter) DateTimeStartProperty {
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
			PropName:   registries.DateTimeStartProp,
			Value:      types.NewDateTimeValue(value, format),
			Parameters: params,
		}
	case string(registries.Date):
		return &datePropertyType{
			PropName:   registries.DateTimeStartProp,
			Value:      types.NewDateValue(value),
			Parameters: params,
		}
	default:
		return nil
	}
}

func NewDateTimeStartPropertyFromString(
	value string,
	format types.DTFormat,
	params ...parameters.Parameter) DateTimeStartProperty {
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
			PropName:   registries.DateTimeStartProp,
			Value:      types.NewDateTimeValueFromString(value, format),
			Parameters: params,
		}
	case string(registries.Date):
		return &datePropertyType{
			PropName:   registries.DateTimeStartProp,
			Value:      types.NewDateValueFromString(value),
			Parameters: params,
		}
	default:
		return nil
	}
}
