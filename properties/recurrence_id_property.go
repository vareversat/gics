package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.4.4

import (
	"time"

	"github.com/vareversat/gics/parameters"

	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type RecurrenceIdProperty interface {
	DateTimePropertyType
	DatePropertyType
}

func NewRecurrenceIdProperty(
	value time.Time,
	format types.DTFormat,
	params ...parameters.Parameter) RecurrenceIdProperty {
	// Get the VALUE param
	valueType := string(registry.DATETIME)
	for i := 0; i < len(params); i++ {
		if params[i].GetParamName() == registry.Value {
			valueType = params[i].GetParamValue()
		}
	}
	switch valueType {
	case string(registry.DATETIME):
		return &dateTimePropertyType{
			PropName:   registry.RECURRENCEID,
			Value:      types.NewDateTimeValue(value, format),
			Parameters: params,
		}
	case string(registry.DATE):
		return &datePropertyType{
			PropName:   registry.RECURRENCEID,
			Value:      types.NewDateValue(value),
			Parameters: params,
		}
	default:
		return nil
	}
}

func NewRecurrenceIdPropertyFromString(
	value string,
	format types.DTFormat,
	params ...parameters.Parameter) RecurrenceIdProperty {
	// Get the VALUE param
	valueType := string(registry.DATETIME)
	for i := 0; i < len(params); i++ {
		if params[i].GetParamName() == registry.Value {
			valueType = params[i].GetParamValue()
		}
	}
	switch valueType {
	case string(registry.DATETIME):
		return &dateTimePropertyType{
			PropName:   registry.RECURRENCEID,
			Value:      types.NewDateTimeValueFromString(value, format),
			Parameters: params,
		}
	case string(registry.DATE):
		return &datePropertyType{
			PropName:   registry.RECURRENCEID,
			Value:      types.NewDateValueFromString(value),
			Parameters: params,
		}
	default:
		return nil
	}
}
