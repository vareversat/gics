package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.3

import (
	"time"

	"github.com/vareversat/gics/parameters"

	"github.com/vareversat/gics/registries"
	"github.com/vareversat/gics/types"
)

type DateTimeDueProperty interface {
	DateTimePropertyType
	DatePropertyType
}

func NewDateTimeDueProperty(
	value time.Time,
	format types.DTFormat,
	params ...parameters.Parameter) DateTimeDueProperty {
	// Get the VALUE param
	valueType := string(registries.DateTime)
	for i := 0; i < len(params); i++ {
		if params[i].GetParamName() == registries.Value {
			valueType = params[i].GetParamValue()
		}
	}
	switch valueType {
	case string(registries.DateTime):
		return &dateTimePropertyType{
			PropName:   registries.DateTimeDue,
			Value:      types.NewDateTimeValue(value, format),
			Parameters: params,
		}
	case string(registries.Date):
		return &datePropertyType{
			PropName:   registries.DateTimeDue,
			Value:      types.NewDateValue(value),
			Parameters: params,
		}
	default:
		return nil
	}
}

func NewDateTimeDuePropertyFromString(
	value string,
	format types.DTFormat,
	params ...parameters.Parameter) DateTimeDueProperty {
	// Get the VALUE param
	valueType := string(registries.DateTime)
	for i := 0; i < len(params); i++ {
		if params[i].GetParamName() == registries.Value {
			valueType = params[i].GetParamValue()
		}
	}
	switch valueType {
	case string(registries.DateTime):
		return &dateTimePropertyType{
			PropName:   registries.DateTimeDue,
			Value:      types.NewDateTimeValueFromString(value, format),
			Parameters: params,
		}
	case string(registries.Date):
		return &datePropertyType{
			PropName:   registries.DateTimeDue,
			Value:      types.NewDateValueFromString(value),
			Parameters: params,
		}
	default:
		return nil
	}
}
