package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.2

import (
	"time"

	"github.com/vareversat/gics/parameters"

	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

type DateTimeEndProperty interface {
	DateTimePropertyType
	DatePropertyType
}

func NewDateTimeEndProperty(
	value time.Time,
	format types.DTFormat,
	params ...parameters.Parameter) DateTimeEndProperty {
	// Get the VALUE param
	valueType := string(registry.DateTime)
	for i := 0; i < len(params); i++ {
		if params[i].GetParamName() == registry.Value {
			valueType = params[i].GetParamValue()
		}
	}
	switch valueType {
	case string(registry.DateTime):
		return &dateTimePropertyType{
			PropName:   registry.DTEND,
			Value:      types.NewDateTimeValue(value, format),
			Parameters: params,
		}
	case string(registry.Date):
		return &datePropertyType{
			PropName:   registry.DTEND,
			Value:      types.NewDateValue(value),
			Parameters: params,
		}
	default:
		return nil
	}
}

func NewDateTimeEndPropertyFromString(
	value string,
	format types.DTFormat,
	params ...parameters.Parameter) DateTimeEndProperty {
	// Get the VALUE param
	valueType := string(registry.DateTime)
	for i := 0; i < len(params); i++ {
		if params[i].GetParamName() == registry.Value {
			valueType = params[i].GetParamValue()
		}
	}
	switch valueType {
	case string(registry.DateTime):
		return &dateTimePropertyType{
			PropName:   registry.DTEND,
			Value:      types.NewDateTimeValueFromString(value, format),
			Parameters: params,
		}
	case string(registry.Date):
		return &datePropertyType{
			PropName:   registry.DTEND,
			Value:      types.NewDateValueFromString(value),
			Parameters: params,
		}
	default:
		return nil
	}
}
