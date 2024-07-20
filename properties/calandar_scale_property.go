package properties

// https://datatracker.ietf.org/doc/html/rfc5545#section-3.7.1

import (
	"github.com/vareversat/gics/parameters"
	"github.com/vareversat/gics/registry"
	"github.com/vareversat/gics/types"
)

// CalendarScaleValue This is the default value of the CALSCALE property
const CalendarScaleValue = "GREGORIAN"

type CalendarScaleProperty interface {
	TextPropertyType
}

func NewCalScaleProperty(params ...parameters.Parameter) CalendarScaleProperty {
	return &textPropertyType{
		Parameters: params,
		PropName:   registry.CALSCALE,
		Value: types.NewTextValue(
			CalendarScaleValue),
	}
}
