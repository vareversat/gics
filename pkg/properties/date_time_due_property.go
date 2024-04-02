package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.8.2.3

import (
	"time"

	"github.com/vareversat/gics/pkg/parameters"

	"github.com/vareversat/gics/pkg/registries"
	"github.com/vareversat/gics/pkg/types"
)

type DateTimeDueProperty interface {
	DateTimePropertyType
	DatePropertyType
}

func NewDateTimeDueProperty(
	value time.Time,
	format types.DTFormat,
	tzIdentifier string,
	valueType registries.Type,
) DateTimeDueProperty {
	paramSlice := make(parameters.Parameters, 0)
	paramSlice = append(paramSlice, parameters.NewTimeZoneIdentifierParam(tzIdentifier))
	paramSlice = append(paramSlice, parameters.NewValueParam(valueType))
	switch valueType {
	case registries.DATETIME:
		return &dateTimePropertyType{
			PropName:   registries.DUE,
			Value:      types.NewDateTimeValue(value, format),
			Parameters: paramSlice,
		}
	case registries.DATE:
		return &datePropertyType{
			PropName:   registries.DUE,
			Value:      types.NewDateValue(value),
			Parameters: paramSlice,
		}
	default:
		return nil
	}
}
