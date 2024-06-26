package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.7.2

import (
	"time"

	"github.com/vareversat/gics/pkg/parameters"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
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
	valueType := string(registries.DATETIME)
	for i := 0; i < len(params); i++ {
		if params[i].GetParamName() == registries.VALUE {
			valueType = params[i].GetParamValue()
		}
	}
	switch valueType {
	case string(registries.DATETIME):
		return &dateTimePropertyType{
			PropName:   registries.DTSTART,
			Value:      types.NewDateTimeValue(value, format),
			Parameters: params,
		}
	case string(registries.DATE):
		return &datePropertyType{
			PropName:   registries.DTSTART,
			Value:      types.NewDateValue(value),
			Parameters: params,
		}
	default:
		return nil
	}
}
